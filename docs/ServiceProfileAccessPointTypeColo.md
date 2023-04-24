# ServiceProfileAccessPointTypeColo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | [optional] [default to null]
**Type_** | [***ServiceProfileAccessPointTypeEnum**](ServiceProfileAccessPointTypeEnum.md) |  | [default to null]
**SupportedBandwidths** | [***[]int32**](array.md) |  | [optional] [default to null]
**AllowRemoteConnections** | **bool** | Setting to allow or prohibit remote connections to the service profile. | [optional] [default to false]
**AllowCustomBandwidth** | **bool** | Setting to enable or disable the ability of the buyer to customize the bandwidth. | [optional] [default to false]
**BandwidthAlertThreshold** | **float64** | percentage of port bandwidth at which an allocation alert is generated - missing on wiki. | [optional] [default to null]
**AllowBandwidthAutoApproval** | **bool** | Setting to enable or disable the ability of the buyer to change connection bandwidth without approval of the seller. | [optional] [default to false]
**AllowBandwidthUpgrade** | **bool** | Availability of a bandwidth upgrade. The default is false. | [optional] [default to null]
**LinkProtocolConfig** | [***ServiceProfileLinkProtocolConfig**](ServiceProfileLinkProtocolConfig.md) |  | [optional] [default to null]
**EnableAutoGenerateServiceKey** | **bool** | for verizon only. | [optional] [default to null]
**ConnectionRedundancyRequired** | **bool** | Mandate redundant connections | [optional] [default to false]
**ApiConfig** | [***ApiConfig**](ApiConfig.md) |  | [optional] [default to null]
**ConnectionLabel** | **string** | custom name for \&quot;Connection\&quot; | [optional] [default to null]
**AuthenticationKey** | [***AuthenticationKey**](AuthenticationKey.md) |  | [optional] [default to null]
**Metadata** | [***ServiceProfileMetadata**](ServiceProfileMetadata.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

