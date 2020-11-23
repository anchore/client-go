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
// VulnerableImage A record of an image vulnerable to some known vulnerability. Includes vulnerable package information
type VulnerableImage struct {
	Image ImageReference `json:"image,omitempty"`
	AffectedPackages []VulnerablePackageReference `json:"affected_packages,omitempty"`
}
