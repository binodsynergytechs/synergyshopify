package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const inventoryItemsBasePath = "inventory_items"

// InventoryItemRepository is an interface for interacting with the
// inventory items endpoints of the Shopify API
// See https://help.shopify.com/en/api/reference/inventory/inventoryitem
type InventoryItemRepository interface {
	ListInventoryItem(interface{}) ([]InventoryItem, error)
	GetInventoryItem(int64, interface{}) (*InventoryItem, error)
	UpdateInventoryItem(InventoryItem) (*InventoryItem, error)
}

// InventoryItemClient is the default implementation of the InventoryItemRepository interface
type InventoryItemClient struct {
	client *Client
}

// InventoryItem represents a Shopify inventory item
type InventoryItem struct {
	ID                           int64            `json:"id,omitempty"`
	SKU                          string           `json:"sku,omitempty"`
	CreatedAt                    *time.Time       `json:"created_at,omitempty"`
	UpdatedAt                    *time.Time       `json:"updated_at,omitempty"`
	Cost                         *decimal.Decimal `json:"cost,omitempty"`
	Tracked                      *bool            `json:"tracked,omitempty"`
	AdminGraphqlApiId            string           `json:"admin_graphql_api_id,omitempty"`
	RequireShipping              bool             `json:"requires_shipping"`
	ProvinceCodeOfOrigin         string           `json:"province_code_of_origin"`         //TODO: Field Available In Latest Shopify Model
	HarmonizedSystemCode         string           `json:"harmonized_system_code"`          //TODO: Field Available In Latest Shopify Model
	CountryHarmonizedSystemCodes []interface{}    `json:"country_harmonized_system_codes"` //TODO: Field Available In Latest Shopify Model
	CountryCodeOfOrigin          string           `json:"country_code_of_origin"`          //TODO: Field Available In Latest Shopify Model
}

// SingleInventoryItemResponse is used for handling single item requests and responses
type SingleInventoryItemResponse struct {
	InventoryItem *InventoryItem `json:"inventory_item"`
}

// MultipleInventoryItemsResponse is used for handling multiple item responsees
type MultipleInventoryItemsResponse struct {
	InventoryItems []InventoryItem `json:"inventory_items"`
}

// List inventory items
func (ic *InventoryItemClient) ListInventoryItem(options interface{}) ([]InventoryItem, error) {
	path := fmt.Sprintf("%s.json", inventoryItemsBasePath)
	resource := new(MultipleInventoryItemsResponse)
	err := ic.client.Get(path, resource, options)
	return resource.InventoryItems, err
}

// Get a inventory item
func (ic *InventoryItemClient) GetInventoryItem(id int64, options interface{}) (*InventoryItem, error) {
	path := fmt.Sprintf("%s/%d.json", inventoryItemsBasePath, id)
	resource := new(SingleInventoryItemResponse)
	err := ic.client.Get(path, resource, options)
	return resource.InventoryItem, err
}

// Update a inventory item
func (ic *InventoryItemClient) UpdateInventoryItem(item InventoryItem) (*InventoryItem, error) {
	path := fmt.Sprintf("%s/%d.json", inventoryItemsBasePath, item.ID)
	wrappedData := SingleInventoryItemResponse{InventoryItem: &item}
	resource := new(SingleInventoryItemResponse)
	err := ic.client.Put(path, wrappedData, resource)
	return resource.InventoryItem, err
}
