package synergyshopify

import (
	"fmt"
	"time"
)

// create interface for fulfillment order service
type FulfillmentOrderService interface {
	Get(interface{}) (*FulfillmentOrder, error)
	Create(FulfillmentRequest) (*FulfillmentOrder, error)
}

type FulfillmentOrder struct {
	ID                  int64          `json:"id,omitempty"`
	ShopID              int64          `json:"shop_id,omitempty"`
	OrderID             int64          `json:"order_id,omitempty"`
	AssignedLocationID  int64          `json:"assigned_location_id,omitempty"`
	RequestStatus       string         `json:"request_status,omitempty"`
	Status              string         `json:"status,omitempty"`
	SupportedActions    []string       `json:"supported_actions,omitempty"`
	Destination         interface{}    `json:"destination,omitempty"`
	LineItems           []LineItem     `json:"line_items,omitempty"`
	FulfillAt           string         `json:"fulfill_at,omitempty"`
	FulfillBy           interface{}    `json:"fulfill_by,omitempty"`
	InternationalDuties interface{}    `json:"international_duties,omitempty"`
	FulfillmentHolds    []interface{}  `json:"fulfillment_holds,omitempty"`
	DeliveryMethod      DeliveryMethod `json:"delivery_method,omitempty"`
	CreatedAt           string         `json:"created_at,omitempty"`
	UpdatedAt           string         `json:"updated_at,omitempty"`
	AssignedLocation    Location       `json:"assigned_location,omitempty"`
	MerchantRequests    []interface{}  `json:"merchant_requests,omitempty"`
}

type FulfillmentRequest struct {
	Message                  string                     `json:"message,omitempty"`
	NotifyCustomer           bool                       `json:"notify_customer,omitempty"`
	TrackingInfo             TrackingInfo               `json:"tracking_info,omitempty"`
	LineItemsByFulfillmentID []LineItemsByFulfillmentID `json:"line_items_by_fulfillment_order,omitempty"`
}

type TrackingInfo struct {
	Number  string `json:"number,omitempty"`
	Company string `json:"company,omitempty"`
}

type LineItemsByFulfillmentID struct {
	FulfillmentOrderID        int64                      `json:"fulfillment_order_id,omitempty"`
	FulfillmentOrderLineItems []FulfillmentOrderLineItem `json:"fulfillment_order_line_items,omitempty"`
}

type FulfillmentOrderLineItem struct {
	ID       int64 `json:"id,omitempty"`
	Quantity int64 `json:"quantity,omitempty"`
}

type DeliveryMethod struct {
	ID                  int64      `json:"id,omitempty"`
	MethodType          string     `json:"method_type,omitempty"`
	MinDeliveryDateTime *time.Time `json:"min_delivery_date_time,omitempty"`
	MaxDeliveryDateTime *time.Time `json:"max_delivery_date_time,omitempty"`
}

// FulfillmentOrderResource represents the result from the fulfillment_orders/X.json endpoint
type FulfillmentOrderResource struct {
	FulfillmentOrder []FulfillmentOrder `json:"fulfillment_orders,omitempty"`
}

type FulfillmentOrderResourceResp struct {
	FulfillmentOrder FulfillmentOrder `json:"fulfillment,omitempty"`
}

type FulfillmentRequestResource struct {
	FulfillmentRequest *FulfillmentRequest `json:"fulfillment,omitempty"`
}

// FulfillmentOrderServiceOp handles communication with the fulfillment
type FulfillmentOrderServiceOp struct {
	client     *Client
	resource   string
	resourceID int64
}

// Get Receive a single FulfillmentOrder
func (s *FulfillmentOrderServiceOp) Get(options interface{}) ([]FulfillmentOrder, error) {
	path := fmt.Sprintf("%s/%d/fulfillment_orders.json", s.resource, s.resourceID)
	resource := new(FulfillmentOrderResource)
	err := s.client.Get(path, resource, options, true)
	return resource.FulfillmentOrder, err
}

// Create Create a new FulfillmentOrder
func (s *FulfillmentOrderServiceOp) Create(fulfillmentRequest FulfillmentRequest) (FulfillmentOrder, error) {
	path := "/fulfillments.json"
	fulfillment := FulfillmentRequestResource{FulfillmentRequest: &fulfillmentRequest}
	resource := new(FulfillmentOrderResourceResp)
	err := s.client.Post(path, fulfillment, resource)
	return resource.FulfillmentOrder, err
}

// Update Update a Fulfillment tracking info
func (s *FulfillmentOrderServiceOp) UpdateTracking(fulfillmentId int64, fulfillmentRequest FulfillmentRequest) (FulfillmentOrder, error) {
	path := fmt.Sprintf("/fulfillments/%d/update_tracking.json", fulfillmentId)
	fulfillment := FulfillmentRequestResource{FulfillmentRequest: &fulfillmentRequest}
	resource := new(FulfillmentOrderResourceResp)
	err := s.client.Post(path, fulfillment, resource)
	return resource.FulfillmentOrder, err
}
