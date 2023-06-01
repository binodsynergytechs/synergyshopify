package goshopify

import (
	"time"
)

// ShopService is an interface for interfacing with the shop endpoint of the
// Shopify API.
// See: https://help.shopify.com/api/reference/shop
type ShopService interface {
	Get(options interface{}) (*Shop, error)
}

// ShopServiceOp handles communication with the shop related methods of the
// Shopify API.
type ShopServiceOp struct {
	client *Client
}

// Shop represents a Shopify shop
type Shop struct {
	ID                                   int64      `json:"id"`
	Name                                 string     `json:"name"`
	ShopOwner                            string     `json:"shop_owner"`
	Email                                string     `json:"email"`
	CustomerEmail                        string     `json:"customer_email"`
	CreatedAt                            *time.Time `json:"created_at"`
	CookieConsentLevel                   string     `json:"cookie_consent_level"` //TODO: Newly Added Field
	UpdatedAt                            *time.Time `json:"updated_at"`
	Address1                             string     `json:"address1"`
	Address2                             string     `json:"address2"`
	City                                 string     `json:"city"`
	Country                              string     `json:"country"`
	CountryCode                          string     `json:"country_code"`
	CountryName                          string     `json:"country_name"`
	CountryTaxes                         bool       `json:"county_taxes"`
	Currency                             string     `json:"currency"`
	Domain                               string     `json:"domain"`
	EnabledPresentmentCurrencies         []string   `json:"enabled_presentment_currencies"`    //TODO: Newly Added Field
	EligibleForCardReaderGiveaway        bool       `json:"eligible_for_card_reader_giveaway"` //TODO: Newly Added Field
	Latitude                             float64    `json:"latitude"`
	Longitude                            float64    `json:"longitude"`
	Phone                                string     `json:"phone"`
	Province                             string     `json:"province"`
	ProvinceCode                         string     `json:"province_code"`
	Zip                                  string     `json:"zip"`
	MoneyFormat                          string     `json:"money_format"`
	MoneyWithCurrencyFormat              string     `json:"money_with_currency_format"`
	WeightUnit                           string     `json:"weight_unit"`
	MyShopifyDomain                      string     `json:"myshopify_domain"`
	PlanName                             string     `json:"plan_name"`
	PlanDisplayName                      string     `json:"plan_display_name"`
	PasswordEnabled                      bool       `json:"password_enabled"`
	PrimaryLocale                        string     `json:"primary_locale"`
	PrimaryLocationId                    int64      `json:"primary_location_id"` //FIXME: deprecated field
	Timezone                             string     `json:"timezone"`
	IanaTimezone                         string     `json:"iana_timezone"`
	ForceSSL                             bool       `json:"force_ssl"` //FIXME: deprecated field
	TaxShipping                          bool       `json:"tax_shipping"`
	TaxesIncluded                        bool       `json:"taxes_included"`
	HasStorefront                        bool       `json:"has_storefront"`
	HasDiscounts                         bool       `json:"has_discounts"`
	HasGiftCards                         bool       `json:"has_gift_cards"`
	SetupRequire                         bool       `json:"setup_required"`
	CheckoutAPISupported                 bool       `json:"checkout_api_supported"`
	Source                               string     `json:"source"`
	GoogleAppsDomain                     string     `json:"google_apps_domain"`
	GoogleAppsLoginEnabled               bool       `json:"google_apps_login_enabled"`
	MoneyInEmailsFormat                  string     `json:"money_in_emails_format"`
	MoneyWithCurrencyInEmailsFormat      string     `json:"money_with_currency_in_emails_format"`
	EligibleForPayments                  bool       `json:"eligible_for_payments"`
	RequiresExtraPaymentsAgreement       bool       `json:"requires_extra_payments_agreement"`
	PreLaunchEnabled                     bool       `json:"pre_launch_enabled"`
	TransactionalSmsDisabled             bool       `json:"transactional_sms_disabled"`                //TODO: Newly Added Field
	MarketingSmsConsentEnabledAtCheckout bool       `json:"marketing_sms_consent_enabled_at_checkout"` //TODO: Newly Added Field
}

// Represents the result from the admin/shop.json endpoint
type ShopResource struct {
	Shop *Shop `json:"shop"`
}

// Get shop
func (s *ShopServiceOp) Get(options interface{}) (*Shop, error) {
	resource := new(ShopResource)
	err := s.client.Get("shop.json", resource, options)
	return resource.Shop, err
}
