# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNexusNetwork**](NexusNetworksApi.md#CreateNexusNetwork) | **Post** /nexus/i4/networks | Create Network
[**DeleteNexusNetworkByUuid**](NexusNetworksApi.md#DeleteNexusNetworkByUuid) | **Delete** /nexus/i4/networks/{networkId} | Delete Network By ID
[**GetNexusNetworkByUuid**](NexusNetworksApi.md#GetNexusNetworkByUuid) | **Get** /nexus/i4/networks/{networkId} | Get Network By ID

# **CreateNexusNetwork**
> NexusNetwork CreateNexusNetwork(ctx, body, optional)
Create Network

This API provides capability to create user's Fabric Platform Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NexusNetworkPostRequest**](NexusNetworkPostRequest.md)|  | 
 **optional** | ***NexusNetworksApiCreateNexusNetworkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NexusNetworksApiCreateNexusNetworkOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHAPIKEY** | **optional.**| Api Key | 

### Return type

[**NexusNetwork**](NexusNetwork.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNexusNetworkByUuid**
> NexusNetwork DeleteNexusNetworkByUuid(ctx, networkId, optional)
Delete Network By ID

This API provides capability to delete user's Fabric Platform Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | [**string**](.md)| Network UUID | 
 **optional** | ***NexusNetworksApiDeleteNexusNetworkByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NexusNetworksApiDeleteNexusNetworkByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHAPIKEY** | **optional.String**| Api Key | 

### Return type

[**NexusNetwork**](NexusNetwork.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNexusNetworkByUuid**
> NexusNetwork GetNexusNetworkByUuid(ctx, networkId, optional)
Get Network By ID

This API provides capability to retrieve user's Fabric Platform Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | [**string**](.md)| Network UUID | 
 **optional** | ***NexusNetworksApiGetNexusNetworkByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NexusNetworksApiGetNexusNetworkByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHAPIKEY** | **optional.String**| Api Key | 

### Return type

[**NexusNetwork**](NexusNetwork.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

