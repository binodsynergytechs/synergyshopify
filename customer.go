package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const customersBasePath = "customers"
const MultipleCustomersResponseName = "customers"

// CustomerRepository is an interface for interfacing with the customers endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/customer
type CustomerRepository interface {
	ListCustomer(interface{}) ([]Customer, error)
	ListWithPaginationCustomer(options interface{}) ([]Customer, *Pagination, error)
	CountCustomer(interface{}) (int, error)
	GetCustomer(int64, interface{}) (*Customer, error)
	SearchCustomer(interface{}) ([]Customer, error)
	CreateCustomer(Customer) (*Customer, error)
	UpdateCustomer(Customer) (*Customer, error)
	DeleteCustomer(int64) error
	ListOrders(int64, interface{}) ([]Order, error)
	ListTags(interface{}) ([]string, error)

	// MetaFieldsRepository used for Customer resource to communicate with MetaFields resource
	MetaFieldRepository
}

// CustomerClient handles communication with the product related methods of
// the Shopify API.
type CustomerClient struct {
	client *Client
}

// Customer represents a Shopify customer
type Customer struct {
	ID                    int64              `json:"id,omitempty"`
	Email                 string             `json:"email,omitempty"`
	FirstName             string             `json:"first_name,omitempty"`
	LastName              string             `json:"last_name,omitempty"`
	State                 string             `json:"state,omitempty"`
	Note                  string             `json:"note,omitempty"`
	VerifiedEmail         bool               `json:"verified_email,omitempty"`
	MultiPassIdentifier   string             `json:"multipass_identifier,omitempty"`
	OrdersCount           int                `json:"orders_count,omitempty"`
	TaxExempt             bool               `json:"tax_exempt,omitempty"`
	TotalSpent            *decimal.Decimal   `json:"total_spent,omitempty"`
	Phone                 string             `json:"phone,omitempty"`
	Tags                  string             `json:"tags,omitempty"`
	LastOrderId           int64              `json:"last_order_id,omitempty"`
	LastOrderName         string             `json:"last_order_name,omitempty"`
	AcceptsMarketing      bool               `json:"accepts_marketing,omitempty"` // FIXME: Field Not Available Or Deprecated In Latest Shopify Model 23/04
	DefaultAddress        *CustomerAddress   `json:"default_address,omitempty"`
	Addresses             []*CustomerAddress `json:"addresses,omitempty"`
	CreatedAt             *time.Time         `json:"created_at,omitempty"`
	UpdatedAt             *time.Time         `json:"updated_at,omitempty"`
	MetaFields            []MetaField        `json:"metaFields,omitempty"`
	Currency              string             `json:"currency"`                 // TODO: Field Available In Latest Shopify Model 23/04
	Password              string             `json:"password"`                 // TODO: Field Available In Latest Shopify Model 23/04
	PasswordConfirmation  string             `json:"password_confirmation"`    // TODO: Field Available In Latest Shopify Model 23/04
	EmailMarketingConsent interface{}        `json:"email_marketing_consent"`  // TODO: Field Available In Latest Shopify Model 23/04
	SmsMarketingConsent   interface{}        `json:"sms_marketing_consent"`    // TODO: Field Available In Latest Shopify Model 23/04
	TaxExemptions         []string           `json:"tax_exemptions,omitempty"` // TODO: Field Available In Latest Shopify Model 23/04
}

// Represents the result from the customers/X.json endpoint
type SingleCustomerResponse struct {
	Customer *Customer `json:"customer"`
}

// Represents the result from the customers.json endpoint
type MultipleCustomersResponse struct {
	Customers []Customer `json:"customers"`
}

// Represents the result from the customers/tags.json endpoint
type CustomerTagsResponse struct {
	Tags []string `json:"tags"`
}

// Represents the options available when searching for a customer
type CustomerSearchOptions struct {
	Page   int    `url:"page,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Fields string `url:"fields,omitempty"`
	Order  string `url:"order,omitempty"`
	Query  string `url:"query,omitempty"`
}

// List customers
func (cc *CustomerClient) ListCustomer(options interface{}) ([]Customer, error) {
	path := fmt.Sprintf("%s.json", customersBasePath)
	resource := new(MultipleCustomersResponse)
	err := cc.client.Get(path, resource, options)
	return resource.Customers, err
}

// ListWithPagination lists customers and return pagination to retrieve next/previous results.
func (cc *CustomerClient) ListWithPaginationCustomer(options interface{}) ([]Customer, *Pagination, error) {
	path := fmt.Sprintf("%s.json", customersBasePath)
	resource := new(MultipleCustomersResponse)

	pagination, err := cc.client.ListWithPagination(path, resource, options)
	if err != nil {
		return nil, nil, err
	}

	return resource.Customers, pagination, nil
}

// Count customers
func (cc *CustomerClient) CountCustomer(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", customersBasePath)
	return cc.client.Count(path, options)
}

// Get customer
func (cc *CustomerClient) GetCustomer(customerID int64, options interface{}) (*Customer, error) {
	path := fmt.Sprintf("%s/%v.json", customersBasePath, customerID)
	resource := new(SingleCustomerResponse)
	err := cc.client.Get(path, resource, options)
	return resource.Customer, err
}

// Create a new customer
func (cc *CustomerClient) CreateCustomer(customer Customer) (*Customer, error) {
	path := fmt.Sprintf("%s.json", customersBasePath)
	wrappedData := SingleCustomerResponse{Customer: &customer}
	resource := new(SingleCustomerResponse)
	err := cc.client.Post(path, wrappedData, resource)
	return resource.Customer, err
}

// Update an existing customer
func (cc *CustomerClient) UpdateCustomer(customer Customer) (*Customer, error) {
	path := fmt.Sprintf("%s/%d.json", customersBasePath, customer.ID)
	wrappedData := SingleCustomerResponse{Customer: &customer}
	resource := new(SingleCustomerResponse)
	err := cc.client.Put(path, wrappedData, resource)
	return resource.Customer, err
}

// Delete an existing customer
func (cc *CustomerClient) DeleteCustomer(customerID int64) error {
	path := fmt.Sprintf("%s/%d.json", customersBasePath, customerID)
	return cc.client.Delete(path)
}

// Search customers
func (cc *CustomerClient) SearchCustomer(options interface{}) ([]Customer, error) {
	path := fmt.Sprintf("%s/search.json", customersBasePath)
	resource := new(MultipleCustomersResponse)
	err := cc.client.Get(path, resource, options)
	return resource.Customers, err
}

// ListOrders retrieves all orders from a customer
func (cc *CustomerClient) ListOrders(customerID int64, options interface{}) ([]Order, error) {
	path := fmt.Sprintf("%s/%d/orders.json", customersBasePath, customerID)
	resource := new(OrdersResource)
	err := cc.client.Get(path, resource, options)
	return resource.Orders, err
}

// ListTags retrieves all unique tags across all customers
func (cc *CustomerClient) ListTags(options interface{}) ([]string, error) {
	path := fmt.Sprintf("%s/tags.json", customersBasePath)
	resource := new(CustomerTagsResponse)
	err := cc.client.Get(path, resource, options)
	return resource.Tags, err
}

// List metaFields for a customer
func (cc *CustomerClient) ListMetaFields(customerID int64, options interface{}) ([]MetaField, error) {
	metaFieldService := &MetaFieldClient{client: cc.client, resource: MultipleCustomersResponseName, resourceID: customerID}
	return metaFieldService.ListMetaFields(options)
}

// Count metaFields for a customer
func (cc *CustomerClient) CountMetaFields(customerID int64, options interface{}) (int, error) {
	metaFieldService := &MetaFieldClient{client: cc.client, resource: MultipleCustomersResponseName, resourceID: customerID}
	return metaFieldService.CountMetaFields(options)
}

// Get individual metaField for a customer
func (cc *CustomerClient) GetMetaField(customerID int64, metaFieldID int64, options interface{}) (*MetaField, error) {
	metaFieldService := &MetaFieldClient{client: cc.client, resource: MultipleCustomersResponseName, resourceID: customerID}
	return metaFieldService.GetMetaFields(metaFieldID, options)
}

// Create a new metaField for a customer
func (cc *CustomerClient) CreateMetaField(customerID int64, metaField MetaField) (*MetaField, error) {
	metaFieldService := &MetaFieldClient{client: cc.client, resource: MultipleCustomersResponseName, resourceID: customerID}
	return metaFieldService.CreateMetaFields(metaField)
}

// Update an existing metaField for a customer
func (cc *CustomerClient) UpdateMetaField(customerID int64, metaField MetaField) (*MetaField, error) {
	metaFieldService := &MetaFieldClient{client: cc.client, resource: MultipleCustomersResponseName, resourceID: customerID}
	return metaFieldService.UpdateMetaFields(metaField)
}

// // Delete an existing metaField for a customer
func (cc *CustomerClient) DeleteMetaField(customerID int64, metaFieldID int64) error {
	metaFieldService := &MetaFieldClient{client: cc.client, resource: MultipleCustomersResponseName, resourceID: customerID}
	return metaFieldService.DeleteMetaFields(metaFieldID)
}
