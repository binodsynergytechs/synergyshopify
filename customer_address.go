package goshopify

import "fmt"

// const SinglecustomerAddressResponseName = "customer-addresses"

// CustomerAddressRepository is an interface for interfacing with the customer address endpoints of the Shopify API.
// See: https://help.shopify.com/en/api/reference/customers/customer_address
type CustomerAddressRepository interface {
	ListCustomerAddress(int64, interface{}) ([]CustomerAddress, error)
	GetCustomerAddress(int64, int64, interface{}) (*CustomerAddress, error)
	CreateCustomerAddress(int64, CustomerAddress) (*CustomerAddress, error)
	UpdateCustomerAddress(int64, CustomerAddress) (*CustomerAddress, error)
	DeleteCustomerAddress(int64, int64) error
}

// CustomerAddressClient handles communication with the customer address related methods of
// the Shopify API.
type CustomerAddressClient struct {
	client *Client
}

// CustomerAddress represents a Shopify customer address
type CustomerAddress struct {
	ID           int64  `json:"id,omitempty"`
	CustomerID   int64  `json:"customer_id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Company      string `json:"company,omitempty"`
	Address1     string `json:"address1,omitempty"`
	Address2     string `json:"address2,omitempty"`
	City         string `json:"city,omitempty"`
	Province     string `json:"province,omitempty"`
	Country      string `json:"country,omitempty"`
	Zip          string `json:"zip,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Name         string `json:"name,omitempty"`
	ProvinceCode string `json:"province_code,omitempty"`
	CountryCode  string `json:"country_code,omitempty"`
	CountryName  string `json:"country_name,omitempty"`
	Default      bool   `json:"default,omitempty"` // FIXME: Field Not Available Or Deprecated In Latest Shopify Model 23/04
}

// CustomerAddressResponse represents the result from the addresses/X.json endpoint
type SingleCustomerAddressResponse struct {
	Address *CustomerAddress `json:"customer_address"`
}

// CustomerAddressResponse represents the result from the customers/X/addresses.json endpoint
type MultipleCustomerAddressesResponse struct {
	Addresses []CustomerAddress `json:"addresses"`
}

// List addresses
func (cac *CustomerAddressClient) ListCustomerAddress(customerID int64, options interface{}) ([]CustomerAddress, error) {
	path := fmt.Sprintf("%s/%d/addresses.json", customersBasePath, customerID)
	resource := new(MultipleCustomerAddressesResponse)
	err := cac.client.Get(path, resource, options)
	return resource.Addresses, err
}

// Get address
func (cac *CustomerAddressClient) GetCustomerAddress(customerID, addressID int64, options interface{}) (*CustomerAddress, error) {
	path := fmt.Sprintf("%s/%d/addresses/%d.json", customersBasePath, customerID, addressID)
	resource := new(SingleCustomerAddressResponse)
	err := cac.client.Get(path, resource, options)
	return resource.Address, err
}

// Create a new address for given customer
func (cac *CustomerAddressClient) CreateCustomerAddress(customerID int64, address CustomerAddress) (*CustomerAddress, error) {
	path := fmt.Sprintf("%s/%d/addresses.json", customersBasePath, customerID)
	wrappedData := SingleCustomerAddressResponse{Address: &address}
	resource := new(SingleCustomerAddressResponse)
	err := cac.client.Post(path, wrappedData, resource)
	return resource.Address, err
}

// Create a new address for given customer
func (cac *CustomerAddressClient) UpdateCustomerAddress(customerID int64, address CustomerAddress) (*CustomerAddress, error) {
	path := fmt.Sprintf("%s/%d/addresses/%d.json", customersBasePath, customerID, address.ID)
	wrappedData := SingleCustomerAddressResponse{Address: &address}
	resource := new(SingleCustomerAddressResponse)
	err := cac.client.Put(path, wrappedData, resource)
	return resource.Address, err
}

// Delete an existing address
func (cac *CustomerAddressClient) DeleteCustomerAddress(customerID, addressID int64) error {
	return cac.client.Delete(fmt.Sprintf("%s/%d/addresses/%d.json", customersBasePath, customerID, addressID))
}
