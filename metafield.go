package goshopify

import (
	"fmt"
	"time"
)

// MetafieldsService is an interface for other Shopify resources
// to interface with the metafield endpoints of the Shopify API.
// https://help.shopify.com/api/reference/metafield
type MetaFieldsRepository interface {
	ListMetaFields(int64, interface{}) ([]MetaField, error)
	CountMetaFields(int64, interface{}) (int, error)
	GetMetaField(int64, int64, interface{}) (*MetaField, error)
	CreateMetaField(int64, MetaField) (*MetaField, error)
	UpdateMetaField(int64, MetaField) (*MetaField, error)
	DeleteMetaField(int64, int64) error
}

// MetaFieldClient handles communication with the metafield
// related methods of the Shopify API.
type MetaFieldClient struct {
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
	AdminGraphqlApiID string      `json:"admin_graphql_api_id,omitempty"` // FIXME: Not Available In Latest Shopify Model
}

// FIXME: Not Available In Latest Shopify Model
// TODO: Available In Latest Shopify Model 23/04

// SingleMetaFieldResource represents the result from the metafields/X.json endpoint
type SingleMetaFieldResource struct {
	MetaField *MetaField `json:"metafield"`
}

// MultipleMetafieldsResource represents the result from the metafieldmc.json endpoint
type MultipleMetaFieldsResource struct {
	MetaFields []MetaField `json:"metafields"`
}

// List metafields
func (mc *MetaFieldClient) List(options interface{}) ([]MetaField, error) {
	prefix := MetaFieldPathPrefix(mc.resource, mc.resourceID)
	path := fmt.Sprintf("%mc.json", prefix)
	resource := new(MultipleMetaFieldsResource)
	err := mc.client.Get(path, resource, options)
	return resource.MetaFields, err
}

// Count metafields
func (mc *MetaFieldClient) Count(options interface{}) (int, error) {
	prefix := MetaFieldPathPrefix(mc.resource, mc.resourceID)
	path := fmt.Sprintf("%s/count.json", prefix)
	return mc.client.Count(path, options)
}

// Get individual metafield
func (mc *MetaFieldClient) Get(metafieldID int64, options interface{}) (*MetaField, error) {
	prefix := MetaFieldPathPrefix(mc.resource, mc.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, metafieldID)
	resource := new(SingleMetaFieldResource)
	err := mc.client.Get(path, resource, options)
	return resource.MetaField, err
}

// Create a new metafield
func (mc *MetaFieldClient) Create(metafield MetaField) (*MetaField, error) {
	prefix := MetaFieldPathPrefix(mc.resource, mc.resourceID)
	path := fmt.Sprintf("%mc.json", prefix)
	wrappedData := SingleMetaFieldResource{MetaField: &metafield}
	resource := new(SingleMetaFieldResource)
	err := mc.client.Post(path, wrappedData, resource)
	return resource.MetaField, err
}

// Update an existing metafield
func (mc *MetaFieldClient) Update(metafield MetaField) (*MetaField, error) {
	prefix := MetaFieldPathPrefix(mc.resource, mc.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, metafield.ID)
	wrappedData := SingleMetaFieldResource{MetaField: &metafield}
	resource := new(SingleMetaFieldResource)
	err := mc.client.Put(path, wrappedData, resource)
	return resource.MetaField, err
}

// Delete an existing metafield
func (mc *MetaFieldClient) Delete(metafieldID int64) error {
	prefix := MetaFieldPathPrefix(mc.resource, mc.resourceID)
	return mc.client.Delete(fmt.Sprintf("%s/%d.json", prefix, metafieldID))
}
