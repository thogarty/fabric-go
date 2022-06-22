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

type RoutingProtocolBgpType struct {
	// Routing protocol type
	Type_   string                    `json:"type"`
	Name    string                    `json:"name,omitempty"`
	BgpIpv4 *BgpConnectionIpv4Request `json:"bgpIpv4,omitempty"`
	BgpIpv6 *BgpConnectionIpv6Request `json:"bgpIpv6,omitempty"`
	// Customer asn
	CustomerAsn int64 `json:"customerAsn,omitempty"`
	// BGP authorization key
	BgpAuthKey string              `json:"bgpAuthKey,omitempty"`
	Bfd        *RoutingProtocolBfd `json:"bfd,omitempty"`
}
