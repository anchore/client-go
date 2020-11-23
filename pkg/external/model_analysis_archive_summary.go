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
// AnalysisArchiveSummary A summarization of the analysis archive, including size, counts, etc. This archive stores image analysis only, never the actual image content or layers.
type AnalysisArchiveSummary struct {
	// The number of unique images (digests) in the archive
	TotalImageCount int32 `json:"total_image_count,omitempty"`
	// The number of tag records (registry/repo:tag pull strings) in the archive. This may include repeated tags but will always have a unique tag->digest mapping per record.
	TotalTagCount int32 `json:"total_tag_count,omitempty"`
	// The total sum of all the bytes stored to the backing storage. Accounts for anchore-applied compression, but not compression by the underlying storage system.
	TotalDataBytes int32 `json:"total_data_bytes,omitempty"`
	// The timestamp of the most recent archived image
	LastUpdated time.Time `json:"last_updated,omitempty"`
}
