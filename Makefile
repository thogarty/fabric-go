.PHONY: all gen patch fetch

CURRENT_UID := $(shell id -u)
CURRENT_GID := $(shell id -g)

# https://github.com/OpenAPITools/openapi-generator-cli
# SPEC_URL:="https://developer.equinix.com/sites/default/files/fabric-v4-catalog-fabric_v4_3.yaml"
SPEC_URL:="https://api.swaggerhub.com/apis/equinix-api/fabric/4.7/swagger.yaml"

SPEC_FETCHED_FILE:=spec.fetched.yaml
SPEC_PATCHED_FILE:=spec.patched.yaml
IMAGE=swaggerapi/swagger-codegen-cli-v3:3.0.34
# IMAGE=openapitools/openapi-generator-cli # to use, change gen: to gen-openapitools
VALIDATE_IMAGE=openapitools/openapi-generator-cli

GIT_ORG=equinix-labs
GIT_REPO=fabric-go
PACKAGE_PREFIX=fabric
PACKAGE_MAJOR=v4
CRI=docker # nerdctl

SWAGGER=${CRI} run --rm -u ${CURRENT_UID}:${CURRENT_GID} -v $(CURDIR):/local ${IMAGE}
VALIDATE=${CRI} run --rm -u ${CURRENT_UID}:${CURRENT_GID} -v $(CURDIR):/local ${VALIDATE_IMAGE}

GOLANGCI_LINT=golangci-lint

#all: pull fetch patch clean gen mod docs move-other patch-post fmt test stage
all: pull fetch patch clean gen mod fmt patch-post docs move-other test stage

pull:
	${CRI} pull ${IMAGE}

fetch:
	curl \
		-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36' \
		-o ${SPEC_FETCHED_FILE} ${SPEC_URL}

fix-tags:
	- jq '. | select(((.paths[][].tags| type=="array"), length) > 1).paths[][].tags |= [.[0]]' ${SPEC_FETCHED_FILE} | diff -d -U6 ${SPEC_FETCHED_FILE} - >  patches/01-tag-from-last-in-path.patch

patch:
	# patch is idempotent, always starting with the fetched
	# fetched file to create the patched file.
	cp ${SPEC_FETCHED_FILE} ${SPEC_PATCHED_FILE}
	ARGS="-o ${SPEC_PATCHED_FILE} ${SPEC_FETCHED_FILE}"; \
	for diff in $(shell find patches/${SPEC_FETCHED_FILE} -name \*.patch | sort -n); do \
		patch --no-backup-if-mismatch -N -t $$ARGS $$diff; \
		ARGS=${SPEC_PATCHED_FILE}; \
	done

patch-post:
	# patch is idempotent, always starting with the generated files
	for diff in $(shell find patches/post -name \*.patch | sort -n); do \
		patch --no-backup-if-mismatch -N -t -p1 -i $$diff; \
	done

clean: clean-docs
	rm -rf ${PACKAGE_PREFIX}

gen-openapitools:
	${SWAGGER} generate -g go \
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
	${SWAGGER} generate -l go \
		--additional-properties packageName=${PACKAGE_MAJOR} \
		--model-package types \
		--api-package models \
		--git-user-id ${GIT_ORG} \
		--git-repo-id ${GIT_REPO} \
		-o /local/${PACKAGE_PREFIX}/${PACKAGE_MAJOR} \
		-i /local/${SPEC_PATCHED_FILE}

gen: gen-swagger

validate:
	${VALIDATE} validate \
		--recommend \
		-i /local/${SPEC_PATCHED_FILE}

spectral:
	spectral --help > /dev/null 2>&1 || echo "Spectral is not installed"
	spectral lint $(SPEC_PATCHED_FILE)

mod:
	rm -f go.mod go.sum
	go mod init github.com/${GIT_ORG}/${GIT_REPO}
	go mod tidy

test:
	go test -v ./...

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
lint:
	@$(GOLANGCI_LINT) run -v --no-config --fast=false --fix --disable-all --enable goimports $(PACKAGE_PREFIX)

fmt:
	go run mvdan.cc/gofumpt@v0.3.1 -l -w $(PACKAGE_PREFIX)

stage:
	test -d .git && git add --intent-to-add README.md docs ${PACKAGE_PREFIX} go.mod go.sum
