package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const applicationChargesBasePath = "application_charges"

// ApplicationChargeRepository is an interface for managing application charges in the Shopify API.
// See https://help.shopify.com/api/reference/billing/applicationcharge
type ApplicationChargeRepository interface {
	CreateApplicationCharge(ApplicationCharge) (*ApplicationCharge, error)
	GetApplicationCharge(int64, interface{}) (*ApplicationCharge, error)
	ListApplicationCharges(interface{}) ([]ApplicationCharge, error)
	ActivateApplicationCharge(ApplicationCharge) (*ApplicationCharge, error)
}

// ApplicationChargeClient handles communication with the ApplicationCharge endpoints of the Shopify API.
type ApplicationChargeClient struct {
	client *Client
}

// ApplicationCharge represents an application charge in Shopify.
type ApplicationCharge struct {
	ID                 int64            `json:"id"`
	Name               string           `json:"name"`
	APIClientID        int64            `json:"api_client_id"`
	Price              *decimal.Decimal `json:"price"`
	Status             string           `json:"status"`
	ReturnURL          string           `json:"return_url"`
	Test               *bool            `json:"test"`
	CreatedAt          *time.Time       `json:"created_at"`
	UpdatedAt          *time.Time       `json:"updated_at"`
	ChargeType         *string          `json:"charge_type"`          // TODO: Field Not Available Or Deprecated In Latest Shopify Model 23/04
	DecoratedReturnURL string           `json:"decorated_return_url"` // TODO: Field Not Available Or Deprecated In Latest Shopify Model 23/04
	ConfirmationURL    string           `json:"confirmation_url"`
}

// ApplicationChargeResponse represents the response from the application charge endpoints.
type ApplicationChargeResponse struct {
	Charge *ApplicationCharge `json:"application_charge"`
}

// ApplicationChargesResponse represents the response from the application_charges endpoint.
type ApplicationChargesResponse struct {
	Charges []ApplicationCharge `json:"application_charges"`
}

// CreateApplicationCharge creates a new application charge.
// It takes an ApplicationCharge parameter and returns a pointer to the created ApplicationCharge and an error, if any.
func (a ApplicationChargeClient) CreateApplicationCharge(charge ApplicationCharge) (*ApplicationCharge, error) {
	path := fmt.Sprintf("%s.json", applicationChargesBasePath)
	resource := &ApplicationChargeResponse{}
	return resource.Charge, a.client.Post(path, ApplicationChargeResponse{Charge: &charge}, resource)
}

// GetApplicationCharge retrieves an individual application charge.
// It takes the chargeID as an int64 and options as an interface{} parameter.
// It returns a pointer to the retrieved ApplicationCharge and an error, if any.
func (a ApplicationChargeClient) GetApplicationCharge(chargeID int64, options interface{}) (*ApplicationCharge, error) {
	path := fmt.Sprintf("%s/%d.json", applicationChargesBasePath, chargeID)
	resource := &ApplicationChargeResponse{}
	return resource.Charge, a.client.Get(path, resource, options)
}

// ListApplicationCharges retrieves all application charges.
// It takes options as an interface{} parameter.
// It returns a slice of ApplicationCharge and an error, if any.
func (a ApplicationChargeClient) ListApplicationCharges(options interface{}) ([]ApplicationCharge, error) {
	path := fmt.Sprintf("%s.json", applicationChargesBasePath)
	resource := &ApplicationChargesResponse{}
	return resource.Charges, a.client.Get(path, resource, options)
}

// ActivateApplicationCharge activates an application charge.
// It takes an ApplicationCharge parameter and returns a pointer to the activated ApplicationCharge and an error, if any.
func (a ApplicationChargeClient) ActivateApplicationCharge(charge ApplicationCharge) (*ApplicationCharge, error) {
	path := fmt.Sprintf("%s/%d/activate.json", applicationChargesBasePath, charge.ID)
	resource := &ApplicationChargeResponse{}
	return resource.Charge, a.client.Post(path, ApplicationChargeResponse{Charge: &charge}, resource)
}
