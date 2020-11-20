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
// AnalysisUpdateNotificationPayloadAllOf struct for AnalysisUpdateNotificationPayloadAllOf
type AnalysisUpdateNotificationPayloadAllOf struct {
	CurrEval AnalysisUpdateEval `json:"curr_eval,omitempty"`
	LastEval AnalysisUpdateEval `json:"last_eval,omitempty"`
	// List of Corresponding Image Annotations
	Annotations *map[string]interface{} `json:"annotations,omitempty"`
}
