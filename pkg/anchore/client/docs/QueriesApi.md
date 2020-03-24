# \QueriesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImageContentByType**](QueriesApi.md#GetImageContentByType) | **Get** /images/{imageDigest}/content/{ctype} | Get the content of an image by type
[**GetImageContentByTypeFiles**](QueriesApi.md#GetImageContentByTypeFiles) | **Get** /images/{imageDigest}/content/files | Get the content of an image by type files
[**GetImageContentByTypeImageId**](QueriesApi.md#GetImageContentByTypeImageId) | **Get** /images/by_id/{imageId}/content/{ctype} | Get the content of an image by type
[**GetImageContentByTypeImageIdFiles**](QueriesApi.md#GetImageContentByTypeImageIdFiles) | **Get** /images/by_id/{imageId}/content/files | Get the content of an image by type files
[**GetImageContentByTypeImageIdJavapackage**](QueriesApi.md#GetImageContentByTypeImageIdJavapackage) | **Get** /images/by_id/{imageId}/content/java | Get the content of an image by type java
[**GetImageContentByTypeJavapackage**](QueriesApi.md#GetImageContentByTypeJavapackage) | **Get** /images/{imageDigest}/content/java | Get the content of an image by type java
[**GetImageMetadataByType**](QueriesApi.md#GetImageMetadataByType) | **Get** /images/{imageDigest}/metadata/{mtype} | Get the metadata of an image by type
[**GetImageVulnerabilitiesByType**](QueriesApi.md#GetImageVulnerabilitiesByType) | **Get** /images/{imageDigest}/vuln/{vtype} | Get vulnerabilities by type
[**GetImageVulnerabilitiesByTypeImageId**](QueriesApi.md#GetImageVulnerabilitiesByTypeImageId) | **Get** /images/by_id/{imageId}/vuln/{vtype} | Get vulnerabilities by type
[**GetImageVulnerabilityTypes**](QueriesApi.md#GetImageVulnerabilityTypes) | **Get** /images/{imageDigest}/vuln | Get vulnerability types
[**GetImageVulnerabilityTypesByImageId**](QueriesApi.md#GetImageVulnerabilityTypesByImageId) | **Get** /images/by_id/{imageId}/vuln | Get vulnerability types
[**ListImageContent**](QueriesApi.md#ListImageContent) | **Get** /images/{imageDigest}/content | List image content types
[**ListImageContentByImageid**](QueriesApi.md#ListImageContentByImageid) | **Get** /images/by_id/{imageId}/content | List image content types
[**ListImageMetadata**](QueriesApi.md#ListImageMetadata) | **Get** /images/{imageDigest}/metadata | List image metadata types
[**QueryImagesByPackage**](QueriesApi.md#QueryImagesByPackage) | **Get** /query/images/by_package | List of images containing given package
[**QueryImagesByVulnerability**](QueriesApi.md#QueryImagesByVulnerability) | **Get** /query/images/by_vulnerability | List images vulnerable to the specific vulnerability ID.
[**QueryVulnerabilities**](QueriesApi.md#QueryVulnerabilities) | **Get** /query/vulnerabilities | Listing information about given vulnerability



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


## GetImageMetadataByType

> MetadataResponse GetImageMetadataByType(ctx, imageDigest, mtype, optional)
Get the metadata of an image by type

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**|  | 
**mtype** | **string**|  | 
 **optional** | ***GetImageMetadataByTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImageMetadataByTypeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**MetadataResponse**](MetadataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageVulnerabilitiesByType

> VulnerabilityResponse GetImageVulnerabilitiesByType(ctx, imageDigest, vtype, optional)
Get vulnerabilities by type

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**|  | 
**vtype** | **string**|  | 
 **optional** | ***GetImageVulnerabilitiesByTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImageVulnerabilitiesByTypeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceRefresh** | **optional.Bool**|  | 
 **vendorOnly** | **optional.Bool**|  | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**VulnerabilityResponse**](VulnerabilityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageVulnerabilitiesByTypeImageId

> VulnerabilityResponse GetImageVulnerabilitiesByTypeImageId(ctx, imageId, vtype, optional)
Get vulnerabilities by type

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string**|  | 
**vtype** | **string**|  | 
 **optional** | ***GetImageVulnerabilitiesByTypeImageIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImageVulnerabilitiesByTypeImageIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**VulnerabilityResponse**](VulnerabilityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageVulnerabilityTypes

> []string GetImageVulnerabilityTypes(ctx, imageDigest, optional)
Get vulnerability types

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**|  | 
 **optional** | ***GetImageVulnerabilityTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImageVulnerabilityTypesOpts struct


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


## GetImageVulnerabilityTypesByImageId

> []string GetImageVulnerabilityTypesByImageId(ctx, imageId, optional)
Get vulnerability types

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string**|  | 
 **optional** | ***GetImageVulnerabilityTypesByImageIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImageVulnerabilityTypesByImageIdOpts struct


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


## ListImageMetadata

> []string ListImageMetadata(ctx, imageDigest, optional)
List image metadata types

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**|  | 
 **optional** | ***ListImageMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListImageMetadataOpts struct


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


## QueryImagesByVulnerability

> PaginatedVulnerableImageList QueryImagesByVulnerability(ctx, vulnerabilityId, optional)
List images vulnerable to the specific vulnerability ID.

Returns a listing of images and their respective packages vulnerable to the given vulnerability ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vulnerabilityId** | **string**| The ID of the vulnerability to search for within all images stored in anchore-engine (e.g. CVE-1999-0001) | 
 **optional** | ***QueryImagesByVulnerabilityOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryImagesByVulnerabilityOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **optional.String**| Filter results to images within the given vulnerability namespace (e.g. debian:8, ubuntu:14.04) | 
 **affectedPackage** | **optional.String**| Filter results to images with vulnable packages with the given package name (e.g. libssl) | 
 **severity** | **optional.String**| Filter results to vulnerable package/vulnerability with the given severity | 
 **vendorOnly** | **optional.Bool**| Filter results to include only vulnerabilities that are not marked as invalid by upstream OS vendor data | [default to true]
 **page** | **optional.Int32**| The page of results to fetch. Pages start at 1 | 
 **limit** | **optional.Int32**| Limit the number of records for the requested page. If omitted or set to 0, return all results in a single page | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**PaginatedVulnerableImageList**](PaginatedVulnerableImageList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryVulnerabilities

> PaginatedVulnerabilityList QueryVulnerabilities(ctx, id, optional)
Listing information about given vulnerability

List (w/filters) vulnerability records known by the system, with affected packages information if present

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**[]string**](string.md)| The ID of the vulnerability (e.g. CVE-1999-0001) | 
 **optional** | ***QueryVulnerabilitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a QueryVulnerabilitiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **affectedPackage** | **optional.String**| Filter results by specified package name (e.g. sed) | 
 **affectedPackageVersion** | **optional.String**| Filter results by specified package version (e.g. 4.4-1) | 
 **page** | **optional.String**| The page of results to fetch. Pages start at 1 | [default to 1]
 **limit** | **optional.Int32**| Limit the number of records for the requested page. If omitted or set to 0, return all results in a single page | 
 **namespace** | [**optional.Interface of []string**](string.md)| Namespace(s) to filter vulnerability records by | 

### Return type

[**PaginatedVulnerabilityList**](PaginatedVulnerabilityList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

