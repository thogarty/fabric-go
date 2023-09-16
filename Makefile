.PHONY: all gen patch fetch

CURRENT_UID := $(shell id -u)
CURRENT_GID := $(shell id -g)

# https://github.com/OpenAPITools/openapi-generator-cli
SPEC_URL:="https://api.swaggerhub.com/apis/equinix-api/fabric/4.9/swagger.yaml"

SPEC_FETCHED_FILE:=spec.fetched.yaml
SPEC_PATCHED_FILE:=spec.patched.yaml
OPENAPI_CODEGEN_TAG=v6.4.0
OPENAPI_CODEGEN_IMAGE=openapitools/openapi-generator-cli:${OPENAPI_CODEGEN_TAG} # to use, change gen: to gen-openapitools
SWAGGER_CODEGEN_IMAGE=swaggerapi/swagger-codegen-cli-v3:3.0.34
VALIDATE_IMAGE=openapitools/openapi-generator-cli

GIT_ORG=equinix-labs
GIT_REPO=fabric-go
PACKAGE_PREFIX=fabric
PACKAGE_MAJOR=v4
CRI=docker

DOCKER=${CRI} run --rm -u ${CURRENT_UID}:${CURRENT_GID} -v $(CURDIR):/local
DOCKER_SWAGGER=${DOCKER} ${SWAGGER_CODEGEN_IMAGE}
DOCKER_OPENAPI=${DOCKER} ${OPENAPI_CODEGEN_IMAGE}

IMAGE=${SWAGGER_CODEGEN_IMAGE}

GOLANGCI_LINT=golangci-lint


# Update the dependency on the gen job to switch between code generator methods.
# Options will switch between:
# Open API
# Swagger Codegen Docker
# Swagger Code Gen Java Jar for OSX Machines with Apple Silicon CPU M1/M2
all: pull fetch patch clean gen mod fmt patch-post docs move-other test stage


# Update the IMAGE variable above to either ${OPENAPI_CODEGEN_IMAGE} or ${SWAGGER_CODEGEN_IMAGE}
# depending on the job that you will be using to generate the Fabric Go SDK
pull:
	${CRI} pull ${IMAGE}

fetch:
	curl \
		-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36' \
		-o ${SPEC_FETCHED_FILE} ${SPEC_URL}

patch:
	# patch is idempotent, always starting with the fetched
	# fetched file to create the patched file.
	cp ${SPEC_FETCHED_FILE} ${SPEC_PATCHED_FILE}
	ARGS="-o ${SPEC_PATCHED_FILE} ${SPEC_FETCHED_FILE}"; \
	for diff in $(shell find patches/${SPEC_FETCHED_FILE} -name \*.patch | sort -n); do \
		patch --no-backup-if-mismatch -N -t $$ARGS $$diff; \
		ARGS=${SPEC_PATCHED_FILE}; \
	done

clean: clean-docs
	rm -rf ${PACKAGE_PREFIX}

gen-openapitools:
	${DOCKER_OPENAPI} generate -g go \
		--package-name ${PACKAGE_MAJOR} \
		--model-package types \
		--api-package models \
		--git-user-id ${GIT_ORG} \
		--git-repo-id ${GIT_REPO} \
		-o /local/${PACKAGE_PREFIX}/${PACKAGE_MAJOR} \
		-i /local/${SPEC_PATCHED_FILE}

gen-swagger:
	# see swaggerapi/swagger-codegen-cli-v3 config-help -l go
	# and swaggerapi/swagger-codegen-cli-v3 generate -help
	# bugs: https://github.com/swagger-api/swagger-codegen/issues?q=is%3Aopen+is%3Aissue+label%3A%22Client%3A+Go%22
	${DOCKER_SWAGGER} generate -l go \
		--additional-properties packageName=${PACKAGE_MAJOR} \
		--model-package types \
		--api-package models \
		--git-user-id ${GIT_ORG} \
		--git-repo-id ${GIT_REPO} \
		-o ./${PACKAGE_PREFIX}/${PACKAGE_MAJOR} \
		-i ./${SPEC_PATCHED_FILE}

# Requires downloading the specific version of the swagger-codegen-cli.jar from the artifact website
# and putting it in the root of the repository. It will be .gitignore'd but it needs to be present.
# Command below matches the Docker image Swagger Codegen CLI version above 3.0.34
# curl -o swagger-codegen-cli.jar https://repo1.maven.org/maven2/io/swagger/codegen/v3/swagger-codegen-cli/3.0.34/swagger-codegen-cli-3.0.34.jar
gen-swagger-apple-chip-cpu:
	java -jar swagger-codegen-cli.jar generate -l go \
		--additional-properties packageName=${PACKAGE_MAJOR} \
		--model-package types \
		--api-package models \
		--git-user-id ${GIT_ORG} \
		--git-repo-id ${GIT_REPO} \
		-o ./${PACKAGE_PREFIX}/${PACKAGE_MAJOR} \
		-i ./${SPEC_PATCHED_FILE}

gen: gen-swagger

mod:
	rm -f go.mod go.sum
	go mod init github.com/${GIT_ORG}/${GIT_REPO}
	go mod tidy

fmt:
	go run mvdan.cc/gofumpt@v0.3.1 -l -w $(PACKAGE_PREFIX)

patch-post:
	# patch is idempotent, always starting with the generated files
	for diff in $(shell find patches/post -name \*.patch | sort -n); do \
		patch --no-backup-if-mismatch -N -t -p1 -i $$diff; \
	done

clean-docs:
	rm -rf README.md docs

move-docs:
	mv ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/README.md .
	mv ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/docs .

docs: clean-docs move-docs

move-other:
	rm -rf api .travis.yml git_push.sh
	rm -f ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/.travis.yml
	mv ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/api .
	mv ${PACKAGE_PREFIX}/${PACKAGE_MAJOR}/git_push.sh .

test:
	go test -v ./...

stage:
	test -d .git && git add --intent-to-add README.md docs ${PACKAGE_PREFIX} go.mod go.sum

# https://github.com/OpenAPITools/openapi-generator/issues/741#issuecomment-569791780
remove-dupe-requests: ## Removes duplicate Request structs from the generated code
	@for struct in $$(grep -h 'type .\{1,\} struct' $(PACKAGE_MAJOR)/*.go | grep Request  | sort | uniq -c | grep -v '^      1' | awk '{print $$3}'); do \
	  for f in $$(/bin/ls $(PACKAGE_MAJOR)/*.go); do \
	    if grep -qF "type $${struct} struct" "$${f}"; then \
	      if eval "test -z \$${$${struct}}"; then \
	        echo "skipping first appearance of $${struct} in file $${f}"; \
	        eval "export $${struct}=1"; \
	      else \
	        echo "removing dupe $${struct} from file $${f}"; \
	        tr '\n' '\r' <"$${f}" | sed 's~// '"$${struct}"'.\{1,\}type '"$${struct}"' struct {[^}]\{1,\}}~~' | tr '\r' '\n' >"$${f}.tmp"; \
	        mv -f "$${f}.tmp" "$${f}"; \
	      fi; \
	    fi \
	  done \
	done

fix-tags:
	- jq '. | select(((.paths[][].tags| type=="array"), length) > 1).paths[][].tags |= [.[0]]' ${SPEC_FETCHED_FILE} | diff -d -U6 ${SPEC_FETCHED_FILE} - >  patches/01-tag-from-last-in-path.patch


lint:
	@$(GOLANGCI_LINT) run -v --no-config --fast=false --fix --disable-all --enable goimports $(PACKAGE_PREFIX)
