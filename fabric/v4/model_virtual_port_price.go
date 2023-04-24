/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * API version: 4.7
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// Preferences and settings for a virtual port connected to an internet service provider (ISP) or other Equinix platform entity.
type VirtualPortPrice struct {
	// Unique identifier assigned to the virtual port. Either the uuid or the remaining attributes must be supplied.
	Uuid     string                `json:"uuid,omitempty"`
	Type_    *VirtualPortType      `json:"type,omitempty"`
	Location *VirtualPortLocation  `json:"location,omitempty"`
	Lag      *LinkAggregationGroup `json:"lag,omitempty"`
	// Number of physical ports requested. The defaults is 1.
	PhysicalPortsQuantity int32 `json:"physicalPortsQuantity,omitempty"`
	// Aggregated data transfer capacity,  expressed as follows <br> -> Mbps, megabits (1 million bits) per second <br> -> Gbps, gigabits (1 billion bits) per second <br> Bandwidth must be divisible by physicalPortsQuantity.
	Bandwidth          int32                     `json:"bandwidth,omitempty"`
	Redundancy         *VirtualPortRedundancy    `json:"redundancy,omitempty"`
	ConnectivitySource *ConnectivitySource       `json:"connectivitySource,omitempty"`
	ServiceType        *VirtualPortServiceType   `json:"serviceType,omitempty"`
	Settings           *VirtualPortConfiguration `json:"settings,omitempty"`
}
