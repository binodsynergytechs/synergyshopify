package goshopify

import (
	"fmt"
	"time"
)

const collectsBasePath = "collects"

// CollectRepository is an interface for interfacing with the collect endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/products/collect
type CollectRepository interface {
	ListCollect(interface{}) ([]Collect, error)
	CountCollect(interface{}) (int, error)
	GetCollect(int64, interface{}) (*Collect, error)
	CreateCollect(Collect) (*Collect, error)
	DeleteCollect(int64) error
}

// CollectClient handles communication with the collect related methods of
// the Shopify API.
type CollectClient struct {
	client *Client
}

// Collect represents a Shopify collect
type Collect struct {
	ID           int64      `json:"id,omitempty"`
	CollectionID int64      `json:"collection_id,omitempty"`
	ProductID    int64      `json:"product_id,omitempty"`
	Featured     bool       `json:"featured,omitempty"` // FIXME: Field Not Available Or Deprecated In Latest Shopify Model 23/04
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	Position     int        `json:"position,omitempty"`
	SortValue    string     `json:"sort_value,omitempty"`
}

// Represents the result from the collects/X.json endpoint
type SingleCollectResponse struct {
	Collect *Collect `json:"collect"`
}

// Represents the result from the collects.json endpoint
type MultipleCollectsResponse struct {
	Collects []Collect `json:"collects"`
}

// List collects
func (s *CollectClient) ListCollect(options interface{}) ([]Collect, error) {
	path := fmt.Sprintf("%s.json", collectsBasePath)
	resource := new(MultipleCollectsResponse)
	err := s.client.Get(path, resource, options)
	return resource.Collects, err
}

// Count collects
func (s *CollectClient) CountCollect(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", collectsBasePath)
	return s.client.Count(path, options)
}

// Get individual collect
func (s *CollectClient) GetCollect(collectID int64, options interface{}) (*Collect, error) {
	path := fmt.Sprintf("%s/%d.json", collectsBasePath, collectID)
	resource := new(SingleCollectResponse)
	err := s.client.Get(path, resource, options)
	return resource.Collect, err
}

// Create collects
func (s *CollectClient) CreateCollect(collect Collect) (*Collect, error) {
	path := fmt.Sprintf("%s.json", collectsBasePath)
	wrappedData := SingleCollectResponse{Collect: &collect}
	resource := new(SingleCollectResponse)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Collect, err
}

// Delete an existing collect
func (s *CollectClient) DeleteCollect(collectID int64) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", collectsBasePath, collectID))
}
