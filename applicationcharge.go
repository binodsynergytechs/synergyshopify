package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const applicationChargesBasePath = "application_charges"

// ApplicationChargeService is an interface for interacting with the
// ApplicationCharge endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/billing/applicationcharge
type ApplicationChargeService interface {
	CreateApplicationCharge(ApplicationCharge) (*ApplicationCharge, error)
	GetApplicationCharge(int64, interface{}) (*ApplicationCharge, error)
	ListApplicationCharges(interface{}) ([]ApplicationCharge, error)
	ActivateApplicationCharge(ApplicationCharge) (*ApplicationCharge, error)
}

type ApplicationChargeServiceOp struct {
	client *Client
}

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
	ChargeType         *string          `json:"charge_type"`          // FIXME: Field Not Available Or Deprecated In Latest Shopify Model 23/04
	DecoratedReturnURL string           `json:"decorated_return_url"` // FIXME: Field Not Available Or Deprecated In Latest Shopify Model 23/04
	ConfirmationURL    string           `json:"confirmation_url"`
}

// ApplicationChargeResource represents the result from the
// admin/application_charges{/X{/activate.json}.json}.json endpoints.
type ApplicationChargeResource struct {
	Charge *ApplicationCharge `json:"application_charge"`
}

// ApplicationChargesResource represents the result from the
// admin/application_charges.json endpoint.
type ApplicationChargesResource struct {
	Charges []ApplicationCharge `json:"application_charges"`
}

// CreateApplicationCharge creates a new application charge.
// It takes an ApplicationCharge parameter and returns a pointer to the created ApplicationCharge and an error, if any.
func (a ApplicationChargeServiceOp) CreateApplicationCharge(charge ApplicationCharge) (*ApplicationCharge, error) {
	path := fmt.Sprintf("%s.json", applicationChargesBasePath)
	resource := &ApplicationChargeResource{}
	return resource.Charge, a.client.Post(path, ApplicationChargeResource{Charge: &charge}, resource)
}

// GetApplicationCharge retrieves an individual application charge.
// It takes the chargeID as an int64 and options as an interface{} parameter.
// It returns a pointer to the retrieved ApplicationCharge and an error, if any.
func (a ApplicationChargeServiceOp) GetApplicationCharge(chargeID int64, options interface{}) (*ApplicationCharge, error) {
	path := fmt.Sprintf("%s/%d.json", applicationChargesBasePath, chargeID)
	resource := &ApplicationChargeResource{}
	return resource.Charge, a.client.Get(path, resource, options)
}

// ListApplicationCharges retrieves all application charges.
// It takes options as an interface{} parameter.
// It returns a slice of ApplicationCharge and an error, if any.
func (a ApplicationChargeServiceOp) ListApplicationCharges(options interface{}) ([]ApplicationCharge, error) {
	path := fmt.Sprintf("%s.json", applicationChargesBasePath)
	resource := &ApplicationChargesResource{}
	return resource.Charges, a.client.Get(path, resource, options)
}

// ActivateApplicationCharge activates an application charge.
// It takes an ApplicationCharge parameter and returns a pointer to the activated ApplicationCharge and an error, if any.
func (a ApplicationChargeServiceOp) ActivateApplicationCharge(charge ApplicationCharge) (*ApplicationCharge, error) {
	path := fmt.Sprintf("%s/%d/activate.json", applicationChargesBasePath, charge.ID)
	resource := &ApplicationChargeResource{}
	return resource.Charge, a.client.Post(path, ApplicationChargeResource{Charge: &charge}, resource)
}
