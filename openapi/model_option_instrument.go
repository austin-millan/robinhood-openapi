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
// OptionInstrument struct for OptionInstrument
type OptionInstrument struct {
	ChainId string `json:"chain_id,omitempty"`
	ChainSymbol string `json:"chain_symbol,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	ExpirationDate string `json:"expiration_date,omitempty"`
	Id string `json:"id,omitempty"`
	IssueDate string `json:"issue_date,omitempty"`
	MinTicks MinTicks `json:"min_ticks,omitempty"`
	RhsTradability string `json:"rhs_tradability,omitempty"`
	State string `json:"state,omitempty"`
	StrikePrice string `json:"strike_price,omitempty"`
	Tradability string `json:"tradability,omitempty"`
	Type string `json:"type,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Url string `json:"url,omitempty"`
}