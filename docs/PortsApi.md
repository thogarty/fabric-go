# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePort**](PortsApi.md#CreatePort) | **Post** /fabric/v4/ports | Create Port
[**GetPortByUuid**](PortsApi.md#GetPortByUuid) | **Get** /fabric/v4/ports/{portId} | Get Port by uuid
[**GetPorts**](PortsApi.md#GetPorts) | **Get** /fabric/v4/ports | Get All Ports
[**SearchPorts**](PortsApi.md#SearchPorts) | **Post** /fabric/i4/ports/search | Search ports

# **CreatePort**
> AllPortsResponse CreatePort(ctx, body, optional)
Create Port

Create Port creates Equinix Fabric™ Port.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Port**](Port.md)|  | 
 **optional** | ***PortsApiCreatePortOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortsApiCreatePortOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.**| User name | 

### Return type

[**AllPortsResponse**](AllPortsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPortByUuid**
> Port GetPortByUuid(ctx, portId)
Get Port by uuid

Get Port By uuid returns details of assigned and available Equinix Fabric port for the specified user credentials. The metro code attribute in the response shows the origin of the proposed connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portId** | [**string**](.md)| Port UUID | 

### Return type

[**Port**](Port.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPorts**
> AllPortsResponse GetPorts(ctx, optional)
Get All Ports

Get All Ports returns details of all assigned and available ports for the specified user credentials. The metro attribute in the response shows the origin of the proposed connection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortsApiGetPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortsApiGetPortsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| port name to be provided if specific port(s) to be retrieved | 

### Return type

[**AllPortsResponse**](AllPortsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchPorts**
> AllPortsResponse SearchPorts(ctx, body, optional)
Search ports

The API provides capability to get list of user's virtual ports using search criteria, including optional filtering, pagination and sorting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PortSearchRequest**](PortSearchRequest.md)|  | 
 **optional** | ***PortsApiSearchPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortsApiSearchPortsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.**| Correlation identifier | 

### Return type

[**AllPortsResponse**](AllPortsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

