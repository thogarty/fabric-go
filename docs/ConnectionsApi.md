# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnection**](ConnectionsApi.md#CreateConnection) | **Post** /fabric/v4/connections | Create Connection
[**CreateConnectionAction**](ConnectionsApi.md#CreateConnectionAction) | **Post** /fabric/v4/connections/{connectionId}/actions | Connection Actions
[**CreateConnectionsInBulk**](ConnectionsApi.md#CreateConnectionsInBulk) | **Post** /fabric/v4/connections/bulk | Bulk Connections
[**DeleteConnectionByUuid**](ConnectionsApi.md#DeleteConnectionByUuid) | **Delete** /fabric/v4/connections/{connectionId} | Delete by ID
[**GetConnectionByUuid**](ConnectionsApi.md#GetConnectionByUuid) | **Get** /fabric/v4/connections/{connectionId} | Get Connection by ID
[**PutConnectionByUuid**](ConnectionsApi.md#PutConnectionByUuid) | **Put** /fabric/v4/connections/{connectionId} | Replace by ID
[**SearchConnections**](ConnectionsApi.md#SearchConnections) | **Post** /fabric/v4/connections/search | Search connections
[**UpdateConnectionByUuid**](ConnectionsApi.md#UpdateConnectionByUuid) | **Patch** /fabric/v4/connections/{connectionId} | Update by ID
[**UpdateConnectionsByUuids**](ConnectionsApi.md#UpdateConnectionsByUuids) | **Patch** /fabric/v4/connections/bulk | Update Connections
[**ValidateConnections**](ConnectionsApi.md#ValidateConnections) | **Post** /fabric/v4/connections/validate | Validate AuthKey

# **CreateConnection**
> Connection CreateConnection(ctx, body, optional)
Create Connection

This API provides capability to create user's virtual connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectionPostRequest**](ConnectionPostRequest.md)|  | 
 **optional** | ***ConnectionsApiCreateConnectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectionsApiCreateConnectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.**| User name | 
 **xSOURCE** | **optional.**| source | 
 **accountSubCustomerUcmId** | **optional.**| subCustomerUcmId | 

### Return type

[**Connection**](Connection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConnectionAction**
> Connection CreateConnectionAction(ctx, body, connectionId, optional)
Connection Actions

This API provides capability to accept/reject user's virtual connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectionActionRequest**](ConnectionActionRequest.md)|  | 
  **connectionId** | [**string**](.md)| Connection Id | 
 **optional** | ***ConnectionsApiCreateConnectionActionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectionsApiCreateConnectionActionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.**| User name | 
 **xSOURCE** | **optional.**| source | 

### Return type

[**Connection**](Connection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConnectionsInBulk**
> ConnectionBulk CreateConnectionsInBulk(ctx, body, optional)
Bulk Connections

This API provides capability to create bulk virtual connections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectionBulkPostRequest**](ConnectionBulkPostRequest.md)|  | 
 **optional** | ***ConnectionsApiCreateConnectionsInBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectionsApiCreateConnectionsInBulkOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.**| User name | 
 **xSOURCE** | **optional.**| source | 
 **accountSubCustomerUcmId** | **optional.**| subCustomerUcmId | 

### Return type

[**ConnectionBulk**](ConnectionBulk.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConnectionByUuid**
> Connection DeleteConnectionByUuid(ctx, connectionId, optional)
Delete by ID

Delete Connection by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | [**string**](.md)| Connection UUID | 
 **optional** | ***ConnectionsApiDeleteConnectionByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectionsApiDeleteConnectionByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 
 **xSOURCE** | **optional.String**| source | 

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

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 
 **direction** | [**optional.Interface of ConnectionDirection**](.md)| Connection Direction | 

### Return type

[**Connection**](Connection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutConnectionByUuid**
> Connection PutConnectionByUuid(ctx, body, connectionId, optional)
Replace by ID

The API provides capability to replace virtual connection details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectionPutRequest**](ConnectionPutRequest.md)|  | 
  **connectionId** | [**string**](.md)| Connection UUID | 
 **optional** | ***ConnectionsApiPutConnectionByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectionsApiPutConnectionByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.**| User name | 
 **accountSubCustomerUcmId** | **optional.**| subCustomerUcmId | 
 **xSOURCE** | **optional.**| source | 

### Return type

[**Connection**](Connection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchConnections**
> ConnectionSearchResponse SearchConnections(ctx, body, optional)
Search connections

The API provides capability to get list of user's virtual connections using search criteria, including optional filtering, pagination and sorting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchRequest**](SearchRequest.md)|  | 
 **optional** | ***ConnectionsApiSearchConnectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectionsApiSearchConnectionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.**| User name | 

### Return type

[**ConnectionSearchResponse**](ConnectionSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConnectionByUuid**
> Connection UpdateConnectionByUuid(ctx, body, connectionId, optional)
Update by ID

Update Connection by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]ConnectionChangeOperation**](ConnectionChangeOperation.md)|  | 
  **connectionId** | [**string**](.md)| Connection Id | 
 **optional** | ***ConnectionsApiUpdateConnectionByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectionsApiUpdateConnectionByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.**| User name | 
 **xSOURCE** | **optional.**| source | 

### Return type

[**Connection**](Connection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json-patch+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConnectionsByUuids**
> Connection UpdateConnectionsByUuids(ctx, body, uuid, optional)
Update Connections

The API provides bulk capability for Admin users to update virtual connection details (for e.g., Aside Access Points) by it's UUID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]ConnectionChangeOperation**](ConnectionChangeOperation.md)|  | 
  **uuid** | [**string**](.md)| Connection UUID | 
 **optional** | ***ConnectionsApiUpdateConnectionsByUuidsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectionsApiUpdateConnectionsByUuidsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.**| User name | 
 **xSOURCE** | **optional.**| source | 

### Return type

[**Connection**](Connection.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json-patch+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateConnections**
> ConnectionResponse ValidateConnections(ctx, body, optional)
Validate AuthKey

This API provides capability to validate by auth key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ValidateByAuthKeyPostRequest**](ValidateByAuthKeyPostRequest.md)|  | 
 **optional** | ***ConnectionsApiValidateConnectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectionsApiValidateConnectionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.**| User name | 

### Return type

[**ConnectionResponse**](ConnectionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

