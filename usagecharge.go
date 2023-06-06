package goshopify

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const usageChargesPath = "usage_charges"

// UsageChargeRepository is an interface for interacting with the UsageCharge endpoints of the Shopify API.
// See https://help.shopify.com/en/api/reference/billing/usagecharge#endpoints
type UsageChargeRepository interface {
	CreateUsageCharge(int64, UsageCharge) (*UsageCharge, error)
	GetUsageCharge(int64, int64, interface{}) (*UsageCharge, error)
	ListUsageCharge(int64, interface{}) ([]UsageCharge, error)
}

// UsageChargeClient handles communication with the
// UsageCharge related methods of the Shopify API.
type UsageChargeClient struct {
	client *Client
}

// UsageCharge represents a Shopify UsageCharge.
type UsageCharge struct {
	BalanceRemaining             *decimal.Decimal `json:"balance_remaining,omitempty"` //FIXME: Not Available In Shopify
	BalanceUsed                  *decimal.Decimal `json:"balance_used,omitempty"`      //FIXME: Not Available In Shopify
	BillingOn                    *time.Time       `json:"billing_on,omitempty"`        //FIXME: Not Available In Shopify
	RiskLevel                    *decimal.Decimal `json:"risk_level,omitempty"`        //FIXME: Not Available In Shopify
	CreatedAt                    *time.Time       `json:"created_at,omitempty"`
	Description                  string           `json:"description,omitempty"`
	ID                           int64            `json:"id,omitempty"`
	Price                        *decimal.Decimal `json:"price,omitempty"`
	UpdatedAt                    *time.Time       `json:"Updated_at,omitempty"`            //TODO: New Field Added In Shopify
	RecurringApplicationChargeID int64            `json:"recurring_application_charge_id"` //TODO: New Field Added In Shopify
	Currency                     string           `json:"currency"`                        //TODO: New Field Added In Shopify
}

func (u *UsageCharge) UnmarshalJSON(data []byte) error {
	// This is a workaround for the API returning BillingOn date in the format of "YYYY-MM-DD"
	// https://help.shopify.com/en/api/reference/billing/usagecharge#endpoints
	// For a longer explanation of the hack check:
	// http://choly.ca/post/go-json-marshalling/
	type alias UsageCharge
	aux := &struct {
		BillingOn *string `json:"billing_on"`
		*alias
	}{alias: (*alias)(u)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if err := parse(&u.BillingOn, aux.BillingOn); err != nil {
		return err
	}
	return nil
}

// UsageChargeResource represents the result from the /admin/recurring_application_charges/X/usage_charges/X.json endpoints
type UsageChargeResource struct {
	Charge *UsageCharge `json:"usage_charge"`
}

// UsageChargesResource represents the result from the
// admin/recurring_application_charges/X/usage_charges.json endpoint.
type UsageChargesResource struct {
	Charges []UsageCharge `json:"usage_charges"`
}

// Create creates new usage charge given a recurring charge. *required fields: price and description
func (uc *UsageChargeClient) CreateUsageCharge(chargeID int64, usageCharge UsageCharge) (
	*UsageCharge, error) {

	path := fmt.Sprintf("%s/%d/%s.json", recurringApplicationChargesBasePath, chargeID, usageChargesPath)
	wrappedData := UsageChargeResource{Charge: &usageCharge}
	resource := &UsageChargeResource{}
	err := uc.client.Post(path, wrappedData, resource)
	return resource.Charge, err
}

// Get gets individual usage charge.
func (uc *UsageChargeClient) GetUsageCharge(chargeID int64, usageChargeID int64, options interface{}) (
	*UsageCharge, error) {

	path := fmt.Sprintf("%s/%d/%s/%d.json", recurringApplicationChargesBasePath, chargeID, usageChargesPath, usageChargeID)
	resource := &UsageChargeResource{}
	err := uc.client.Get(path, resource, options)
	return resource.Charge, err
}

// List gets all usage charges associated with the recurring charge.
func (uc *UsageChargeClient) ListUsageCharge(chargeID int64, options interface{}) (
	[]UsageCharge, error) {

	path := fmt.Sprintf("%s/%d/%s.json", recurringApplicationChargesBasePath, chargeID, usageChargesPath)
	resource := &UsageChargesResource{}
	err := uc.client.Get(path, resource, options)
	return resource.Charges, err
}
