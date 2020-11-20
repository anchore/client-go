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
// TagUpdateNotificationPayload struct for TagUpdateNotificationPayload
type TagUpdateNotificationPayload struct {
	UserId string `json:"userId,omitempty"`
	SubscriptionKey string `json:"subscription_key,omitempty"`
	SubscriptionType string `json:"subscription_type,omitempty"`
	NotificationId string `json:"notificationId,omitempty"`
	// A list containing the current image digest
	CurrEval []map[string]interface{} `json:"curr_eval,omitempty"`
	// A list containing the previous image digests
	LastEval []map[string]interface{} `json:"last_eval,omitempty"`
	// List of Corresponding Image Annotations
	Annotations *map[string]interface{} `json:"annotations,omitempty"`
}
