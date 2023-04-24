# ServiceToken

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | [***ServiceTokenType**](ServiceTokenType.md) |  | [optional] [default to null]
**Href** | **string** | An absolute URL that is the subject of the link&#x27;s context. | [optional] [default to null]
**Uuid** | **string** | Equinix-assigned service token identifier | [default to null]
**Name** | **string** | Customer-provided service token name | [optional] [default to null]
**Description** | **string** | Customer-provided service token description | [optional] [default to null]
**ExpirationDateTime** | [**time.Time**](time.Time.md) | Expiration date and time of the service token. | [optional] [default to null]
**Connection** | [***ServiceTokenConnection**](ServiceTokenConnection.md) |  | [optional] [default to null]
**State** | [***ServiceTokenState**](ServiceTokenState.md) |  | [optional] [default to null]
**Notifications** | [**[]SimplifiedNotification**](SimplifiedNotification.md) | Service token related notifications | [optional] [default to null]
**Account** | [***SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] [default to null]
**Changelog** | [***Changelog**](Changelog.md) |  | [optional] [default to null]
**Project** | [***Project**](Project.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

