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
// ImageImportLayerMetadata struct for ImageImportLayerMetadata
type ImageImportLayerMetadata struct {
	Digest string `json:"digest,omitempty"`
	Size int32 `json:"size,omitempty"`
	Location string `json:"location,omitempty"`
}
