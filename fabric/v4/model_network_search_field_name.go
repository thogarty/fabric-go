/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * API version: 4.10
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// NetworkSearchFieldName : Possible field names to use on filters
type NetworkSearchFieldName string

// List of NetworkSearchFieldName
const (
	NAME_NetworkSearchFieldName                    NetworkSearchFieldName = "/name"
	UUID_NetworkSearchFieldName                    NetworkSearchFieldName = "/uuid"
	SCOPE_NetworkSearchFieldName                   NetworkSearchFieldName = "/scope"
	TYPE__NetworkSearchFieldName                   NetworkSearchFieldName = "/type"
	OPERATIONEQUINIX_STATUS_NetworkSearchFieldName NetworkSearchFieldName = "/operation/equinixStatus"
	LOCATIONREGION_NetworkSearchFieldName          NetworkSearchFieldName = "/location/region"
	PROJECTPROJECT_ID_NetworkSearchFieldName       NetworkSearchFieldName = "/project/projectId"
	ACCOUNTGLOBAL_CUST_ID_NetworkSearchFieldName   NetworkSearchFieldName = "/account/globalCustId"
	ACCOUNTORG_ID_NetworkSearchFieldName           NetworkSearchFieldName = "/account/orgId"
	DELETED_DATE_NetworkSearchFieldName            NetworkSearchFieldName = "/deletedDate"
	STAR_NetworkSearchFieldName                    NetworkSearchFieldName = "/_*"
)
