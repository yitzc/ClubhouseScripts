/*
 * Clubhouse API
 *
 * Clubhouse API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// StoryContentsTask struct for StoryContentsTask
type StoryContentsTask struct {
	// True/false boolean indicating whether the Task has been completed.
	Complete bool `json:"complete,omitempty"`
	// Full text of the Task.
	Description string `json:"description"`
	// This field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId *string `json:"external_id,omitempty"`
	// An array of UUIDs of the Owners of this Task.
	OwnerIds []string `json:"owner_ids,omitempty"`
	// The number corresponding to the Task's position within a list of Tasks on a Story.
	Position int64 `json:"position,omitempty"`
}
