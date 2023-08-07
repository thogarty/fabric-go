# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchPrices**](PricesApi.md#SearchPrices) | **Post** /fabric/v4/prices/search | Get Prices

# **SearchPrices**
> PriceSearchResponse SearchPrices(ctx, body)
Get Prices

Search prices by search criteria

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FilterBody**](FilterBody.md)|  | 

### Return type

[**PriceSearchResponse**](PriceSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

