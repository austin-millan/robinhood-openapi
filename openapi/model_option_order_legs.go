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
// OptionOrderLegs struct for OptionOrderLegs
type OptionOrderLegs struct {
	Executions []OptionOrderExecutions `json:"executions,omitempty"`
	Id string `json:"id,omitempty"`
	Option string `json:"option,omitempty"`
	PositionEffect string `json:"position_effect,omitempty"`
	RatioQuantity float32 `json:"ratio_quantity,omitempty"`
	Side string `json:"side,omitempty"`
}
