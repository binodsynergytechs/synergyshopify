package goshopify

import (
	"github.com/shopspring/decimal"
)

// ShippingZoneRepository is an interface for interfacing with the shipping zones endpoint
// of the Shopify API.
// See: https://help.shopify.com/api/reference/store-properties/shippingzone
type ShippingZoneRepository interface {
	ListShippingZone() ([]ShippingZone, error)
}

// ShippingZoneClient handles communication with the shipping zone related methods
// of the Shopify API.
type ShippingZoneClient struct {
	client *Client
}

// ShippingZone represents a Shopify shipping zone
type ShippingZone struct {
	ID                           int64                         `json:"id,omitempty"`
	Name                         string                        `json:"name,omitempty"`
	ProfileID                    string                        `json:"profile_id,omitempty"`
	LocationGroupID              string                        `json:"location_group_id,omitempty"`
	AdminGraphqlApiID            string                        `json:"admin_graphql_api_id,omitempty"`
	Countries                    []ShippingCountry             `json:"countries,omitempty"`
	WeightBasedShippingRates     []WeightBasedShippingRate     `json:"weight_based_shipping_rates,omitempty"`
	PriceBasedShippingRates      []PriceBasedShippingRate      `json:"price_based_shipping_rates,omitempty"`
	CarrierShippingRateProviders []CarrierShippingRateProvider `json:"carrier_shipping_rate_providers,omitempty"`
	Provinces                    []ShippingProvince            `json:"provinces,omitempty"` // TODO: Latest Added From Shopify 23/04
}

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
func (s *ShippingZoneClient) ListShippingZone() ([]ShippingZone, error) {
	resource := new(ShippingZonesResource)
	err := s.client.Get("shipping_zones.json", resource, nil)
	return resource.ShippingZones, err
}
