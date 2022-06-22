/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * API version: 4.3
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// GatewayActionType : Gateway action type
type GatewayActionType string

// List of GatewayActionType
const (
	BGP_SESSION_STATUS_UPDATE_GatewayActionType GatewayActionType = "BGP_SESSION_STATUS_UPDATE"
	ROUTE_TABLE_ENTRY_UPDATE_GatewayActionType  GatewayActionType = "ROUTE_TABLE_ENTRY_UPDATE"
)
