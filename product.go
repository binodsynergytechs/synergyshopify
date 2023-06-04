package goshopify

import (
	"fmt"
	"regexp"
	"time"
)

const productsBasePath = "products"
const MultipleproductsResponseName = "products"

// linkRegex is used to extract pagination links from product search results.
var linkRegex = regexp.MustCompile(`^ *<([^>]+)>; rel="(previous|next)" *$`)

// ProductRepository is an interface for interfacing with the product endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/product
type ProductRepository interface {
	ListProduct(interface{}) ([]Product, error)
	ListProductWithPagination(interface{}) ([]Product, *Pagination, error)
	CountProduct(interface{}) (int, error)
	GetProduct(int64, interface{}) (*Product, error)
	CreateProduct(Product) (*Product, error)
	UpdateProduct(Product) (*Product, error)
	DeleteProduct(int64) error
	MetaFieldsRepository
}

// ProductClient handles communication with the product related methods of
// the Shopify API.
type ProductClient struct {
	client *Client
}

// Product represents a Shopify product
type Product struct {
	ID                             int64           `json:"id,omitempty"`
	Title                          string          `json:"title,omitempty"`
	BodyHTML                       string          `json:"body_html,omitempty"`
	Vendor                         string          `json:"vendor,omitempty"`
	ProductType                    string          `json:"product_type,omitempty"`
	Handle                         string          `json:"handle,omitempty"`
	CreatedAt                      *time.Time      `json:"created_at,omitempty"`
	UpdatedAt                      *time.Time      `json:"updated_at,omitempty"`
	PublishedAt                    *time.Time      `json:"published_at,omitempty"`
	PublishedScope                 string          `json:"published_scope,omitempty"`
	Tags                           string          `json:"tags,omitempty"`
	Status                         string          `json:"status,omitempty"`
	Options                        []ProductOption `json:"options,omitempty"`
	Variants                       []Variant       `json:"variants,omitempty"`
	Image                          Image           `json:"image,omitempty"`
	Images                         []Image         `json:"images,omitempty"`
	TemplateSuffix                 string          `json:"template_suffix,omitempty"`
	MetaFieldsGlobalTitleTag       string          `json:"metafields_global_title_tag,omitempty"`
	MetaFieldsGlobalDescriptionTag string          `json:"metafields_global_description_tag,omitempty"`
	MetaFields                     []MetaField     `json:"metafields,omitempty"`
	AdminGraphqlApiID              string          `json:"admin_graphql_api_id,omitempty"`
}

// The options provided by Shopify
type ProductOption struct {
	ID        int64    `json:"id,omitempty"`
	ProductID int64    `json:"product_id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Position  int      `json:"position,omitempty"`
	Values    []string `json:"values,omitempty"`
}

type ProductListOptions struct {
	ListOptions
	CollectionID          int64     `url:"collection_id,omitempty"`
	ProductType           string    `url:"product_type,omitempty"`
	Vendor                string    `url:"vendor,omitempty"`
	Handle                string    `url:"handle,omitempty"`
	PublishedAtMin        time.Time `url:"published_at_min,omitempty"`
	PublishedAtMax        time.Time `url:"published_at_max,omitempty"`
	PublishedStatus       string    `url:"published_status,omitempty"`
	PresentmentCurrencies string    `url:"presentment_currencies,omitempty"`
}

// Represents the result from the products/X.json endpoint
type SingleProductResponse struct {
	Product *Product `json:"product"`
}

// Represents the result from the products.json endpoint
type MultipleProductsResponse struct {
	Products []Product `json:"products"`
}

// Pagination of results
type Pagination struct {
	NextPageOptions     *ListOptions
	PreviousPageOptions *ListOptions
}

// List products
func (s *ProductClient) ListProduct(options interface{}) ([]Product, error) {
	products, _, err := s.ListProductWithPagination(options)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// ListWithPagination lists products and return pagination to retrieve next/previous results.
func (s *ProductClient) ListProductWithPagination(options interface{}) ([]Product, *Pagination, error) {
	path := fmt.Sprintf("%s.json", productsBasePath)
	resource := new(MultipleProductsResponse)

	pagination, err := s.client.ListWithPagination(path, resource, options)
	if err != nil {
		return nil, nil, err
	}

	return resource.Products, pagination, nil
}

// Count products
func (s *ProductClient) CountProduct(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", productsBasePath)
	return s.client.Count(path, options)
}

// Get individual product
func (s *ProductClient) GetProduct(productID int64, options interface{}) (*Product, error) {
	path := fmt.Sprintf("%s/%d.json", productsBasePath, productID)
	resource := new(SingleProductResponse)
	err := s.client.Get(path, resource, options)
	return resource.Product, err
}

// Create a new product
func (s *ProductClient) CreateProduct(product Product) (*Product, error) {
	path := fmt.Sprintf("%s.json", productsBasePath)
	wrappedData := SingleProductResponse{Product: &product}
	resource := new(SingleProductResponse)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Product, err
}

// Update an existing product
func (s *ProductClient) UpdateProduct(product Product) (*Product, error) {
	path := fmt.Sprintf("%s/%d.json", productsBasePath, product.ID)
	wrappedData := SingleProductResponse{Product: &product}
	resource := new(SingleProductResponse)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Product, err
}

func (s *ProductClient) DeleteProduct(productID int64) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", productsBasePath, productID))
}

func (s *ProductClient) ListMetaFields(productID int64, options interface{}) ([]MetaField, error) {
	mField := &MetaFieldClient{client: s.client, resource: MultipleproductsResponseName, resourceID: productID}
	return mField.ListMetaField(options)
}

func (s *ProductClient) CountMetaFields(productID int64, options interface{}) (int, error) {
	mField := &MetaFieldClient{client: s.client, resource: MultipleproductsResponseName, resourceID: productID}
	return mField.CountMetaField(options)
}

func (s *ProductClient) GetMetaField(productID int64, metaFieldID int64, options interface{}) (*MetaField, error) {
	mField := &MetaFieldClient{client: s.client, resource: MultipleproductsResponseName, resourceID: productID}
	return mField.GetMetaField(metaFieldID, options)
}

func (s *ProductClient) CreateMetaField(productID int64, metaField MetaField) (*MetaField, error) {
	mField := &MetaFieldClient{client: s.client, resource: MultipleproductsResponseName, resourceID: productID}
	return mField.CreateMetaField(metaField)
}

func (s *ProductClient) UpdateMetaField(productID int64, metaField MetaField) (*MetaField, error) {
	mField := &MetaFieldClient{client: s.client, resource: MultipleproductsResponseName, resourceID: productID}
	return mField.UpdateMetaField(metaField)
}

func (s *ProductClient) DeleteMetaField(productID int64, metaFieldID int64) error {
	mField := &MetaFieldClient{client: s.client, resource: MultipleproductsResponseName, resourceID: productID}
	return mField.DeleteMetaField(metaFieldID)
}
