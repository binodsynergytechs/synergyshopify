package goshopify

import (
	"fmt"
	"time"
)

const discountCodeBasePath = "price_rules/%d/discount_codes"

// DiscountCodeRepository is an interface for interfacing with the discount endpoints of the Shopify API.
// See: https://help.shopify.com/en/api/reference/discounts/PriceRuleDiscountCode
type DiscountCodeRepository interface {
	CreateDiscountCode(int64, PriceRuleDiscountCode) (*PriceRuleDiscountCode, error)
	UpdateDiscountCode(int64, PriceRuleDiscountCode) (*PriceRuleDiscountCode, error)
	ListDiscountCode(int64) ([]PriceRuleDiscountCode, error)
	GetDiscountCode(int64, int64) (*PriceRuleDiscountCode, error)
	DeleteDiscountCode(int64, int64) error
}

// DiscountCodeClient handles communication with the discount code
// related methods of the Shopify API.
type DiscountCodeClient struct {
	client *Client
}

// PriceRuleDiscountCode represents a Shopify Discount Code
type PriceRuleDiscountCode struct {
	ID          int64      `json:"id,omitempty"`
	PriceRuleID int64      `json:"price_rule_id,omitempty"`
	Code        string     `json:"code,omitempty"`
	UsageCount  int        `json:"usage_count,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	Errors      []string   `json:"// FIXME: Field Not Available In Model 23/04"`
}

// FIXME: Field Not Available In Model 23/04
// TODO: Latest Field Available In Model 23/04

// MultipleDiscountCodesResponse is the result from the discount_codes.json endpoint
type MultipleDiscountCodesResponse struct {
	DiscountCodes []PriceRuleDiscountCode `json:"discount_codes"`
}

// SingleDiscountCodeResponse represents the result from the discount_codes/X.json endpoint
type SingleDiscountCodeResponse struct {
	PriceRuleDiscountCode *PriceRuleDiscountCode `json:"discount_code"`
}

// Create a discount code
func (dcc *DiscountCodeClient) CreateDiscountCode(priceRuleID int64, dc PriceRuleDiscountCode) (*PriceRuleDiscountCode, error) {
	path := fmt.Sprintf(discountCodeBasePath+".json", priceRuleID)
	wrappedData := SingleDiscountCodeResponse{PriceRuleDiscountCode: &dc}
	resource := new(SingleDiscountCodeResponse)
	err := dcc.client.Post(path, wrappedData, resource)
	return resource.PriceRuleDiscountCode, err
}

// Update an existing discount code
func (dcc *DiscountCodeClient) UpdateDiscountCode(priceRuleID int64, dc PriceRuleDiscountCode) (*PriceRuleDiscountCode, error) {
	path := fmt.Sprintf(discountCodeBasePath+"/%d.json", priceRuleID, dc.ID)
	wrappedData := SingleDiscountCodeResponse{PriceRuleDiscountCode: &dc}
	resource := new(SingleDiscountCodeResponse)
	err := dcc.client.Put(path, wrappedData, resource)
	return resource.PriceRuleDiscountCode, err
}

// List of discount codes
func (dcc *DiscountCodeClient) ListDiscountCode(priceRuleID int64) ([]PriceRuleDiscountCode, error) {
	path := fmt.Sprintf(discountCodeBasePath+".json", priceRuleID)
	resource := new(MultipleDiscountCodesResponse)
	err := dcc.client.Get(path, resource, nil)
	return resource.DiscountCodes, err
}

// Get a single discount code
func (dcc *DiscountCodeClient) GetDiscountCode(priceRuleID int64, discountCodeID int64) (*PriceRuleDiscountCode, error) {
	path := fmt.Sprintf(discountCodeBasePath+"/%d.json", priceRuleID, discountCodeID)
	resource := new(SingleDiscountCodeResponse)
	err := dcc.client.Get(path, resource, nil)
	return resource.PriceRuleDiscountCode, err
}

// Delete a discount code
func (dcc *DiscountCodeClient) DeleteDiscountCode(priceRuleID int64, discountCodeID int64) error {
	return dcc.client.Delete(fmt.Sprintf(discountCodeBasePath+"/%d.json", priceRuleID, discountCodeID))
}
