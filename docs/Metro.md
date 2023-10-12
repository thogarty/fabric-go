# Metro

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | The Canonical URL at which the resource resides. | [optional] [default to null]
**Type_** | **string** | Indicator of a Fabric Metro | [optional] [default to null]
**Code** | **string** | Code Assigned to an Equinix IBX data center in a specified metropolitan area. | [optional] [default to null]
**Region** | **string** | Board geographic area in which the data center is located | [optional] [default to null]
**Name** | **string** | Name of the region in which the data center is located. | [optional] [default to null]
**EquinixAsn** | **int64** | Autonomous system number (ASN) for a specified Fabric metro. The ASN is a unique identifier that carries the network routing protocol and exchanges that data with other internal systems via border gateway protocol. | [optional] [default to null]
**LocalVCBandwidthMax** | **int64** | This field holds Max Connection speed with in the metro | [optional] [default to null]
**GeoCoordinates** | [***GeoCoordinates**](GeoCoordinates.md) |  | [optional] [default to null]
**ConnectedMetros** | [**[]ConnectedMetro**](ConnectedMetro.md) |  | [optional] [default to null]
**GeoScopes** | [**[]GeoScopeType**](GeoScopeType.md) | List of supported geographic boundaries of a Fabric Metro. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

