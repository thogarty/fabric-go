# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkByUuidInternal**](InternalNetworksApi.md#GetNetworkByUuidInternal) | **Get** /fabric/i4/networks/{networkId} | Get Network By ID
[**SearchNetworksInternal**](InternalNetworksApi.md#SearchNetworksInternal) | **Post** /fabric/i4/networks/search | Search Network

# **GetNetworkByUuidInternal**
> Network GetNetworkByUuidInternal(ctx, networkId, optional)
Get Network By ID

This API provides capability to retrieve user's Fabric Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | [**string**](.md)| Network UUID | 
 **optional** | ***InternalNetworksApiGetNetworkByUuidInternalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InternalNetworksApiGetNetworkByUuidInternalOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 

### Return type

[**Network**](Network.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchNetworksInternal**
> NetworkSearchResponse SearchNetworksInternal(ctx, body, optional)
Search Network

The API provides capability to get list of user's Fabric Network using search criteria, including optional filtering, pagination and sorting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NetworkSearchRequest**](NetworkSearchRequest.md)|  | 
 **optional** | ***InternalNetworksApiSearchNetworksInternalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InternalNetworksApiSearchNetworksInternalOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.**| Correlation identifier | 

### Return type

[**NetworkSearchResponse**](NetworkSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

