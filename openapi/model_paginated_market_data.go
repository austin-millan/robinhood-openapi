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
// PaginatedMarketData struct for PaginatedMarketData
type PaginatedMarketData struct {
	Count string `json:"count,omitempty"`
	Next string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results []MarketData `json:"results,omitempty"`
}
