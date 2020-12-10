# RegistryDigestSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pullstring** | **string** | A digest-based pullstring (e.g. docker.io/nginx@sha256:123abc) | 
**Tag** | **string** | A valid docker tag reference (e.g. docker.io/nginx:latest) that will be associated with the image but not used to pull the image. | 
**CreationTimestampOverride** | [**time.Time**](time.Time.md) | Optional override of the image creation time to support proper tag history construction in cases of out-of-order analysis compared to registry history for the tag | [optional] 
**Dockerfile** | **string** | Base64 encoded content of the dockerfile used to build the image, if available. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


