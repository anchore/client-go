/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.16
 * Contact: nurmi@anchore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package external
// ImportDistribution struct for ImportDistribution
type ImportDistribution struct {
	Name string `json:"name"`
	Version string `json:"version"`
	IdLike string `json:"idLike"`
}
