# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGateway**](GatewaysApi.md#CreateGateway) | **Post** /fabric/v4/gateways | Create Gateway
[**CreateGatewayAction**](GatewaysApi.md#CreateGatewayAction) | **Post** /fabric/v4/gateways/{gatewayId}/actions | Route table actions
[**DeleteGatewayByUuid**](GatewaysApi.md#DeleteGatewayByUuid) | **Delete** /fabric/v4/gateways/{gatewayId} | Delete Gateway
[**GetGatewayActions**](GatewaysApi.md#GetGatewayActions) | **Get** /fabric/v4/gateways/{gatewayId}/actions | Get actions
[**GetGatewayByUuid**](GatewaysApi.md#GetGatewayByUuid) | **Get** /fabric/v4/gateways/{gatewayId} | Get Gateway
[**GetGatewayPackageByCode**](GatewaysApi.md#GetGatewayPackageByCode) | **Get** /fabric/v4/gatewayPackages/{gatewayPackageCode} | Get Package Details
[**GetGatewayPackages**](GatewaysApi.md#GetGatewayPackages) | **Get** /fabric/v4/gatewayPackages | List Packages
[**SearchGateways**](GatewaysApi.md#SearchGateways) | **Post** /fabric/v4/gateways/search | Search Gateway
[**UpdateGatewayByUuid**](GatewaysApi.md#UpdateGatewayByUuid) | **Patch** /fabric/v4/gateways/{gatewayId} | Update Gateway

# **CreateGateway**
> VirtualGateway CreateGateway(ctx, body, optional)
Create Gateway

This API provides capability to create user's Fabric Gateways

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VirtualGatewayPostRequest**](VirtualGatewayPostRequest.md)|  | 
 **optional** | ***GatewaysApiCreateGatewayOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewaysApiCreateGatewayOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.**| User name | 

### Return type

[**VirtualGateway**](VirtualGateway.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGatewayAction**
> GatewayActionResponse CreateGatewayAction(ctx, body, gatewayId, optional)
Route table actions

This API provides capability to refresh route table and bgp session summary information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GatewayActionRequest**](GatewayActionRequest.md)|  | 
  **gatewayId** | [**string**](.md)| Gateway UUID | 
 **optional** | ***GatewaysApiCreateGatewayActionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewaysApiCreateGatewayActionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.**| User name | 
 **xSOURCE** | **optional.**| source | 

### Return type

[**GatewayActionResponse**](GatewayActionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGatewayByUuid**
> DeleteGatewayByUuid(ctx, gatewayId, optional)
Delete Gateway

This API provides capability to delete user's Fabric Gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | [**string**](.md)| Gateway UUID | 
 **optional** | ***GatewaysApiDeleteGatewayByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewaysApiDeleteGatewayByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGatewayActions**
> GatewayActionResponse GetGatewayActions(ctx, gatewayId, optional)
Get actions

This API provides capability to fetch action status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | [**string**](.md)| Gateway UUID | 
 **optional** | ***GatewaysApiGetGatewayActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewaysApiGetGatewayActionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | [**optional.Interface of GatewayActionState**](.md)| Action state | 
 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 
 **xSOURCE** | **optional.String**| source | 

### Return type

[**GatewayActionResponse**](GatewayActionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGatewayByUuid**
> VirtualGateway GetGatewayByUuid(ctx, gatewayId, optional)
Get Gateway

This API provides capability to retrieve user's Fabric Gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | [**string**](.md)| Gateway UUID | 
 **optional** | ***GatewaysApiGetGatewayByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewaysApiGetGatewayByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 

### Return type

[**VirtualGateway**](VirtualGateway.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGatewayPackageByCode**
> VirtualGatewayPackage GetGatewayPackageByCode(ctx, gatewayPackageCode, optional)
Get Package Details

This API provides capability to retrieve user's Fabric Gateway Package Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayPackageCode** | [**GatewayPackageCode**](.md)| Equinix-assigned gateway package identifier | 
 **optional** | ***GatewaysApiGetGatewayPackageByCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewaysApiGetGatewayPackageByCodeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 

### Return type

[**VirtualGatewayPackage**](VirtualGatewayPackage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGatewayPackages**
> PackageResponse GetGatewayPackages(ctx, optional)
List Packages

This API provides capability to retrieve user's Fabric Gateway Packages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GatewaysApiGetGatewayPackagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewaysApiGetGatewayPackagesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 
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
> SearchResponse SearchGateways(ctx, body, optional)
Search Gateway

The API provides capability to get list of user's Fabric Gateways using search criteria, including optional filtering, pagination and sorting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VirtualGatewaySearchRequest**](VirtualGatewaySearchRequest.md)|  | 
 **optional** | ***GatewaysApiSearchGatewaysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewaysApiSearchGatewaysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.**| User name | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGatewayByUuid**
> VirtualGateway UpdateGatewayByUuid(ctx, body, gatewayId, optional)
Update Gateway

This API provides capability to update user's Fabric Gateway

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]VirtualGatewayChangeOperation**](VirtualGatewayChangeOperation.md)|  | 
  **gatewayId** | [**string**](.md)| Gateway UUID | 
 **optional** | ***GatewaysApiUpdateGatewayByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewaysApiUpdateGatewayByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.**| User name | 

### Return type

[**VirtualGateway**](VirtualGateway.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json-patch+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

