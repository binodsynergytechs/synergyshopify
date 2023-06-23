package synergyshopify

import "fmt"

const orderRiskBasePath = "orders"

// OrderRiskService is an interface for interacting with the order risk
// endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/order-risk
type OrderRiskService interface {
	List(int64, interface{}) ([]OrderRisk, error)
	Get(int64, int64, interface{}) (*OrderRisk, error)
	Create(OrderRisk, int64) (*OrderRisk, error)
	Update(int64, OrderRisk) (*OrderRisk, error)
	Delete(int64, int64) error
}

// OrderRiksServiceOp handles communication with the order risk
// related methods of the Shopify API.
type OrderRiksServiceOp struct {
	client *Client
}

// OrderRisk represents a Shopify Order Risk.
type OrderRisk struct {
	ID              int64  `json:"id"`
	OrderID         int64  `json:"order_id"`
	CheckoutID      int64  `json:"checkout_id"`
	Source          string `json:"source"`
	Score           string `json:"score"`
	Recommedation   string `json:"recommedation"`
	Display         bool   `json:"display"`
	CauseCancel     bool   `json:"cause_cancel"`
	Message         string `json:"message"`
	MerchantMessage string `json:"merchant_message"`
}

// OrderRiskResource represents the result form the risks/X.json endpoint
type OrderRiskResource struct {
	Risk *OrderRisk `json:"risk"`
}

// OrderRisksResource represents the result from the risks.json endpoint
type OrderRisksResource struct {
	Risks []OrderRisk `json:"risks"`
}

// List OrderRisks
func (s *OrderRiksServiceOp) List(orderId int64, options interface{}) ([]OrderRisk, error) {
	path := fmt.Sprintf("%s/%d/risks.json", orderRiskBasePath, orderId)
	resource := new(OrderRisksResource)
	err := s.client.Get(path, resource, options, true)
	return resource.Risks, err
}

// Get individual order risk
func (s *OrderRiksServiceOp) Get(orderID int64, riskID int64, options interface{}) (*OrderRisk, error) {
	path := fmt.Sprintf("%s/%d/risks/%d.json", orderRiskBasePath, orderID, riskID)
	resource := new(OrderRiskResource)
	err := s.client.Get(path, resource, options, true)
	return resource.Risk, err
}

// Create a new order risk

func (s *OrderRiksServiceOp) Create(orderrisk OrderRisk, orderID int64) (*OrderRisk, error) {
	path := fmt.Sprintf("%s/%d/risks.json", orderRiskBasePath, orderID)
	wrappedData := OrderRiskResource{Risk: &orderrisk}
	resource := new(OrderRiskResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Risk, err
}

// Update an existing order risk
func (s *OrderRiksServiceOp) Update(orderID int64, orderrisk OrderRisk) (*OrderRisk, error) {
	path := fmt.Sprintf("%s/%d/risks/%d.json", orderRiskBasePath, orderID, orderrisk.ID)
	wrappedData := OrderRiskResource{Risk: &orderrisk}
	resource := new(OrderRiskResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Risk, err
}

// Delete an existing order risk.
func (s *OrderRiksServiceOp) Delete(OrderID, OrderRiskID int64) error {
	return s.client.Delete(fmt.Sprintf("%s/%d/risks/%d.json", orderRiskBasePath, OrderID, OrderRiskID))
}
