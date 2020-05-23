/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// OrderState the model 'OrderState'
type OrderState string

// List of OrderState
const (
	QUEUED OrderState = "queued"
	UNCONFIRMED OrderState = "unconfirmed"
	CONFIRMED OrderState = "confirmed"
	PARTIALLY_FILLED OrderState = "partially_filled"
	FILLED OrderState = "filled"
	REJECTED OrderState = "rejected"
	CANCELED OrderState = "canceled"
	FAILED OrderState = "failed"
)
