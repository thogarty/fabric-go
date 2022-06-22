# ServiceProfileRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | Service Profile URI response attribute | [optional] [default to null]
**Type_** | [***ServiceProfileTypeEnum**](ServiceProfileTypeEnum.md) |  | [default to null]
**Name** | **string** | Customer-assigned service profile name | [default to null]
**Uuid** | **string** | Equinix-assigned service profile identifier | [optional] [default to null]
**Description** | **string** | User-provided service description | [default to null]
**Notifications** | [**[]SimplifiedNotification**](SimplifiedNotification.md) | Recipients of notifications on service profile change | [optional] [default to null]
**Tags** | [***[]string**](array.md) |  | [optional] [default to null]
**Visibility** | [***ServiceProfileVisibilityEnum**](ServiceProfileVisibilityEnum.md) |  | [default to null]
**AllowedEmails** | **[]string** |  | [optional] [default to null]
**AccessPointTypeConfigs** | [**[]ServiceProfileAccessPointType**](ServiceProfileAccessPointType.md) |  | [default to null]
**CustomFields** | [**[]CustomField**](CustomField.md) |  | [optional] [default to null]
**MarketingInfo** | [***MarketingInfo**](MarketingInfo.md) |  | [optional] [default to null]
**Ports** | [**[]ServiceProfileAccessPointColo**](ServiceProfileAccessPointCOLO.md) |  | [optional] [default to null]
**VirtualDevices** | [**[]ServiceProfileAccessPointVd**](ServiceProfileAccessPointVD.md) |  | [optional] [default to null]
**Metros** | [**[]ServiceMetro**](ServiceMetro.md) | Derived response attribute. | [optional] [default to null]
**SelfProfile** | **bool** | response attribute indicates whether the profile belongs to the same organization as the api-invoker. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

