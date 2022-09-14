# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetwork**](NetworksApi.md#CreateNetwork) | **Post** /fabric/v4/networks | Create Network
[**DeleteNetworkByUuid**](NetworksApi.md#DeleteNetworkByUuid) | **Delete** /fabric/v4/networks/{networkId} | Delete Network By ID
[**GetConnectionsByNetworkUuid**](NetworksApi.md#GetConnectionsByNetworkUuid) | **Get** /fabric/v4/networks/{networkId}/connections | Get Connections
[**GetNetworkByUuid**](NetworksApi.md#GetNetworkByUuid) | **Get** /fabric/v4/networks/{networkId} | Get Network By ID
[**GetNetworkChangeByUuid**](NetworksApi.md#GetNetworkChangeByUuid) | **Get** /fabric/v4/networks/{networkId}/changes/{changeId} | Get Change By ID
[**GetNetworkChanges**](NetworksApi.md#GetNetworkChanges) | **Get** /fabric/v4/networks/{networkId}/changes | Get Network Changes
[**SearchNetworks**](NetworksApi.md#SearchNetworks) | **Post** /fabric/v4/networks/search | Search Network
[**UpdateNetworkByUuid**](NetworksApi.md#UpdateNetworkByUuid) | **Patch** /fabric/v4/networks/{networkId} | Update Network By ID

# **CreateNetwork**
> Network CreateNetwork(ctx, body)
Create Network

This API provides capability to create user's Fabric Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NetworkPostRequest**](NetworkPostRequest.md)|  | 

### Return type

[**Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNetworkByUuid**
> Network DeleteNetworkByUuid(ctx, networkId)
Delete Network By ID

This API provides capability to delete user's Fabric Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | [**string**](.md)| Network UUID | 

### Return type

[**Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectionsByNetworkUuid**
> NetworkConnections GetConnectionsByNetworkUuid(ctx, networkId)
Get Connections

The API provides capability to get list of user's Fabric Network connections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | [**string**](.md)| Network UUID | 

### Return type

[**NetworkConnections**](NetworkConnections.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkByUuid**
> Network GetNetworkByUuid(ctx, networkId)
Get Network By ID

This API provides capability to retrieve user's Fabric Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | [**string**](.md)| Network UUID | 

### Return type

[**Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkChangeByUuid**
> NetworkChange GetNetworkChangeByUuid(ctx, networkId, changeId)
Get Change By ID

This API provides capability to retrieve user's Fabric Network Change

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | [**string**](.md)| Network UUID | 
  **changeId** | [**string**](.md)| Network Change UUID | 

### Return type

[**NetworkChange**](NetworkChange.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkChanges**
> NetworkChangeResponse GetNetworkChanges(ctx, networkId)
Get Network Changes

The API provides capability to get list of user's Fabric Network changes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | [**string**](.md)| Network UUID | 

### Return type

[**NetworkChangeResponse**](NetworkChangeResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchNetworks**
> NetworkSearchResponse SearchNetworks(ctx, body)
Search Network

The API provides capability to get list of user's Fabric Network using search criteria, including optional filtering, pagination and sorting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NetworkSearchRequest**](NetworkSearchRequest.md)|  | 

### Return type

[**NetworkSearchResponse**](NetworkSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNetworkByUuid**
> Network UpdateNetworkByUuid(ctx, body, networkId)
Update Network By ID

This API provides capability to update user's Fabric Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]NetworkChangeOperation**](NetworkChangeOperation.md)|  | 
  **networkId** | [**string**](.md)| Network UUID | 

### Return type

[**Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json-patch+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

