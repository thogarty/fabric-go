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

import (
	"time"
)

// Bandwidth utilization statistics for a specified interval.
type Metrics struct {
	// Interval end timestamp
	IntervalEndTimestamp time.Time `json:"intervalEndTimestamp,omitempty"`
	// Max bandwidth within statistics object time interval, represented in units specified by response \"units\" field
	Max float32 `json:"max,omitempty"`
	// Mean bandwidth within statistics object time interval, represented in units specified by response \"units\" field
	Mean float32 `json:"mean,omitempty"`
}
