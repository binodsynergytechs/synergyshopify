package goshopify

import (
	"fmt"
	"time"
)

// FulfillmentsService is an interface for other Shopify resources
// to interface with the fulfillment endpoints of the Shopify API.
// https://help.shopify.com/api/reference/fulfillment
type FulfillmentRepository interface {
	ListFulfillments(int64, interface{}) ([]Fulfillment, error)
	CountFulfillments(int64, interface{}) (int, error)
	GetFulfillments(int64, int64, interface{}) (*Fulfillment, error)
	CreateFulfillments(int64, Fulfillment) (*Fulfillment, error)
	UpdateFulfillments(int64, Fulfillment) (*Fulfillment, error)
	CompleteFulfillments(int64, int64) (*Fulfillment, error)
	TransitionFulfillments(int64, int64) (*Fulfillment, error)
	CancelFulfillments(int64, int64) (*Fulfillment, error)
}

// FulfillmentClient handles communication with the fulfillment
// related methods of the Shopify API.
type FulfillmentClient struct {
	client     *Client
	resource   string
	resourceID int64
}

// Fulfillment represents a Shopify fulfillment.
type Fulfillment struct {
	ID                         int64      `json:"id,omitempty"`
	OrderID                    int64      `json:"order_id,omitempty"`
	LocationID                 int64      `json:"location_id,omitempty"`
	Status                     string     `json:"status,omitempty"`
	CreatedAt                  *time.Time `json:"created_at,omitempty"`
	Service                    string     `json:"service,omitempty"`
	UpdatedAt                  *time.Time `json:"updated_at,omitempty"`
	TrackingCompany            string     `json:"tracking_company,omitempty"`
	ShipmentStatus             string     `json:"shipment_status,omitempty"`
	TrackingNumber             string     `json:"tracking_number,omitempty"`
	TrackingNumbers            []string   `json:"tracking_numbers,omitempty"`
	TrackingUrl                string     `json:"tracking_url,omitempty"`
	TrackingUrls               []string   `json:"tracking_urls,omitempty"`
	Receipt                    Receipt    `json:"receipt,omitempty"`
	LineItems                  []LineItem `json:"line_items,omitempty"`
	NotifyCustomer             bool       `json:"notify_customer"`
	Name                       string     `json:"name"`                          // TODO: Latest From Shopify Model 23/04
	VariantInventoryManagement string     `jso0n:"variant_inventory_management"` // TODO: Latest From Shopify Model 23/04
	OriginAddress              Address    `json:"origin_address"`                // TODO: Latest From Shopify Model 23/04
}

// Receipt represents a Shopify receipt.
type Receipt struct {
	TestCase      bool   `json:"testcase,omitempty"`
	Authorization string `json:"authorization,omitempty"`
	LocationID    int64  `json:"location_id"`    // TODO: Latest From Shopify Model 23/04
	UserID        int64  `json:"user_id"`        // TODO: Latest From Shopify Model 23/04
	ReceiptType   string `json:"receipt_type"`   // TODO: Latest From Shopify Model 23/04
	ReceiptHeader string `json:"receipt_header"` // TODO: Latest From Shopify Model 23/04
	ReceiptFooter string `json:"receipt_footer"` // TODO: Latest From Shopify Model 23/04
	CreatedAt     string `json:"created_at"`     // TODO: Latest From Shopify Model 23/04
	UpdatedAt     string `json:"updated_at"`     // TODO: Latest From Shopify Model 23/04
	ReceiptNumber string `json:"receipt_number"` // TODO: Latest From Shopify Model 23/04
}

// SingleFulfillmentResponse represents the result from the fulfillments/X.json endpoint
type SingleFulfillmentResponse struct {
	Fulfillment *Fulfillment `json:"fulfillment"`
}

// MultipleFulfillmentsResponse represents the result from the fullfilmentfc.json endpoint
type MultipleFulfillmentsResponse struct {
	Fulfillments []Fulfillment `json:"fulfillments"`
}

// List fulfillments
func (fc *FulfillmentClient) ListFulfillments(options interface{}) ([]Fulfillment, error) {
	prefix := FulfillmentPathPrefix(fc.resource, fc.resourceID)
	path := fmt.Sprintf("%fc.json", prefix)
	resource := new(MultipleFulfillmentsResponse)
	err := fc.client.Get(path, resource, options)
	return resource.Fulfillments, err
}

// Count fulfillments
func (fc *FulfillmentClient) CountFulfillments(options interface{}) (int, error) {
	prefix := FulfillmentPathPrefix(fc.resource, fc.resourceID)
	path := fmt.Sprintf("%s/count.json", prefix)
	return fc.client.Count(path, options)
}

// Get individual fulfillment
func (fc *FulfillmentClient) GetFulfillments(fulfillmentID int64, options interface{}) (*Fulfillment, error) {
	prefix := FulfillmentPathPrefix(fc.resource, fc.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, fulfillmentID)
	resource := new(SingleFulfillmentResponse)
	err := fc.client.Get(path, resource, options)
	return resource.Fulfillment, err
}

// Create a new fulfillment
func (fc *FulfillmentClient) CreateFulfillments(fulfillment Fulfillment) (*Fulfillment, error) {
	prefix := FulfillmentPathPrefix(fc.resource, fc.resourceID)
	path := fmt.Sprintf("%fc.json", prefix)
	wrappedData := SingleFulfillmentResponse{Fulfillment: &fulfillment}
	resource := new(SingleFulfillmentResponse)
	err := fc.client.Post(path, wrappedData, resource)
	return resource.Fulfillment, err
}

// Update an existing fulfillment
func (fc *FulfillmentClient) UpdateFulfillments(fulfillment Fulfillment) (*Fulfillment, error) {
	prefix := FulfillmentPathPrefix(fc.resource, fc.resourceID)
	path := fmt.Sprintf("%s/%d.json", prefix, fulfillment.ID)
	wrappedData := SingleFulfillmentResponse{Fulfillment: &fulfillment}
	resource := new(SingleFulfillmentResponse)
	err := fc.client.Put(path, wrappedData, resource)
	return resource.Fulfillment, err
}

// Complete an existing fulfillment
func (fc *FulfillmentClient) CompleteFulfillments(fulfillmentID int64) (*Fulfillment, error) {
	prefix := FulfillmentPathPrefix(fc.resource, fc.resourceID)
	path := fmt.Sprintf("%s/%d/complete.json", prefix, fulfillmentID)
	resource := new(SingleFulfillmentResponse)
	err := fc.client.Post(path, nil, resource)
	return resource.Fulfillment, err
}

// Transition an existing fulfillment
func (fc *FulfillmentClient) TransitionFulfillments(fulfillmentID int64) (*Fulfillment, error) {
	prefix := FulfillmentPathPrefix(fc.resource, fc.resourceID)
	path := fmt.Sprintf("%s/%d/open.json", prefix, fulfillmentID)
	resource := new(SingleFulfillmentResponse)
	err := fc.client.Post(path, nil, resource)
	return resource.Fulfillment, err
}

// Cancel an existing fulfillment
func (fc *FulfillmentClient) CancelFulfillments(fulfillmentID int64) (*Fulfillment, error) {
	prefix := FulfillmentPathPrefix(fc.resource, fc.resourceID)
	path := fmt.Sprintf("%s/%d/cancel.json", prefix, fulfillmentID)
	resource := new(SingleFulfillmentResponse)
	err := fc.client.Post(path, nil, resource)
	return resource.Fulfillment, err
}
