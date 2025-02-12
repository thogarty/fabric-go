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

// Fabric Cloud Router object
type CloudRouter struct {
	Type_ string `json:"type,omitempty"`
	// Customer-provided Cloud Router name
	Name     string                        `json:"name,omitempty"`
	Location *SimplifiedLocationWithoutIbx `json:"location,omitempty"`
	Package_ *CloudRouterPackageType       `json:"package,omitempty"`
	Order    *Order                        `json:"order,omitempty"`
	Project  *Project                      `json:"project,omitempty"`
	Account  *SimplifiedAccount            `json:"account,omitempty"`
	// Preferences for notifications on connection configuration or status changes
	Notifications []SimplifiedNotification `json:"notifications,omitempty"`
	// Cloud Routers URI
	Href string `json:"href,omitempty"`
	// Equinix-assigned access point identifier
	Uuid  string                       `json:"uuid,omitempty"`
	State *CloudRouterAccessPointState `json:"state,omitempty"`
	// Equinix ASN
	EquinixAsn int64 `json:"equinixAsn,omitempty"`
	// Access point used and maximum number of IPv4 BGP routes
	BgpIpv4RoutesCount int32 `json:"bgpIpv4RoutesCount,omitempty"`
	// Access point used and maximum number of IPv6 BGP routes
	BgpIpv6RoutesCount int32 `json:"bgpIpv6RoutesCount,omitempty"`
	// Number of connections associated with this Access point
	ConnectionsCount int32 `json:"connectionsCount,omitempty"`
	// Number of distinct ipv4 routes
	DistinctIpv4PrefixesCount int32 `json:"distinctIpv4PrefixesCount,omitempty"`
	// Number of distinct ipv6 routes
	DistinctIpv6PrefixesCount int32              `json:"distinctIpv6PrefixesCount,omitempty"`
	ChangeLog                 *Changelog         `json:"changeLog,omitempty"`
	Change                    *CloudRouterChange `json:"change,omitempty"`
}
