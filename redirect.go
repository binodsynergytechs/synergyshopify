package goshopify

import (
	"fmt"
)

const redirectsBasePath = "redirects"

// RedirectRepository is an interface for interacting with the redirects
// endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/online_store/redirect
type RedirectRepository interface {
	ListRedirect(interface{}) ([]Redirect, error)
	CountRedirect(interface{}) (int, error)
	GetRedirect(int64, interface{}) (*Redirect, error)
	CreateRedirect(Redirect) (*Redirect, error)
	UpdateRedirect(Redirect) (*Redirect, error)
	DeleteRedirect(int64) error
}

// RedirectClient handles communication with the redirect related methods of the
// Shopify API.
type RedirectClient struct {
	client *Client
}

// Redirect represents a Shopify redirect.
type Redirect struct {
	ID     int64  `json:"id"`
	Path   string `json:"path"`
	Target string `json:"target"`
}

// RedirectResource represents the result from the redirects/X.json endpoint
type RedirectResource struct {
	Redirect *Redirect `json:"redirect"`
}

// RedirectsResource represents the result from the redirects.json endpoint
type RedirectsResource struct {
	Redirects []Redirect `json:"redirects"`
}

// List redirects
func (rc *RedirectClient) ListRedirect(options interface{}) ([]Redirect, error) {
	path := fmt.Sprintf("%s.json", redirectsBasePath)
	resource := new(RedirectsResource)
	err := rc.client.Get(path, resource, options)
	return resource.Redirects, err
}

// Count redirects
func (rc *RedirectClient) CountRedirect(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", redirectsBasePath)
	return rc.client.Count(path, options)
}

// Get individual redirect
func (rc *RedirectClient) GetRedirect(redirectID int64, options interface{}) (*Redirect, error) {
	path := fmt.Sprintf("%s/%d.json", redirectsBasePath, redirectID)
	resource := new(RedirectResource)
	err := rc.client.Get(path, resource, options)
	return resource.Redirect, err
}

// Create a new redirect
func (rc *RedirectClient) CreateRedirect(redirect Redirect) (*Redirect, error) {
	path := fmt.Sprintf("%s.json", redirectsBasePath)
	wrappedData := RedirectResource{Redirect: &redirect}
	resource := new(RedirectResource)
	err := rc.client.Post(path, wrappedData, resource)
	return resource.Redirect, err
}

// Update an existing redirect
func (rc *RedirectClient) UpdateRedirect(redirect Redirect) (*Redirect, error) {
	path := fmt.Sprintf("%s/%d.json", redirectsBasePath, redirect.ID)
	wrappedData := RedirectResource{Redirect: &redirect}
	resource := new(RedirectResource)
	err := rc.client.Put(path, wrappedData, resource)
	return resource.Redirect, err
}

// Delete an existing redirect.
func (rc *RedirectClient) DeleteRedirect(redirectID int64) error {
	return rc.client.Delete(fmt.Sprintf("%s/%d.json", redirectsBasePath, redirectID))
}
