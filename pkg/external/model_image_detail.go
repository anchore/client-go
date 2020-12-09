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
// ImageDetail A metadata detail record for a specific image. Multiple detail records may map a single catalog image.
type ImageDetail struct {
	CreatedAt float32 `json:"created_at,omitempty"`
	LastUpdated float32 `json:"last_updated,omitempty"`
	// Full docker-pullable tag string referencing the image
	Fulltag string `json:"fulltag,omitempty"`
	// Full docker-pullable digest string including the registry url and repository necessary get the image
	Fulldigest string `json:"fulldigest,omitempty"`
	UserId string `json:"userId,omitempty"`
	ImageId string `json:"imageId,omitempty"`
	Registry string `json:"registry,omitempty"`
	Repo string `json:"repo,omitempty"`
	Dockerfile *string `json:"dockerfile,omitempty"`
	// The parent Anchore Image record to which this detail maps
	ImageDigest string `json:"imageDigest,omitempty"`
}
