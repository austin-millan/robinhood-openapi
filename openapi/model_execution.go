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
// Execution struct for Execution
type Execution struct {
	Id string `json:"id,omitempty"`
	Price float64 `json:"price,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
	SettlementDate string `json:"settlement_date,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}