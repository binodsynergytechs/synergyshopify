package goshopify

import (
	"fmt"
	"time"
)

const storefrontAccessTokensBasePath = "storefront_access_tokens"

// StorefrontAccessTokenRepository is an interface for interfacing with the storefront access
// token endpoints of the Shopify API.
// See: https://help.shopify.com/api/reference/access/storefrontaccesstoken
type StorefrontAccessTokenRepository interface {
	ListStorefrontAccessToken(interface{}) ([]StorefrontAccessToken, error)
	CreateStorefrontAccessToken(StorefrontAccessToken) (*StorefrontAccessToken, error)
	DeleteStorefrontAccessToken(int64) error
}

// StorefrontAccessTokenClient handles communication with the storefront access token
// related methods of the Shopify API.
type StorefrontAccessTokenClient struct {
	client *Client
}

// StorefrontAccessToken represents a Shopify storefront access token
type StorefrontAccessToken struct {
	ID                int64      `json:"id,omitempty"`
	Title             string     `json:"title,omitempty"`
	AccessToken       string     `json:"access_token,omitempty"`
	AccessScope       string     `json:"access_scope,omitempty"`
	AdminGraphqlApiID string     `json:"admin_graphql_api_id,omitempty"` // FIXME: NOT AVAILABLE IN LATEST SHOPIFY UPDATES
	CreatedAt         *time.Time `json:"created_at,omitempty"`
}

// StorefrontAccessTokenResource represents the result from the admin/storefront_access_tokens.json endpoint
type StorefrontAccessTokenResource struct {
	StorefrontAccessToken *StorefrontAccessToken `json:"storefront_access_token"`
}

// StorefrontAccessTokensResource is the root object for a storefront access tokens get request.
type StorefrontAccessTokensResource struct {
	StorefrontAccessTokens []StorefrontAccessToken `json:"storefront_access_tokens"`
}

// List storefront access tokens
func (s *StorefrontAccessTokenClient) ListStorefrontAccessToken(options interface{}) ([]StorefrontAccessToken, error) {
	path := fmt.Sprintf("%s.json", storefrontAccessTokensBasePath)
	resource := new(StorefrontAccessTokensResource)
	err := s.client.Get(path, resource, options)
	return resource.StorefrontAccessTokens, err
}

// Create a new storefront access token
func (s *StorefrontAccessTokenClient) CreateStorefrontAccessToken(storefrontAccessToken StorefrontAccessToken) (*StorefrontAccessToken, error) {
	path := fmt.Sprintf("%s.json", storefrontAccessTokensBasePath)
	wrappedData := StorefrontAccessTokenResource{StorefrontAccessToken: &storefrontAccessToken}
	resource := new(StorefrontAccessTokenResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.StorefrontAccessToken, err
}

// Delete an existing storefront access token
func (s *StorefrontAccessTokenClient) DeleteStorefrontAccessToken(ID int64) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", storefrontAccessTokensBasePath, ID))
}
