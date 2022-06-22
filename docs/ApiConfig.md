# ApiConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiAvailable** | **bool** | Setting indicating whether the API is available (true) or not (false). | [optional] [default to false]
**IntegrationId** | **string** |  | [optional] [default to null]
**EquinixManagedPort** | **bool** | Setting indicating that the port is managed by Equinix (true) or not (false). | [optional] [default to false]
**EquinixManagedVlan** | **bool** | Setting indicating that the VLAN is managed by Equinix (true) or not (false). | [optional] [default to false]
**AllowOverSubscription** | **bool** | Setting showing that oversubscription support is available (true) or not (false). The default is false. Oversubscription is the sale of more than the available network bandwidth. This practice is common and legitimate. After all, many customers use less bandwidth than they&#x27;ve purchased. And network users don&#x27;t consume bandwidth all at the same time. The leftover bandwidth can be sold to other customers. When demand surges, operational and engineering resources can be shifted to accommodate the load.  | [optional] [default to false]
**OverSubscriptionLimit** | **int32** | A cap on oversubscription. | [optional] [default to 1]
**BandwidthFromApi** | **bool** |  | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

