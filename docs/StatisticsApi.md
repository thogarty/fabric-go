# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConnectionStatsByPortUuid**](StatisticsApi.md#GetConnectionStatsByPortUuid) | **Get** /fabric/v4/connections/{connectionId}/stats | Get Stats by uuid
[**GetPortStats**](StatisticsApi.md#GetPortStats) | **Get** /fabric/v4/ports/stats | Top Port Statistics
[**GetPortStatsByPortUuid**](StatisticsApi.md#GetPortStatsByPortUuid) | **Get** /fabric/v4/ports/{portId}/stats | Get Stats by uuid

# **GetConnectionStatsByPortUuid**
> Statistics GetConnectionStatsByPortUuid(ctx, connectionId, startDateTime, endDateTime, viewPoint)
Get Stats by uuid

This API provides service-level metrics so that you can view access and gather key information required to manage service subscription sizing and capacity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | [**string**](.md)| Connection UUID | 
  **startDateTime** | [**time.Time**](.md)| startDateTime | 
  **endDateTime** | [**time.Time**](.md)| endDateTime | 
  **viewPoint** | [**ViewPoint**](.md)| viewPoint | 

### Return type

[**Statistics**](Statistics.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPortStats**
> TopUtilizedStatistics GetPortStats(ctx, metros, optional)
Top Port Statistics

This API provides top utilized service-level traffic metrics so that you can view access and gather key information required to manage service subscription sizing and capacity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **metros** | [**[]string**](.md)| Two-letter prefix indicating the metropolitan area in which a specified Equinix asset is located. | 
 **optional** | ***StatisticsApiGetPortStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatisticsApiGetPortStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | [**optional.Interface of Sort**](.md)| Key or set of keys that organizes the search payload by property (such as createdDate or metroCode) or by direction. Ascending (ASC) is the default value. The \&quot;‒\&quot; prefix indicates descending (DESC) order. | 
 **top** | [**optional.Interface of int32**](.md)| Filter returning only the specified number of most heavily trafficked ports. The standard value is [1...10], and the default is 5. | 
 **duration** | [**optional.Interface of Duration**](.md)| duration | 
 **direction** | [**optional.Interface of QueryDirection**](.md)| Direction of traffic from the requester&#x27;s viewpoint. The default is outbound. | 
 **metricInterval** | [**optional.Interface of MetricInterval**](.md)| metricInterval | 
 **projectId** | [**optional.Interface of string**](.md)| projectId | 

### Return type

[**TopUtilizedStatistics**](TopUtilizedStatistics.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPortStatsByPortUuid**
> Statistics GetPortStatsByPortUuid(ctx, portId, startDateTime, endDateTime)
Get Stats by uuid

This API provides service-level traffic metrics so that you can view access and gather key information required to manage service subscription sizing and capacity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portId** | [**string**](.md)| Port UUID | 
  **startDateTime** | [**time.Time**](.md)| startDateTime | 
  **endDateTime** | [**time.Time**](.md)| endDateTime | 

### Return type

[**Statistics**](Statistics.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

