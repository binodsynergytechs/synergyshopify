package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const variantsBasePath = "variants"
const MultipleVariantsResponseName = "variants"

// VariantRepository is an interface for interacting with the variant endpoints
// of the Shopify API.
// See https://help.shopify.com/api/reference/product_variant
type VariantRepository interface {
	ListVariant(int64, interface{}) ([]Variant, error)
	CountVariant(int64, interface{}) (int, error)
	GetVariant(int64, interface{}) (*Variant, error)
	CreateVariant(int64, Variant) (*Variant, error)
	UpdateVariant(Variant) (*Variant, error)
	DeleteVariant(int64, int64) error
	// MetaFieldsService used for Variant resource to communicate with MetaFields resource
	ListVariantMetaFields(variantID int64, options interface{}) ([]MetaField, error)
	CountVariantMetaFields(variantID int64, options interface{}) (int, error)
	GetVariantMetaField(variantID int64, metaFieldID int64, options interface{}) (*MetaField, error)
	CreateVariantMetaField(variantID int64, metaFielD MetaField) (*MetaField, error)
	UpdateVariantMetaField(variantID int64, metaFielD MetaField) (*MetaField, error)
	DeleteVariantMetaField(variantID int64, metaFieldID int64) error
}

// VariantClient handles communication with the variant related methods of
// the Shopify API.
type VariantClient struct {
	client *Client
}

// Variant represents a Shopify variant
type Variant struct {
	ID                   int64            `json:"id,omitempty"`
	ProductID            int64            `json:"product_id,omitempty"`
	Title                string           `json:"title,omitempty"`
	Sku                  string           `json:"sku,omitempty"`
	Position             int              `json:"position,omitempty"`
	Grams                int              `json:"grams,omitempty"`
	InventoryPolicy      string           `json:"inventory_policy,omitempty"`
	Price                *decimal.Decimal `json:"price,omitempty"`
	CompareAtPrice       *decimal.Decimal `json:"compare_at_price,omitempty"`
	FulfillmentService   string           `json:"fulfillment_service,omitempty"`
	InventoryManagement  string           `json:"inventory_management,omitempty"`
	InventoryItemId      int64            `json:"inventory_item_id,omitempty"`
	Option1              string           `json:"option1,omitempty"`
	Option2              string           `json:"option2,omitempty"`
	Option3              string           `json:"option3,omitempty"`
	CreatedAt            *time.Time       `json:"created_at,omitempty"`
	UpdatedAt            *time.Time       `json:"updated_at,omitempty"`
	Taxable              bool             `json:"taxable,omitempty"`
	TaxCode              string           `json:"tax_code,omitempty"`
	Barcode              string           `json:"barcode,omitempty"`
	ImageID              int64            `json:"image_id,omitempty"`
	InventoryQuantity    int              `json:"inventory_quantity,omitempty"`
	Weight               *decimal.Decimal `json:"weight,omitempty"`
	WeightUnit           string           `json:"weight_unit,omitempty"`
	OldInventoryQuantity int              `json:"old_inventory_quantity,omitempty"`
	RequireShipping      bool             `json:"requires_shipping"`
	AdminGraphqlApiID    string           `json:"admin_graphql_api_id,omitempty"`
	MetaFields           []MetaField      `json:"metafields,omitempty"`
}

// SingleVariantResponse represents the result from the variants/X.json endpoint
type SingleVariantResponse struct {
	Variant *Variant `json:"variant"`
}

// MultipleVariantsResponse represents the result from the products/X/variants.json endpoint
type MultipleVariantsResponse struct {
	Variants []Variant `json:"variants"`
}

// List variants
func (vc *VariantClient) ListVariant(productID int64, options interface{}) ([]Variant, error) {
	path := fmt.Sprintf("%s/%d/variants.json", productsBasePath, productID)
	resource := new(MultipleVariantsResponse)
	err := vc.client.Get(path, resource, options)
	return resource.Variants, err
}

// Count variants
func (vc *VariantClient) CountVariant(productID int64, options interface{}) (int, error) {
	path := fmt.Sprintf("%s/%d/variants/count.json", productsBasePath, productID)
	return vc.client.Count(path, options)
}

// Get individual variant
func (vc *VariantClient) GetVariant(variantID int64, options interface{}) (*Variant, error) {
	path := fmt.Sprintf("%s/%d.json", variantsBasePath, variantID)
	resource := new(SingleVariantResponse)
	err := vc.client.Get(path, resource, options)
	return resource.Variant, err
}

// Create a new variant
func (vc *VariantClient) CreateVariant(productID int64, variant Variant) (*Variant, error) {
	path := fmt.Sprintf("%s/%d/variants.json", productsBasePath, productID)
	wrappedData := SingleVariantResponse{Variant: &variant}
	resource := new(SingleVariantResponse)
	err := vc.client.Post(path, wrappedData, resource)
	return resource.Variant, err
}

// Update existing variant
func (vc *VariantClient) UpdateVariant(variant Variant) (*Variant, error) {
	path := fmt.Sprintf("%s/%d.json", variantsBasePath, variant.ID)
	wrappedData := SingleVariantResponse{Variant: &variant}
	resource := new(SingleVariantResponse)
	err := vc.client.Put(path, wrappedData, resource)
	return resource.Variant, err
}

// Delete an existing variant
func (vc *VariantClient) DeleteVariant(productID int64, variantID int64) error {
	return vc.client.Delete(fmt.Sprintf("%s/%d/variants/%d.json", productsBasePath, productID, variantID))
}

// ListMetafields for a variant
func (vc *VariantClient) ListVariantMetaFields(variantID int64, options interface{}) ([]MetaField, error) {
	metaField := &MetaFieldClient{client: vc.client, resource: MultipleVariantsResponseName, resourceID: variantID}
	return metaField.ListMetaField(options)
}

// CountMetafields for a variant
func (vc *VariantClient) CountVariantMetaFields(variantID int64, options interface{}) (int, error) {
	metaField := &MetaFieldClient{client: vc.client, resource: MultipleVariantsResponseName, resourceID: variantID}
	return metaField.CountMetaField(options)
}

// GetMetafield for a variant
func (vc *VariantClient) GetVariantMetaField(variantID int64, metaFieldID int64, options interface{}) (*MetaField, error) {
	metaField := &MetaFieldClient{client: vc.client, resource: MultipleVariantsResponseName, resourceID: variantID}
	return metaField.GetMetaField(metaFieldID, options)
}

// CreateMetafield for a variant
func (vc *VariantClient) CreateVariantMetaField(variantID int64, metaFielD MetaField) (*MetaField, error) {
	metaField := &MetaFieldClient{client: vc.client, resource: MultipleVariantsResponseName, resourceID: variantID}
	return metaField.CreateMetaField(metaFielD)
}

// UpdateMetafield for a variant
func (vc *VariantClient) UpdateVariantMetaField(variantID int64, metaFielD MetaField) (*MetaField, error) {
	metaField := &MetaFieldClient{client: vc.client, resource: MultipleVariantsResponseName, resourceID: variantID}
	return metaField.UpdateMetaField(metaFielD)
}

// DeleteMetafield for a variant
func (vc *VariantClient) DeleteVariantMetaField(variantID int64, metaFieldID int64) error {
	metaField := &MetaFieldClient{client: vc.client, resource: MultipleVariantsResponseName, resourceID: variantID}
	return metaField.DeleteMetaField(metaFieldID)
}
