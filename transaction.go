package goshopify

import "fmt"

// TransactionRepository is an interface for interfacing with the transactions endpoints of the Shopify API.
// See: https://help.shopify.com/api/reference/transaction
type TransactionRepository interface {
	ListTransaction(int64, interface{}) ([]Transaction, error)
	CountTransaction(int64, interface{}) (int, error)
	GetTransaction(int64, int64, interface{}) (*Transaction, error)
	CreateTransaction(int64, Transaction) (*Transaction, error)
}

// TransactionClient handles communication with the transaction related methods of the Shopify API.
type TransactionClient struct {
	client *Client
}

// TransactionResource represents the result from the orders/X/transactions/Y.json endpoint
type TransactionResource struct {
	Transaction *Transaction `json:"transaction"`
}

// TransactionsResource represents the result from the orders/X/transactions.json endpoint
type TransactionsResource struct {
	Transactions []Transaction `json:"transactions"`
}

// List transactions
func (tc *TransactionClient) ListTransaction(orderID int64, options interface{}) ([]Transaction, error) {
	path := fmt.Sprintf("%s/%d/transactions.json", ordersBasePath, orderID)
	resource := new(TransactionsResource)
	err := tc.client.Get(path, resource, options)
	return resource.Transactions, err
}

// Count transactions
func (tc *TransactionClient) CountTransaction(orderID int64, options interface{}) (int, error) {
	path := fmt.Sprintf("%s/%d/transactions/count.json", ordersBasePath, orderID)
	return tc.client.Count(path, options)
}

// Get individual transaction
func (tc *TransactionClient) GetTransaction(orderID int64, transactionID int64, options interface{}) (*Transaction, error) {
	path := fmt.Sprintf("%s/%d/transactions/%d.json", ordersBasePath, orderID, transactionID)
	resource := new(TransactionResource)
	err := tc.client.Get(path, resource, options)
	return resource.Transaction, err
}

// Create a new transaction
func (tc *TransactionClient) CreateTransaction(orderID int64, transaction Transaction) (*Transaction, error) {
	path := fmt.Sprintf("%s/%d/transactions.json", ordersBasePath, orderID)
	wrappedData := TransactionResource{Transaction: &transaction}
	resource := new(TransactionResource)
	err := tc.client.Post(path, wrappedData, resource)
	return resource.Transaction, err
}
