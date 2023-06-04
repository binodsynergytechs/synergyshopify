package goshopify

import (
	"fmt"
	"time"
)

const smartCollectionsBasePath = "smart_collections"
const smartCollectionsResourceName = "collections"

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

	// MetafieldsService used for SmartCollection resource to communicate with Metafields resource
	MetaFieldRepository
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
	Published      bool        `json:"published"` // FIXME: Not Available In Latest Shopify Model 23/04
	PublishedAt    *time.Time  `json:"published_at"`
	PublishedScope string      `json:"published_scope"`
	Rules          []Rule      `json:"rules"`
	Disjunctive    bool        `json:"disjunctive"`
	MetaFields     []MetaField `json:"metafields,omitempty"` // FIXME: Not Available In Latest Shopify Model 23/04
}

//TODO: latest from shopify 23/04

// SmartCollectionResource represents the result from the smart_collections/X.json endpoint
type SmartCollectionResource struct {
	Collection *SmartCollection `json:"smart_collection"`
}

// SmartCollectionsResource represents the result from the smart_collections.json endpoint
type SmartCollectionsResource struct {
	Collections []SmartCollection `json:"smart_collections"`
}

// List smart collections
func (s *SmartCollectionClient) ListSmartCollection(options interface{}) ([]SmartCollection, error) {
	path := fmt.Sprintf("%s.json", smartCollectionsBasePath)
	resource := new(SmartCollectionsResource)
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
	resource := new(SmartCollectionResource)
	err := s.client.Get(path, resource, options)
	return resource.Collection, err
}

// Create a new smart collection
// See Image for the details of the Image creation for a collection.
func (s *SmartCollectionClient) CreateSmartCollection(collection SmartCollection) (*SmartCollection, error) {
	path := fmt.Sprintf("%s.json", smartCollectionsBasePath)
	wrappedData := SmartCollectionResource{Collection: &collection}
	resource := new(SmartCollectionResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Collection, err
}

// Update an existing smart collection
func (s *SmartCollectionClient) UpdateSmartCollection(collection SmartCollection) (*SmartCollection, error) {
	path := fmt.Sprintf("%s/%d.json", smartCollectionsBasePath, collection.ID)
	wrappedData := SmartCollectionResource{Collection: &collection}
	resource := new(SmartCollectionResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Collection, err
}

// Delete an existing smart collection.
func (s *SmartCollectionClient) DeleteSmartCollection(collectionID int64) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", smartCollectionsBasePath, collectionID))
}

// List metafields for a smart collection
func (s *SmartCollectionClient) ListMetaFields(smartCollectionID int64, options interface{}) ([]MetaField, error) {
	metaFieldService := &MetaFieldClient{client: s.client, resource: smartCollectionsResourceName, resourceID: smartCollectionID}
	return metaFieldService.ListMetaFields(options)
}

// Count metaFields for a smart collection
func (s *SmartCollectionClient) CountMetaFields(smartCollectionID int64, options interface{}) (int, error) {
	metaFieldService := &MetaFieldClient{client: s.client, resource: smartCollectionsResourceName, resourceID: smartCollectionID}
	return metaFieldService.CountMetaFields(options)
}

// Get individual metaField for a smart collection
func (s *SmartCollectionClient) GetMetaField(smartCollectionID int64, metaFieldID int64, options interface{}) (*MetaField, error) {
	metaFieldService := &MetaFieldClient{client: s.client, resource: smartCollectionsResourceName, resourceID: smartCollectionID}
	return metaFieldService.GetMetaFields(metaFieldID, options)
}

// Create a new metaField for a smart collection
func (s *SmartCollectionClient) CreateMetaField(smartCollectionID int64, metaField MetaField) (*MetaField, error) {
	metaFieldService := &MetaFieldClient{client: s.client, resource: smartCollectionsResourceName, resourceID: smartCollectionID}
	return metaFieldService.CreateMetaFields(metaField)
}

// Update an existing metaField for a smart collection
func (s *SmartCollectionClient) UpdateMetaField(smartCollectionID int64, metaField MetaField) (*MetaField, error) {
	metaFieldService := &MetaFieldClient{client: s.client, resource: smartCollectionsResourceName, resourceID: smartCollectionID}
	return metaFieldService.UpdateMetaFields(metaField)
}

// // Delete an existing metaField for a smart collection
func (s *SmartCollectionClient) DeleteMetaField(smartCollectionID int64, metaFieldID int64) error {
	metaFieldService := &MetaFieldClient{client: s.client, resource: smartCollectionsResourceName, resourceID: smartCollectionID}
	return metaFieldService.DeleteMetaFields(metaFieldID)
}
