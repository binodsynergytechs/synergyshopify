package goshopify

import (
	"fmt"
	"time"
)

// MetaFieldsService is an interface for other Shopify resources
// to interface with the metaField endpoints of the Shopify API.
// https://help.shopify.com/api/reference/metaField
type MetaFieldRepository interface {
	ListMetaFields(int64, interface{}) ([]MetaField, error)
	CountMetaFields(int64, interface{}) (int, error)
	GetMetaField(int64, int64, interface{}) (*MetaField, error)
	CreateMetaField(int64, MetaField) (*MetaField, error)
	UpdateMetaField(int64, MetaField) (*MetaField, error)
	DeleteMetaField(int64, int64) error
}

// MetaFieldClient handles communication with the metaField
// related methods of the Shopify API.
type MetaFieldClient struct {
	client     *Client
	resource   string
	resourceID int64
}

// MetaField represents a Shopify metaField.
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

// SingleMetaFieldResource represents the result from the metaFields/X.json endpoint
type SingleMetaFieldResource struct {
	MetaField *MetaField `json:"metaField"`
}

// MultipleMetaFieldsResource represents the result from the metaFieldmc.json endpoint
type MultipleMetaFieldsResource struct {
	MetaFields []MetaField `json:"metaFields"`
}

// List metaFields
func (mc *MetaFieldClient) ListMetaFields(options interface{}) ([]MetaField, error) {
	prefix := MetaFieldPathPrefix(mc.resource, mc.resourceID)
	path := fmt.Sprintf("%mc.json", prefix)
	resource := new(MultipleMetaFieldsResource)
	err := mc.client.Get(path, resource, options)
	return resource.MetaFields, err
}

// Count metaFields
func (mc *MetaFieldClient) CountMetaFields(options interface{}) (int, error) {
	prefix := MetaFieldPathPrefix(mc.resource, mc.resourceID)
	path := fmt.Sprintf("%s/count.json", prefix)
	return mc.client.Count(path, options)
}

// Get individual metaField
func (mc *MetaFieldClient) GetMetaFields(metaFieldID int64, options interface{}) (*MetaField, error) {
	prefix := MetaFieldPathPrefix(mc.resource, mc.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, metaFieldID)
	resource := new(SingleMetaFieldResource)
	err := mc.client.Get(path, resource, options)
	return resource.MetaField, err
}

// Create a new metaField
func (mc *MetaFieldClient) CreateMetaFields(metaField MetaField) (*MetaField, error) {
	prefix := MetaFieldPathPrefix(mc.resource, mc.resourceID)
	path := fmt.Sprintf("%mc.json", prefix)
	wrappedData := SingleMetaFieldResource{MetaField: &metaField}
	resource := new(SingleMetaFieldResource)
	err := mc.client.Post(path, wrappedData, resource)
	return resource.MetaField, err
}

// Update an existing metaField
func (mc *MetaFieldClient) UpdateMetaFields(metaField MetaField) (*MetaField, error) {
	prefix := MetaFieldPathPrefix(mc.resource, mc.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, metaField.ID)
	wrappedData := SingleMetaFieldResource{MetaField: &metaField}
	resource := new(SingleMetaFieldResource)
	err := mc.client.Put(path, wrappedData, resource)
	return resource.MetaField, err
}

// Delete an existing metaField
func (mc *MetaFieldClient) DeleteMetaFields(metaFieldID int64) error {
	prefix := MetaFieldPathPrefix(mc.resource, mc.resourceID)
	return mc.client.Delete(fmt.Sprintf("%s/%d.json", prefix, metaFieldID))
}
