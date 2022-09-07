# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNexusGatewayByUuid**](NexusGatewayApi.md#CreateNexusGatewayByUuid) | **Post** /nexus/i4/gateway | Get Gateway by ID
[**DeleteNexusGatewayByUuid**](NexusGatewayApi.md#DeleteNexusGatewayByUuid) | **Delete** /nexus/i4/gateway/{gatewayId} | Delete by ID
[**GetNexusGatewayByUuid**](NexusGatewayApi.md#GetNexusGatewayByUuid) | **Get** /nexus/i4/gateway/{gatewayId} | Get Gateway by ID
[**UpdateNexusGatewayByUuid**](NexusGatewayApi.md#UpdateNexusGatewayByUuid) | **Patch** /nexus/i4/gateway/{gatewayId} | Update by ID

# **CreateNexusGatewayByUuid**
> NexusGatewayResponse CreateNexusGatewayByUuid(ctx, body, optional)
Get Gateway by ID

The API provides capability to get user's virtual gateway details  by it's gateway ID (UUID)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NexusGatewayRequest**](NexusGatewayRequest.md)|  | 
 **optional** | ***NexusGatewayApiCreateNexusGatewayByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NexusGatewayApiCreateNexusGatewayByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHAPIKEY** | **optional.**| API Key | 

### Return type

[**NexusGatewayResponse**](NexusGatewayResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNexusGatewayByUuid**
> NexusGatewayResponse DeleteNexusGatewayByUuid(ctx, gatewayId, optional)
Delete by ID

Delete Gateway by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | [**string**](.md)| Gateway UUID | 
 **optional** | ***NexusGatewayApiDeleteNexusGatewayByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NexusGatewayApiDeleteNexusGatewayByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHAPIKEY** | **optional.String**| API Key | 

### Return type

[**NexusGatewayResponse**](NexusGatewayResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNexusGatewayByUuid**
> NexusGatewayResponse GetNexusGatewayByUuid(ctx, gatewayId, optional)
Get Gateway by ID

The API provides capability to get user's virtual gateway details by it's gateway ID (UUID)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | [**string**](.md)| Gateway UUID | 
 **optional** | ***NexusGatewayApiGetNexusGatewayByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NexusGatewayApiGetNexusGatewayByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHAPIKEY** | **optional.String**| API Key | 

### Return type

[**NexusGatewayResponse**](NexusGatewayResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNexusGatewayByUuid**
> NexusGatewayResponse UpdateNexusGatewayByUuid(ctx, body, gatewayId, optional)
Update by ID

Update Gateway by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NexusGatewayUpdate**](NexusGatewayUpdate.md)|  | 
  **gatewayId** | [**string**](.md)| Gateway UUID | 
 **optional** | ***NexusGatewayApiUpdateNexusGatewayByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NexusGatewayApiUpdateNexusGatewayByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHAPIKEY** | **optional.**| API Key | 

### Return type

[**NexusGatewayResponse**](NexusGatewayResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

