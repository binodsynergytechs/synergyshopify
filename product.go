package goshopify

import (
	"fmt"
	"regexp"
	"time"
)

const productsBasePath = "products"
const MultipleProductsResponseName = "products"

// linkRegex is used to extract pagination links from product search results.
var linkRegex = regexp.MustCompile(`^ *<([^>]+)>; rel="(previous|next)" *$`)

// ProductService is an interface for interfacing with the product endpoints
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

	// MetaFieldsService used for Product resource to communicate with Metafields resource
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
	TemplateSuffix                 string          `json:"template_suffix,omitempty"`
	Images                         []Image         `json:"images,omitempty"`                            // FIXME: Not Available In Latest Shopify Update
	MetaFieldsGlobalTitleTag       string          `json:"metafields_global_title_tag,omitempty"`       // FIXME: Not Available In Latest Shopify Update
	MetaFieldsGlobalDescriptionTag string          `json:"metafields_global_description_tag,omitempty"` // FIXME: Not Available In Latest Shopify Update
	MetaFields                     []MetaField     `json:"metafields,omitempty"`                        // FIXME: Not Available In Latest Shopify Update
	AdminGraphqlApiId              string          `json:"admin_graphql_api_id,omitempty"`              // FIXME: Not Available In Latest Shopify Update
}

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
func (pc *ProductClient) List(options interface{}) ([]Product, error) {
	products, _, err := pc.ListWithPagination(options)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// ListWithPagination lists products and return pagination to retrieve next/previous results.
func (pc *ProductClient) ListWithPagination(options interface{}) ([]Product, *Pagination, error) {
	path := fmt.Sprintf("%s.json", productsBasePath)
	resource := new(MultipleProductsResponse)

	pagination, err := pc.client.ListWithPagination(path, resource, options)
	if err != nil {
		return nil, nil, err
	}

	return resource.Products, pagination, nil
}

// Count products
func (pc *ProductClient) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", productsBasePath)
	return pc.client.Count(path, options)
}

// Get individual product
func (pc *ProductClient) Get(productID int64, options interface{}) (*Product, error) {
	path := fmt.Sprintf("%s/%d.json", productsBasePath, productID)
	resource := new(SingleProductResponse)
	err := pc.client.Get(path, resource, options)
	return resource.Product, err
}

// Create a new product
func (pc *ProductClient) Create(product Product) (*Product, error) {
	path := fmt.Sprintf("%s.json", productsBasePath)
	wrappedData := SingleProductResponse{Product: &product}
	resource := new(SingleProductResponse)
	err := pc.client.Post(path, wrappedData, resource)
	return resource.Product, err
}

// Update an existing product
func (pc *ProductClient) Update(product Product) (*Product, error) {
	path := fmt.Sprintf("%s/%d.json", productsBasePath, product.ID)
	wrappedData := SingleProductResponse{Product: &product}
	resource := new(SingleProductResponse)
	err := pc.client.Put(path, wrappedData, resource)
	return resource.Product, err
}

// Delete an existing product
func (pc *ProductClient) Delete(productID int64) error {
	return pc.client.Delete(fmt.Sprintf("%s/%d.json", productsBasePath, productID))
}

// ListMetafields for a product
func (pc *ProductClient) ListMetafields(productID int64, options interface{}) ([]MetaField, error) {
	metafieldService := &MetaFieldClient{client: pc.client, resource: MultipleProductsResponseName, resourceID: productID}
	return metafieldService.List(options)
}

// Count metafields for a product
func (pc *ProductClient) CountMetafields(productID int64, options interface{}) (int, error) {
	metafieldService := &MetaFieldClient{client: pc.client, resource: MultipleProductsResponseName, resourceID: productID}
	return metafieldService.Count(options)
}

// GetMetafield for a product
func (pc *ProductClient) GetMetafield(productID int64, metafieldID int64, options interface{}) (*MetaField, error) {
	metafieldService := &MetaFieldClient{client: pc.client, resource: MultipleProductsResponseName, resourceID: productID}
	return metafieldService.Get(metafieldID, options)
}

// CreateMetafield for a product
func (pc *ProductClient) CreateMetafield(productID int64, metafield MetaField) (*MetaField, error) {
	metafieldService := &MetaFieldClient{client: pc.client, resource: MultipleProductsResponseName, resourceID: productID}
	return metafieldService.Create(metafield)
}

// UpdateMetafield for a product
func (pc *ProductClient) UpdateMetafield(productID int64, metafield MetaField) (*MetaField, error) {
	metafieldService := &MetaFieldClient{client: pc.client, resource: MultipleProductsResponseName, resourceID: productID}
	return metafieldService.Update(metafield)
}

// DeleteMetafield for a product
func (pc *ProductClient) DeleteMetafield(productID int64, metafieldID int64) error {
	metafieldService := &MetaFieldClient{client: pc.client, resource: MultipleProductsResponseName, resourceID: productID}
	return metafieldService.Delete(metafieldID)
}
