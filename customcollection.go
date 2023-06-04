package goshopify

import (
	"fmt"
	"time"
)

const customCollectionsBasePath = "custom_collections"
const MultipleCustomCollectionsResponseName = "collections"

// CustomCollectionRepository is an interface for interacting with the custom
// collection endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/customcollection
type CustomCollectionRepository interface {
	ListCustomCollection(interface{}) ([]CustomCollection, error)
	CountCustomCollection(interface{}) (int, error)
	GetCustomCollection(int64, interface{}) (*CustomCollection, error)
	CreateCustomCollection(CustomCollection) (*CustomCollection, error)
	UpdateCustomCollection(CustomCollection) (*CustomCollection, error)
	DeleteCustomCollection(int64) error

	// MetaFieldsRepository used for CustomCollection resource to communicate with MetaFields resource
	MetaFieldsRepository
}

// CustomCollectionClient handles communication with the custom collection
// related methods of the Shopify API.
type CustomCollectionClient struct {
	client *Client
}

// CustomCollection represents a Shopify custom collection.
type CustomCollection struct {
	ID             int64       `json:"id"`
	Handle         string      `json:"handle"`
	Title          string      `json:"title"`
	UpdatedAt      *time.Time  `json:"updated_at"`
	BodyHTML       string      `json:"body_html"`
	SortOrder      string      `json:"sort_order"`
	TemplateSuffix string      `json:"template_suffix"`
	Image          Image       `json:"image"`
	Published      bool        `json:"published"`
	PublishedAt    *time.Time  `json:"published_at"`
	PublishedScope string      `json:"published_scope"`
	MetaFields     []MetaField `json:"metaFields,omitempty"` // FIXME: Field Not Available Or Deprecated In Latest Shopify Model 23/04
}

// SingleCustomCollectionResponse represents the result form the custom_collections/X.json endpoint
type SingleCustomCollectionResponse struct {
	Collection *CustomCollection `json:"custom_collection"`
}

// MultipleCustomCollectionsResponse represents the result from the custom_collections.json endpoint
type MultipleCustomCollectionsResponse struct {
	Collections []CustomCollection `json:"custom_collections"`
}

// List custom collections
func (ccc *CustomCollectionClient) ListCustomCollection(options interface{}) ([]CustomCollection, error) {
	path := fmt.Sprintf("%s.json", customCollectionsBasePath)
	resource := new(MultipleCustomCollectionsResponse)
	err := ccc.client.Get(path, resource, options)
	return resource.Collections, err
}

// Count custom collections
func (ccc *CustomCollectionClient) CountCustomCollection(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", customCollectionsBasePath)
	return ccc.client.Count(path, options)
}

// Get individual custom collection
func (ccc *CustomCollectionClient) GetCustomCollection(collectionID int64, options interface{}) (*CustomCollection, error) {
	path := fmt.Sprintf("%s/%d.json", customCollectionsBasePath, collectionID)
	resource := new(SingleCustomCollectionResponse)
	err := ccc.client.Get(path, resource, options)
	return resource.Collection, err
}

// Create a new custom collection
// See Image for the details of the Image creation for a collection.
func (ccc *CustomCollectionClient) CreateCustomCollection(collection CustomCollection) (*CustomCollection, error) {
	path := fmt.Sprintf("%s.json", customCollectionsBasePath)
	wrappedData := SingleCustomCollectionResponse{Collection: &collection}
	resource := new(SingleCustomCollectionResponse)
	err := ccc.client.Post(path, wrappedData, resource)
	return resource.Collection, err
}

// Update an existing custom collection
func (ccc *CustomCollectionClient) UpdateCustomCollection(collection CustomCollection) (*CustomCollection, error) {
	path := fmt.Sprintf("%s/%d.json", customCollectionsBasePath, collection.ID)
	wrappedData := SingleCustomCollectionResponse{Collection: &collection}
	resource := new(SingleCustomCollectionResponse)
	err := ccc.client.Put(path, wrappedData, resource)
	return resource.Collection, err
}

// Delete an existing custom collection.
func (ccc *CustomCollectionClient) Delete(collectionID int64) error {
	return ccc.client.Delete(fmt.Sprintf("%s/%d.json", customCollectionsBasePath, collectionID))
}

// List metaFields for a custom collection
func (ccc *CustomCollectionClient) ListMetaFields(customCollectionID int64, options interface{}) ([]MetaField, error) {
	metaFieldService := &MetaFieldClient{client: ccc.client, resource: MultipleCustomCollectionsResponseName, resourceID: customCollectionID}
	return metaFieldService.List(options)
}

// Count metaFields for a custom collection
func (ccc *CustomCollectionClient) CountMetaFields(customCollectionID int64, options interface{}) (int, error) {
	metaFieldService := &MetaFieldClient{client: ccc.client, resource: MultipleCustomCollectionsResponseName, resourceID: customCollectionID}
	return metaFieldService.Count(options)
}

// Get individual metaField for a custom collection
func (ccc *CustomCollectionClient) GetMetaField(customCollectionID int64, metaFieldID int64, options interface{}) (*MetaField, error) {
	metaFieldService := &MetaFieldClient{client: ccc.client, resource: MultipleCustomCollectionsResponseName, resourceID: customCollectionID}
	return metaFieldService.Get(metaFieldID, options)
}

// Create a new metaField for a custom collection
func (ccc *CustomCollectionClient) CreateMetaField(customCollectionID int64, metaField MetaField) (*MetaField, error) {
	metaFieldService := &MetaFieldClient{client: ccc.client, resource: MultipleCustomCollectionsResponseName, resourceID: customCollectionID}
	return metaFieldService.Create(metaField)
}

// Update an existing metaField for a custom collection
func (ccc *CustomCollectionClient) UpdateMetaField(customCollectionID int64, metaField MetaField) (*MetaField, error) {
	metaFieldService := &MetaFieldClient{client: ccc.client, resource: MultipleCustomCollectionsResponseName, resourceID: customCollectionID}
	return metaFieldService.Update(metaField)
}

// // Delete an existing metaField for a custom collection
func (ccc *CustomCollectionClient) DeleteMetaField(customCollectionID int64, metaFieldID int64) error {
	metaFieldService := &MetaFieldClient{client: ccc.client, resource: MultipleCustomCollectionsResponseName, resourceID: customCollectionID}
	return metaFieldService.Delete(metaFieldID)
}
