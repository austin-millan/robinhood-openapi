/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// MultipleOfWatchListsData struct for MultipleOfWatchListsData
type MultipleOfWatchListsData struct {
	Count int32 `json:"count,omitempty"`
	Next string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results []WatchListsData `json:"results,omitempty"`
}
