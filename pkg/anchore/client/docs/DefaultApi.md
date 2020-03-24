# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOauthToken**](DefaultApi.md#GetOauthToken) | **Post** /oauth/token | 
[**HealthNoop**](DefaultApi.md#HealthNoop) | **Get** /health | 
[**ListFileContentSearchResults**](DefaultApi.md#ListFileContentSearchResults) | **Get** /images/{imageDigest}/artifacts/file_content_search | Return a list of analyzer artifacts of the specified type
[**ListRetrievedFiles**](DefaultApi.md#ListRetrievedFiles) | **Get** /images/{imageDigest}/artifacts/retrieved_files | Return a list of analyzer artifacts of the specified type
[**ListSecretSearchResults**](DefaultApi.md#ListSecretSearchResults) | **Get** /images/{imageDigest}/artifacts/secret_search | Return a list of analyzer artifacts of the specified type
[**Ping**](DefaultApi.md#Ping) | **Get** / | 
[**VersionNoop**](DefaultApi.md#VersionNoop) | **Get** /version | 



## GetOauthToken

> TokenResponse GetOauthToken(ctx, )


Request a jwt token for subsequent operations, this request is authenticated with normal HTTP auth

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthNoop

> HealthNoop(ctx, )


Health check, returns 200 and no body if service is running

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFileContentSearchResults

> []FileContentSearchResult ListFileContentSearchResults(ctx, imageDigest)
Return a list of analyzer artifacts of the specified type

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**|  | 

### Return type

[**[]FileContentSearchResult**](FileContentSearchResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRetrievedFiles

> []RetrievedFile ListRetrievedFiles(ctx, imageDigest)
Return a list of analyzer artifacts of the specified type

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**|  | 

### Return type

[**[]RetrievedFile**](RetrievedFile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecretSearchResults

> []SecretSearchResult ListSecretSearchResults(ctx, imageDigest)
Return a list of analyzer artifacts of the specified type

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**|  | 

### Return type

[**[]SecretSearchResult**](SecretSearchResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Ping

> string Ping(ctx, )


Simple status check

### Required Parameters

This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionNoop

> ServiceVersion VersionNoop(ctx, )


Returns the version object for the service, including db schema version info

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ServiceVersion**](ServiceVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

