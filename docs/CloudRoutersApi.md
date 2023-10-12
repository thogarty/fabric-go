# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCloudRouter**](CloudRoutersApi.md#CreateCloudRouter) | **Post** /fabric/v4/routers | Create Routers
[**CreateCloudRouterAction**](CloudRoutersApi.md#CreateCloudRouterAction) | **Post** /fabric/v4/routers/{routerId}/actions | Route table actions
[**DeleteCloudRouterByUuid**](CloudRoutersApi.md#DeleteCloudRouterByUuid) | **Delete** /fabric/v4/routers/{routerId} | Delete Routers
[**GetCloudRouterActions**](CloudRoutersApi.md#GetCloudRouterActions) | **Get** /fabric/v4/routers/{routerId}/actions | Get actions
[**GetCloudRouterByUuid**](CloudRoutersApi.md#GetCloudRouterByUuid) | **Get** /fabric/v4/routers/{routerId} | Get Routers
[**GetCloudRouterPackageByCode**](CloudRoutersApi.md#GetCloudRouterPackageByCode) | **Get** /fabric/v4/routerPackages/{routerPackageCode} | Get Package Details
[**GetCloudRouterPackages**](CloudRoutersApi.md#GetCloudRouterPackages) | **Get** /fabric/v4/routerPackages | List Packages
[**SearchCloudRouterRoutes**](CloudRoutersApi.md#SearchCloudRouterRoutes) | **Post** /fabric/v4/routers/{routerId}/routes/search | Search Route Table
[**SearchCloudRouters**](CloudRoutersApi.md#SearchCloudRouters) | **Post** /fabric/v4/routers/search | Search Routers
[**UpdateCloudRouterByUuid**](CloudRoutersApi.md#UpdateCloudRouterByUuid) | **Patch** /fabric/v4/routers/{routerId} | Update Routers

# **CreateCloudRouter**
> CloudRouter CreateCloudRouter(ctx, body)
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

# **CreateCloudRouterAction**
> CloudRouterActionResponse CreateCloudRouterAction(ctx, body, routerId)
Route table actions

This API provides capability to refresh route table and bgp session summary information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CloudRouterActionRequest**](CloudRouterActionRequest.md)|  | 
  **routerId** | [**string**](.md)| Router UUID | 

### Return type

[**CloudRouterActionResponse**](CloudRouterActionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudRouterByUuid**
> DeleteCloudRouterByUuid(ctx, routerId)
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

# **GetCloudRouterActions**
> CloudRouterActionResponse GetCloudRouterActions(ctx, routerId, optional)
Get actions

This API provides capability to fetch action status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routerId** | [**string**](.md)| Router UUID | 
 **optional** | ***CloudRoutersApiGetCloudRouterActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudRoutersApiGetCloudRouterActionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | [**optional.Interface of CloudRouterActionState**](.md)| Action state | 

### Return type

[**CloudRouterActionResponse**](CloudRouterActionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudRouterByUuid**
> CloudRouter GetCloudRouterByUuid(ctx, routerId)
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

# **GetCloudRouterPackageByCode**
> CloudRouterPackage GetCloudRouterPackageByCode(ctx, routerPackageCode)
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

# **GetCloudRouterPackages**
> PackageResponse GetCloudRouterPackages(ctx, optional)
List Packages

This API provides capability to retrieve user's Cloud Routers Packages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CloudRoutersApiGetCloudRouterPackagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CloudRoutersApiGetCloudRouterPackagesOpts struct
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

# **SearchCloudRouterRoutes**
> RouteTableEntrySearchResponse SearchCloudRouterRoutes(ctx, body, routerId)
Search Route Table

The API provides capability to get list of user's Fabric Cloud Router route table entries using search criteria, including optional filtering, pagination and sorting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RouteTableEntrySearchRequest**](RouteTableEntrySearchRequest.md)|  | 
  **routerId** | [**string**](.md)| Router UUID | 

### Return type

[**RouteTableEntrySearchResponse**](RouteTableEntrySearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchCloudRouters**
> SearchResponse SearchCloudRouters(ctx, body)
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

# **UpdateCloudRouterByUuid**
> CloudRouter UpdateCloudRouterByUuid(ctx, body, routerId)
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

