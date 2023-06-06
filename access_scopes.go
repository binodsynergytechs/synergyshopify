package goshopify

// AccessScopesRepository represents a repository for managing access scopes.
type AccessScopesRepository interface {
	ListAccessScopes(interface{}) ([]AccessScope, error)
}

// AccessScope represents an access scope in Shopify.
type AccessScope struct {
	Handle      string `json:"handle,omitempty"`
	Description string `json:"description,omitempty"` // TODO: Field Not Available Or Deprecated In Latest Shopify Model 23/04
}

// AccessScopesResponse represents the response from the oauth/access_scopes.json endpoint.
type AccessScopesResponse struct {
	AccessScopes []AccessScope `json:"access_scopes,omitempty"`
}

// AccessScopesClient handles communication with the Access Scopes
// related methods of the Shopify API.
type AccessScopesClient struct {
	client *Client
}

// ListAccessScopes retrieves access scopes based on the used OAuth token.
func (asc *AccessScopesClient) ListAccessScopes(options interface{}) ([]AccessScope, error) {
	// Assuming the path is relative to some base URL.
	path := "/oauth/access_scopes.json"

	// Create a new instance of the resource struct to hold the response.
	resource := new(AccessScopesResponse)

	// Use the HTTP client to send a GET request to the specified path.
	err := asc.client.Get(path, resource, options)

	// Return the access scopes and any potential error from the request.
	return resource.AccessScopes, err
}
