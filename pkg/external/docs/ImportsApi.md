# \ImportsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOperation**](ImportsApi.md#CreateOperation) | **Post** /imports/images | Begin the import of an image analyzed by Syft into the system
[**GetOperation**](ImportsApi.md#GetOperation) | **Get** /imports/images/{operation_id} | Get detail on a single import
[**ImportImageConfig**](ImportsApi.md#ImportImageConfig) | **Post** /imports/images/{operation_id}/image_config | Import a docker or OCI image config to associate with the image
[**ImportImageDockerfile**](ImportsApi.md#ImportImageDockerfile) | **Post** /imports/images/{operation_id}/dockerfile | Begin the import of an image analyzed by Syft into the system
[**ImportImageManifest**](ImportsApi.md#ImportImageManifest) | **Post** /imports/images/{operation_id}/manifest | Import a docker or OCI distribution manifest to associate with the image
[**ImportImagePackages**](ImportsApi.md#ImportImagePackages) | **Post** /imports/images/{operation_id}/packages | Begin the import of an image analyzed by Syft into the system
[**ImportImageParentManifest**](ImportsApi.md#ImportImageParentManifest) | **Post** /imports/images/{operation_id}/parent_manifest | Import a docker or OCI distribution manifest list to associate with the image
[**InvalidateOperation**](ImportsApi.md#InvalidateOperation) | **Delete** /imports/images/{operation_id} | Invalidate operation ID so it can be garbage collected
[**ListImportDockerfiles**](ImportsApi.md#ListImportDockerfiles) | **Get** /imports/images/{operation_id}/dockerfile | List uploaded dockerfiles
[**ListImportImageConfigs**](ImportsApi.md#ListImportImageConfigs) | **Get** /imports/images/{operation_id}/image_config | List uploaded image configs
[**ListImportImageManifests**](ImportsApi.md#ListImportImageManifests) | **Get** /imports/images/{operation_id}/manifest | List uploaded image manifests
[**ListImportPackages**](ImportsApi.md#ListImportPackages) | **Get** /imports/images/{operation_id}/packages | List uploaded package manifests
[**ListImportParentManifests**](ImportsApi.md#ListImportParentManifests) | **Get** /imports/images/{operation_id}/parent_manifest | List uploaded parent manifests (manifest lists for a tag)
[**ListOperations**](ImportsApi.md#ListOperations) | **Get** /imports/images | Lists in-progress imports



## CreateOperation

> ImageImportOperation CreateOperation(ctx, )

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


## GetOperation

> ImageImportOperation GetOperation(ctx, operationId)

Get detail on a single import

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string**|  | 

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


## ImportImageConfig

> ImageImportContentResponse ImportImageConfig(ctx, operationId, contents)

Import a docker or OCI image config to associate with the image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string**|  | 
**contents** | **interface{}**|  | 

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


## ImportImageDockerfile

> ImageImportContentResponse ImportImageDockerfile(ctx, operationId, contents)

Begin the import of an image analyzed by Syft into the system

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string**|  | 
**contents** | **string**|  | 

### Return type

[**ImageImportContentResponse**](ImageImportContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain; utf-8
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportImageManifest

> ImageImportContentResponse ImportImageManifest(ctx, operationId, contents)

Import a docker or OCI distribution manifest to associate with the image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string**|  | 
**contents** | **interface{}**|  | 

### Return type

[**ImageImportContentResponse**](ImageImportContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.oci.image.manifest.v1+json, application/vnd.docker.distribution.manifest.v2+json, application/vnd.docker.distribution.manifest.v1+json, application/vnd.docker.distribution.manifest.v1+prettyjws
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportImagePackages

> ImageImportContentResponse ImportImagePackages(ctx, operationId, sbom)

Begin the import of an image analyzed by Syft into the system

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string**|  | 
**sbom** | [**ImagePackageManifest**](ImagePackageManifest.md)|  | 

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


## ImportImageParentManifest

> ImageImportContentResponse ImportImageParentManifest(ctx, operationId, contents)

Import a docker or OCI distribution manifest list to associate with the image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string**|  | 
**contents** | **interface{}**|  | 

### Return type

[**ImageImportContentResponse**](ImageImportContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.docker.distribution.manifest.list.v2+json, application/vnd.oci.image.index.v1+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvalidateOperation

> ImageImportOperation InvalidateOperation(ctx, operationId)

Invalidate operation ID so it can be garbage collected

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string**|  | 

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


## ListImportDockerfiles

> []string ListImportDockerfiles(ctx, operationId)

List uploaded dockerfiles

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


## ListImportImageConfigs

> []string ListImportImageConfigs(ctx, operationId)

List uploaded image configs

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


## ListImportImageManifests

> []string ListImportImageManifests(ctx, operationId)

List uploaded image manifests

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


## ListImportPackages

> []string ListImportPackages(ctx, operationId)

List uploaded package manifests

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


## ListImportParentManifests

> []string ListImportParentManifests(ctx, operationId)

List uploaded parent manifests (manifest lists for a tag)

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


## ListOperations

> []ImageImportOperation ListOperations(ctx, )

Lists in-progress imports

### Required Parameters

This endpoint does not need any parameter.

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

