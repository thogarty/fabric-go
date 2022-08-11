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

// Access point object
type AccessPoint struct {
	Type_ *AccessPointType `json:"type"`
	Account *SimplifiedAccount `json:"account,omitempty"`
	Location *SimplifiedLocation `json:"location"`
	Port *Port `json:"port,omitempty"`
	Profile *SimplifiedServiceProfile `json:"profile,omitempty"`
	Gateway *VirtualGateway `json:"gateway,omitempty"`
	LinkProtocol *SimplifiedLinkProtocol `json:"linkProtocol,omitempty"`
	VirtualDevice *VirtualDevice `json:"virtualDevice,omitempty"`
	Interface_ *ModelInterface `json:"interface,omitempty"`
	// Access point seller region
	SellerRegion string `json:"sellerRegion,omitempty"`
	PeeringType *PeeringType `json:"peeringType,omitempty"`
	// Access point authentication key
	AuthenticationKey string `json:"authenticationKey,omitempty"`
	// Access point routing protocols configuration
	RoutingProtocols []RoutingProtocol `json:"routingProtocols,omitempty"`
	AdditionalInfo *AdditionalInfo `json:"additionalInfo,omitempty"`
	// Provider assigned Connection Id
	ProviderConnectionId string `json:"providerConnectionId,omitempty"`
}
