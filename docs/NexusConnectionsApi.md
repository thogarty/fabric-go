# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNexusConnectionByUuid**](NexusConnectionsApi.md#DeleteNexusConnectionByUuid) | **Delete** /nexus/i4/connections/{connectionId} | Delete by ID
[**GetNexusConnectionByUuid**](NexusConnectionsApi.md#GetNexusConnectionByUuid) | **Get** /nexus/i4/connections/{connectionId} | Get Connection by ID
[**NexusCreateConnection**](NexusConnectionsApi.md#NexusCreateConnection) | **Post** /nexus/i4/connections | Create Connection
[**UpdateNexusConnectionByUuid**](NexusConnectionsApi.md#UpdateNexusConnectionByUuid) | **Patch** /nexus/i4/connections/{connectionId} | Update by ID

# **DeleteNexusConnectionByUuid**
> NexusConnection DeleteNexusConnectionByUuid(ctx, connectionId, optional)
Delete by ID

Delete Connection by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | [**string**](.md)| Connection Id | 
 **optional** | ***NexusConnectionsApiDeleteNexusConnectionByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NexusConnectionsApiDeleteNexusConnectionByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHAPIKEY** | **optional.String**| API Key | 

### Return type

[**NexusConnection**](NexusConnection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNexusConnectionByUuid**
> NexusConnection GetNexusConnectionByUuid(ctx, connectionId, optional)
Get Connection by ID

The API provides capability to get user's virtual connection details (Access Points, Link Protocols, etc) by it's connection ID (UUID)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | [**string**](.md)| Connection Id | 
 **optional** | ***NexusConnectionsApiGetNexusConnectionByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NexusConnectionsApiGetNexusConnectionByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHAPIKEY** | **optional.String**| API Key | 

### Return type

[**NexusConnection**](NexusConnection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NexusCreateConnection**
> NexusConnection NexusCreateConnection(ctx, body, optional)
Create Connection

This API provides capability to create user's virtual connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NexusConnectionPostRequest**](NexusConnectionPostRequest.md)|  | 
 **optional** | ***NexusConnectionsApiNexusCreateConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NexusConnectionsApiNexusCreateConnectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHAPIKEY** | **optional.**| API Key | 

### Return type

[**NexusConnection**](NexusConnection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNexusConnectionByUuid**
> NexusConnection UpdateNexusConnectionByUuid(ctx, body, connectionId, optional)
Update by ID

Update Connection by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectionsConnectionIdBody**](ConnectionsConnectionIdBody.md)|  | 
  **connectionId** | [**string**](.md)| Connection Id | 
 **optional** | ***NexusConnectionsApiUpdateNexusConnectionByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NexusConnectionsApiUpdateNexusConnectionByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHAPIKEY** | **optional.**| API Key | 

### Return type

[**NexusConnection**](NexusConnection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

