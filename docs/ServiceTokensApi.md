# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceToken**](ServiceTokensApi.md#CreateServiceToken) | **Post** /fabric/v4/serviceTokens | Create Service Token
[**DeleteServiceTokenByUuid**](ServiceTokensApi.md#DeleteServiceTokenByUuid) | **Delete** /fabric/v4/serviceTokens/{serviceTokenId} | Delete Token by uuid
[**GetServiceTokenByUuid**](ServiceTokensApi.md#GetServiceTokenByUuid) | **Get** /fabric/v4/serviceTokens/{serviceTokenId} | Get Token by uuid
[**GetServiceTokens**](ServiceTokensApi.md#GetServiceTokens) | **Get** /fabric/v4/serviceTokens | Get All Tokens

# **CreateServiceToken**
> ServiceToken CreateServiceToken(ctx, body)
Create Service Token

Create Service Tokens generates Equinix Fabric? service tokens. These tokens authorize users to access protected resources and services.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceToken**](ServiceToken.md)|  | 

### Return type

[**ServiceToken**](ServiceToken.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceTokenByUuid**
> DeleteServiceTokenByUuid(ctx, serviceTokenId)
Delete Token by uuid

Delete Service Tokens removes an Equinix Fabric service token corresponding to the specified uuid which are in INACTIVE state.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceTokenId** | [**string**](.md)| Service Token UUID | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceTokenByUuid**
> ServiceToken GetServiceTokenByUuid(ctx, serviceTokenId)
Get Token by uuid

Get Specified Service Tokens uses the uuid of an Equinix Fabric service token to return details about the token's type, state, location, bandwidth, and other key properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceTokenId** | [**string**](.md)| Service Token UUID | 

### Return type

[**ServiceToken**](ServiceToken.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceTokens**
> ServiceTokens GetServiceTokens(ctx, optional)
Get All Tokens

Get All Service Tokens creates a list of all Equinix Fabric service tokens associated with the subscriber's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServiceTokensApiGetServiceTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceTokensApiGetServiceTokensOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Float64**| offset | 
 **limit** | **optional.Float64**| number of records to fetch | 

### Return type

[**ServiceTokens**](ServiceTokens.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

