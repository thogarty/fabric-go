# Connection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | [***ConnectionType**](ConnectionType.md) |  | [default to null]
**Href** | **string** | Connection URI | [optional] [default to null]
**Uuid** | **string** | Equinix-assigned connection identifier | [optional] [default to null]
**Name** | **string** | Customer-provided connection name | [default to null]
**Description** | **string** | Customer-provided connection description | [optional] [default to null]
**State** | [***ConnectionState**](ConnectionState.md) |  | [optional] [default to null]
**Change** | [***Change**](Change.md) |  | [optional] [default to null]
**Operation** | [***ConnectionOperation**](ConnectionOperation.md) |  | [optional] [default to null]
**Order** | [***Order**](Order.md) |  | [optional] [default to null]
**Notifications** | [**[]SimplifiedNotification**](SimplifiedNotification.md) | Preferences for notifications on connection configuration or status changes | [optional] [default to null]
**Account** | [***SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] [default to null]
**ChangeLog** | [***Changelog**](Changelog.md) |  | [optional] [default to null]
**Bandwidth** | **int32** | Connection bandwidth in Mbps | [default to null]
**Redundancy** | [***ConnectionRedundancy**](ConnectionRedundancy.md) |  | [optional] [default to null]
**IsRemote** | **bool** | Connection property derived from access point locations | [optional] [default to null]
**Direction** | [***ConnectionDirection**](ConnectionDirection.md) |  | [optional] [default to null]
**ASide** | [***ConnectionSide**](ConnectionSide.md) |  | [default to null]
**ZSide** | [***ConnectionSide**](ConnectionSide.md) |  | [default to null]
**AdditionalInfo** | [**[]ConnectionSideAdditionalInfo**](ConnectionSideAdditionalInfo.md) | Connection additional information | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

