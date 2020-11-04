/*
 * Clubhouse API
 *
 * Clubhouse API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// CreateStoryCommentParams Request parameters for creating a Comment on a Clubhouse Story.
type CreateStoryCommentParams struct {
	// The Member ID of the Comment's author. Defaults to the user identified by the API token.
	AuthorId string `json:"author_id,omitempty"`
	// Defaults to the time/date the comment is created, but can be set to reflect another date.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// This field can be set to another unique ID. In the case that the comment has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId string `json:"external_id,omitempty"`
	// The comment text.
	Text string `json:"text"`
	// Defaults to the time/date the comment is last updated, but can be set to reflect another date.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
