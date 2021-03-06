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
// PaginatedInstrumentData struct for PaginatedInstrumentData
type PaginatedInstrumentData struct {
	Count string `json:"count,omitempty"`
	Next string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results []InstrumentData `json:"results,omitempty"`
}
