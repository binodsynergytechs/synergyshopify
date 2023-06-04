package goshopify

import (
	"fmt"
	"time"
)

const pagesBasePath = "pages"
const pagesResourceName = "pages"

// PagesPageRepository is an interface for interacting with the pages
// endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/online_store/page
type PageRepository interface {
	ListPage(interface{}) ([]Page, error)
	CountPage(interface{}) (int, error)
	GetPage(int64, interface{}) (*Page, error)
	CreatePage(Page) (*Page, error)
	UpdatePage(Page) (*Page, error)
	DeletePage(int64) error

	// MetafieldsService used for Pages resource to communicate with Metafields
	// resource
	MetaFieldsRepository
}

// PageClient handles communication with the page related methods of the
// Shopify API.
type PageClient struct {
	client *Client
}

// Page represents a Shopify page.
type Page struct {
	ID             int64       `json:"id,omitempty"`
	Author         string      `json:"author,omitempty"`
	Handle         string      `json:"handle,omitempty"`
	Title          string      `json:"title,omitempty"`
	CreatedAt      *time.Time  `json:"created_at,omitempty"`
	UpdatedAt      *time.Time  `json:"updated_at,omitempty"`
	BodyHTML       string      `json:"body_html,omitempty"`
	TemplateSuffix string      `json:"template_suffix,omitempty"`
	PublishedAt    *time.Time  `json:"published_at,omitempty"`
	ShopID         int64       `json:"shop_id,omitempty"`
	MetaFields     []MetaField `json:"metafields,omitempty"`
}

// PageResource represents the result from the pages/X.json endpoint
type PageResource struct {
	Page *Page `json:"page"`
}

// PagesResource represents the result from the pages.json endpoint
type PagesResource struct {
	Pages []Page `json:"pages"`
}

// List pages
func (pc *PageClient) ListPage(options interface{}) ([]Page, error) {
	path := fmt.Sprintf("%s.json", pagesBasePath)
	resource := new(PagesResource)
	err := pc.client.Get(path, resource, options)
	return resource.Pages, err
}

// Count pages
func (pc *PageClient) CountPage(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", pagesBasePath)
	return pc.client.Count(path, options)
}

// Get individual page
func (pc *PageClient) GetPage(pageID int64, options interface{}) (*Page, error) {
	path := fmt.Sprintf("%s/%d.json", pagesBasePath, pageID)
	resource := new(PageResource)
	err := pc.client.Get(path, resource, options)
	return resource.Page, err
}

// Create a new page
func (pc *PageClient) CreatePage(page Page) (*Page, error) {
	path := fmt.Sprintf("%s.json", pagesBasePath)
	wrappedData := PageResource{Page: &page}
	resource := new(PageResource)
	err := pc.client.Post(path, wrappedData, resource)
	return resource.Page, err
}

// Update an existing page
func (pc *PageClient) UpdatePage(page Page) (*Page, error) {
	path := fmt.Sprintf("%s/%d.json", pagesBasePath, page.ID)
	wrappedData := PageResource{Page: &page}
	resource := new(PageResource)
	err := pc.client.Put(path, wrappedData, resource)
	return resource.Page, err
}

// Delete an existing page.
func (pc *PageClient) DeletePage(pageID int64) error {
	return pc.client.Delete(fmt.Sprintf("%s/%d.json", pagesBasePath, pageID))
}

// List metaFields for a page
func (pc *PageClient) ListMetaFields(pageID int64, options interface{}) ([]MetaField, error) {
	metaFieldService := &MetaFieldClient{client: pc.client, resource: pagesResourceName, resourceID: pageID}
	return metaFieldService.ListMetaField(options)
}

// Count metaFields for a page
func (pc *PageClient) CountMetaFields(pageID int64, options interface{}) (int, error) {
	metaFieldService := &MetaFieldClient{client: pc.client, resource: pagesResourceName, resourceID: pageID}
	return metaFieldService.CountMetaField(options)
}

// Get individual metafield for a page
func (pc *PageClient) GetMetaField(pageID int64, metaFieldID int64, options interface{}) (*MetaField, error) {
	metaFieldService := &MetaFieldClient{client: pc.client, resource: pagesResourceName, resourceID: pageID}
	return metaFieldService.GetMetaField(metaFieldID, options)
}

// Create a new metafield for a page
func (pc *PageClient) CreateMetaField(pageID int64, metaField MetaField) (*MetaField, error) {
	metaFieldService := &MetaFieldClient{client: pc.client, resource: pagesResourceName, resourceID: pageID}
	return metaFieldService.CreateMetaField(metaField)
}

// Update an existing metafield for a page
func (pc *PageClient) UpdateMetaField(pageID int64, metaField MetaField) (*MetaField, error) {
	metaFieldService := &MetaFieldClient{client: pc.client, resource: pagesResourceName, resourceID: pageID}
	return metaFieldService.UpdateMetaField(metaField)
}

// Delete an existing metafield for a page
func (pc *PageClient) DeleteMetaField(pageID int64, metaFieldID int64) error {
	metaFieldService := &MetaFieldClient{client: pc.client, resource: pagesResourceName, resourceID: pageID}
	return metaFieldService.DeleteMetaField(metaFieldID)
}
