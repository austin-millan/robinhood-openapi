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
// Order struct for Order
type Order struct {
	Account string `json:"account,omitempty"`
	Action OrderAction `json:"action,omitempty"`
	AveragePrice string `json:"average_price,omitempty"`
	CancelUrl string `json:"cancel_url,omitempty"`
	ClientId string `json:"client_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	CumulativeQuantity string `json:"cumulative_quantity,omitempty"`
	DollarBasedAmount string `json:"dollar_based_amount,omitempty"`
	Executions []Execution `json:"executions,omitempty"`
	TotalNotional OrderTotalNotional `json:"total_notional,omitempty"`
	ExecutedNotional OrderTotalNotional `json:"executed_notional,omitempty"`
	ExtendedHours bool `json:"extended_hours,omitempty"`
	Fees string `json:"fees,omitempty"`
	Id string `json:"id,omitempty"`
	InvestmentScheduleId string `json:"investment_schedule_id,omitempty"`
	Instrument string `json:"instrument,omitempty"`
	LastTrailPrice string `json:"last_trail_price,omitempty"`
	LastTrailPriceUpdatedAt string `json:"last_trail_price_updated_at,omitempty"`
	LastTransactionAt string `json:"last_transaction_at,omitempty"`
	OverrideDayTradeChecks bool `json:"override_day_trade_checks,omitempty"`
	OverrideDtbpChecks bool `json:"override_dtbp_checks,omitempty"`
	Position string `json:"position,omitempty"`
	Price string `json:"price,omitempty"`
	Quantity string `json:"quantity,omitempty"`
	RefId string `json:"ref_id,omitempty"`
	RejectReason string `json:"reject_reason,omitempty"`
	Side Side `json:"side,omitempty"`
	State OrderState `json:"state,omitempty"`
	StopPrice string `json:"stop_price,omitempty"`
	StopTriggeredAt string `json:"stop_triggered_at,omitempty"`
	Symbol string `json:"symbol,omitempty"`
	TimeInForce TimeInForce `json:"time_in_force,omitempty"`
	Trigger Trigger `json:"trigger,omitempty"`
	Type ExecutionType `json:"type,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Url string `json:"url,omitempty"`
}
