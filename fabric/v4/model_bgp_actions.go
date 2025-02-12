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

// BgpActions : BGP action type
type BgpActions string

// List of BGPActions
const (
	CLEAR_BGPIPV4_BgpActions         BgpActions = "CLEAR_BGPIPV4"
	CLEAR_BGPIPV6_BgpActions         BgpActions = "CLEAR_BGPIPV6"
	CLEAR_BGPIPV4_INBOUND_BgpActions BgpActions = "CLEAR_BGPIPV4_INBOUND"
	CLEAR_BGPIPV6_INBOUND_BgpActions BgpActions = "CLEAR_BGPIPV6_INBOUND"
	RESET_BGPIPV4_BgpActions         BgpActions = "RESET_BGPIPV4"
	RESET_BGPIPV6_BgpActions         BgpActions = "RESET_BGPIPV6"
)
