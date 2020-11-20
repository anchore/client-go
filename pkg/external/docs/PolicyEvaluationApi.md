# \PolicyEvaluationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImagePolicyCheck**](PolicyEvaluationApi.md#GetImagePolicyCheck) | **Get** /images/{imageDigest}/check | Check policy evaluation status for image
[**GetImagePolicyCheckByImageId**](PolicyEvaluationApi.md#GetImagePolicyCheckByImageId) | **Get** /images/by_id/{imageId}/check | Check policy evaluation status for image



## GetImagePolicyCheck

> []map[string]interface{} GetImagePolicyCheck(ctx, imageDigest, tag, optional)
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

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImagePolicyCheckByImageId

> []map[string]interface{} GetImagePolicyCheckByImageId(ctx, imageId, tag, optional)
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

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

