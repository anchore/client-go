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
// VulnUpdateNotificationData struct for VulnUpdateNotificationData
type VulnUpdateNotificationData struct {
	NotificationUser string `json:"notification_user,omitempty"`
	NotificationUserEmail string `json:"notification_user_email,omitempty"`
	NotificationType string `json:"notification_type,omitempty"`
	NotificationPayload VulnUpdateNotificationPayload `json:"notification_payload,omitempty"`
}
