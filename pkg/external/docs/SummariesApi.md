# \SummariesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListImagetags**](SummariesApi.md#ListImagetags) | **Get** /summaries/imagetags | List all visible image digests and tags



## ListImagetags

> []AnchoreImageTagSummary ListImagetags(ctx, optional)

List all visible image digests and tags

List all image tags visible to the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListImagetagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListImagetagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageStatus** | [**optional.Interface of []string**](string.md)| Filter images in one or more states such as active, deleting. Defaults to active images only if unspecified | [default to [&quot;active&quot;]]
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]AnchoreImageTagSummary**](AnchoreImageTagSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

