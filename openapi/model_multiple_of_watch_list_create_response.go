/*
 * Robinhood API
 *
 * Robinhood API Documentation
 *
 * API version: 3.0.1
 * Contact: austin.millan@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// MultipleOfWatchListCreateResponse struct for MultipleOfWatchListCreateResponse
type MultipleOfWatchListCreateResponse struct {
	Count string `json:"count,omitempty"`
	Next string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results []WatchListCreateResponse `json:"results,omitempty"`
}
