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
	ID                              int64      `json:"id"`
	Name                            string     `json:"name"`
	ShopOwner                       string     `json:"shop_owner"`
	Email                           string     `json:"email"`
	CustomerEmail                   string     `json:"customer_email"`
	CreatedAt                       *time.Time `json:"created_at"`
	UpdatedAt                       *time.Time `json:"updated_at"`
	Address1                        string     `json:"address1"`
	Address2                        string     `json:"address2"`
	City                            string     `json:"city"`
	Country                         string     `json:"country"`
	CountryCode                     string     `json:"country_code"`
	CountryName                     string     `json:"country_name"`
	Currency                        string     `json:"currency"`
	Domain                          string     `json:"domain"`
	Latitude                        float64    `json:"latitude"`
	Longitude                       float64    `json:"longitude"`
	Phone                           string     `json:"phone"`
	Province                        string     `json:"province"`
	ProvinceCode                    string     `json:"province_code"`
	Zip                             string     `json:"zip"`
	MoneyFormat                     string     `json:"money_format"`
	MoneyWithCurrencyFormat         string     `json:"money_with_currency_format"`
	WeightUnit                      string     `json:"weight_unit"`
	MyshopifyDomain                 string     `json:"myshopify_domain"`
	PlanName                        string     `json:"plan_name"`
	PlanDisplayName                 string     `json:"plan_display_name"`
	PasswordEnabled                 bool       `json:"password_enabled"`
	PrimaryLocale                   string     `json:"primary_locale"`
	PrimaryLocationId               int64      `json:"primary_location_id"`
	Timezone                        string     `json:"timezone"`
	IanaTimezone                    string     `json:"iana_timezone"`
	ForceSSL                        bool       `json:"force_ssl"`
	TaxShipping                     bool       `json:"tax_shipping"`
	TaxesIncluded                   bool       `json:"taxes_included"`
	HasStorefront                   bool       `json:"has_storefront"`
	HasDiscounts                    bool       `json:"has_discounts"`
	HasGiftcards                    bool       `json:"has_gift_cards"`
	SetupRequire                    bool       `json:"setup_required"`
	CountyTaxes                     bool       `json:"county_taxes"`
	CheckoutAPISupported            bool       `json:"checkout_api_supported"`
	Source                          string     `json:"source"`
	GoogleAppsDomain                string     `json:"google_apps_domain"`
	GoogleAppsLoginEnabled          bool       `json:"google_apps_login_enabled"`
	MoneyInEmailsFormat             string     `json:"money_in_emails_format"`
	MoneyWithCurrencyInEmailsFormat string     `json:"money_with_currency_in_emails_format"`
	EligibleForPayments             bool       `json:"eligible_for_payments"`
	RequiresExtraPaymentsAgreement  bool       `json:"requires_extra_payments_agreement"`
	PreLaunchEnabled                bool       `json:"pre_launch_enabled"`
}

//TODO: latest from shopify 23/04

// type Shop struct {
// 	Alerts                              []ShopAlert                       `json:"alerts"`
// 	AllProductCategories                []ProductCategory                 `json:"allProductCategories"`
// 	BillingAddress                      ShopAddress                       `json:"billingAddress"`
// 	ChannelDefinitionsForInstalledChannels []AvailableChannelDefinitionsByChannel `json:"channelDefinitionsForInstalledChannels"`
// 	CheckoutApiSupported                bool                              `json:"checkoutApiSupported"`
// 	ContactEmail                        string                            `json:"contactEmail"`
// 	CountriesInShippingZones            CountriesInShippingZones          `json:"countriesInShippingZones"`
// 	CurrencyCode                        CurrencyCode                      `json:"currencyCode"`
// 	CurrencyFormats                     CurrencyFormats                   `json:"currencyFormats"`
// 	CustomerAccounts                    ShopCustomerAccountsSetting       `json:"customerAccounts"`
// 	Description                         *string                           `json:"description"`
// 	Email                               string                            `json:"email"`
// 	EnabledPresentmentCurrencies        []CurrencyCode                    `json:"enabledPresentmentCurrencies"`
// 	Features                            ShopFeatures                      `json:"features"`
// 	FulfillmentServices                 []FulfillmentService              `json:"fulfillmentServices"`
// 	IanaTimezone                        string                            `json:"ianaTimezone"`
// 	ID                                  string                            `json:"id"`
// 	LimitedPendingOrderCount            LimitedPendingOrderCount          `json:"limitedPendingOrderCount"`
// 	MerchantApprovalSignals             *MerchantApprovalSignals          `json:"merchantApprovalSignals"`
// 	Metafield                           *Metafield                        `json:"metafield"`
// 	MyshopifyDomain                     string                            `json:"myshopifyDomain"`
// 	Name                                string                            `json:"name"`
// 	NavigationSettings                  []NavigationItem                  `json:"navigationSettings"`
// 	OrderNumberFormatPrefix             string                            `json:"orderNumberFormatPrefix"`
// 	OrderNumberFormatSuffix             string                            `json:"orderNumberFormatSuffix"`
// 	PaymentSettings                     PaymentSettings                   `json:"paymentSettings"`
// 	Plan                                ShopPlan                          `json:"plan"`
// 	PrimaryDomain                       Domain                            `json:"primaryDomain"`
// 	PublicationCount                    int                               `json:"publicationCount"`
// 	ResourceLimits                      ShopResourceLimits                `json:"resourceLimits"`
// 	RichTextEditorURL                   string                            `json:"richTextEditorUrl"`
// 	SearchFilters                       SearchFilterOptions               `json:"searchFilters"`
// 	SetupRequired                       bool                              `json:"setupRequired"`
// 	ShipsToCountries                    []CountryCode                     `json:"shipsToCountries"`
// 	ShopPolicies                        []ShopPolicy                      `json:"shopPolicies"`
// 	TaxShipping                         bool                              `json:"taxShipping"`
// 	TaxesIncluded                       bool                              `json:"taxesIncluded"`
// 	TimezoneAbbreviation                string                            `json:"timezoneAbbreviation"`
// 	TimezoneOffset                      string                            `json:"timezoneOffset"`
// 	TimezoneOffsetMinutes               int                               `json:"timezoneOffsetMinutes"`
// 	TransactionalSMSDisabled            bool                              `json:"transactionalSmsDisabled"`
// 	Translations                        []Translation                     `json:"translations"`
// 	UnitSystem                          UnitSystem                        `json:"unitSystem"`
// 	UploadedImagesByIds                 []Image                           `json:"uploadedImagesByIds"`
// }

//TODO: latest from shopify 23/04

// type Shop struct {
// 	Alerts                   []struct {
// 		*ShopAlert `json:"alerts"`
// 	} `json:"alerts"`
// 	AllProductCategories      []struct {
// 		*ProductCategory `json:"allProductCategories"`
// 	} `json:"allProductCategories"`
// 	BillingAddress           struct {
// 		*ShopAddress `json:"billingAddress"`
// 	} `json:"billingAddress"`
// 	ChannelDefinitionsForInstalledChannels []struct {
// 		*AvailableChannelDefinitionsByChannel `json:"availableChannelDefinitions"`
// 	} `json:"channelDefinitionsForInstalledChannels"`
// 	CheckoutApiSupported           bool    `json:"checkoutApiSupported"`
// 	CustomerAccounts           bool    `json:"customerAccounts"`
// 	EnabledPresentmentCurrencies bool    `json:"enabledPresentmentCurrencies"`
// 	Email                       string  `json:"email"`
// 	EnabledFeatures             []struct {
// 		*Feature `json:"feature"`
// 	} `json:"enabledFeatures"`
// 	Features                    struct {
// 		*ShopFeatures `json:"features"`
// 	} `json:"features"`
// 	FulfillmentServices         []struct {
// 		*FulfillmentService `json:"fulfillmentService"`
// 	} `json:"fulfillmentServices"`
// 	IanaTimezone              string    `json:"ianaTimezone"`
// 	ID                         string    `json:"id"`
// 	Name                        string    `json:"name"`
// 	NavigationSettings         struct {
// 		*NavigationItem `json:"navigationSettings"`
// 	} `json:"navigationSettings"`
// 	OrderNumberFormatPrefix    string    `json:"orderNumberFormatPrefix"`
// 	OrderNumberFormatSuffix    string    `json:"orderNumberFormatSuffix"`
// 	PaymentSettings           struct {
// 		*PaymentSettings `json:"paymentSettings"`
// 	} `json:"paymentSettings"`
// 	Plan                       struct {
// 		*ShopPlan `json:"plan"`
// 	} `json:"plan"`
// 	PublicationCount           int      `json:"publicationCount"`
// 	ResourceLimits           struct {
// 		*ShopResourceLimits `json:"resourceLimits"`
// 	} `json:"resourceLimits"`
// 	RichTextEditorUrl         string    `json:"richTextEditorUrl"`
// 	SearchFilters             []struct {
// 		*Translation `json:"translation"`
// 	} `json:"searchFilters"`
// 	SetupRequired             bool    `json:"setupRequired"`
// 	ShipsToCountries         []struct {
// 		*CountryCode `json:"countryCode"`
// 	} `json:"shipsToCountries"`
// 	ShopPolicies             struct {
// 		*ShopPolicy `json:"shopPolicies"`
// 	} `json:"shopPolicies"`
// 	TaxShipping                bool    `json:"taxShipping"`
// 	TaxesIncluded             bool    `json:"taxesIncluded"`
// 	TransactionalSmsDisabled    bool    `json:"transactionalSmsDisabled"`
// 	Translations              []struct {
// 		*Translation `json:"translation"`
// 	} `json:"translations"`
// 	UnitSystem                 string    `json:"unitSystem"`
// 	UploadedImagesByIds       []struct {
// 		*Image                 `json:"uploadedImagesByIds"`
// 	} `json:"uploadedImagesByIds"`
// }

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
