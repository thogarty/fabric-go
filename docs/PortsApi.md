# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddToLag**](PortsApi.md#AddToLag) | **Post** /fabric/v4/ports/{portId}/physicalPorts/bulk | Add to Lag
[**CreateBulkPort**](PortsApi.md#CreateBulkPort) | **Post** /fabric/v4/ports/bulk | Create Port
[**CreatePort**](PortsApi.md#CreatePort) | **Post** /fabric/v4/ports | Create Port
[**GetPortByUuid**](PortsApi.md#GetPortByUuid) | **Get** /fabric/v4/ports/{portId} | Get Port by uuid
[**GetPorts**](PortsApi.md#GetPorts) | **Get** /fabric/v4/ports | Get All Ports
[**GetVlans**](PortsApi.md#GetVlans) | **Get** /fabric/v4/ports/{portUuid}/linkProtocols | Get Vlans

# **AddToLag**
> AllPhysicalPortsResponse AddToLag(ctx, body, portId)
Add to Lag

Add Physical Ports to Virtual Port.<font color=\"red\"> <sup color='red'>Preview</sup></font>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkPhysicalPort**](BulkPhysicalPort.md)|  | 
  **portId** | [**string**](.md)| Port UUID | 

### Return type

[**AllPhysicalPortsResponse**](AllPhysicalPortsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBulkPort**
> AllPortsResponse CreateBulkPort(ctx, body)
Create Port

Create Port creates Equinix Fabric? Port.<font color=\"red\"> <sup color='red'>Preview</sup></font>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkPort**](BulkPort.md)|  | 

### Return type

[**AllPortsResponse**](AllPortsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePort**
> Port CreatePort(ctx, body)
Create Port

Creates Equinix Fabric? Port.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Port**](Port.md)|  | 

### Return type

[**Port**](Port.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

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

# **GetVlans**
> LinkProtocolGetResponse GetVlans(ctx, portUuid)
Get Vlans

The API provides capability to retrieve Vlans for a Port.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portUuid** | [**string**](.md)| Port UUID | 

### Return type

[**LinkProtocolGetResponse**](LinkProtocolGetResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

