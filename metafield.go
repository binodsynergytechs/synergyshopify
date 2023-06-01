package goshopify

import (
	"fmt"
	"time"
)

// MetafieldService is an interface for interfacing with the metafield endpoints
// of the Shopify API.
// https://help.shopify.com/api/reference/metafield
type MetafieldService interface {
	List(interface{}) ([]MetaField, error)
	Count(interface{}) (int, error)
	Get(int64, interface{}) (*MetaField, error)
	Create(MetaField) (*MetaField, error)
	Update(MetaField) (*MetaField, error)
	Delete(int64) error
}

// MetafieldsService is an interface for other Shopify resources
// to interface with the metafield endpoints of the Shopify API.
// https://help.shopify.com/api/reference/metafield
type MetafieldsService interface {
	ListMetafields(int64, interface{}) ([]MetaField, error)
	CountMetafields(int64, interface{}) (int, error)
	GetMetafield(int64, int64, interface{}) (*MetaField, error)
	CreateMetafield(int64, MetaField) (*MetaField, error)
	UpdateMetafield(int64, MetaField) (*MetaField, error)
	DeleteMetafield(int64, int64) error
}

// MetafieldServiceOp handles communication with the metafield
// related methods of the Shopify API.
type MetafieldServiceOp struct {
	client     *Client
	resource   string
	resourceID int64
}

// MetaField represents a Shopify metafield.
type MetaField struct {
	ID                int64       `json:"id,omitempty"`
	Key               string      `json:"key,omitempty"`
	Value             interface{} `json:"value,omitempty"`
	ValueType         string      `json:"value_type,omitempty"` // FIXME: Not Available In Latest Shopify Model
	Type              string      `json:"type,omitempty"`
	Namespace         string      `json:"namespace,omitempty"`
	Description       string      `json:"description,omitempty"`
	OwnerId           int64       `json:"owner_id,omitempty"`
	CreatedAt         *time.Time  `json:"created_at,omitempty"`
	UpdatedAt         *time.Time  `json:"updated_at,omitempty"`
	OwnerResource     string      `json:"owner_resource,omitempty"`
	AdminGraphqlAPIID string      `json:"admin_graphql_api_id,omitempty"` // FIXME: Not Available In Latest Shopify Model
}

// FIXME: Not Available In Latest Shopify Model
// TODO: Available In Latest Shopify Model 23/04

// MetafieldResource represents the result from the metafields/X.json endpoint
type MetafieldResource struct {
	Metafield *MetaField `json:"metafield"`
}

// MetafieldsResource represents the result from the metafields.json endpoint
type MetafieldsResource struct {
	Metafields []MetaField `json:"metafields"`
}

// List metafields
func (s *MetafieldServiceOp) List(options interface{}) ([]MetaField, error) {
	prefix := MetafieldPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s.json", prefix)
	resource := new(MetafieldsResource)
	err := s.client.Get(path, resource, options)
	return resource.Metafields, err
}

// Count metafields
func (s *MetafieldServiceOp) Count(options interface{}) (int, error) {
	prefix := MetafieldPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/count.json", prefix)
	return s.client.Count(path, options)
}

// Get individual metafield
func (s *MetafieldServiceOp) Get(metafieldID int64, options interface{}) (*MetaField, error) {
	prefix := MetafieldPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, metafieldID)
	resource := new(MetafieldResource)
	err := s.client.Get(path, resource, options)
	return resource.Metafield, err
}

// Create a new metafield
func (s *MetafieldServiceOp) Create(metafield MetaField) (*MetaField, error) {
	prefix := MetafieldPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s.json", prefix)
	wrappedData := MetafieldResource{Metafield: &metafield}
	resource := new(MetafieldResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Metafield, err
}

// Update an existing metafield
func (s *MetafieldServiceOp) Update(metafield MetaField) (*MetaField, error) {
	prefix := MetafieldPathPrefix(s.resource, s.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, metafield.ID)
	wrappedData := MetafieldResource{Metafield: &metafield}
	resource := new(MetafieldResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Metafield, err
}

// Delete an existing metafield
func (s *MetafieldServiceOp) Delete(metafieldID int64) error {
	prefix := MetafieldPathPrefix(s.resource, s.resourceID)
	return s.client.Delete(fmt.Sprintf("%s/%d.json", prefix, metafieldID))
}
