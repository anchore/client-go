# \QueryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryImagesByPackage**](QueryApi.md#QueryImagesByPackage) | **Get** /query/images/by_package | List of images containing given package
[**QueryImagesByVulnerability**](QueryApi.md#QueryImagesByVulnerability) | **Get** /query/images/by_vulnerability | List images vulnerable to the specific vulnerability ID.
[**QueryVulnerabilities**](QueryApi.md#QueryVulnerabilities) | **Get** /query/vulnerabilities | Listing information about given vulnerability



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

