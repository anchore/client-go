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
// FeedMetadata Metadata on a single feed based on what the engine finds from querying the endpoints
type FeedMetadata struct {
	// name of the feed
	Name string `json:"name,omitempty"`
	// Date the metadata record was created in engine (first seen on source)
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Date the metadata was last updated
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Groups []FeedGroupMetadata `json:"groups,omitempty"`
	LastFullSync time.Time `json:"last_full_sync,omitempty"`
}
