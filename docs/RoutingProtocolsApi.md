# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectionRoutingProtocol**](RoutingProtocolsApi.md#CreateConnectionRoutingProtocol) | **Post** /fabric/v4/connections/{connectionId}/routingProtocols | Create Protocol
[**CreateConnectionRoutingProtocolsInBulk**](RoutingProtocolsApi.md#CreateConnectionRoutingProtocolsInBulk) | **Post** /fabric/v4/connections/{connectionId}/routingProtocols/bulk | Bulk Create Protocol
[**DeleteConnectionRoutingProtocolByUuid**](RoutingProtocolsApi.md#DeleteConnectionRoutingProtocolByUuid) | **Delete** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId} | Delete Protocol
[**GetConnectionRoutingProtocolAllBgpActions**](RoutingProtocolsApi.md#GetConnectionRoutingProtocolAllBgpActions) | **Get** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId}/actions | Get BGP Actions
[**GetConnectionRoutingProtocolByUuid**](RoutingProtocolsApi.md#GetConnectionRoutingProtocolByUuid) | **Get** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId} | Get Protocol
[**GetConnectionRoutingProtocols**](RoutingProtocolsApi.md#GetConnectionRoutingProtocols) | **Get** /fabric/v4/connections/{connectionId}/routingProtocols | GetRoutingProtocols
[**GetConnectionRoutingProtocolsBgpActionByUuid**](RoutingProtocolsApi.md#GetConnectionRoutingProtocolsBgpActionByUuid) | **Get** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId}/actions/{actionId} | Get BGP Action
[**GetConnectionRoutingProtocolsChangeByUuid**](RoutingProtocolsApi.md#GetConnectionRoutingProtocolsChangeByUuid) | **Get** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId}/changes/{changeId} | Get Change By ID
[**GetConnectionRoutingProtocolsChanges**](RoutingProtocolsApi.md#GetConnectionRoutingProtocolsChanges) | **Get** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId}/changes | Get Changes
[**PatchConnectionRoutingProtocolByUuid**](RoutingProtocolsApi.md#PatchConnectionRoutingProtocolByUuid) | **Patch** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId} | Patch Protocol
[**PostConnectionRoutingProtocolBgpActionByUuid**](RoutingProtocolsApi.md#PostConnectionRoutingProtocolBgpActionByUuid) | **Post** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId}/actions | Clear/Reset BGP
[**ReplaceConnectionRoutingProtocolByUuid**](RoutingProtocolsApi.md#ReplaceConnectionRoutingProtocolByUuid) | **Put** /fabric/v4/connections/{connectionId}/routingProtocols/{routingProtocolId} | Replace Protocol
[**ValidateRoutingProtocol**](RoutingProtocolsApi.md#ValidateRoutingProtocol) | **Post** /fabric/v4/routers/{routerId}/validate | Validate Subnet

# **CreateConnectionRoutingProtocol**
> RoutingProtocolData CreateConnectionRoutingProtocol(ctx, body, connectionId)
Create Protocol

This API provides capability to create Routing Protocol for connections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutingProtocolBase**](RoutingProtocolBase.md)|  | 
  **connectionId** | [**string**](.md)| Connection Id | 

### Return type

[**RoutingProtocolData**](RoutingProtocolData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConnectionRoutingProtocolsInBulk**
> GetResponse CreateConnectionRoutingProtocolsInBulk(ctx, body, connectionId)
Bulk Create Protocol

This API provides capability to create Routing Protocol for connections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectionRoutingProtocolPostRequest**](ConnectionRoutingProtocolPostRequest.md)|  | 
  **connectionId** | [**string**](.md)| Connection Id | 

### Return type

[**GetResponse**](GetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConnectionRoutingProtocolByUuid**
> RoutingProtocolData DeleteConnectionRoutingProtocolByUuid(ctx, routingProtocolId, connectionId)
Delete Protocol

This API provides capability to delete Routing Protocols on virtual connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routingProtocolId** | [**string**](.md)| Routing Protocol Id | 
  **connectionId** | [**string**](.md)| Connection Id | 

### Return type

[**RoutingProtocolData**](RoutingProtocolData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectionRoutingProtocolAllBgpActions**
> BgpActionsBulkData GetConnectionRoutingProtocolAllBgpActions(ctx, routingProtocolId, connectionId, optional)
Get BGP Actions

This API provides capability to get all BGP actions status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routingProtocolId** | [**string**](.md)| Routing Protocol Id | 
  **connectionId** | [**string**](.md)| Connection Id | 
 **optional** | ***RoutingProtocolsApiGetConnectionRoutingProtocolAllBgpActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoutingProtocolsApiGetConnectionRoutingProtocolAllBgpActionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **optional.Int32**| offset | 
 **limit** | **optional.Int32**| number of records to fetch | 

### Return type

[**BgpActionsBulkData**](BGPActionsBulkData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectionRoutingProtocolByUuid**
> RoutingProtocolData GetConnectionRoutingProtocolByUuid(ctx, routingProtocolId, connectionId)
Get Protocol

This API provides capability to accept/reject user's virtual connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routingProtocolId** | [**string**](.md)| Routing Protocol Id | 
  **connectionId** | [**string**](.md)| Connection Id | 

### Return type

[**RoutingProtocolData**](RoutingProtocolData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectionRoutingProtocols**
> GetResponse GetConnectionRoutingProtocols(ctx, connectionId, optional)
GetRoutingProtocols

This API provides capability to get Routing Protocols for connections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | [**string**](.md)| Connection Id | 
 **optional** | ***RoutingProtocolsApiGetConnectionRoutingProtocolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoutingProtocolsApiGetConnectionRoutingProtocolsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| offset | 
 **limit** | **optional.Int32**| number of records to fetch | 

### Return type

[**GetResponse**](GetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectionRoutingProtocolsBgpActionByUuid**
> BgpActionData GetConnectionRoutingProtocolsBgpActionByUuid(ctx, connectionId, routingProtocolId, actionId)
Get BGP Action

This API provides capability to retrieve specific BGP action

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | [**string**](.md)| Connection Id | 
  **routingProtocolId** | [**string**](.md)| Routing Protocol Id | 
  **actionId** | [**string**](.md)| BGP Action UUID | 

### Return type

[**BgpActionData**](BGPActionData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectionRoutingProtocolsChangeByUuid**
> RoutingProtocolChangeData GetConnectionRoutingProtocolsChangeByUuid(ctx, connectionId, routingProtocolId, changeId)
Get Change By ID

This API provides capability to retrieve specific Routing Protocol Change

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | [**string**](.md)| Connection Id | 
  **routingProtocolId** | [**string**](.md)| Routing Protocol Id | 
  **changeId** | [**string**](.md)| Routing Protocol Change UUID | 

### Return type

[**RoutingProtocolChangeData**](RoutingProtocolChangeData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectionRoutingProtocolsChanges**
> RoutingProtocolChangeDataResponse GetConnectionRoutingProtocolsChanges(ctx, connectionId, routingProtocolId, optional)
Get Changes

This API provides capability to retrieve user's Routing Protocol Changes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | [**string**](.md)| Connection Id | 
  **routingProtocolId** | [**string**](.md)| Routing Protocol Id | 
 **optional** | ***RoutingProtocolsApiGetConnectionRoutingProtocolsChangesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoutingProtocolsApiGetConnectionRoutingProtocolsChangesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **optional.Int32**| offset | 
 **limit** | **optional.Int32**| number of records to fetch | 

### Return type

[**RoutingProtocolChangeDataResponse**](RoutingProtocolChangeDataResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchConnectionRoutingProtocolByUuid**
> RoutingProtocolData PatchConnectionRoutingProtocolByUuid(ctx, body, routingProtocolId, connectionId)
Patch Protocol

This API provides capability to partially update Routing Protocols on a virtual connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]ConnectionChangeOperation**](ConnectionChangeOperation.md)|  | 
  **routingProtocolId** | [**string**](.md)| Routing Protocol Id | 
  **connectionId** | [**string**](.md)| Connection Id | 

### Return type

[**RoutingProtocolData**](RoutingProtocolData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostConnectionRoutingProtocolBgpActionByUuid**
> BgpActionData PostConnectionRoutingProtocolBgpActionByUuid(ctx, body, routingProtocolId, connectionId)
Clear/Reset BGP

This API provides capability to clear/reset Routing Protocols BGP session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BgpActionRequest**](BgpActionRequest.md)|  | 
  **routingProtocolId** | [**string**](.md)| Routing Protocol Id | 
  **connectionId** | [**string**](.md)| Connection Id | 

### Return type

[**BgpActionData**](BGPActionData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceConnectionRoutingProtocolByUuid**
> RoutingProtocolData ReplaceConnectionRoutingProtocolByUuid(ctx, body, routingProtocolId, connectionId)
Replace Protocol

This API provides capability to replace complete Routing Protocols on a virtual connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutingProtocolBase**](RoutingProtocolBase.md)|  | 
  **routingProtocolId** | [**string**](.md)| Routing Protocol Id | 
  **connectionId** | [**string**](.md)| Connection Id | 

### Return type

[**RoutingProtocolData**](RoutingProtocolData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateRoutingProtocol**
> ValidateSubnetResponse ValidateRoutingProtocol(ctx, body, routerId)
Validate Subnet

This API provides capability to validate all subnets associated with any connection in the given FCR

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ValidateSubnetRequest**](ValidateSubnetRequest.md)|  | 
  **routerId** | [**string**](.md)| Cloud Router UUID | 

### Return type

[**ValidateSubnetResponse**](ValidateSubnetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

