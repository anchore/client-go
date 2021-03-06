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
// AnalysisUpdateEval Evaluation Results for an entity (current or last)
type AnalysisUpdateEval struct {
	AnalysisStatus string `json:"analysis_status,omitempty"`
	Annotations interface{} `json:"annotations,omitempty"`
	ImageDigest string `json:"image_digest,omitempty"`
}
