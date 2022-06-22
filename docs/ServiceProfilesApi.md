# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceProfile**](ServiceProfilesApi.md#CreateServiceProfile) | **Post** /fabric/v4/serviceProfiles | Create Profile
[**DeleteServiceProfileByUuid**](ServiceProfilesApi.md#DeleteServiceProfileByUuid) | **Delete** /fabric/v4/serviceProfiles/{serviceProfileId} | Delete Profile
[**GetServiceProfileByUuid**](ServiceProfilesApi.md#GetServiceProfileByUuid) | **Get** /fabric/v4/serviceProfiles/{serviceProfileId} | Get Profile
[**GetServiceProfiles**](ServiceProfilesApi.md#GetServiceProfiles) | **Get** /fabric/v4/serviceProfiles | Get all Profiles
[**PutServiceProfileByUuid**](ServiceProfilesApi.md#PutServiceProfileByUuid) | **Put** /fabric/v4/serviceProfiles/{serviceProfileId} | Replace Profile
[**SearchServiceProfiles**](ServiceProfilesApi.md#SearchServiceProfiles) | **Post** /fabric/v4/serviceProfiles/search | Profile Search
[**UpdateServiceProfileByUuid**](ServiceProfilesApi.md#UpdateServiceProfileByUuid) | **Patch** /fabric/v4/serviceProfiles/{serviceProfileId} | Update Profile

# **CreateServiceProfile**
> ServiceProfile CreateServiceProfile(ctx, body)
Create Profile

Create Service Profile creates Equinix Fabric? Service Profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceProfileRequest**](ServiceProfileRequest.md)|  | 

### Return type

[**ServiceProfile**](ServiceProfile.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceProfileByUuid**
> ServiceProfile DeleteServiceProfileByUuid(ctx, serviceProfileId)
Delete Profile

delete Service Profile by UUID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceProfileId** | [**string**](.md)| Service Profile UUID | 

### Return type

[**ServiceProfile**](ServiceProfile.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceProfileByUuid**
> ServiceProfile GetServiceProfileByUuid(ctx, serviceProfileId, optional)
Get Profile

Get service profile by UUID. View Point parameter if set to zSide will give seller's view of the profile otherwise buyer's view.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceProfileId** | [**string**](.md)| Service Profile UUID | 
 **optional** | ***ServiceProfilesApiGetServiceProfileByUuidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceProfilesApiGetServiceProfileByUuidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **viewPoint** | **optional.String**| flips view between buyer and seller representation | [default to aSide]

### Return type

[**ServiceProfile**](ServiceProfile.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceProfiles**
> ServiceProfiles GetServiceProfiles(ctx, optional)
Get all Profiles

The API request returns all Equinix Fabric Service Profiles in accordance with the view point requested.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServiceProfilesApiGetServiceProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceProfilesApiGetServiceProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| offset | 
 **limit** | **optional.Int32**| number of records to fetch | 
 **viewPoint** | **optional.String**| flips view between buyer and seller representation | [default to aSide]

### Return type

[**ServiceProfiles**](ServiceProfiles.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json; charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutServiceProfileByUuid**
> ServiceProfile PutServiceProfileByUuid(ctx, body, serviceProfileId, ifMatch)
Replace Profile

This API request replaces a service profile definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceProfileRequest**](ServiceProfileRequest.md)|  | 
  **serviceProfileId** | [**string**](.md)| Service Profile UUID | 
  **ifMatch** | **string**| conditional request | 

### Return type

[**ServiceProfile**](ServiceProfile.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchServiceProfiles**
> ServiceProfiles SearchServiceProfiles(ctx, body)
Profile Search

Search service profiles by search criteria

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceProfileSearchRequest**](ServiceProfileSearchRequest.md)|  | 

### Return type

[**ServiceProfiles**](ServiceProfiles.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json; charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceProfileByUuid**
> ServiceProfile UpdateServiceProfileByUuid(ctx, body, serviceProfileId, ifMatch)
Update Profile

Update Service Profile by UUID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]JsonPatchOperation**](JsonPatchOperation.md)|  | 
  **serviceProfileId** | [**string**](.md)| Service Profile UUID | 
  **ifMatch** | **string**| conditional request | 

### Return type

[**ServiceProfile**](ServiceProfile.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json
 - **Accept**: application/json; charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

