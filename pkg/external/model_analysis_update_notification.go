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
// AnalysisUpdateNotification struct for AnalysisUpdateNotification
type AnalysisUpdateNotification struct {
	QueueId string `json:"queueId,omitempty"`
	UserId string `json:"userId,omitempty"`
	DataId string `json:"dataId,omitempty"`
	CreatedAt int32 `json:"created_at,omitempty"`
	LastUpdated int32 `json:"last_updated,omitempty"`
	RecordStateKey string `json:"record_state_key,omitempty"`
	RecordStateVal *string `json:"record_state_val,omitempty"`
	Tries int32 `json:"tries,omitempty"`
	MaxTries int32 `json:"max_tries,omitempty"`
	Data AnalysisUpdateNotificationData `json:"data,omitempty"`
}