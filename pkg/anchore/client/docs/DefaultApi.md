# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCredential**](DefaultApi.md#AddCredential) | **Post** /user/credentials | add/replace credential
[**AddImage**](DefaultApi.md#AddImage) | **Post** /images | Submit a new image for analysis by the engine
[**AddPolicy**](DefaultApi.md#AddPolicy) | **Post** /policies | Add a new policy
[**AddRepository**](DefaultApi.md#AddRepository) | **Post** /repositories | Add repository to watch
[**AddSubscription**](DefaultApi.md#AddSubscription) | **Post** /subscriptions | Add a subscription of a specific type
[**ArchiveImageAnalysis**](DefaultApi.md#ArchiveImageAnalysis) | **Post** /archives/images | 
[**CreateAccount**](DefaultApi.md#CreateAccount) | **Post** /accounts | Create a new user. Only avaialble to admin user.
[**CreateAnalysisArchiveRule**](DefaultApi.md#CreateAnalysisArchiveRule) | **Post** /archives/rules | 
[**CreateRegistry**](DefaultApi.md#CreateRegistry) | **Post** /registries | Add a new registry
[**CreateUser**](DefaultApi.md#CreateUser) | **Post** /accounts/{accountname}/users | Create a new user
[**CreateUserCredential**](DefaultApi.md#CreateUserCredential) | **Post** /accounts/{accountname}/users/{username}/credentials | add/replace credential
[**DeleteAccount**](DefaultApi.md#DeleteAccount) | **Delete** /accounts/{accountname} | Delete the specified account, only allowed if the account is in the disabled state. All users will be deleted along with the account and all resources will be garbage collected
[**DeleteAnalysisArchiveRule**](DefaultApi.md#DeleteAnalysisArchiveRule) | **Delete** /archives/rules/{ruleId} | 
[**DeleteArchivedAnalysis**](DefaultApi.md#DeleteArchivedAnalysis) | **Delete** /archives/images/{imageDigest} | 
[**DeleteEvent**](DefaultApi.md#DeleteEvent) | **Delete** /events/{eventId} | Delete Event
[**DeleteEvents**](DefaultApi.md#DeleteEvents) | **Delete** /events | Delete Events
[**DeleteFeed**](DefaultApi.md#DeleteFeed) | **Delete** /system/feeds/{feed} | 
[**DeleteFeedGroup**](DefaultApi.md#DeleteFeedGroup) | **Delete** /system/feeds/{feed}/{group} | 
[**DeleteImage**](DefaultApi.md#DeleteImage) | **Delete** /images/{imageDigest} | Delete an image analysis
[**DeleteImageByImageId**](DefaultApi.md#DeleteImageByImageId) | **Delete** /images/by_id/{imageId} | Delete image by docker imageId
[**DeletePolicy**](DefaultApi.md#DeletePolicy) | **Delete** /policies/{policyId} | Delete policy
[**DeleteRegistry**](DefaultApi.md#DeleteRegistry) | **Delete** /registries/{registry} | Delete a registry configuration
[**DeleteService**](DefaultApi.md#DeleteService) | **Delete** /system/services/{servicename}/{hostid} | Delete the service config
[**DeleteSubscription**](DefaultApi.md#DeleteSubscription) | **Delete** /subscriptions/{subscriptionId} | Delete subscriptions of a specific type
[**DeleteUser**](DefaultApi.md#DeleteUser) | **Delete** /accounts/{accountname}/users/{username} | Delete a specific user credential by username of the credential. Cannot be the credential used to authenticate the request.
[**DeleteUserCredential**](DefaultApi.md#DeleteUserCredential) | **Delete** /accounts/{accountname}/users/{username}/credentials | Delete a credential by type
[**DescribeErrorCodes**](DefaultApi.md#DescribeErrorCodes) | **Get** /system/error_codes | Describe anchore engine error codes.
[**DescribePolicy**](DefaultApi.md#DescribePolicy) | **Get** /system/policy_spec | Describe the policy language spec implemented by this service.
[**GetAccount**](DefaultApi.md#GetAccount) | **Get** /accounts/{accountname} | Get info about an user. Only available to admin user. Uses the main user Id, not a username.
[**GetAccountUser**](DefaultApi.md#GetAccountUser) | **Get** /accounts/{accountname}/users/{username} | Get a specific user in the specified account
[**GetAnalysisArchiveRule**](DefaultApi.md#GetAnalysisArchiveRule) | **Get** /archives/rules/{ruleId} | 
[**GetArchivedAnalysis**](DefaultApi.md#GetArchivedAnalysis) | **Get** /archives/images/{imageDigest} | 
[**GetCredentials**](DefaultApi.md#GetCredentials) | **Get** /user/credentials | Get current credential summary
[**GetEvent**](DefaultApi.md#GetEvent) | **Get** /events/{eventId} | Get Event
[**GetImage**](DefaultApi.md#GetImage) | **Get** /images/{imageDigest} | Get image metadata
[**GetImageByImageId**](DefaultApi.md#GetImageByImageId) | **Get** /images/by_id/{imageId} | Lookup image by docker imageId
[**GetImageContentByType**](DefaultApi.md#GetImageContentByType) | **Get** /images/{imageDigest}/content/{ctype} | Get the content of an image by type
[**GetImageContentByTypeFiles**](DefaultApi.md#GetImageContentByTypeFiles) | **Get** /images/{imageDigest}/content/files | Get the content of an image by type files
[**GetImageContentByTypeImageId**](DefaultApi.md#GetImageContentByTypeImageId) | **Get** /images/by_id/{imageId}/content/{ctype} | Get the content of an image by type
[**GetImageContentByTypeImageIdFiles**](DefaultApi.md#GetImageContentByTypeImageIdFiles) | **Get** /images/by_id/{imageId}/content/files | Get the content of an image by type files
[**GetImageContentByTypeImageIdJavapackage**](DefaultApi.md#GetImageContentByTypeImageIdJavapackage) | **Get** /images/by_id/{imageId}/content/java | Get the content of an image by type java
[**GetImageContentByTypeJavapackage**](DefaultApi.md#GetImageContentByTypeJavapackage) | **Get** /images/{imageDigest}/content/java | Get the content of an image by type java
[**GetImageMetadataByType**](DefaultApi.md#GetImageMetadataByType) | **Get** /images/{imageDigest}/metadata/{mtype} | Get the metadata of an image by type
[**GetImagePolicyCheck**](DefaultApi.md#GetImagePolicyCheck) | **Get** /images/{imageDigest}/check | Check policy evaluation status for image
[**GetImagePolicyCheckByImageId**](DefaultApi.md#GetImagePolicyCheckByImageId) | **Get** /images/by_id/{imageId}/check | Check policy evaluation status for image
[**GetImageVulnerabilitiesByType**](DefaultApi.md#GetImageVulnerabilitiesByType) | **Get** /images/{imageDigest}/vuln/{vtype} | Get vulnerabilities by type
[**GetImageVulnerabilitiesByTypeImageId**](DefaultApi.md#GetImageVulnerabilitiesByTypeImageId) | **Get** /images/by_id/{imageId}/vuln/{vtype} | Get vulnerabilities by type
[**GetImageVulnerabilityTypes**](DefaultApi.md#GetImageVulnerabilityTypes) | **Get** /images/{imageDigest}/vuln | Get vulnerability types
[**GetImageVulnerabilityTypesByImageId**](DefaultApi.md#GetImageVulnerabilityTypesByImageId) | **Get** /images/by_id/{imageId}/vuln | Get vulnerability types
[**GetOauthToken**](DefaultApi.md#GetOauthToken) | **Post** /oauth/token | 
[**GetPolicy**](DefaultApi.md#GetPolicy) | **Get** /policies/{policyId} | Get specific policy
[**GetRegistry**](DefaultApi.md#GetRegistry) | **Get** /registries/{registry} | Get a specific registry configuration
[**GetServiceDetail**](DefaultApi.md#GetServiceDetail) | **Get** /system | System status
[**GetServicesByName**](DefaultApi.md#GetServicesByName) | **Get** /system/services/{servicename} | Get a service configuration and state
[**GetServicesByNameAndHost**](DefaultApi.md#GetServicesByNameAndHost) | **Get** /system/services/{servicename}/{hostid} | Get service config for a specific host
[**GetStatus**](DefaultApi.md#GetStatus) | **Get** /status | Service status
[**GetSubscription**](DefaultApi.md#GetSubscription) | **Get** /subscriptions/{subscriptionId} | Get a specific subscription set
[**GetSystemFeeds**](DefaultApi.md#GetSystemFeeds) | **Get** /system/feeds | list feeds operations and information
[**GetUser**](DefaultApi.md#GetUser) | **Get** /user | List authenticated user info
[**GetUsersAccount**](DefaultApi.md#GetUsersAccount) | **Get** /account | List the account for the authenticated user
[**HealthNoop**](DefaultApi.md#HealthNoop) | **Get** /health | 
[**ImportImageArchive**](DefaultApi.md#ImportImageArchive) | **Post** /import/images | Import an anchore image tar.gz archive file.
[**ListAccounts**](DefaultApi.md#ListAccounts) | **Get** /accounts | List user summaries. Only available to the system admin user.
[**ListAnalysisArchive**](DefaultApi.md#ListAnalysisArchive) | **Get** /archives/images | 
[**ListAnalysisArchiveRules**](DefaultApi.md#ListAnalysisArchiveRules) | **Get** /archives/rules | 
[**ListArchives**](DefaultApi.md#ListArchives) | **Get** /archives | 
[**ListEventTypes**](DefaultApi.md#ListEventTypes) | **Get** /event_types | List Event Types
[**ListEvents**](DefaultApi.md#ListEvents) | **Get** /events | List Events
[**ListFileContentSearchResults**](DefaultApi.md#ListFileContentSearchResults) | **Get** /images/{imageDigest}/artifacts/file_content_search | Return a list of analyzer artifacts of the specified type
[**ListImageContent**](DefaultApi.md#ListImageContent) | **Get** /images/{imageDigest}/content | List image content types
[**ListImageContentByImageid**](DefaultApi.md#ListImageContentByImageid) | **Get** /images/by_id/{imageId}/content | List image content types
[**ListImageMetadata**](DefaultApi.md#ListImageMetadata) | **Get** /images/{imageDigest}/metadata | List image metadata types
[**ListImages**](DefaultApi.md#ListImages) | **Get** /images | List all visible images
[**ListImagetags**](DefaultApi.md#ListImagetags) | **Get** /summaries/imagetags | List all visible image digests and tags
[**ListPolicies**](DefaultApi.md#ListPolicies) | **Get** /policies | List policies
[**ListRegistries**](DefaultApi.md#ListRegistries) | **Get** /registries | List configured registries
[**ListRetrievedFiles**](DefaultApi.md#ListRetrievedFiles) | **Get** /images/{imageDigest}/artifacts/retrieved_files | Return a list of analyzer artifacts of the specified type
[**ListSecretSearchResults**](DefaultApi.md#ListSecretSearchResults) | **Get** /images/{imageDigest}/artifacts/secret_search | Return a list of analyzer artifacts of the specified type
[**ListServices**](DefaultApi.md#ListServices) | **Get** /system/services | List system services
[**ListSubscriptions**](DefaultApi.md#ListSubscriptions) | **Get** /subscriptions | List all subscriptions
[**ListUserCredentials**](DefaultApi.md#ListUserCredentials) | **Get** /accounts/{accountname}/users/{username}/credentials | Get current credential summary
[**ListUsers**](DefaultApi.md#ListUsers) | **Get** /accounts/{accountname}/users | List accounts for the user
[**Ping**](DefaultApi.md#Ping) | **Get** / | 
[**PostSystemFeeds**](DefaultApi.md#PostSystemFeeds) | **Post** /system/feeds | trigger feeds operations
[**QueryImagesByPackage**](DefaultApi.md#QueryImagesByPackage) | **Get** /query/images/by_package | List of images containing given package
[**QueryImagesByVulnerability**](DefaultApi.md#QueryImagesByVulnerability) | **Get** /query/images/by_vulnerability | List images vulnerable to the specific vulnerability ID.
[**QueryVulnerabilities**](DefaultApi.md#QueryVulnerabilities) | **Get** /query/vulnerabilities | Listing information about given vulnerability
[**ToggleFeedEnabled**](DefaultApi.md#ToggleFeedEnabled) | **Put** /system/feeds/{feed} | 
[**ToggleGroupEnabled**](DefaultApi.md#ToggleGroupEnabled) | **Put** /system/feeds/{feed}/{group} | 
[**UpdateAccountState**](DefaultApi.md#UpdateAccountState) | **Put** /accounts/{accountname}/state | Update the state of an account to either enabled or disabled. For deletion use the DELETE route
[**UpdatePolicy**](DefaultApi.md#UpdatePolicy) | **Put** /policies/{policyId} | Update policy
[**UpdateRegistry**](DefaultApi.md#UpdateRegistry) | **Put** /registries/{registry} | Update/replace a registry configuration
[**UpdateSubscription**](DefaultApi.md#UpdateSubscription) | **Put** /subscriptions/{subscriptionId} | Update an existing and specific subscription
[**VersionNoop**](DefaultApi.md#VersionNoop) | **Get** /version | 



## AddCredential

> User AddCredential(ctx, credential)
add/replace credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credential** | [**AccessCredential**](AccessCredential.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddImage

> []AnchoreImage AddImage(ctx, image, optional)
Submit a new image for analysis by the engine

Creates a new analysis task that is executed asynchronously

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**image** | [**ImageAnalysisRequest**](ImageAnalysisRequest.md)|  | 
 **optional** | ***AddImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Override any existing entry in the system | 
 **autosubscribe** | **optional.Bool**| Instruct engine to automatically begin watching the added tag for updates from registry | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]AnchoreImage**](AnchoreImage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddPolicy

> PolicyBundleRecord AddPolicy(ctx, bundle, optional)
Add a new policy

Adds a new policy bundle to the system

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundle** | [**PolicyBundle**](PolicyBundle.md)|  | 
 **optional** | ***AddPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddPolicyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**PolicyBundleRecord**](PolicyBundleRecord.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRepository

> []Subscription AddRepository(ctx, repository, optional)
Add repository to watch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repository** | **string**| full repository to add e.g. docker.io/library/alpine | 
 **optional** | ***AddRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddRepositoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autosubscribe** | **optional.Bool**| flag to enable/disable auto tag_update activation when new images from a repo are added | 
 **lookuptag** | **optional.String**| use specified existing tag to perform repo scan (default is &#39;latest&#39;) | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]Subscription**](Subscription.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSubscription

> []Subscription AddSubscription(ctx, subscription, optional)
Add a subscription of a specific type

Create a new subscription to watch a tag and get notifications of changes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscription** | [**SubscriptionRequest**](SubscriptionRequest.md)|  | 
 **optional** | ***AddSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]Subscription**](Subscription.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchiveImageAnalysis

> []AnalysisArchiveAddResult ArchiveImageAnalysis(ctx, imageReferences)


### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageReferences** | [**[]string**](string.md)|  | 

### Return type

[**[]AnalysisArchiveAddResult**](AnalysisArchiveAddResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccount

> Account CreateAccount(ctx, account)
Create a new user. Only avaialble to admin user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | [**AccountCreationRequest**](AccountCreationRequest.md)|  | 

### Return type

[**Account**](Account.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAnalysisArchiveRule

> AnalysisArchiveTransitionRule CreateAnalysisArchiveRule(ctx, rule)


### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rule** | [**AnalysisArchiveTransitionRule**](AnalysisArchiveTransitionRule.md)|  | 

### Return type

[**AnalysisArchiveTransitionRule**](AnalysisArchiveTransitionRule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRegistry

> []RegistryConfiguration CreateRegistry(ctx, registrydata, optional)
Add a new registry

Adds a new registry to the system

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrydata** | [**RegistryConfigurationRequest**](RegistryConfigurationRequest.md)|  | 
 **optional** | ***CreateRegistryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateRegistryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validate** | **optional.Bool**| flag to determine whether or not to validate registry/credential at registry add time | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]RegistryConfiguration**](RegistryConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> User CreateUser(ctx, accountname, user)
Create a new user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string**|  | 
**user** | [**UserCreationRequest**](UserCreationRequest.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserCredential

> User CreateUserCredential(ctx, accountname, username, credential)
add/replace credential

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string**|  | 
**username** | **string**|  | 
**credential** | [**AccessCredential**](AccessCredential.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccount

> DeleteAccount(ctx, accountname)
Delete the specified account, only allowed if the account is in the disabled state. All users will be deleted along with the account and all resources will be garbage collected

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnalysisArchiveRule

> AnalysisArchiveTransitionRule DeleteAnalysisArchiveRule(ctx, ruleId)


### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string**|  | 

### Return type

[**AnalysisArchiveTransitionRule**](AnalysisArchiveTransitionRule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArchivedAnalysis

> ArchivedAnalysis DeleteArchivedAnalysis(ctx, imageDigest, optional)


Performs a synchronous archive deletion

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**|  | 
 **optional** | ***DeleteArchivedAnalysisOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteArchivedAnalysisOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**|  | 

### Return type

[**ArchivedAnalysis**](ArchivedAnalysis.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEvent

> DeleteEvent(ctx, eventId, optional)
Delete Event

Delete an event by its event ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| Event ID of the event to be deleted | 
 **optional** | ***DeleteEventOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteEventOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEvents

> []string DeleteEvents(ctx, optional)
Delete Events

Delete all or a subset of events filtered using the optional query parameters

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **before** | **optional.String**| Delete events that occurred before the timestamp | 
 **since** | **optional.String**| Delete events that occurred after the timestamp | 
 **level** | **optional.String**| Delete events that match the level - INFO or ERROR | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImage

> DeleteImage(ctx, imageDigest, optional)
Delete an image analysis

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**|  | 
 **optional** | ***DeleteImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**|  | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImageByImageId

> DeleteImageByImageId(ctx, imageId, optional)
Delete image by docker imageId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string**|  | 
 **optional** | ***DeleteImageByImageIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteImageByImageIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**|  | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicy

> DeletePolicy(ctx, policyId, optional)
Delete policy

Delete the specified policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**|  | 
 **optional** | ***DeletePolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeletePolicyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRegistry

> DeleteRegistry(ctx, registry, optional)
Delete a registry configuration

Delete a registry configuration record from the system. Does not remove any images.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registry** | **string**|  | 
 **optional** | ***DeleteRegistryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteRegistryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscription

> DeleteSubscription(ctx, subscriptionId, optional)
Delete subscriptions of a specific type

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string**|  | 
 **optional** | ***DeleteSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, accountname, username)
Delete a specific user credential by username of the credential. Cannot be the credential used to authenticate the request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string**|  | 
**username** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserCredential

> DeleteUserCredential(ctx, accountname, username, credentialType)
Delete a credential by type

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string**|  | 
**username** | **string**|  | 
**credentialType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> Account GetAccount(ctx, accountname)
Get info about an user. Only available to admin user. Uses the main user Id, not a username.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string**|  | 

### Return type

[**Account**](Account.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountUser

> User GetAccountUser(ctx, accountname, username)
Get a specific user in the specified account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string**|  | 
**username** | **string**|  | 

### Return type

[**User**](User.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalysisArchiveRule

> AnalysisArchiveTransitionRule GetAnalysisArchiveRule(ctx, ruleId)


### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string**|  | 

### Return type

[**AnalysisArchiveTransitionRule**](AnalysisArchiveTransitionRule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArchivedAnalysis

> ArchivedAnalysis GetArchivedAnalysis(ctx, imageDigest)


Returns the archive metadata record identifying the image and tags for the analysis in the archive.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**| The image digest to identify the image analysis | 

### Return type

[**ArchivedAnalysis**](ArchivedAnalysis.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredentials

> []AccessCredential GetCredentials(ctx, )
Get current credential summary

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]AccessCredential**](AccessCredential.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvent

> EventResponse GetEvent(ctx, eventId, optional)
Get Event

Lookup an event by its event ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| Event ID of the event for lookup | 
 **optional** | ***GetEventOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEventOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImage

> []AnchoreImage GetImage(ctx, imageDigest, optional)
Get image metadata

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageDigest** | **string**|  | 
 **optional** | ***GetImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]AnchoreImage**](AnchoreImage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageByImageId

> []AnchoreImage GetImageByImageId(ctx, imageId, optional)
Lookup image by docker imageId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string**|  | 
 **optional** | ***GetImageByImageIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImageByImageIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]AnchoreImage**](AnchoreImage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOauthToken

> TokenResponse GetOauthToken(ctx, )


Request a jwt token for subsequent operations, this request is authenticated with normal HTTP auth

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicy

> []PolicyBundleRecord GetPolicy(ctx, policyId, optional)
Get specific policy

Get the policy bundle content

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**|  | 
 **optional** | ***GetPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPolicyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **detail** | **optional.Bool**| Include policy bundle detail in the form of the full bundle content for each entry | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]PolicyBundleRecord**](PolicyBundleRecord.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistry

> RegistryConfiguration GetRegistry(ctx, registry, optional)
Get a specific registry configuration

Get information on a specific registry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registry** | **string**|  | 
 **optional** | ***GetRegistryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRegistryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**RegistryConfiguration**](RegistryConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscription

> []Subscription GetSubscription(ctx, subscriptionId, optional)
Get a specific subscription set

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string**|  | 
 **optional** | ***GetSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]Subscription**](Subscription.md)

### Authorization

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> User GetUser(ctx, )
List authenticated user info

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersAccount

> Account GetUsersAccount(ctx, )
List the account for the authenticated user

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**Account**](Account.md)

### Authorization

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportImageArchive

> []AnchoreImage ImportImageArchive(ctx, archiveFile)
Import an anchore image tar.gz archive file.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**archiveFile** | ***os.File*****os.File**| anchore image tar archive. | 

### Return type

[**[]AnchoreImage**](AnchoreImage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccounts

> []Account ListAccounts(ctx, optional)
List user summaries. Only available to the system admin user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAccountsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **optional.String**| Filter accounts by state | 

### Return type

[**[]Account**](Account.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAnalysisArchive

> []ArchivedAnalysis ListAnalysisArchive(ctx, )


### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ArchivedAnalysis**](ArchivedAnalysis.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAnalysisArchiveRules

> []AnalysisArchiveTransitionRule ListAnalysisArchiveRules(ctx, optional)


### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAnalysisArchiveRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAnalysisArchiveRulesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemGlobal** | **optional.Bool**| If true include system global rules (owned by admin) even for non-admin users. Defaults to true if not set. Can be set to false to exclude globals | 

### Return type

[**[]AnalysisArchiveTransitionRule**](AnalysisArchiveTransitionRule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArchives

> ArchiveSummary ListArchives(ctx, )


### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ArchiveSummary**](ArchiveSummary.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEventTypes

> []EventCategory ListEventTypes(ctx, )
List Event Types

Returns list of event types in the category hierarchy

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]EventCategory**](EventCategory.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvents

> EventsList ListEvents(ctx, optional)
List Events

Returns a paginated list of events in the descending order of their occurrence. Optional query parameters may be used for filtering results

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceServicename** | **optional.String**| Filter events by the originating service | 
 **sourceHostid** | **optional.String**| Filter events by the originating host ID | 
 **eventType** | **optional.String**| Filter events by a prefix match on the event type (e.g. \&quot;user.image.\&quot;) | 
 **resourceType** | **optional.String**| Filter events by the type of resource - tag, imageDigest, repository etc | 
 **resourceId** | **optional.String**| Filter events by the id of the resource | 
 **level** | **optional.String**| Filter events by the level - INFO or ERROR | 
 **since** | **optional.String**| Return events that occurred after the timestamp | 
 **before** | **optional.String**| Return events that occurred before the timestamp | 
 **page** | **optional.Int32**| Pagination controls - return the nth page of results. Defaults to first page if left empty | [default to 1]
 **limit** | **optional.Int32**| Number of events in the result set. Defaults to 100 if left empty | [default to 100]
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**EventsList**](EventsList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImages

> []AnchoreImage ListImages(ctx, optional)
List all visible images

List all images visible to the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListImagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListImagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **history** | **optional.Bool**| Include image history in the response | 
 **fulltag** | **optional.String**| Full docker-pull string to filter results by (e.g. docker.io/library/nginx:latest, or myhost.com:5000/testimages:v1.1.1) | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 
 **imageToGet** | [**optional.Interface of ImageFilter**](ImageFilter.md)|  | 

### Return type

[**[]AnchoreImage**](AnchoreImage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]AnchoreImageTagSummary**](AnchoreImageTagSummary.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicies

> []PolicyBundleRecord ListPolicies(ctx, optional)
List policies

List all saved policy bundles

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **detail** | **optional.Bool**| Include policy bundle detail in the form of the full bundle content for each entry | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]PolicyBundleRecord**](PolicyBundleRecord.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRegistries

> []RegistryConfiguration ListRegistries(ctx, optional)
List configured registries

List all configured registries the system can/will watch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListRegistriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRegistriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]RegistryConfiguration**](RegistryConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscriptions

> []Subscription ListSubscriptions(ctx, optional)
List all subscriptions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSubscriptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionKey** | **optional.String**| filter only subscriptions matching key | 
 **subscriptionType** | **optional.String**| filter only subscriptions matching type | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]Subscription**](Subscription.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserCredentials

> []AccessCredential ListUserCredentials(ctx, accountname, username)
Get current credential summary

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string**|  | 
**username** | **string**|  | 

### Return type

[**[]AccessCredential**](AccessCredential.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> []User ListUsers(ctx, accountname)
List accounts for the user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string**|  | 

### Return type

[**[]User**](User.md)

### Authorization

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

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

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountState

> AccountStatus UpdateAccountState(ctx, accountname, desiredState)
Update the state of an account to either enabled or disabled. For deletion use the DELETE route

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string**|  | 
**desiredState** | [**AccountStatus**](AccountStatus.md)|  | 

### Return type

[**AccountStatus**](AccountStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicy

> []PolicyBundleRecord UpdatePolicy(ctx, policyId, bundle, optional)
Update policy

Update/replace and existing policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**|  | 
**bundle** | [**PolicyBundleRecord**](PolicyBundleRecord.md)|  | 
 **optional** | ***UpdatePolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdatePolicyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **active** | **optional.Bool**| Mark policy as active | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]PolicyBundleRecord**](PolicyBundleRecord.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRegistry

> []RegistryConfiguration UpdateRegistry(ctx, registry, registrydata, optional)
Update/replace a registry configuration

Replaces an existing registry record with the given record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registry** | **string**|  | 
**registrydata** | [**RegistryConfigurationRequest**](RegistryConfigurationRequest.md)|  | 
 **optional** | ***UpdateRegistryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateRegistryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validate** | **optional.Bool**| flag to determine whether or not to validate registry/credential at registry update time | 
 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]RegistryConfiguration**](RegistryConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscription

> []Subscription UpdateSubscription(ctx, subscriptionId, subscription, optional)
Update an existing and specific subscription

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string**|  | 
**subscription** | [**SubscriptionUpdate**](SubscriptionUpdate.md)|  | 
 **optional** | ***UpdateSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAnchoreAccount** | **optional.String**| An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]Subscription**](Subscription.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
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

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

