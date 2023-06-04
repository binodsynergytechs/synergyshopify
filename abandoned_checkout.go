package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const abandonedCheckoutsBasePath = "checkouts"

// AbandonedCheckoutRepository is an interface for interfacing with the abandonedCheckouts endpoints
// of the Shopify API.
// See: https://shopify.dev/docs/api/admin-rest/latest/resources/abandoned-checkouts
type AbandonedCheckoutRepository interface {
	ListAbandonedCheckout(interface{}) ([]AbandonedCheckout, error)
}

// AbandonedCheckoutClient handles communication with the checkout related methods of
// the Shopify API.
type AbandonedCheckoutClient struct {
	client *Client
}

// Represents the result from the checkouts.json endpoint
type AbandonedCheckoutsResource struct {
	AbandonedCheckouts []AbandonedCheckout `json:"checkouts,omitempty"`
}

// AbandonedCheckout represents a Shopify abandoned checkout
type AbandonedCheckout struct {
	ID                       int64                `json:"id,omitempty"`
	Token                    string               `json:"token,omitempty"` //FIXME: Field Not Available Or Deprecated In Latest Shopify Model 23/04
	CartToken                string               `json:"cart_token,omitempty"`
	Email                    string               `json:"email,omitempty"`
	Gateway                  string               `json:"gateway,omitempty"`
	BuyerAcceptsMarketing    bool                 `json:"buyer_accepts_marketing,omitempty"`
	CreatedAt                *time.Time           `json:"created_at,omitempty"`
	UpdatedAt                *time.Time           `json:"updated_at,omitempty"`
	LandingSite              string               `json:"landing_site,omitempty"`
	Note                     string               `json:"note,omitempty"`
	NoteAttributes           []NoteAttribute      `json:"note_attributes,omitempty"` //FIXME: Field Not Available Or Deprecated In Latest Shopify Model 23/04
	ReferringSite            string               `json:"referring_site,omitempty"`
	ShippingLines            []ShippingLines      `json:"shipping_lines,omitempty"`
	TaxesIncluded            bool                 `json:"taxes_included,omitempty"`
	TotalWeight              int                  `json:"total_weight,omitempty"`
	Currency                 string               `json:"currency,omitempty"`
	CompletedAt              *time.Time           `json:"completed_at,omitempty"`
	ClosedAt                 *time.Time           `json:"closed_at,omitempty"`
	UserID                   int64                `json:"user_id,omitempty"`
	SourceIdentifier         string               `json:"source_identifier,omitempty"` //FIXME: Field Not Available Or Deprecated In Latest Shopify Model 23/04
	SourceUrl                string               `json:"source_url,omitempty"`        //FIXME: Field Not Available Or Deprecated In Latest Shopify Model 23/04
	DeviceID                 int64                `json:"device_id,omitempty"`
	Phone                    string               `json:"phone,omitempty"`
	CustomerLocale           string               `json:"customer_locale,omitempty"`
	Name                     string               `json:"name,omitempty"`   //FIXME: Field Not Available Or Deprecated In Latest Shopify Model 23/04
	Source                   string               `json:"source,omitempty"` //FIXME: Field Not Available Or Deprecated In Latest Shopify Model 23/04
	AbandonedCheckoutUrl     string               `json:"abandoned_checkout_url,omitempty"`
	DiscountCodes            []DiscountCode       `json:"discount_codes,omitempty"`
	TaxLines                 []TaxLine            `json:"tax_lines,omitempty"` //FIXME: Field Not Available Or Deprecated In Latest Shopify Model 23/04
	SourceName               string               `json:"source_name,omitempty"`
	PresentmentCurrency      string               `json:"presentment_currency,omitempty"`
	BuyerAcceptsSmsMarketing bool                 `json:"buyer_accepts_sms_marketing,omitempty"`
	SmsMarketingPhone        string               `json:"sms_marketing_phone,omitempty"`
	TotalDiscounts           *decimal.Decimal     `json:"total_discounts,omitempty"`
	LineItems                []LineItem           `json:"line_items,omitempty"` //TODO: Field Available In Latest Shopify Model 23/04
	LocationID               int64                `json:"location_id"`          //TODO: Field Available In Latest Shopify Model 23/04
	TaxIncluded              bool                 `json:"tax_included"`         //TODO: Field Available In Latest Shopify Model 23/04
	TotalTax                 *decimal.Decimal     `json:"total_tax"`            //TODO: Field Available In Latest Shopify Model 23/04
	TotalLineItemsPrice      *decimal.Decimal     `json:"total_line_items_price,omitempty"`
	TotalPrice               *decimal.Decimal     `json:"total_price,omitempty"`
	SubtotalPrice            *decimal.Decimal     `json:"subtotal_price,omitempty"`
	TotalDuties              string               `json:"total_duties,omitempty"`
	BillingAddress           *Address             `json:"billing_address,omitempty"`
	ShippingAddress          *Address             `json:"shipping_address,omitempty"`
	Customer                 *Customer            `json:"customer,omitempty"`
	SmsMarketingConsent      *SmsMarketingConsent `json:"sms_marketing_consent,omitempty"` //FIXME: Field Not Available Or Deprecated In Latest Shopify Model 23/04
	AdminGraphqlApiID        string               `json:"admin_graphql_api_id,omitempty"`  //FIXME: Field Not Available Or Deprecated In Latest Shopify Model 23/04
	DefaultAddress           *CustomerAddress     `json:"default_address,omitempty"`       //FIXME: Field Not Available Or Deprecated In Latest Shopify Model 23/04
}

type SmsMarketingConsent struct {
	State                string     `json:"state,omitempty"`
	OptInLevel           string     `json:"opt_in_level,omitempty"`
	ConsentUpdatedAt     *time.Time `json:"consent_updated_at,omitempty"`
	ConsentCollectedFrom string     `json:"consent_collected_from,omitempty"`
}

// Get abandoned checkout list
func (acc *AbandonedCheckoutClient) ListAbandonedCheckout(options interface{}) ([]AbandonedCheckout, error) {
	// Construct the path for the abandoned checkouts endpoint using fmt.Sprintf
	path := fmt.Sprintf("%s.json", abandonedCheckoutsBasePath)

	// Create a new instance of the resource struct to hold the response
	resource := new(AbandonedCheckoutsResource)

	// Use the HTTP client to send a GET request to the specified path
	err := acc.client.Get(path, resource, options)

	// Return the abandoned checkouts and any potential error from the request
	return resource.AbandonedCheckouts, err
}
