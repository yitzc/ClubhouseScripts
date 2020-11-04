/*
 * Clubhouse API
 *
 * Clubhouse API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// MaxSearchResultsExceededError Error returned when total maximum supported results have been reached.
type MaxSearchResultsExceededError struct {
	// The name for this type of error, `maximum-results-exceeded`
	Error string `json:"error"`
	// The maximum number of search results supported, `1000`
	MaximumResults int64 `json:"maximum-results"`
	// An explanatory message: \"A maximum of 1000 search results are supported.\"
	Message string `json:"message"`
}
