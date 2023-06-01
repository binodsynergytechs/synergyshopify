package goshopify

import (
	"github.com/shopspring/decimal"
)

// ShippingZoneService is an interface for interfacing with the shipping zones endpoint
// of the Shopify API.
// See: https://help.shopify.com/api/reference/store-properties/shippingzone
type ShippingZoneService interface {
	List() ([]ShippingZone, error)
}

// ShippingZoneServiceOp handles communication with the shipping zone related methods
// of the Shopify API.
type ShippingZoneServiceOp struct {
	client *Client
}

// ShippingZone represents a Shopify shipping zone
type ShippingZone struct {
	ID                           int64                         `json:"id,omitempty"`
	Name                         string                        `json:"name,omitempty"`
	ProfileID                    string                        `json:"profile_id,omitempty"`
	LocationGroupID              string                        `json:"location_group_id,omitempty"`
	AdminGraphqlAPIID            string                        `json:"admin_graphql_api_id,omitempty"`
	Countries                    []ShippingCountry             `json:"countries,omitempty"`
	WeightBasedShippingRates     []WeightBasedShippingRate     `json:"weight_based_shipping_rates,omitempty"`
	PriceBasedShippingRates      []PriceBasedShippingRate      `json:"price_based_shipping_rates,omitempty"`
	CarrierShippingRateProviders []CarrierShippingRateProvider `json:"carrier_shipping_rate_providers,omitempty"`
}

// TODO: Latest From Shopify 23/04

// type ScriptTag struct {
// 	ID                  int         `json:"id"`
// 	Name                string      `json:"name"`
// 	ProfileID           string      `json:"profile_id"`
// 	LocationGroupID      string      `json:"location_group_id"`
// 	Countries            []Country   `json:"countries"`
// 	Provinces           []Province  `json:"provinces"`
// 	CarrierShippingRateProviders []interface{} `json:"carrier_shipping_rate_providers"`
// 	PriceBasedShippingRates []PriceBasedShippingRate `json:"price_based_shipping_rates"`
// 	WeightBasedShippingRates []WeightBasedShippingRate `json:"weight_based_shipping_rates"`
// }

// TODO: Latest From Shopify 23/04
// type Country struct {
// 	ID          int    `json:"id"`
// 	ShippingZoneID int    `json:"shipping_zone_id"`
// 	Name        string `json:"name"`
// 	Tax         float64 `json:"tax"`
// 	Code        string `json:"code"`
// 	TaxName     string `json:"tax_name"`
// 	Provinces  []interface{} `json:"provinces"`
// }

// TODO: Latest From Shopify 23/04
// type Province struct {
// 	Code        string `json:"code"`
// 	CountryID   int    `json:"country_id"`
// 	ShippingZoneID int    `json:"shipping_zone_id"`
// 	Name        string `json:"name"`
// 	Tax         float64 `json:"tax"`
// 	TaxName     string `json:"tax_name"`
// 	TaxPercent   float64 `json:"tax_percent"`
// }

// TODO: Latest From Shopify 23/04
// type PriceBasedShippingRate struct {
// 	ID          int    `json:"id"`
// 	Name        string `json:"name"`
// 	Price       float64 `json:"price"`
// 	MinOrderSubtotal float64 `json:"min_order_subtotal"`
// 	MaxOrderSubtotal float64 `json:"max_order_subtotal"`
// }

// TODO: Latest From Shopify 23/04
// type WeightBasedShippingRate struct {
// 	ID          int    `json:"id"`
// 	Name        string `json:"name"`
// 	Price       float64 `json:"price"`
// 	WeightLow    float64 `json:"weight_low"`
// 	WeightHigh   float64 `json:"weight_high"`
// 	WeightRange float64 `json:"weight_range"`
// }

// ShippingCountry represents a Shopify shipping country
type ShippingCountry struct {
	ID             int64              `json:"id,omitempty"`
	ShippingZoneID int64              `json:"shipping_zone_id,omitempty"`
	Name           string             `json:"name,omitempty"`
	Tax            *decimal.Decimal   `json:"tax,omitempty"`
	Code           string             `json:"code,omitempty"`
	TaxName        string             `json:"tax_name,omitempty"`
	Provinces      []ShippingProvince `json:"provinces,omitempty"`
}

// ShippingProvince represents a Shopify shipping province
type ShippingProvince struct {
	ID             int64            `json:"id,omitempty"`
	CountryID      int64            `json:"country_id,omitempty"`
	ShippingZoneID int64            `json:"shipping_zone_id,omitempty"`
	Name           string           `json:"name,omitempty"`
	Code           string           `json:"code,omitempty"`
	Tax            *decimal.Decimal `json:"tax,omitempty"`
	TaxName        string           `json:"tax_name,omitempty"`
	TaxType        string           `json:"tax_type,omitempty"`
	TaxPercentage  *decimal.Decimal `json:"tax_percentage,omitempty"`
}

// WeightBasedShippingRate represents a Shopify weight-constrained shipping rate
type WeightBasedShippingRate struct {
	ID             int64            `json:"id,omitempty"`
	ShippingZoneID int64            `json:"shipping_zone_id,omitempty"`
	Name           string           `json:"name,omitempty"`
	Price          *decimal.Decimal `json:"price,omitempty"`
	WeightLow      *decimal.Decimal `json:"weight_low,omitempty"`
	WeightHigh     *decimal.Decimal `json:"weight_high,omitempty"`
}

// PriceBasedShippingRate represents a Shopify subtotal-constrained shipping rate
type PriceBasedShippingRate struct {
	ID               int64            `json:"id,omitempty"`
	ShippingZoneID   int64            `json:"shipping_zone_id,omitempty"`
	Name             string           `json:"name,omitempty"`
	Price            *decimal.Decimal `json:"price,omitempty"`
	MinOrderSubtotal *decimal.Decimal `json:"min_order_subtotal,omitempty"`
	MaxOrderSubtotal *decimal.Decimal `json:"max_order_subtotal,omitempty"`
}

// CarrierShippingRateProvider represents a Shopify carrier-constrained shipping rate
type CarrierShippingRateProvider struct {
	ID               int64             `json:"id,omitempty"`
	CarrierServiceID int64             `json:"carrier_service_id,omitempty"`
	ShippingZoneID   int64             `json:"shipping_zone_id,omitempty"`
	FlatModifier     *decimal.Decimal  `json:"flat_modifier,omitempty"`
	PercentModifier  *decimal.Decimal  `json:"percent_modifier,omitempty"`
	ServiceFilter    map[string]string `json:"service_filter,omitempty"`
}

// Represents the result from the shipping_zones.json endpoint
type ShippingZonesResource struct {
	ShippingZones []ShippingZone `json:"shipping_zones"`
}

// List shipping zones
func (s *ShippingZoneServiceOp) List() ([]ShippingZone, error) {
	resource := new(ShippingZonesResource)
	err := s.client.Get("shipping_zones.json", resource, nil)
	return resource.ShippingZones, err
}
