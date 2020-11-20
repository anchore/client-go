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
// BaseNotificationData Every notification has a payload, which follows this basic structure
type BaseNotificationData struct {
	NotificationUser string `json:"notification_user,omitempty"`
	NotificationUserEmail string `json:"notification_user_email,omitempty"`
	NotificationType string `json:"notification_type,omitempty"`
}
