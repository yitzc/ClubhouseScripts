/*
 * Clubhouse API
 *
 * Clubhouse API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateProject struct for UpdateProject
type UpdateProject struct {
	// The Project abbreviation used in Story summaries. Should be kept to 3 characters at most.
	Abbreviation string `json:"abbreviation,omitempty"`
	// A true/false boolean indicating whether the Story is in archived state.
	Archived bool `json:"archived,omitempty"`
	// The color that represents the Project in the UI.
	Color string `json:"color,omitempty"`
	// The number of days before the thermometer appears in the Story summary.
	DaysToThermometer int64 `json:"days_to_thermometer,omitempty"`
	// The Project's description.
	Description string `json:"description,omitempty"`
	// An array of UUIDs for any Members you want to add as Followers.
	FollowerIds []string `json:"follower_ids,omitempty"`
	// The Project's name.
	Name string `json:"name,omitempty"`
	// Configuration to enable or disable thermometers in the Story summary.
	ShowThermometer bool `json:"show_thermometer,omitempty"`
	// The ID of the team the project belongs to.
	TeamId int64 `json:"team_id,omitempty"`
}