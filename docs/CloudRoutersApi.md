# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGateway**](CloudRoutersApi.md#CreateGateway) | **Post** /fabric/v4/routers | Create Routers
[**CreateGatewayAction**](CloudRoutersApi.md#CreateGatewayAction) | **Post** /fabric/v4/routers/{routerId}/actions | Route table actions
[**DeleteGatewayByUuid**](CloudRoutersApi.md#DeleteGatewayByUuid) | **Delete** /fabric/v4/routers/{routerId} | Delete Routers
[**GetGatewayActions**](CloudRoutersApi.md#GetGatewayActions) | **Get** /fabric/v4/routers/{routerId}/actions | Get actions
[**GetGatewayByUuid**](CloudRoutersApi.md#GetGatewayByUuid) | **Get** /fabric/v4/routers/{routerId} | Get Routers
[**GetGatewayPackageByCode**](CloudRoutersApi.md#GetGatewayPackageByCode) | **Get** /fabric/v4/routerPackages/{routerPackageCode} | Get Package Details
[**GetGatewayPackages**](CloudRoutersApi.md#GetGatewayPackages) | **Get** /fabric/v4/routerPackages | List Packages
[**SearchGateways**](CloudRoutersApi.md#SearchGateways) | **Post** /fabric/v4/routers/search | Search Routers
[**UpdateGatewayByUuid**](CloudRoutersApi.md#UpdateGatewayByUuid) | **Patch** /fabric/v4/routers/{routerId} | Update Routers

# **CreateGateway**
> CloudRouter CreateGateway(ctx, body)
Create Routers

This API provides capability to create user's Cloud Routers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CloudRouterPostRequest**](CloudRouterPostRequest.md)|  | 

### Return type

[**CloudRouter**](CloudRouter.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGatewayAction**
> GatewayActionResponse CreateGatewayAction(ctx, body, routerId)
Route table actions

This API provides capability to refresh route table and bgp session summary information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GatewayActionRequest**](GatewayActionRequest.md)|  | 
  **routerId** | [**string**](.md)| Router UUID | 

### Return type

[**GatewayActionResponse**](GatewayActionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGatewayByUuid**
> DeleteGatewayByUuid(ctx, routerId)
Delete Routers

This API provides capability to delete user's Cloud Routers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routerId** | [**string**](.md)| Cloud Router UUID | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGatewayActions**
> GatewayActionResponse GetGatewayActions(ctx, routerId, optional)
Get actions

This API provides capability to fetch action status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routerId** | [**string**](.md)| Router UUID | 
 **optional** | ***CloudRoutersApiGetGatewayActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudRoutersApiGetGatewayActionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | [**optional.Interface of GatewayActionState**](.md)| Action state | 

### Return type

[**GatewayActionResponse**](GatewayActionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGatewayByUuid**
> CloudRouter GetGatewayByUuid(ctx, routerId)
Get Routers

This API provides capability to retrieve user's Cloud Routers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routerId** | [**string**](.md)| Cloud Router UUID | 

### Return type

[**CloudRouter**](CloudRouter.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGatewayPackageByCode**
> CloudRouterPackage GetGatewayPackageByCode(ctx, routerPackageCode)
Get Package Details

This API provides capability to retrieve user's Cloud Routers Package Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routerPackageCode** | [**RouterPackageCode**](.md)| Equinix-assigned Cloud Router package identifier | 

### Return type

[**CloudRouterPackage**](CloudRouterPackage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGatewayPackages**
> PackageResponse GetGatewayPackages(ctx, optional)
List Packages

This API provides capability to retrieve user's Cloud Routers Packages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CloudRoutersApiGetGatewayPackagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudRoutersApiGetGatewayPackagesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| offset | 
 **limit** | **optional.Int32**| number of records to fetch | 

### Return type

[**PackageResponse**](PackageResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchGateways**
> SearchResponse SearchGateways(ctx, body)
Search Routers

The API provides capability to get list of user's Cloud Routers using search criteria, including optional filtering, pagination and sorting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CloudRouterSearchRequest**](CloudRouterSearchRequest.md)|  | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGatewayByUuid**
> CloudRouter UpdateGatewayByUuid(ctx, body, routerId)
Update Routers

This API provides capability to update user's Cloud Routers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]CloudRouterChangeOperation**](CloudRouterChangeOperation.md)|  | 
  **routerId** | [**string**](.md)| Cloud Router UUID | 

### Return type

[**CloudRouter**](CloudRouter.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json-patch+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

