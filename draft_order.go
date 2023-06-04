package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const (
	draftOrdersBasePath     = "draft_orders"
	draftOrdersResourceName = "draft_orders"
)

// DraftOrderRepository is an interface for interfacing with the draft orders endpoints of
// the Shopify API.
// See: https://help.shopify.com/api/reference/orders/draftorder
type DraftOrderRepository interface {
	ListDraftOrder(interface{}) ([]DraftOrder, error)
	CountDraftOrder(interface{}) (int, error)
	GetDraftOrder(int64, interface{}) (*DraftOrder, error)
	CreateDraftOrder(DraftOrder) (*DraftOrder, error)
	UpdateDraftOrder(DraftOrder) (*DraftOrder, error)
	DeleteDraftOrder(int64) error
	InvoiceDraftOrder(int64, DraftOrderInvoice) (*DraftOrderInvoice, error)
	CompleteDraftOrder(int64, bool) (*DraftOrder, error)

	// MetaFieldsRepository used for DrafT Order resource to communicate with MetaFields resource
	MetaFieldRepository
}

// DraftOrderClient handles communication with the draft order related methods of the
// Shopify API.
type DraftOrderClient struct {
	client *Client
}

// DraftOrder represents a shopify draft order
type DraftOrder struct {
	ID              int64            `json:"id,omitempty"`
	OrderID         int64            `json:"order_id,omitempty"`
	Name            string           `json:"name,omitempty"`
	Customer        *Customer        `json:"customer,omitempty"`
	ShippingAddress *Address         `json:"shipping_address,omitempty"`
	BillingAddress  *Address         `json:"billing_address,omitempty"`
	Note            string           `json:"note,omitempty"`
	NoteAttributes  []NoteAttribute  `json:"note_attributes,omitempty"`
	Email           string           `json:"email,omitempty"`
	Currency        string           `json:"currency,omitempty"`
	InvoiceSentAt   *time.Time       `json:"invoice_sent_at,omitempty"`
	InvoiceURL      string           `json:"invoice_url,omitempty"`
	PaymentTerm     interface{}      `json:"payment_terms"` // TODO: Latest Field Available In Model 23/04
	SourceName      string           `json:"source_name"`   // TODO: Latest Field Available In Model 23/04
	LineItems       []LineItem       `json:"line_items,omitempty"`
	ShippingLine    *ShippingLines   `json:"shipping_line,omitempty"`
	Tags            string           `json:"tags,omitempty"`
	TaxLines        []TaxLine        `json:"tax_lines,omitempty"`
	AppliedDiscount *AppliedDiscount `json:"applied_discount,omitempty"`
	TaxesIncluded   bool             `json:"taxes_included,omitempty"`
	TotalTax        string           `json:"total_tax,omitempty"`
	TaxExempt       *bool            `json:"tax_exempt,omitempty"`
	TaxExemptions   []string         `json:"tax_exemptions"` // TODO: Latest Field Available In Model 23/04
	TotalPrice      string           `json:"total_price,omitempty"`
	SubtotalPrice   *decimal.Decimal `json:"subtotal_price,omitempty"`
	CompletedAt     *time.Time       `json:"completed_at,omitempty"`
	CreatedAt       *time.Time       `json:"created_at,omitempty"`
	UpdatedAt       *time.Time       `json:"updated_at,omitempty"`
	Status          string           `json:"status,omitempty"`
	// only in request to flag using the customer's default address
	UseCustomerDefaultAddress bool `json:"use_customer_default_address,omitempty"`
}

// type DraftOrder struct {
// 	AppliedDiscount string                 `json:"appliedDiscount"`
// 	BillingAddress   MailingAddress         `json:"billingAddress"`
// 	BillingAddressMatchesShippingAddress bool   `json:"billingAddressMatchesShippingAddress"`
// 	CompletedAt       time.Time            `json:"completedAt"`
// 	CreatedAt         time.Time            `json:"createdAt"`
// 	CurrencyCode      string                `json:"currencyCode"`
// 	CustomAttributes   []Attribute           `json:"customAttributes"`
// 	Customer           Customer               `json:"customer"`
// 	DraftOrderID       ID                     `json:"id"`
// 	Email              string                `json:"email"`
// 	HasTimelineComment  bool                   `json:"hasTimelineComment"`
// 	Id                  ID                     `json:"id"`
// 	InvoiceEmailTemplateSubject string              `json:"invoiceEmailTemplateSubject"`
// 	InvoiceSentAt       time.Time            `json:"invoiceSentAt"`
// 	InvoiceUrl          URL                     `json:"invoiceUrl"`
// 	LegacyResourceID    UnsignedInt64           `json:"legacyResourceId"`
// 	LineItemsSubtotalPrice MoneyBag                `json:"lineItemsSubtotalPrice"`
// 	MarketName          string                `json:"marketName"`
// 	MarketRegionCountryCode CountryCode             `json:"marketRegionCountryCode"`
// 	MetaData              map[string]interface{} `json:"metaData"`
// 	Name                  string                `json:"name"`
// 	Note2                 string                `json:"note2"`
// 	Order                Order                   `json:"order"`
// 	PaymentTerms           PaymentTerms           `json:"paymentTerms"`
// 	Phone                 string                `json:"phone"`
// 	PresentmentCurrencyCode CurrencyCode             `json:"presentmentCurrencyCode"`
// 	PurchasingEntity       PurchasingEntity       `json:"purchasingEntity"`
// 	Ready                 bool                   `json:"ready"`
// 	ReserveInventoryUntil time.Time            `json:"reserveInventoryUntil"`
// 	ShippingAddress       MailingAddress         `json:"shippingAddress"`
// 	ShippingLine          ShippingLine           `json:"shippingLine"`
// 	Status                 DraftOrderStatus        `json:"status"`
// 	SubtotalPrice         Money                     `json:"subtotalPrice"`
// 	SubtotalPriceSet        MoneyBag               `json:"subtotalPriceSet"`
// 	TotalDiscountsSet        MoneyBag               `json:"totalDiscountsSet"`
// 	TotalPrice             Money                     `json:"totalPrice"`
// 	TotalPriceSet           MoneyBag               `json:"totalPriceSet"`
// 	taxesIncluded         bool                   `json:"taxesIncluded"`
// 	taxLines             []TaxLine               `json:"taxLines"`
// 	TaxesIncluded         bool                   `json:"taxesIncluded"`
// 	TotalLineItemsPriceSet MoneyBag                `json:"totalLineItemsPriceSet"`
// 	TotalPrice               Money                     `json:"totalPrice"`
// 	TotalPriceSet           MoneyBag               `json:"totalPriceSet"`
// }

// TODO: latest from shopify 23/04

// type MoneyBag struct {
// 	Amount float64 `json:"amount"`
// 	Currency string  `json:"currency"`
// }

// TODO: latest from shopify 23/04
// type MoneyV2 struct {
// 	Amount float64 `json:"amount"`
// }

// TODO: latest from shopify 23/04
// type DraftOrderAppliedDiscount struct {
// 	Amount   MoneyBag `json:"amountSet"`
// 	AmountV2 MoneyV2 `json:"amountV2"`
// 	Description string `json:"description"`
// 	Title       string `json:"title"`
// 	Value       float64 `json:"value"`
// 	ValueType   string `json:"valueType"`
// }

// AppliedDiscount is the discount applied to the line item or the draft order object.
type AppliedDiscount struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Value       string `json:"value,omitempty"`
	ValueType   string `json:"value_type,omitempty"`
	Amount      string `json:"amount,omitempty"`
}

// DraftOrderInvoice is the struct used to create an invoice for a draft order
type DraftOrderInvoice struct {
	To            string   `json:"to,omitempty"`
	From          string   `json:"from,omitempty"`
	Subject       string   `json:"subject,omitempty"`
	CustomMessage string   `json:"custom_message,omitempty"`
	Bcc           []string `json:"bcc,omitempty"`
}

type DraftOrdersResource struct {
	DraftOrders []DraftOrder `json:"draft_orders"`
}

type DraftOrderResource struct {
	DraftOrder *DraftOrder `json:"draft_order"`
}

type DraftOrderInvoiceResource struct {
	DraftOrderInvoice *DraftOrderInvoice `json:"draft_order_invoice,omitempty"`
}

// DraftOrderListOptions represents the possible options that can be used
// to further query the list draft orders endpoint
type DraftOrderListOptions struct {
	Fields       string     `url:"fields,omitempty"`
	Limit        int        `url:"limit,omitempty"`
	SinceID      int64      `url:"since_id,omitempty"`
	UpdatedAtMin *time.Time `url:"updated_at_min,omitempty"`
	UpdatedAtMax *time.Time `url:"updated_at_max,omitempty"`
	IDs          string     `url:"ids,omitempty"`
	Status       string     `url:"status,omitempty"`
}

// DraftOrderCountOptions represents the possible options to the count draft orders endpoint
type DraftOrderCountOptions struct {
	Fields  string `url:"fields,omitempty"`
	Limit   int    `url:"limit,omitempty"`
	SinceID int64  `url:"since_id,omitempty"`
	IDs     string `url:"ids,omitempty"`
	Status  string `url:"status,omitempty"`
}

// Create draft order
func (s *DraftOrderClient) CreateDraftOrder(draftOrder DraftOrder) (*DraftOrder, error) {
	path := fmt.Sprintf("%s.json", draftOrdersBasePath)
	wrappedData := DraftOrderResource{DraftOrder: &draftOrder}
	resource := new(DraftOrderResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.DraftOrder, err
}

// List draft orders
func (s *DraftOrderClient) ListDraftOrder(options interface{}) ([]DraftOrder, error) {
	path := fmt.Sprintf("%s.json", draftOrdersBasePath)
	resource := new(DraftOrdersResource)
	err := s.client.Get(path, resource, options)
	return resource.DraftOrders, err
}

// Count draft orders
func (s *DraftOrderClient) CountDraftOrder(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", draftOrdersBasePath)
	return s.client.Count(path, options)
}

// Delete draft orders
func (s *DraftOrderClient) DeleteDraftOrder(draftOrderID int64) error {
	path := fmt.Sprintf("%s/%d.json", draftOrdersBasePath, draftOrderID)
	return s.client.Delete(path)
}

// Invoice a draft order
func (s *DraftOrderClient) InvoiceDraftOrder(draftOrderID int64, draftOrderInvoice DraftOrderInvoice) (*DraftOrderInvoice, error) {
	path := fmt.Sprintf("%s/%d/send_invoice.json", draftOrdersBasePath, draftOrderID)
	wrappedData := DraftOrderInvoiceResource{DraftOrderInvoice: &draftOrderInvoice}
	resource := new(DraftOrderInvoiceResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.DraftOrderInvoice, err
}

// Get individual draft order
func (s *DraftOrderClient) GetDraftOrder(draftOrderID int64, options interface{}) (*DraftOrder, error) {
	path := fmt.Sprintf("%s/%d.json", draftOrdersBasePath, draftOrderID)
	resource := new(DraftOrderResource)
	err := s.client.Get(path, resource, options)
	return resource.DraftOrder, err
}

// Update draft order
func (s *DraftOrderClient) UpdateDraftOrder(draftOrder DraftOrder) (*DraftOrder, error) {
	path := fmt.Sprintf("%s/%d.json", draftOrdersBasePath, draftOrder.ID)
	wrappedData := DraftOrderResource{DraftOrder: &draftOrder}
	resource := new(DraftOrderResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.DraftOrder, err
}

// Complete draft order
func (s *DraftOrderClient) CompleteDraftOrder(draftOrderID int64, paymentPending bool) (*DraftOrder, error) {
	path := fmt.Sprintf("%s/%d/complete.json?payment_pending=%t", draftOrdersBasePath, draftOrderID, paymentPending)
	resource := new(DraftOrderResource)
	err := s.client.Put(path, nil, resource)
	return resource.DraftOrder, err
}

// List metaFields for an order
func (s *DraftOrderClient) ListMetaFields(draftOrderID int64, options interface{}) ([]MetaField, error) {
	metaFieldService := &MetaFieldClient{client: s.client, resource: draftOrdersResourceName, resourceID: draftOrderID}
	return metaFieldService.ListMetaFields(options)
}

// Count metaFields for an order
func (s *DraftOrderClient) CountMetaFields(draftOrderID int64, options interface{}) (int, error) {
	metaFieldService := &MetaFieldClient{client: s.client, resource: draftOrdersResourceName, resourceID: draftOrderID}
	return metaFieldService.CountMetaFields(options)
}

// Get individual metaField for an order
func (s *DraftOrderClient) GetMetaField(draftOrderID int64, metaFieldID int64, options interface{}) (*MetaField, error) {
	metaFieldService := &MetaFieldClient{client: s.client, resource: draftOrdersResourceName, resourceID: draftOrderID}
	return metaFieldService.GetMetaFields(metaFieldID, options)
}

// Create a new metaField for an order
func (s *DraftOrderClient) CreateMetaField(draftOrderID int64, metaField MetaField) (*MetaField, error) {
	metaFieldService := &MetaFieldClient{client: s.client, resource: draftOrdersResourceName, resourceID: draftOrderID}
	return metaFieldService.CreateMetaFields(metaField)
}

// Update an existing metaField for an order
func (s *DraftOrderClient) UpdateMetaField(draftOrderID int64, metaField MetaField) (*MetaField, error) {
	metaFieldService := &MetaFieldClient{client: s.client, resource: draftOrdersResourceName, resourceID: draftOrderID}
	return metaFieldService.UpdateMetaFields(metaField)
}

// Delete an existing metaField for an order
func (s *DraftOrderClient) DeleteMetaField(draftOrderID int64, metaFieldID int64) error {
	metaFieldService := &MetaFieldClient{client: s.client, resource: draftOrdersResourceName, resourceID: draftOrderID}
	return metaFieldService.DeleteMetaFields(metaFieldID)
}
