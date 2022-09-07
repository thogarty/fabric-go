# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFilters**](FiltersApi.md#GetFilters) | **Get** /fabric/v4/filters | Get Search Filters

# **GetFilters**
> FilterResponse GetFilters(ctx, type_, optional)
Get Search Filters

The API provides capability to get list of user's Fabric Resource search criteria

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | [**ResourceType**](.md)| Resource Type | 
 **optional** | ***FiltersApiGetFiltersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FiltersApiGetFiltersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 

### Return type

[**FilterResponse**](FilterResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

