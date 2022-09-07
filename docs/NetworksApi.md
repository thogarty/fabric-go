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
> Network CreateNetwork(ctx, body, optional)
Create Network

This API provides capability to create user's Fabric Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NetworkPostRequest**](NetworkPostRequest.md)|  | 
 **optional** | ***NetworksApiCreateNetworkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksApiCreateNetworkOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.**| User name | 

### Return type

[**Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNetworkByUuid**
> Network DeleteNetworkByUuid(ctx, networkId, optional)
Delete Network By ID

This API provides capability to delete user's Fabric Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | [**string**](.md)| Network UUID | 
 **optional** | ***NetworksApiDeleteNetworkByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksApiDeleteNetworkByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 

### Return type

[**Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectionsByNetworkUuid**
> NetworkConnections GetConnectionsByNetworkUuid(ctx, networkId, optional)
Get Connections

The API provides capability to get list of user's Fabric Network connections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | [**string**](.md)| Network UUID | 
 **optional** | ***NetworksApiGetConnectionsByNetworkUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksApiGetConnectionsByNetworkUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 

### Return type

[**NetworkConnections**](NetworkConnections.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkByUuid**
> Network GetNetworkByUuid(ctx, networkId, optional)
Get Network By ID

This API provides capability to retrieve user's Fabric Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | [**string**](.md)| Network UUID | 
 **optional** | ***NetworksApiGetNetworkByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksApiGetNetworkByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 

### Return type

[**Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkChangeByUuid**
> NetworkChange GetNetworkChangeByUuid(ctx, networkId, changeId, optional)
Get Change By ID

This API provides capability to retrieve user's Fabric Network Change

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | [**string**](.md)| Network UUID | 
  **changeId** | [**string**](.md)| Network Change UUID | 
 **optional** | ***NetworksApiGetNetworkChangeByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksApiGetNetworkChangeByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 

### Return type

[**NetworkChange**](NetworkChange.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkChanges**
> NetworkChangeResponse GetNetworkChanges(ctx, networkId, optional)
Get Network Changes

The API provides capability to get list of user's Fabric Network changes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | [**string**](.md)| Network UUID | 
 **optional** | ***NetworksApiGetNetworkChangesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksApiGetNetworkChangesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 

### Return type

[**NetworkChangeResponse**](NetworkChangeResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchNetworks**
> NetworkSearchResponse SearchNetworks(ctx, body, optional)
Search Network

The API provides capability to get list of user's Fabric Network using search criteria, including optional filtering, pagination and sorting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NetworkSearchRequest**](NetworkSearchRequest.md)|  | 
 **optional** | ***NetworksApiSearchNetworksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksApiSearchNetworksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.**| User name | 

### Return type

[**NetworkSearchResponse**](NetworkSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNetworkByUuid**
> Network UpdateNetworkByUuid(ctx, body, networkId, optional)
Update Network By ID

This API provides capability to update user's Fabric Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]NetworkChangeOperation**](NetworkChangeOperation.md)|  | 
  **networkId** | [**string**](.md)| Network UUID | 
 **optional** | ***NetworksApiUpdateNetworkByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksApiUpdateNetworkByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.**| User name | 

### Return type

[**Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json-patch+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

