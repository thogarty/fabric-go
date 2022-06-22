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

type Order struct {
	// Purchase order number
	PurchaseOrderNumber string `json:"purchaseOrderNumber,omitempty"`
	// Billing tier for connection bandwidth
	BillingTier string `json:"billingTier,omitempty"`
	// Order Identification
	OrderId string `json:"orderId,omitempty"`
	// Order Reference Number
	OrderNumber string `json:"orderNumber,omitempty"`
}
