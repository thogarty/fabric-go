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

// ConnectionType : Connection type
type ConnectionType string

// List of ConnectionType
const (
	EVPL_VC_ConnectionType       ConnectionType = "EVPL_VC"
	EPL_VC_ConnectionType        ConnectionType = "EPL_VC"
	EC_VC_ConnectionType         ConnectionType = "EC_VC"
	IP_VC_ConnectionType         ConnectionType = "IP_VC"
	ACCESS_EPL_VC_ConnectionType ConnectionType = "ACCESS_EPL_VC"
	EIA_VC_ConnectionType        ConnectionType = "EIA_VC"
	EVPLAN_VC_ConnectionType     ConnectionType = "EVPLAN_VC"
	EPLAN_VC_ConnectionType      ConnectionType = "EPLAN_VC"
	IPWAN_VC_ConnectionType      ConnectionType = "IPWAN_VC"
)
