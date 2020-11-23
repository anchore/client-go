# \ImportsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FinalizeImageImport**](ImportsApi.md#FinalizeImageImport) | **Post** /imports/images/{operation_id}/finalize | Finalize the import, indicating all data is present and the system can load the image
[**GetImageImport**](ImportsApi.md#GetImageImport) | **Get** /imports/images/{operation_id} | Get detail on a single import
[**ImportImagePackages**](ImportsApi.md#ImportImagePackages) | **Post** /imports/images/{operation_id}/packages | Begin the import of an image analyzed by Syft into the system
[**ListImageImportPackages**](ImportsApi.md#ListImageImportPackages) | **Get** /imports/images/{operation_id}/packages | Begin the import of an image analyzed by Syft into the system
[**StartImageImport**](ImportsApi.md#StartImageImport) | **Post** /imports/images | Begin the import of an image analyzed by Syft into the system



## FinalizeImageImport

> FinalizeImageImportResponse FinalizeImageImport(ctx, operationId, manifest)

Finalize the import, indicating all data is present and the system can load the image

Ends the import process upload phase and begins the import async process to load the image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string**|  | 
**manifest** | [**ImageImportManifest**](ImageImportManifest.md)|  | 

### Return type

[**FinalizeImageImportResponse**](FinalizeImageImportResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageImport

> []ImageImportOperation GetImageImport(ctx, operationId)

Get detail on a single import

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string**|  | 

### Return type

[**[]ImageImportOperation**](ImageImportOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportImagePackages

> ImageImportContentResponse ImportImagePackages(ctx, operationId, sbom)

Begin the import of an image analyzed by Syft into the system

Starts the import process

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string**|  | 
**sbom** | [**[]SyftPackage**](SyftPackage.md)|  | 

### Return type

[**ImageImportContentResponse**](ImageImportContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImageImportPackages

> []string ListImageImportPackages(ctx, operationId)

Begin the import of an image analyzed by Syft into the system

Starts the import process

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string**|  | 

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


## StartImageImport

> ImageImportOperation StartImageImport(ctx, )

Begin the import of an image analyzed by Syft into the system

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ImageImportOperation**](ImageImportOperation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

