/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// Order struct for Order
type Order struct {
	Account string `json:"account,omitempty"`
	AveragePrice float64 `json:"average_price,omitempty"`
	Cancel string `json:"cancel,omitempty"`
	ClientId string `json:"client_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	CumulativeQuantity float64 `json:"cumulative_quantity,omitempty"`
	Executions []Execution `json:"executions,omitempty"`
	ExtendedHours bool `json:"extended_hours,omitempty"`
	Fees float64 `json:"fees,omitempty"`
	Id string `json:"id,omitempty"`
	Instrument string `json:"instrument,omitempty"`
	LastTransactionAt time.Time `json:"last_transaction_at,omitempty"`
	OverrideDayTradeChecks bool `json:"override_day_trade_checks,omitempty"`
	OverrideDtbpChecks bool `json:"override_dtbp_checks,omitempty"`
	Position string `json:"position,omitempty"`
	Price float64 `json:"price,omitempty"`
	Quantity int32 `json:"quantity,omitempty"`
	RejectReason string `json:"reject_reason,omitempty"`
	Side string `json:"side,omitempty"`
	State OrderState `json:"state,omitempty"`
	StopPrice float64 `json:"stop_price,omitempty"`
	Symbol string `json:"symbol,omitempty"`
	TimeInForce TimeInForce `json:"time_in_force,omitempty"`
	Trigger string `json:"trigger,omitempty"`
	Type ExecutionType `json:"type,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Url string `json:"url,omitempty"`
}
