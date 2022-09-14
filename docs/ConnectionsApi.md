# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnection**](ConnectionsApi.md#CreateConnection) | **Post** /fabric/v4/connections | Create Connection
[**CreateConnectionAction**](ConnectionsApi.md#CreateConnectionAction) | **Post** /fabric/v4/connections/{connectionId}/actions | Connection Actions
[**CreateConnectionsInBulk**](ConnectionsApi.md#CreateConnectionsInBulk) | **Post** /fabric/v4/connections/bulk | Bulk Connections
[**DeleteConnectionByUuid**](ConnectionsApi.md#DeleteConnectionByUuid) | **Delete** /fabric/v4/connections/{connectionId} | Delete by ID
[**GetConnectionByUuid**](ConnectionsApi.md#GetConnectionByUuid) | **Get** /fabric/v4/connections/{connectionId} | Get Connection by ID
[**SearchConnections**](ConnectionsApi.md#SearchConnections) | **Post** /fabric/v4/connections/search | Search connections
[**UpdateConnectionByUuid**](ConnectionsApi.md#UpdateConnectionByUuid) | **Patch** /fabric/v4/connections/{connectionId} | Update by ID
[**ValidateConnections**](ConnectionsApi.md#ValidateConnections) | **Post** /fabric/v4/connections/validate | Validate AuthKey

# **CreateConnection**
> Connection CreateConnection(ctx, body)
Create Connection

This API provides capability to create user's virtual connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectionPostRequest**](ConnectionPostRequest.md)|  | 

### Return type

[**Connection**](Connection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConnectionAction**
> Connection CreateConnectionAction(ctx, body, connectionId)
Connection Actions

This API provides capability to accept/reject user's virtual connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectionActionRequest**](ConnectionActionRequest.md)|  | 
  **connectionId** | [**string**](.md)| Connection Id | 

### Return type

[**Connection**](Connection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConnectionsInBulk**
> ConnectionBulk CreateConnectionsInBulk(ctx, body)
Bulk Connections

This API provides capability to create bulk virtual connections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectionBulkPostRequest**](ConnectionBulkPostRequest.md)|  | 

### Return type

[**ConnectionBulk**](ConnectionBulk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConnectionByUuid**
> Connection DeleteConnectionByUuid(ctx, connectionId)
Delete by ID

Delete Connection by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | [**string**](.md)| Connection UUID | 

### Return type

[**Connection**](Connection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectionByUuid**
> Connection GetConnectionByUuid(ctx, connectionId, optional)
Get Connection by ID

The API provides capability to get user's virtual connection details (Service Tokens, Access Points, Link Protocols, etc) by it's connection ID (UUID)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | [**string**](.md)| Connection Id | 
 **optional** | ***ConnectionsApiGetConnectionByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectionsApiGetConnectionByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **direction** | [**optional.Interface of ConnectionDirection**](.md)| Connection Direction | 

### Return type

[**Connection**](Connection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchConnections**
> ConnectionSearchResponse SearchConnections(ctx, body)
Search connections

The API provides capability to get list of user's virtual connections using search criteria, including optional filtering, pagination and sorting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchRequest**](SearchRequest.md)|  | 

### Return type

[**ConnectionSearchResponse**](ConnectionSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConnectionByUuid**
> Connection UpdateConnectionByUuid(ctx, body, connectionId)
Update by ID

Update Connection by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]ConnectionChangeOperation**](ConnectionChangeOperation.md)|  | 
  **connectionId** | [**string**](.md)| Connection Id | 

### Return type

[**Connection**](Connection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json-patch+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateConnections**
> ConnectionResponse ValidateConnections(ctx, body)
Validate AuthKey

This API provides capability to validate by auth key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ValidateByAuthKeyPostRequest**](ValidateByAuthKeyPostRequest.md)|  | 

### Return type

[**ConnectionResponse**](ConnectionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

