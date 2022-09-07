# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetroByCode**](MetrosApi.md#GetMetroByCode) | **Get** /fabric/v4/metros/{metroCode} | Get Metro by Code
[**GetMetros**](MetrosApi.md#GetMetros) | **Get** /fabric/v4/metros | Get all Metros

# **GetMetroByCode**
> Metro GetMetroByCode(ctx, metroCode, optional)
Get Metro by Code

GET Metros retrieves all Equinix Fabric metros, as well as latency data between each metro location. .

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **metroCode** | [**string**](.md)| Metro Code | 
 **optional** | ***MetrosApiGetMetroByCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MetrosApiGetMetroByCodeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **correlationId** | **optional.String**| Correlation identifier | 

### Return type

[**Metro**](Metro.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMetros**
> MetroResponse GetMetros(ctx, optional)
Get all Metros

GET All Subscriber Metros with an option query parameter to return all Equinix Fabric metros in which the customer has a presence, as well as latency data for each location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MetrosApiGetMetrosOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MetrosApiGetMetrosOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **presence** | [**optional.Interface of Presence**](.md)| User On Boarded Metros based on Fabric resource availability | 
 **offset** | **optional.Int32**| offset | 
 **limit** | **optional.Int32**| number of records to fetch | 
 **correlationId** | **optional.String**| Correlation identifier | 

### Return type

[**MetroResponse**](MetroResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

