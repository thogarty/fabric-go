# ServiceTokenConnection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | Type of Connection | [default to null]
**Href** | **string** | An absolute URL that is the subject of the link&#x27;s context. | [optional] [default to null]
**Uuid** | **string** | Equinix-assigned connection identifier | [optional] [default to null]
**AllowRemoteConnection** | **bool** | Authorization to connect remotely | [optional] [default to false]
**BandwidthLimit** | **int32** | Connection bandwidth limit in Mbps | [optional] [default to null]
**SupportedBandwidths** | **[]int32** | List of permitted bandwidths. | [optional] [default to null]
**ASide** | [***ServiceTokenSide**](ServiceTokenSide.md) |  | [optional] [default to null]
**ZSide** | [***ServiceTokenSide**](ServiceTokenSide.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

