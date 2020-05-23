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
// InvestmentProfile struct for InvestmentProfile
type InvestmentProfile struct {
	AnnualIncome string `json:"annual_income,omitempty"`
	InvestmentExperience InvestmentExperience `json:"investment_experience,omitempty"`
	InvestmentObjective InvestmentObjective `json:"investment_objective,omitempty"`
	LiquidNetWorth string `json:"liquid_net_worth,omitempty"`
	LiquidityNeeds LiquidityNeeds `json:"liquidity_needs,omitempty"`
	RiskTolerance RiskTolerance `json:"risk_tolerance,omitempty"`
	SourceOfFunds SourceOfFunds `json:"source_of_funds,omitempty"`
	SuitabilityVerified bool `json:"suitability_verified,omitempty"`
	TaxBracket TaxBracket `json:"tax_bracket,omitempty"`
	TimeHorizon TimeHorizon `json:"time_horizon,omitempty"`
	TotalNetWorth TotalNetWorth `json:"total_net_worth,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	User string `json:"user,omitempty"`
}
