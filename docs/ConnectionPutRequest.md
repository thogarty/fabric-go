# ConnectionPutRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | [***ConnectionType**](ConnectionType.md) |  | [optional] [default to null]
**Name** | **string** | Customer-provided connection name | [optional] [default to null]
**Order** | [***Order**](Order.md) |  | [optional] [default to null]
**Notifications** | [**[]SimplifiedNotification**](SimplifiedNotification.md) | Preferences for notifications on connection configuration or status changes | [optional] [default to null]
**Bandwidth** | **int32** | Connection bandwidth in Mbps | [optional] [default to null]
**Redundancy** | [***ConnectionRedundancy**](ConnectionRedundancy.md) |  | [optional] [default to null]
**ASide** | [***ConnectionSide**](ConnectionSide.md) |  | [optional] [default to null]
**ZSide** | [***ConnectionSide**](ConnectionSide.md) |  | [optional] [default to null]
**Project** | [***Project**](Project.md) |  | [optional] [default to null]
**AdditionalInfo** | [**[]ConnectionSideAdditionalInfo**](ConnectionSideAdditionalInfo.md) | Connection additional information | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

