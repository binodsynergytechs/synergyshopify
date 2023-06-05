package goshopify

import (
	"fmt"
	"time"
)

const smartCollectionsBasePath = "smart_collections"
const MultiplesmartCollectionsResponseName = "collections"

// SmartCollectionRepository is an interface for interacting with the smart
// collection endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/smartcollection
type SmartCollectionRepository interface {
	ListSmartCollection(interface{}) ([]SmartCollection, error)
	CountSmartCollection(interface{}) (int, error)
	GetSmartCollection(int64, interface{}) (*SmartCollection, error)
	CreateSmartCollection(SmartCollection) (*SmartCollection, error)
	UpdateSmartCollection(SmartCollection) (*SmartCollection, error)
	DeleteSmartCollection(int64) error
	ListSmartCollectionMetaFields(smartCollectionID int64, options interface{}) ([]MetaField, error)
	CountSmartCollectionMetaFields(smartCollectionID int64, options interface{}) (int, error)
	GetSmartCollectionMetaField(smartCollectionID int64, metaFieldID int64, options interface{}) (*MetaField, error)
	CreateSmartCollectionMetaField(smartCollectionID int64, metaField MetaField) (*MetaField, error)
	UpdateSmartCollectionMetaField(smartCollectionID int64, metaField MetaField) (*MetaField, error)
	DeleteSmartCollectionMetaField(smartCollectionID int64, metaFieldID int64) error
	// MetaFieldsService used for SmartCollection resource to communicate with Metafields resource
	// MetaFieldsRepository
}

// SmartCollectionClient handles communication with the smart collection
// related methods of the Shopify API.
type SmartCollectionClient struct {
	client *Client
}

type Rule struct {
	Column    string `json:"column"`
	Relation  string `json:"relation"`
	Condition string `json:"condition"`
}

// SmartCollection represents a Shopify smart collection.
type SmartCollection struct {
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
	Rules          []Rule      `json:"rules"`
	Disjunctive    bool        `json:"disjunctive"`
	MetaFields     []MetaField `json:"metafields,omitempty"`
}

// SingleSmartCollectionResponse represents the result from the smart_collections/X.json endpoint
type SingleSmartCollectionResponse struct {
	Collection *SmartCollection `json:"smart_collection"`
}

// MultipleSmartCollectionsResponse represents the result from the smart_collections.json endpoint
type MultipleSmartCollectionsResponse struct {
	Collections []SmartCollection `json:"smart_collections"`
}

// List smart collections
func (s *SmartCollectionClient) ListSmartCollection(options interface{}) ([]SmartCollection, error) {
	path := fmt.Sprintf("%s.json", smartCollectionsBasePath)
	resource := new(MultipleSmartCollectionsResponse)
	err := s.client.Get(path, resource, options)
	return resource.Collections, err
}

// Count smart collections
func (s *SmartCollectionClient) CountSmartCollection(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", smartCollectionsBasePath)
	return s.client.Count(path, options)
}

// Get individual smart collection
func (s *SmartCollectionClient) GetSmartCollection(collectionID int64, options interface{}) (*SmartCollection, error) {
	path := fmt.Sprintf("%s/%d.json", smartCollectionsBasePath, collectionID)
	resource := new(SingleSmartCollectionResponse)
	err := s.client.Get(path, resource, options)
	return resource.Collection, err
}

// Create a new smart collection
// See Image for the details of the Image creation for a collection.
func (s *SmartCollectionClient) CreateSmartCollection(collection SmartCollection) (*SmartCollection, error) {
	path := fmt.Sprintf("%s.json", smartCollectionsBasePath)
	wrappedData := SingleSmartCollectionResponse{Collection: &collection}
	resource := new(SingleSmartCollectionResponse)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Collection, err
}

// Update an existing smart collection
func (s *SmartCollectionClient) UpdateSmartCollection(collection SmartCollection) (*SmartCollection, error) {
	path := fmt.Sprintf("%s/%d.json", smartCollectionsBasePath, collection.ID)
	wrappedData := SingleSmartCollectionResponse{Collection: &collection}
	resource := new(SingleSmartCollectionResponse)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Collection, err
}

// Delete an existing smart collection.
func (s *SmartCollectionClient) DeleteSmartCollection(collectionID int64) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", smartCollectionsBasePath, collectionID))
}

// List metaFields for a smart collection
func (s *SmartCollectionClient) ListSmartCollectionMetaFields(smartCollectionID int64, options interface{}) ([]MetaField, error) {
	mField := &MetaFieldClient{client: s.client, resource: MultiplesmartCollectionsResponseName, resourceID: smartCollectionID}
	return mField.ListMetaField(options)
}

// Count metaFields for a smart collection
func (s *SmartCollectionClient) CountSmartCollectionMetaFields(smartCollectionID int64, options interface{}) (int, error) {
	mField := &MetaFieldClient{client: s.client, resource: MultiplesmartCollectionsResponseName, resourceID: smartCollectionID}
	return mField.CountMetaField(options)
}

// Get individual metaField for a smart collection
func (s *SmartCollectionClient) GetSmartCollectionMetaField(smartCollectionID int64, metaFieldID int64, options interface{}) (*MetaField, error) {
	mField := &MetaFieldClient{client: s.client, resource: MultiplesmartCollectionsResponseName, resourceID: smartCollectionID}
	return mField.GetMetaField(metaFieldID, options)
}

// Create a new metaField for a smart collection
func (s *SmartCollectionClient) CreateSmartCollectionMetaField(smartCollectionID int64, metaField MetaField) (*MetaField, error) {
	mField := &MetaFieldClient{client: s.client, resource: MultiplesmartCollectionsResponseName, resourceID: smartCollectionID}
	return mField.CreateMetaField(metaField)
}

// Update an existing metaField for a smart collection
func (s *SmartCollectionClient) UpdateSmartCollectionMetaField(smartCollectionID int64, metaField MetaField) (*MetaField, error) {
	mField := &MetaFieldClient{client: s.client, resource: MultiplesmartCollectionsResponseName, resourceID: smartCollectionID}
	return mField.UpdateMetaField(metaField)
}

// // Delete an existing metaField for a smart collection
func (s *SmartCollectionClient) DeleteSmartCollectionMetaField(smartCollectionID int64, metaFieldID int64) error {
	mField := &MetaFieldClient{client: s.client, resource: MultiplesmartCollectionsResponseName, resourceID: smartCollectionID}
	return mField.DeleteMetaField(metaFieldID)
}
