# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchRoutes**](RoutesApi.md#SearchRoutes) | **Post** /fabric/v4/routers/{routerId}/routes/search | Search Route Table

# **SearchRoutes**
> RouteTableEntrySearchResponse SearchRoutes(ctx, body, routerId)
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

