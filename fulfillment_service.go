package goshopify

import "fmt"

const (
	fulfillmentServiceBasePath = "fulfillment_services"
)

// FulfillmentServiceRepository is an interface for interfacing with the fulfillment service of the Shopify API.
// https://help.shopify.com/api/reference/fulfillmentservice
type FulfillmentServiceRepository interface {
	ListFulfillment(interface{}) ([]FulfillmentService, error)
	GetFulfillment(int64, interface{}) (*FulfillmentService, error)
	CreateFulfillment(FulfillmentService) (*FulfillmentService, error)
	UpdateFulfillment(FulfillmentService) (*FulfillmentService, error)
	DeleteFulfillment(int64) error
}

type FulfillmentService struct {
	Id                     int64  `json:"id,omitempty"`                    // FIXME: Field Not Available In Model 23/04
	IncludePendingStock    bool   `json:"include_pending_stock,omitempty"` // FIXME: Field Not Available In Model 23/04
	Email                  string `json:"email,omitempty"`                 // FIXME: Field Not Available In Model 23/04
	ServiceName            string `json:"service_name,omitempty"`          // FIXME: Field Not Available In Model 23/04
	Name                   string `json:"name,omitempty"`
	Handle                 string `json:"handle,omitempty"`
	FulfillmentOrdersOptIn bool   `json:"fulfillment_orders_opt_in,omitempty"`
	ProviderId             int64  `json:"provider_id,omitempty"`
	LocationId             int64  `json:"location_id,omitempty"`
	CallbackURL            string `json:"callback_url,omitempty"`
	TrackingSupport        bool   `json:"tracking_support,omitempty"`
	InventoryManagement    bool   `json:"inventory_management,omitempty"`
	AdminGraphqlApiId      string `json:"admin_graphql_api_id,omitempty"`
	PermitsSkuSharing      bool   `json:"permits_sku_sharing,omitempty"`
	RequiresShippingMethod bool   `json:"requires_shipping_method,omitempty"`
}

type SingleFulfillmentServiceResponse struct {
	FulfillmentService *FulfillmentService `json:"fulfillment_service,omitempty"`
}

type MultipleFulfillmentServicesResponse struct {
	FulfillmentServices []FulfillmentService `json:"fulfillment_services,omitempty"`
}

type FulfillmentServiceOptions struct {
	Scope string `url:"scope,omitempty"`
}

// FulfillmentServiceClient handles communication with the FulfillmentServices related methods of the Shopify API
type FulfillmentServiceClient struct {
	client *Client
}

// List Receive a list of all FulfillmentService
func (fc *FulfillmentServiceClient) ListFulfillment(options interface{}) ([]FulfillmentService, error) {
	path := fmt.Sprintf("%s.json", fulfillmentServiceBasePath)
	resource := new(MultipleFulfillmentServicesResponse)
	err := fc.client.Get(path, resource, options)
	return resource.FulfillmentServices, err
}

// Get Receive a single FulfillmentService
func (fc *FulfillmentServiceClient) GetFulfillment(fulfillmentServiceId int64, options interface{}) (*FulfillmentService, error) {
	path := fmt.Sprintf("%s/%d.json", fulfillmentServiceBasePath, fulfillmentServiceId)
	resource := new(SingleFulfillmentServiceResponse)
	err := fc.client.Get(path, resource, options)
	return resource.FulfillmentService, err
}

// Create Create a new FulfillmentService
func (fc *FulfillmentServiceClient) CreateFulfillment(fulfillmentService FulfillmentService) (*FulfillmentService, error) {
	path := fmt.Sprintf("%s.json", fulfillmentServiceBasePath)
	wrappedData := SingleFulfillmentServiceResponse{FulfillmentService: &fulfillmentService}
	resource := new(SingleFulfillmentServiceResponse)
	err := fc.client.Post(path, wrappedData, resource)
	return resource.FulfillmentService, err
}

// Update Modify an existing FulfillmentService
func (fc *FulfillmentServiceClient) UpdateFulfillment(fulfillmentService FulfillmentService) (*FulfillmentService, error) {
	path := fmt.Sprintf("%s/%d.json", fulfillmentServiceBasePath, fulfillmentService.Id)
	wrappedData := SingleFulfillmentServiceResponse{FulfillmentService: &fulfillmentService}
	resource := new(SingleFulfillmentServiceResponse)
	err := fc.client.Put(path, wrappedData, resource)
	return resource.FulfillmentService, err
}

// Delete Remove an existing FulfillmentService
func (fc *FulfillmentServiceClient) DeleteFulfillment(fulfillmentServiceId int64) error {
	path := fmt.Sprintf("%s/%d.json", fulfillmentServiceBasePath, fulfillmentServiceId)
	return fc.client.Delete(path)
}
