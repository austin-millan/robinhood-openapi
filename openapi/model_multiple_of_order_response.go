/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// MultipleOfOrderResponse struct for MultipleOfOrderResponse
type MultipleOfOrderResponse struct {
	Count int32 `json:"count,omitempty"`
	Next string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results []OrderResponse `json:"results,omitempty"`
}
