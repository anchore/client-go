/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.15
 * Contact: nurmi@anchore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client
// ImageImportManifestMetadata struct for ImageImportManifestMetadata
type ImageImportManifestMetadata struct {
	Digest string `json:"digest,omitempty"`
	LocalImageId string `json:"local_image_id,omitempty"`
	Layers []ImageImportLayerMetadata `json:"layers,omitempty"`
	Platform ImageImportPlatform `json:"platform,omitempty"`
}