# \SystemApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFeed**](SystemApi.md#DeleteFeed) | **Delete** /system/feeds/{feed} | 
[**DeleteFeedGroup**](SystemApi.md#DeleteFeedGroup) | **Delete** /system/feeds/{feed}/{group} | 
[**DeleteService**](SystemApi.md#DeleteService) | **Delete** /system/services/{servicename}/{hostid} | Delete the service config
[**DescribeErrorCodes**](SystemApi.md#DescribeErrorCodes) | **Get** /system/error_codes | Describe anchore engine error codes.
[**DescribePolicy**](SystemApi.md#DescribePolicy) | **Get** /system/policy_spec | Describe the policy language spec implemented by this service.
[**GetServiceDetail**](SystemApi.md#GetServiceDetail) | **Get** /system | System status
[**GetServicesByName**](SystemApi.md#GetServicesByName) | **Get** /system/services/{servicename} | Get a service configuration and state
[**GetServicesByNameAndHost**](SystemApi.md#GetServicesByNameAndHost) | **Get** /system/services/{servicename}/{hostid} | Get service config for a specific host
[**GetStatus**](SystemApi.md#GetStatus) | **Get** /status | Service status
[**GetSystemFeeds**](SystemApi.md#GetSystemFeeds) | **Get** /system/feeds | list feeds operations and information
[**ListServices**](SystemApi.md#ListServices) | **Get** /system/services | List system services
[**PostSystemFeeds**](SystemApi.md#PostSystemFeeds) | **Post** /system/feeds | trigger feeds operations
[**TestWebhook**](SystemApi.md#TestWebhook) | **Post** /system/webhooks/{webhook_type}/test | Adds the capabilities to test a webhook delivery for the given notification type
[**ToggleFeedEnabled**](SystemApi.md#ToggleFeedEnabled) | **Put** /system/feeds/{feed} | 
[**ToggleGroupEnabled**](SystemApi.md#ToggleGroupEnabled) | **Put** /system/feeds/{feed}/{group} | 



## DeleteFeed

> DeleteFeed(ctx, feed)



Delete the groups and data for the feed and disable the feed itself

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feed** | **string**|  | 

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


## DeleteFeedGroup

> DeleteFeedGroup(ctx, feed, group)



Delete the group data and disable the group itself

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feed** | **string**|  | 
**group** | **string**|  | 

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


## DeleteService

> DeleteService(ctx, servicename, hostid)

Delete the service config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicename** | **string**|  | 
**hostid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeErrorCodes

> []AnchoreErrorCode DescribeErrorCodes(ctx, )

Describe anchore engine error codes.

Describe anchore engine error codes.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]AnchoreErrorCode**](AnchoreErrorCode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribePolicy

> []GateSpec DescribePolicy(ctx, )

Describe the policy language spec implemented by this service.

Get the policy language spec for this service

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]GateSpec**](GateSpec.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceDetail

> SystemStatusResponse GetServiceDetail(ctx, )

System status

Get the system status including queue lengths

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SystemStatusResponse**](SystemStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicesByName

> []Service GetServicesByName(ctx, servicename)

Get a service configuration and state

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicename** | **string**|  | 

### Return type

[**[]Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicesByNameAndHost

> []Service GetServicesByNameAndHost(ctx, servicename, hostid)

Get service config for a specific host

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servicename** | **string**|  | 
**hostid** | **string**|  | 

### Return type

[**[]Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus

> StatusResponse GetStatus(ctx, )

Service status

Get the API service status

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemFeeds

> []FeedMetadata GetSystemFeeds(ctx, )

list feeds operations and information

Return a list of feed and their groups along with update and record count information. This data reflects the state of the policy engine, not the upstream feed service itself.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]FeedMetadata**](FeedMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServices

> []Service ListServices(ctx, )

List system services

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSystemFeeds

> []FeedSyncResult PostSystemFeeds(ctx, optional)

trigger feeds operations

Execute a synchronous feed sync operation. The response will block until complete, then return the result summary.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostSystemFeedsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostSystemFeedsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flush** | **optional.Bool**| instruct system to flush existing data feeds records from anchore-engine | 
 **sync** | **optional.Bool**| instruct system to re-sync data feeds | 

### Return type

[**[]FeedSyncResult**](FeedSyncResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestWebhook

> TestWebhook(ctx, webhookType, optional)

Adds the capabilities to test a webhook delivery for the given notification type

Loads the Webhook configuration for webhook_type, and sends the notification out as a test

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookType** | **string**| The Webhook Type that we should test | 
 **optional** | ***TestWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TestWebhookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationType** | **optional.String**| What kind of Notification to send | [default to tag_update]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleFeedEnabled

> FeedMetadata ToggleFeedEnabled(ctx, feed, enabled)



Disable the feed so that it does not sync on subsequent sync operations

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feed** | **string**|  | 
**enabled** | **bool**|  | 

### Return type

[**FeedMetadata**](FeedMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleGroupEnabled

> []FeedMetadata ToggleGroupEnabled(ctx, feed, group, enabled)



Disable a specific group within a feed to not sync

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feed** | **string**|  | 
**group** | **string**|  | 
**enabled** | **bool**|  | 

### Return type

[**[]FeedMetadata**](FeedMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

