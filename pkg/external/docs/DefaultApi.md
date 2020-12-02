# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImageContentByType**](DefaultApi.md#GetImageContentByType) | **Get** /images/{imageDigest}/content/{ctype} | Get the content of an image by type
[**GetImageContentByTypeFiles**](DefaultApi.md#GetImageContentByTypeFiles) | **Get** /images/{imageDigest}/content/files | Get the content of an image by type files
[**GetImageContentByTypeImageId**](DefaultApi.md#GetImageContentByTypeImageId) | **Get** /images/by_id/{imageId}/content/{ctype} | Get the content of an image by type
[**GetImageContentByTypeImageIdFiles**](DefaultApi.md#GetImageContentByTypeImageIdFiles) | **Get** /images/by_id/{imageId}/content/files | Get the content of an image by type files
[**GetImageContentByTypeImageIdJavapackage**](DefaultApi.md#GetImageContentByTypeImageIdJavapackage) | **Get** /images/by_id/{imageId}/content/java | Get the content of an image by type java
[**GetImageContentByTypeJavapackage**](DefaultApi.md#GetImageContentByTypeJavapackage) | **Get** /images/{imageDigest}/content/java | Get the content of an image by type java
[**GetImageContentByTypeMalware**](DefaultApi.md#GetImageContentByTypeMalware) | **Get** /images/{imageDigest}/content/malware | Get the content of an image by type malware
[**GetImagePolicyCheck**](DefaultApi.md#GetImagePolicyCheck) | **Get** /images/{imageDigest}/check | Check policy evaluation status for image
[**GetImagePolicyCheckByImageId**](DefaultApi.md#GetImagePolicyCheckByImageId) | **Get** /images/by_id/{imageId}/check | Check policy evaluation status for image
[**GetOauthToken**](DefaultApi.md#GetOauthToken) | **Post** /oauth/token | 
[**HealthCheck**](DefaultApi.md#HealthCheck) | **Get** /health | 
[**ListFileContentSearchResults**](DefaultApi.md#ListFileContentSearchResults) | **Get** /images/{imageDigest}/artifacts/file_content_search | Return a list of analyzer artifacts of the specified type
[**ListImageContent**](DefaultApi.md#ListImageContent) | **Get** /images/{imageDigest}/content | List image content types
[**ListImageContentByImageid**](DefaultApi.md#ListImageContentByImageid) | **Get** /images/by_id/{imageId}/content | List image content types
[**ListRetrievedFiles**](DefaultApi.md#ListRetrievedFiles) | **Get** /images/{imageDigest}/artifacts/retrieved_files | Return a list of analyzer artifacts of the specified type
[**ListSecretSearchResults**](DefaultApi.md#ListSecretSearchResults) | **Get** /images/{imageDigest}/artifacts/secret_search | Return a list of analyzer artifacts of the specified type
[**Ping**](DefaultApi.md#Ping) | **Get** / | 
[**QueryImagesByPackage**](DefaultApi.md#QueryImagesByPackage) | **Get** /query/images/by_package | List of images containing given package
[**VersionCheck**](DefaultApi.md#VersionCheck) | **Get** /version | 



## GetImageContentByType

> ContentPackageResponse GetImageContentByType(ctx, imageDigest, ctype, optional)

Get the content of an image by type

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**|  | 
**ctype** | **string**|  | 
 **optional** | ***GetImageContentByTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImageContentByTypeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ContentPackageResponse**](ContentPackageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageContentByTypeFiles

> ContentFilesResponse GetImageContentByTypeFiles(ctx, imageDigest, optional)

Get the content of an image by type files

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**|  | 
 **optional** | ***GetImageContentByTypeFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImageContentByTypeFilesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ContentFilesResponse**](ContentFilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageContentByTypeImageId

> ContentPackageResponse GetImageContentByTypeImageId(ctx, imageId, ctype, optional)

Get the content of an image by type

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string**|  | 
**ctype** | **string**|  | 
 **optional** | ***GetImageContentByTypeImageIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImageContentByTypeImageIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ContentPackageResponse**](ContentPackageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageContentByTypeImageIdFiles

> ContentFilesResponse GetImageContentByTypeImageIdFiles(ctx, imageId, optional)

Get the content of an image by type files

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string**|  | 
 **optional** | ***GetImageContentByTypeImageIdFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImageContentByTypeImageIdFilesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ContentFilesResponse**](ContentFilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageContentByTypeImageIdJavapackage

> ContentJavaPackageResponse GetImageContentByTypeImageIdJavapackage(ctx, imageId, optional)

Get the content of an image by type java

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string**|  | 
 **optional** | ***GetImageContentByTypeImageIdJavapackageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImageContentByTypeImageIdJavapackageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ContentJavaPackageResponse**](ContentJAVAPackageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageContentByTypeJavapackage

> ContentJavaPackageResponse GetImageContentByTypeJavapackage(ctx, imageDigest, optional)

Get the content of an image by type java

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**|  | 
 **optional** | ***GetImageContentByTypeJavapackageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImageContentByTypeJavapackageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ContentJavaPackageResponse**](ContentJAVAPackageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageContentByTypeMalware

> ContentMalwareResponse GetImageContentByTypeMalware(ctx, imageDigest, optional)

Get the content of an image by type malware

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**|  | 
 **optional** | ***GetImageContentByTypeMalwareOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImageContentByTypeMalwareOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**ContentMalwareResponse**](ContentMalwareResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImagePolicyCheck

> []interface{} GetImagePolicyCheck(ctx, imageDigest, tag, optional)

Check policy evaluation status for image

Get the policy evaluation for the given image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**|  | 
**tag** | **string**|  | 
 **optional** | ***GetImagePolicyCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImagePolicyCheckOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policyId** | **optional.String**|  | 
 **detail** | **optional.Bool**|  | 
 **history** | **optional.Bool**|  | 
 **interactive** | **optional.Bool**|  | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImagePolicyCheckByImageId

> []interface{} GetImagePolicyCheckByImageId(ctx, imageId, tag, optional)

Check policy evaluation status for image

Get the policy evaluation for the given image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string**|  | 
**tag** | **string**|  | 
 **optional** | ***GetImagePolicyCheckByImageIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImagePolicyCheckByImageIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policyId** | **optional.String**|  | 
 **detail** | **optional.Bool**|  | 
 **history** | **optional.Bool**|  | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOauthToken

> TokenResponse GetOauthToken(ctx, optional)



Request a jwt token for subsequent operations, this request is authenticated with normal HTTP auth

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetOauthTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOauthTokenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **optional.String**| OAuth Grant type for token | [default to password]
 **username** | **optional.String**| User to assign OAuth token to | 
 **password** | **optional.String**| Password for corresponding user | 
 **clientId** | **optional.String**| The type of client used for the OAuth token | [default to anonymous]

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthCheck

> HealthCheck(ctx, )



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


## ListImageContent

> []string ListImageContent(ctx, imageDigest, optional)

List image content types

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**|  | 
 **optional** | ***ListImageContentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListImageContentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImageContentByImageid

> []string ListImageContentByImageid(ctx, imageId, optional)

List image content types

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string**|  | 
 **optional** | ***ListImageContentByImageidOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListImageContentByImageidOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

**[]string**

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


## QueryImagesByPackage

> PaginatedImageList QueryImagesByPackage(ctx, name, optional)

List of images containing given package

Filterable query interface to search for images containing specified package

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Name of package to search for (e.g. sed) | 
 **optional** | ***QueryImagesByPackageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryImagesByPackageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **packageType** | **optional.String**| Type of package to filter on (e.g. dpkg) | 
 **version** | **optional.String**| Version of named package to filter on (e.g. 4.4-1) | 
 **page** | **optional.String**| The page of results to fetch. Pages start at 1 | 
 **limit** | **optional.Int32**| Limit the number of records for the requested page. If omitted or set to 0, return all results in a single page | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**PaginatedImageList**](PaginatedImageList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionCheck

> ServiceVersion VersionCheck(ctx, )



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

