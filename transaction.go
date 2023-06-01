package goshopify

import "fmt"

// TransactionService is an interface for interfacing with the transactions endpoints of
// the Shopify API.
// See: https://help.shopify.com/api/reference/transaction
type TransactionService interface {
	List(int64, interface{}) ([]Transaction, error)
	Count(int64, interface{}) (int, error)
	Get(int64, int64, interface{}) (*Transaction, error)
	Create(int64, Transaction) (*Transaction, error)
}

//TODO: latest from shopify

// type Transaction struct {
// 	TransactionAmount    float64                 `json:"transaction-amount"`
// 	BuyerPendingPaymentInstructions []PendingPaymentInstructionInput `json:"transaction-buyer_pending_payment_instructions"`
// 	BuyerPendingPaymentNotice      string                  `json:"transaction-buyer_pending_payment_notice"`
// 	CreatedAt                 string                  `json:"transaction-created_at"`
// 	Gateway                   string                  `json:"transaction-gateway"`
// 	GatewayDisplayName        string                  `json:"transaction-gateway_display_name"`
// 	ID                        string                  `json:"transaction-id"`
// 	Kind                      string                  `json:"transaction-kind"`
// 	Name                      string                  `json:"transaction-name"`
// 	PaymentDetails            TransactionPaymentDetails  `json:"transaction-payment_details"`
// 	Receipt                   string                  `json:"transaction-receipt"`
// 	ShowBuyerPendingPaymentInstructions bool `json:"transaction-show_buyer_pending_payment_instructions"`
// 	Status                     string                  `json:"transaction-status"`
// 	StatusLabel                 string                  `json:"transaction-status_label"`
// }



// type PendingPaymentInstructionInput struct {
// 	// Add fields for pending payment instruction input
// }

// type TransactionPaymentDetails struct {
// 	// Add fields for payment details
// }

// TransactionServiceOp handles communication with the transaction related methods of the
// Shopify API.
type TransactionServiceOp struct {
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
func (s *TransactionServiceOp) List(orderID int64, options interface{}) ([]Transaction, error) {
	path := fmt.Sprintf("%s/%d/transactions.json", ordersBasePath, orderID)
	resource := new(TransactionsResource)
	err := s.client.Get(path, resource, options)
	return resource.Transactions, err
}

// Count transactions
func (s *TransactionServiceOp) Count(orderID int64, options interface{}) (int, error) {
	path := fmt.Sprintf("%s/%d/transactions/count.json", ordersBasePath, orderID)
	return s.client.Count(path, options)
}

// Get individual transaction
func (s *TransactionServiceOp) Get(orderID int64, transactionID int64, options interface{}) (*Transaction, error) {
	path := fmt.Sprintf("%s/%d/transactions/%d.json", ordersBasePath, orderID, transactionID)
	resource := new(TransactionResource)
	err := s.client.Get(path, resource, options)
	return resource.Transaction, err
}

// Create a new transaction
func (s *TransactionServiceOp) Create(orderID int64, transaction Transaction) (*Transaction, error) {
	path := fmt.Sprintf("%s/%d/transactions.json", ordersBasePath, orderID)
	wrappedData := TransactionResource{Transaction: &transaction}
	resource := new(TransactionResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Transaction, err
}
