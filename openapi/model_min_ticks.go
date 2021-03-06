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
// MinTicks struct for MinTicks
type MinTicks struct {
	AboveTick string `json:"above_tick,omitempty"`
	BelowTick string `json:"below_tick,omitempty"`
	CutoffPrice string `json:"cutoff_price,omitempty"`
}
