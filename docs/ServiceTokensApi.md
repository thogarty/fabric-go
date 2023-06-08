# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceToken**](ServiceTokensApi.md#CreateServiceToken) | **Post** /fabric/v4/serviceTokens | Create Service Token
[**CreateServiceTokenAction**](ServiceTokensApi.md#CreateServiceTokenAction) | **Post** /fabric/v4/serviceTokens/{serviceTokenId}/actions | ServiceToken Actions
[**DeleteServiceTokenByUuid**](ServiceTokensApi.md#DeleteServiceTokenByUuid) | **Delete** /fabric/v4/serviceTokens/{serviceTokenId} | Delete Token by uuid
[**GetServiceTokenByUuid**](ServiceTokensApi.md#GetServiceTokenByUuid) | **Get** /fabric/v4/serviceTokens/{serviceTokenId} | Get Token by uuid
[**GetServiceTokens**](ServiceTokensApi.md#GetServiceTokens) | **Get** /fabric/v4/serviceTokens | Get All Tokens
[**SearchServiceTokens**](ServiceTokensApi.md#SearchServiceTokens) | **Post** /fabric/v4/serviceTokens/search | Search servicetokens
[**UpdateServiceTokenByUuid**](ServiceTokensApi.md#UpdateServiceTokenByUuid) | **Patch** /fabric/v4/serviceTokens/{serviceTokenId} | Update Token By ID

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

# **CreateServiceTokenAction**
> ServiceToken CreateServiceTokenAction(ctx, body, serviceTokenId)
ServiceToken Actions

This API provides capability to accept/reject user's servicetokens

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceTokenActionRequest**](ServiceTokenActionRequest.md)|  | 
  **serviceTokenId** | [**string**](.md)| Service Token UUID | 

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

Get All ServiceTokens creates a list of all Equinix Fabric service tokens associated with the subscriber's account.

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

# **SearchServiceTokens**
> ServiceTokens SearchServiceTokens(ctx, body)
Search servicetokens

The API provides capability to get list of user's servicetokens using search criteria, including optional filtering, pagination and sorting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceTokenSearchRequest**](ServiceTokenSearchRequest.md)|  | 

### Return type

[**ServiceTokens**](ServiceTokens.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceTokenByUuid**
> ServiceToken UpdateServiceTokenByUuid(ctx, body, serviceTokenId)
Update Token By ID

This API provides capability to update user's Service Token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]ServiceTokenChangeOperation**](ServiceTokenChangeOperation.md)|  | 
  **serviceTokenId** | [**string**](.md)| Service Token UUID | 

### Return type

[**ServiceToken**](ServiceToken.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json-patch+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

