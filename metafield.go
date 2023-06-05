package goshopify

import (
	"fmt"
	"time"
)

// MetaFieldRepository is an interface for interfacing with the metaField endpoints
// of the Shopify API.
// https://help.shopify.com/api/reference/metaField
type MetaFieldRepository interface {
	ListMetaField(interface{}) ([]MetaField, error)
	CountMetaField(interface{}) (int, error)
	GetMetaField(int64, interface{}) (*MetaField, error)
	CreateMetaField(MetaField) (*MetaField, error)
	UpdateMetaField(MetaField) (*MetaField, error)
	DeleteMetaField(int64) error
}

// MetaFieldsRepository is an interface for other Shopify resources
// to interface with the metaField endpoints of the Shopify API.
// https://help.shopify.com/api/reference/metaField
type MetaFieldsRepository interface {
	ListMetaFields(int64, interface{}) ([]MetaField, error)
	CountMetaFields(int64, interface{}) (int, error)
	GetMetaFields(int64, int64, interface{}) (*MetaField, error)
	CreateMetaFields(int64, MetaField) (*MetaField, error)
	UpdateMetaFields(int64, MetaField) (*MetaField, error)
	DeleteMetaFields(int64, int64) error
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
	ValueType         string      `json:"value_type,omitempty"`
	Type              string      `json:"type,omitempty"`
	Namespace         string      `json:"namespace,omitempty"`
	Description       string      `json:"description,omitempty"`
	OwnerId           int64       `json:"owner_id,omitempty"`
	CreatedAt         *time.Time  `json:"created_at,omitempty"`
	UpdatedAt         *time.Time  `json:"updated_at,omitempty"`
	OwnerResource     string      `json:"owner_resource,omitempty"`
	AdminGraphqlApiID string      `json:"admin_graphql_api_id,omitempty"`
}

// SingleMetaFieldResponse represents the result from the metaFields/X.json endpoint
type SingleMetaFieldResponse struct {
	MetaField *MetaField `json:"metaField"`
}

// MultipleMetaFieldsResponse represents the result from the metaFields.json endpoint
type MultipleMetaFieldsResponse struct {
	MetaFields []MetaField `json:"metaFields"`
}

// List metaFields
func (mc *MetaFieldClient) ListMetaField(options interface{}) ([]MetaField, error) {
	prefix := MetaFieldPathPrefix(mc.resource, mc.resourceID)
	path := fmt.Sprintf("%s.json", prefix)
	resource := new(MultipleMetaFieldsResponse)
	err := mc.client.Get(path, resource, options)
	return resource.MetaFields, err
}

// Count metaFields
func (mc *MetaFieldClient) CountMetaField(options interface{}) (int, error) {
	prefix := MetaFieldPathPrefix(mc.resource, mc.resourceID)
	path := fmt.Sprintf("%s/count.json", prefix)
	return mc.client.Count(path, options)
}

// Get individual metaField
func (mc *MetaFieldClient) GetMetaField(metaFieldID int64, options interface{}) (*MetaField, error) {
	prefix := MetaFieldPathPrefix(mc.resource, mc.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, metaFieldID)
	resource := new(SingleMetaFieldResponse)
	err := mc.client.Get(path, resource, options)
	return resource.MetaField, err
}

// Create a new metaField
func (mc *MetaFieldClient) CreateMetaField(metaField MetaField) (*MetaField, error) {
	prefix := MetaFieldPathPrefix(mc.resource, mc.resourceID)
	path := fmt.Sprintf("%s.json", prefix)
	wrappedData := SingleMetaFieldResponse{MetaField: &metaField}
	resource := new(SingleMetaFieldResponse)
	err := mc.client.Post(path, wrappedData, resource)
	return resource.MetaField, err
}

// Update an existing metaField
func (mc *MetaFieldClient) UpdateMetaField(metaField MetaField) (*MetaField, error) {
	prefix := MetaFieldPathPrefix(mc.resource, mc.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, metaField.ID)
	wrappedData := SingleMetaFieldResponse{MetaField: &metaField}
	resource := new(SingleMetaFieldResponse)
	err := mc.client.Put(path, wrappedData, resource)
	return resource.MetaField, err
}

// Delete an existing metaField
func (mc *MetaFieldClient) DeleteMetaField(metaFieldID int64) error {
	prefix := MetaFieldPathPrefix(mc.resource, mc.resourceID)
	return mc.client.Delete(fmt.Sprintf("%s/%d.json", prefix, metaFieldID))
}
