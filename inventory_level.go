package goshopify

import (
	"fmt"
	"time"
)

const inventoryLevelsBasePath = "inventory_levels"

// InventoryLevelRepository is an interface for interacting with the
// inventory items endpoints of the Shopify API
// See https://help.shopify.com/en/api/reference/inventory/inventorylevel
type InventoryLevelRepository interface {
	ListInventoryLevel(interface{}) ([]InventoryLevel, error)
	AdjustInventoryLevel(interface{}) (*InventoryLevel, error)
	DeleteInventoryLevel(int64, int64) error
	ConnectInventoryLevel(InventoryLevel) (*InventoryLevel, error)
	SetInventoryLevel(InventoryLevel) (*InventoryLevel, error)
}

// InventoryLevelClient is the default implementation of the InventoryLevelRepository interface
type InventoryLevelClient struct {
	client *Client
}

// InventoryLevel represents a Shopify inventory level
type InventoryLevel struct {
	InventoryItemId   int64      `json:"inventory_item_id,omitempty"`
	LocationId        int64      `json:"location_id,omitempty"`
	Available         int        `json:"available"`
	CreatedAt         *time.Time `json:"created_at,omitempty"` //FIXME: Field Not Available In Latest Shopify Model
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	AdminGraphqlApiId string     `json:"admin_graphql_api_id,omitempty"` //FIXME: Field Not Available In Latest Shopify Model
}

// InventoryLevelResource is used for handling single level requests and responses
type InventoryLevelResource struct {
	InventoryLevel *InventoryLevel `json:"inventory_level"`
}

// InventoryLevelsResource is used for handling multiple item responsees
type InventoryLevelsResource struct {
	InventoryLevels []InventoryLevel `json:"inventory_levels"`
}

// InventoryLevelListOptions is used for get list
type InventoryLevelListOptions struct {
	InventoryItemIds []int64   `url:"inventory_item_ids,omitempty,comma"`
	LocationIds      []int64   `url:"location_ids,omitempty,comma"`
	Limit            int       `url:"limit,omitempty"`
	UpdatedAtMin     time.Time `url:"updated_at_min,omitempty"`
}

// InventoryLevelAdjustOptions is used for Adjust inventory levels
type InventoryLevelAdjustOptions struct {
	InventoryItemId int64 `json:"inventory_item_id"`
	LocationId      int64 `json:"location_id"`
	Adjust          int   `json:"available_adjustment"`
}

// List inventory levels
func (ic *InventoryLevelClient) ListInventoryLevel(options interface{}) ([]InventoryLevel, error) {
	path := fmt.Sprintf("%s.json", inventoryLevelsBasePath)
	resource := new(InventoryLevelsResource)
	err := ic.client.Get(path, resource, options)
	return resource.InventoryLevels, err
}

// Delete an inventory level
func (ic *InventoryLevelClient) DeleteInventoryLevel(itemId, locationId int64) error {
	path := fmt.Sprintf("%s.json?inventory_item_id=%v&location_id=%v",
		inventoryLevelsBasePath, itemId, locationId)
	return ic.client.Delete(path)
}

// Connect an inventory level
func (ic *InventoryLevelClient) ConnectInventoryLevel(level InventoryLevel) (*InventoryLevel, error) {
	return ic.PostInventoryLevel(fmt.Sprintf("%s/connect.json", inventoryLevelsBasePath), level)
}

// Set an inventory level
func (ic *InventoryLevelClient) SetInventoryLevel(level InventoryLevel) (*InventoryLevel, error) {
	return ic.PostInventoryLevel(fmt.Sprintf("%s/set.json", inventoryLevelsBasePath), level)
}

// Adjust the inventory level of an inventory item at a single location
func (ic *InventoryLevelClient) AdjustInventoryLevel(options interface{}) (*InventoryLevel, error) {
	return ic.PostInventoryLevel(fmt.Sprintf("%s/adjust.json", inventoryLevelsBasePath), options)
}

func (ic *InventoryLevelClient) PostInventoryLevel(path string, options interface{}) (*InventoryLevel, error) {
	resource := new(InventoryLevelResource)
	err := ic.client.Post(path, options, resource)
	return resource.InventoryLevel, err
}
