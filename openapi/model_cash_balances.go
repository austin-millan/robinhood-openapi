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
import (
	"time"
)
// CashBalances struct for CashBalances
type CashBalances struct {
	BuyingPower float64 `json:"buying_power,omitempty"`
	Cash float64 `json:"cash,omitempty"`
	CashAvailableForWithdrawal float64 `json:"cash_available_for_withdrawal,omitempty"`
	CashHeldForOrders float64 `json:"cash_held_for_orders,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UnclearedDeposits float64 `json:"uncleared_deposits,omitempty"`
	UnsettledFunds float64 `json:"unsettled_funds,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
