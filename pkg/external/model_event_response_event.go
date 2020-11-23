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
// EventResponseEvent struct for EventResponseEvent
type EventResponseEvent struct {
	Source EventResponseEventSource `json:"source,omitempty"`
	Resource EventResponseEventResource `json:"resource,omitempty"`
	Type string `json:"type,omitempty"`
	Category string `json:"category,omitempty"`
	Level string `json:"level,omitempty"`
	Message string `json:"message,omitempty"`
	Details interface{} `json:"details,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}
