/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// MultipleOfPosition struct for MultipleOfPosition
type MultipleOfPosition struct {
	Count int32 `json:"count,omitempty"`
	Next string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results []Position `json:"results,omitempty"`
}
