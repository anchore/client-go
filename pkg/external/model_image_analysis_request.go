/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.15
 * Contact: nurmi@anchore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package external
import (
	"time"
)
// ImageAnalysisRequest A request to add an image to be watched and analyzed by the engine. Optionally include the dockerfile content. Either source, digest or tag must be present.
type ImageAnalysisRequest struct {
	// Base64 encoded content of the dockerfile for the image, if available. Deprecated in favor of the 'source' field.
	Dockerfile string `json:"dockerfile,omitempty"`
	// A digest string for an image, maybe a pull string or just a digest. e.g. nginx@sha256:123 or sha256:abc123. If a pull string, it must have same regisry/repo as the tag field. Deprecated in favor of the 'source' field
	Digest string `json:"digest,omitempty"`
	// Full pullable tag reference for image. e.g. docker.io/nginx:latest. Deprecated in favor of the 'source' field
	Tag string `json:"tag,omitempty"`
	// Optional override of the image creation time, only honored when both tag and digest are also supplied  e.g. 2018-10-17T18:14:00Z. Deprecated in favor of the 'source' field
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Optional. The type of image this is adding, defaults to \"docker\". This can be ommitted until multiple image types are supported.
	ImageType string `json:"image_type,omitempty"`
	// Annotations to be associated with the added image in key/value form
	Annotations interface{} `json:"annotations,omitempty"`
	Source ImageSource `json:"source,omitempty"`
}
