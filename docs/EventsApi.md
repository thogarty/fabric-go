# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](EventsApi.md#GetEvent) | **Get** /fabric/v4/events/{eventId} | Get Event
[**GetEventsByAsset**](EventsApi.md#GetEventsByAsset) | **Get** /fabric/v4/events | Get Events of Asset

# **GetEvent**
> GetEvent(ctx, eventId, connectionId, portId, type_)
Get Event

The API provides capability to get list of events for user Asset, including pagination

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventId** | [**string**](.md)| Event UUID | 
  **connectionId** | [**string**](.md)| Connection UUID | 
  **portId** | [**string**](.md)| Port UUID | 
  **type_** | [**interface{}**](.md)| type of events | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventsByAsset**
> GetEventsByAsset(ctx, connectionId, portId, type_)
Get Events of Asset

The API provides capability to get list of events for user Asset, including pagination

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | [**string**](.md)| Connection UUID | 
  **portId** | [**string**](.md)| Port UUID | 
  **type_** | [**interface{}**](.md)| type of events | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

