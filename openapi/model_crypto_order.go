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
// CryptoOrder struct for CryptoOrder
type CryptoOrder struct {
	AccountId string `json:"account_id,omitempty"`
	CurrencyPairId string `json:"currency_pair_id,omitempty"`
	Price string `json:"price,omitempty"`
	Quantity string `json:"quantity,omitempty"`
	RefId string `json:"ref_id,omitempty"`
	Side Side `json:"side,omitempty"`
	TimeInForce TimeInForce `json:"time_in_force,omitempty"`
	Type ExecutionType `json:"type,omitempty"`
}