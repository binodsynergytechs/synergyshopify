package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const giftCardsBasePath = "gift_cards"

// GiftCardRepository is an interface for interfacing with the gift card endpoints of the Shopify API.
// See: https://shopify.dev/docs/api/admin-rest/2023-04/resources/gift-card
type GiftCardRepository interface {
	GetGiftCard(int64) (*GiftCard, error)
	CreateGiftCard(GiftCard) (*GiftCard, error)
	UpdateGiftCard(GiftCard) (*GiftCard, error)
	ListGiftCard() ([]GiftCard, error)
	DisableGiftCard(int64) (*GiftCard, error)
	CountGiftCard(interface{}) (int, error)
}

// giftCardClient handles communication with the gift card related methods of the Shopify API.
type GiftCardClient struct {
	client *Client
}

// giftCard represents a Shopify discount rule
type GiftCard struct {
	ID             int64            `json:"id,omitempty"`
	ApiClientId    int64            `json:"api_client_id,omitempty"`
	Balance        *decimal.Decimal `json:"balance,omitempty"`
	InitialValue   *decimal.Decimal `json:"initial_value,omitempty"`
	Code           string           `json:"code,omitempty"`
	Currency       string           `json:"currency,omitempty"`
	CustomerID     *CustomerID      `json:"customer_id,omitempty"`
	CreatedAt      *time.Time       `json:"created_at,omitempty"`
	DisabledAt     *time.Time       `json:"disabled_at,omitempty"`
	ExpiresOn      string           `json:"expires_on,omitempty"`
	LastCharacters string           `json:"last_characters,omitempty"`
	LineItemID     int64            `json:"line_item_id,omitempty"`
	Note           string           `json:"note,omitempty"`
	OrderID        int64            `json:"order_id,omitempty"`
	TemplateSuffix string           `json:"template_suffix,omitempty"`
	UserID         int64            `json:"user_id,omitempty"`
	UpdatedAt      *time.Time       `json:"updated_at,omitempty"`
}

type CustomerID struct {
	CustomerID int64 `json:"customer_id,omitempty"`
}

// SingleGiftCardResponse represents the result from the gift_cards/X.json endpoint
type SingleGiftCardResponse struct {
	GiftCard *GiftCard `json:"gift_card"`
}

// MultipleGiftCardsResponse represents the result from the gift_cards.json endpoint
type MultipleGiftCardsResponse struct {
	GiftCards []GiftCard `json:"gift_cards"`
}

// Get retrieves a single gift cards
func (gc *GiftCardClient) GetGiftCard(giftCardID int64) (*GiftCard, error) {
	path := fmt.Sprintf("%s/%d.json", giftCardsBasePath, giftCardID)
	resource := new(SingleGiftCardResponse)
	err := gc.client.Get(path, resource, nil)
	return resource.GiftCard, err
}

// List retrieves a list of gift cards
func (gc *GiftCardClient) ListGiftCard() ([]GiftCard, error) {
	path := fmt.Sprintf("%s.json", giftCardsBasePath)
	resource := new(MultipleGiftCardsResponse)
	err := gc.client.Get(path, resource, nil)
	return resource.GiftCards, err
}

// Create creates a gift card
func (gc *GiftCardClient) CreateGiftCard(pr GiftCard) (*GiftCard, error) {
	path := fmt.Sprintf("%s.json", giftCardsBasePath)
	resource := new(SingleGiftCardResponse)
	wrappedData := SingleGiftCardResponse{GiftCard: &pr}
	err := gc.client.Post(path, wrappedData, resource)
	return resource.GiftCard, err
}

// Update updates an existing a gift card
func (gc *GiftCardClient) UpdateGiftCard(pr GiftCard) (*GiftCard, error) {
	path := fmt.Sprintf("%s/%d.json", giftCardsBasePath, pr.ID)
	resource := new(SingleGiftCardResponse)
	wrappedData := SingleGiftCardResponse{GiftCard: &pr}
	err := gc.client.Put(path, wrappedData, resource)
	return resource.GiftCard, err
}

// Disable disables an existing a gift card
func (gc *GiftCardClient) DisableGiftCard(giftCardID int64) (*GiftCard, error) {
	path := fmt.Sprintf("%s/%d/disable.json", giftCardsBasePath, giftCardID)
	resource := new(SingleGiftCardResponse)
	wrappedData := SingleGiftCardResponse{GiftCard: &GiftCard{ID: giftCardID}}
	err := gc.client.Post(path, wrappedData, resource)
	return resource.GiftCard, err
}

// Count retrieves the number of gift cards
func (gc *GiftCardClient) CountGiftCard(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", giftCardsBasePath)
	return gc.client.Count(path, options)
}
