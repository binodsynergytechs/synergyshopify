package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const variantsBasePath = "variants"
const variantsResourceName = "variants"

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

	// MetaFieldsRepository used for Variant resource to communicate with MetaFields resource
	MetaFieldRepository
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
	SKU                  string           `json:"sku,omitempty"`
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
	AdminGraphqlApiId    string           `json:"admin_graphql_api_id,omitempty"`
	MetaFields           []MetaField      `json:"metaFields,omitempty"`
	PresentmentPrices    []Price          `json:"presentment_prices"`               //TODO: New Field Added In Shopify
	OldInventoryQuantity int              `json:"old_inventory_quantity,omitempty"` //FIXME: deprecated field
	RequireShipping      bool             `json:"requires_shipping"`                //FIXME: deprecated field
}

// TODO: New Struct Added In Shopify
type Price struct {
	CurrencyCode string `json:"currency_code"`
	Amount       string `json:"amount"`
}

// VariantResource represents the result from the variants/X.json endpoint
type VariantResource struct {
	Variant *Variant `json:"variant"`
}

// VariantsResource represents the result from the products/X/variants.json endpoint
type VariantsResource struct {
	Variants []Variant `json:"variants"`
}

// List variants
func (vs *VariantClient) ListVariant(productID int64, options interface{}) ([]Variant, error) {
	path := fmt.Sprintf("%s/%d/variants.json", productsBasePath, productID)
	resource := new(VariantsResource)
	err := vs.client.Get(path, resource, options)
	return resource.Variants, err
}

// Count variants
func (vs *VariantClient) CountVariant(productID int64, options interface{}) (int, error) {
	path := fmt.Sprintf("%s/%d/variants/count.json", productsBasePath, productID)
	return vs.client.Count(path, options)
}

// Get individual variant
func (vs *VariantClient) GetVariant(variantID int64, options interface{}) (*Variant, error) {
	path := fmt.Sprintf("%s/%d.json", variantsBasePath, variantID)
	resource := new(VariantResource)
	err := vs.client.Get(path, resource, options)
	return resource.Variant, err
}

// Create a new variant
func (vs *VariantClient) CreateVariant(productID int64, variant Variant) (*Variant, error) {
	path := fmt.Sprintf("%s/%d/variants.json", productsBasePath, productID)
	wrappedData := VariantResource{Variant: &variant}
	resource := new(VariantResource)
	err := vs.client.Post(path, wrappedData, resource)
	return resource.Variant, err
}

// Update existing variant
func (vs *VariantClient) UpdateVariant(variant Variant) (*Variant, error) {
	path := fmt.Sprintf("%s/%d.json", variantsBasePath, variant.ID)
	wrappedData := VariantResource{Variant: &variant}
	resource := new(VariantResource)
	err := vs.client.Put(path, wrappedData, resource)
	return resource.Variant, err
}

// Delete an existing variant
func (vs *VariantClient) DeleteVariant(productID int64, variantID int64) error {
	return vs.client.Delete(fmt.Sprintf("%s/%d/variants/%d.json", productsBasePath, productID, variantID))
}

// ListMetaFields for a variant
func (vs *VariantClient) ListMetaFields(variantID int64, options interface{}) ([]MetaField, error) {
	metaFieldService := &MetaFieldClient{client: vs.client, resource: variantsResourceName, resourceID: variantID}
	return metaFieldService.ListMetaFields(options)
}

// CountMetaFields for a variant
func (vs *VariantClient) CountMetaFields(variantID int64, options interface{}) (int, error) {
	metaFieldService := &MetaFieldClient{client: vs.client, resource: variantsResourceName, resourceID: variantID}
	return metaFieldService.CountMetaFields(options)
}

// GetMetaField for a variant
func (vs *VariantClient) GetMetaField(variantID int64, metaFieldID int64, options interface{}) (*MetaField, error) {
	metaFieldService := &MetaFieldClient{client: vs.client, resource: variantsResourceName, resourceID: variantID}
	return metaFieldService.GetMetaFields(metaFieldID, options)
}

// CreateMetaField for a variant
func (vs *VariantClient) CreateMetaField(variantID int64, metaField MetaField) (*MetaField, error) {
	metaFieldService := &MetaFieldClient{client: vs.client, resource: variantsResourceName, resourceID: variantID}
	return metaFieldService.CreateMetaFields(metaField)
}

// UpdateMetaField for a variant
func (vs *VariantClient) UpdateMetaField(variantID int64, metaField MetaField) (*MetaField, error) {
	metaFieldService := &MetaFieldClient{client: vs.client, resource: variantsResourceName, resourceID: variantID}
	return metaFieldService.UpdateMetaFields(metaField)
}

// DeleteMetaField for a variant
func (vs *VariantClient) DeleteMetaField(variantID int64, metaFieldID int64) error {
	metaFieldService := &MetaFieldClient{client: vs.client, resource: variantsResourceName, resourceID: variantID}
	return metaFieldService.DeleteMetaFields(metaFieldID)
}
