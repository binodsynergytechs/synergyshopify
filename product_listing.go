package goshopify

import (
	"fmt"
	"time"
)

const productListingBasePath = "product_listings"
const productsListingResourceName = "product_listings"

// ProductListingService is an interface for interfacing with the product listing endpoints
// of the Shopify API.
// See: https://shopify.dev/docs/admin-api/rest/reference/sales-channels/productlisting
type ProductListingService interface {
	List(interface{}) ([]ProductListing, error)
	ListWithPagination(interface{}) ([]ProductListing, *Pagination, error)
	Count(interface{}) (int, error)
	Get(int64, interface{}) (*ProductListing, error)
	GetProductIDs(interface{}) ([]int64, error)
	Publish(int64) (*ProductListing, error)
	Delete(int64) error
}

// ProductListingServiceOp handles communication with the product related methods of
// the Shopify API.
type ProductListingServiceOp struct {
	client *Client
}

// ProductListing represents a Shopify product published to your sales channel app
type ProductListing struct {
	ID          int64           `json:"product_id,omitempty"`
	Title       string          `json:"title,omitempty"`
	BodyHTML    string          `json:"body_html,omitempty"`
	Vendor      string          `json:"vendor,omitempty"`
	ProductType string          `json:"product_type,omitempty"`
	Handle      string          `json:"handle,omitempty"`
	CreatedAt   *time.Time      `json:"created_at,omitempty"`
	UpdatedAt   *time.Time      `json:"updated_at,omitempty"`
	PublishedAt *time.Time      `json:"published_at,omitempty"`
	Tags        string          `json:"tags,omitempty"`
	Options     []ProductOption `json:"options,omitempty"`
	Variants    []Variant       `json:"variants,omitempty"`
	Images      []Image         `json:"images,omitempty"`
}

//TODO: latest from shopify 23/04

//	type ProductListing struct {
//		ProductID                    string                 `json:"product_id"`
//		BodyHTML                     string                 `json:"body_html"`
//		CreatedAt                    string                 `json:"created_at"`
//		Handle                       string                 `json:"handle"`
//		Images                        []string               `json:"images"`
//		Options                       []string               `json:"options"`
//		ProductType                   string                 `json:"product_type"`
//		PublishedAt                   string                 `json:"published_at"`
//		Tags                          []string               `json:"tags"`
//		Title                         string                 `json:"title"`
//		UpdatedAt                     string                 `json:"updated_at"`
//		Variants                      []struct {
//			Barcode                  string                 `json:"barcode"`
//			CompareAtPrice            string                 `json:"compare_at_price"`
//			CreatedAt                 string                 `json:"created_at"`
//			FulfillmentService          string                 `json:"fulfillment_service"`
//			Grams                      int                    `json:"grams"`
//			Weight                     int                    `json:"weight"`
//			WeightUnit                  string                 `json:"weight_unit"`
//			ID                          string                 `json:"id"`
//			InventoryManagement          bool                   `json:"inventory_management"`
//			InventoryPolicy              bool                   `json:"inventory_policy"`
//			InventoryQuantity           int                    `json:"inventory_quantity"`
//			Metafield                   string                 `json:"metafield"`
//			Option                      string                 `json:"option"`
//			Position                     int                    `json:"position"`
//			Price                        float64                `json:"price"`
//			ProductID                    string                 `json:"product_id"`
//			RequiresShipping              bool                   `json:"requires_shipping"`
//			Sku                          string                 `json:"sku"`
//			SKU                          string                 `json:"sku"`
//			Taxable                      bool                   `json:"taxable"`
//			Title                         string                 `json:"title"`
//			UpdatedAt                     string                 `json:"updated_at"`
//			Vendor                        string                 `json:"vendor"`
//		} `json:"variants"`
//	}
//
// Represents the result from the product_listings/X.json endpoint
type ProductListingResource struct {
	ProductListing *ProductListing `json:"product_listing"`
}

// Represents the result from the product_listings.json endpoint
type ProductsListingsResource struct {
	ProductListings []ProductListing `json:"product_listings"`
}

// Represents the result from the product_listings/product_ids.json endpoint
type ProductListingIDsResource struct {
	ProductIDs []int64 `json:"product_ids"`
}

// Resource which create product_listing endpoint expects in request body
// e.g.
// PUT /admin/api/2020-07/product_listings/921728736.json
//
//	{
//	  "product_listing": {
//	    "product_id": 921728736
//	  }
//	}
type ProductListingPublishResource struct {
	ProductListing struct {
		ProductID int64 `json:"product_id"`
	} `json:"product_listing"`
}

// List products
func (s *ProductListingServiceOp) List(options interface{}) ([]ProductListing, error) {
	products, _, err := s.ListWithPagination(options)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// ListWithPagination lists products and return pagination to retrieve next/previous results.
func (s *ProductListingServiceOp) ListWithPagination(options interface{}) ([]ProductListing, *Pagination, error) {
	path := fmt.Sprintf("%s.json", productListingBasePath)
	resource := new(ProductsListingsResource)

	pagination, err := s.client.ListWithPagination(path, resource, options)
	if err != nil {
		return nil, nil, err
	}

	return resource.ProductListings, pagination, nil
}

// Count products listings published to your sales channel app
func (s *ProductListingServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", productListingBasePath)
	return s.client.Count(path, options)
}

// Get individual product_listing by product ID
func (s *ProductListingServiceOp) Get(productID int64, options interface{}) (*ProductListing, error) {
	path := fmt.Sprintf("%s/%d.json", productListingBasePath, productID)
	resource := new(ProductListingResource)
	err := s.client.Get(path, resource, options)
	return resource.ProductListing, err
}

// GetProductIDs lists all product IDs that are published to your sales channel
func (s *ProductListingServiceOp) GetProductIDs(options interface{}) ([]int64, error) {
	path := fmt.Sprintf("%s/product_ids.json", productListingBasePath)
	resource := new(ProductListingIDsResource)
	err := s.client.Get(path, resource, options)
	return resource.ProductIDs, err
}

// Publish an existing product listing to your sales channel app
func (s *ProductListingServiceOp) Publish(productID int64) (*ProductListing, error) {
	path := fmt.Sprintf("%s/%v.json", productListingBasePath, productID)
	wrappedData := new(ProductListingPublishResource)
	wrappedData.ProductListing.ProductID = productID
	resource := new(ProductListingResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.ProductListing, err
}

// Delete unpublishes an existing product from your sales channel app.
func (s *ProductListingServiceOp) Delete(productID int64) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", productListingBasePath, productID))
}
