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

type PtpAdvanceConfiguration struct {
	// Time Scale value, ARB denotes Arbitrary and PTP denotes Precision Time Protocol.
	TimeScale string `json:"timeScale,omitempty"`
	Domain    int32  `json:"domain,omitempty"`
	Priority1 int32  `json:"priority1,omitempty"`
	Priority2 int32  `json:"priority2,omitempty"`
	// The mean time interval between Announce messages. A shorter interval makes ptp4l react faster to the changes in the master-slave hierarchy. The interval should be the same in the whole domain. It's specified as a power of two in seconds. The default is 1 (2 seconds).
	LogAnnounceInterval int32 `json:"logAnnounceInterval,omitempty"`
	// The mean time interval between Sync messages. A shorter interval may improve accuracy of the local clock. It's specified as a power of two in seconds. The default is 0 (1 second).
	LogSyncInterval     int32  `json:"logSyncInterval,omitempty"`
	LogDelayReqInterval int32  `json:"logDelayReqInterval,omitempty"`
	TransportMode       string `json:"transportMode,omitempty"`
	// Unicast Grant Time in seconds. For Multicast and Hybrid transport modes, grant time defaults to 300 seconds. For Unicast mode, grant time can be between 30 to 7200.
	GrantTime int32 `json:"grantTime,omitempty"`
}
