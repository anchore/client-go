# ImageImportManifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**ImportContentDigests**](ImportContentDigests.md) |  | [optional] 
**Tags** | **[]string** |  | [optional] 
**Digest** | **string** |  | [optional] 
**ParentDigest** | **string** | The digest of the images&#39;s manifest-list parent if it was accessed from a multi-arch tag where the tag pointed to a manifest-list. This allows preservation of that relationship in the data | [optional] 
**LocalImageId** | **string** | An \&quot;imageId\&quot; as used by Docker if available | [optional] 
**OperationUuid** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


