package goshopify

import (
	"fmt"
	"time"
)

const themesBasePath = "themes"

// Options for theme list
type ThemeListOptions struct {
	Role   string `url:"role,omitempty"`
	Fields string `url:"fields,omitempty"`
}

// ThemeRepository is an interface for interfacing with the themes endpoints of the Shopify API.
// See: https://help.shopify.com/api/reference/theme
type ThemeRepository interface {
	ListTheme(interface{}) ([]Theme, error)
	CreateTheme(Theme) (*Theme, error)
	GetTheme(int64, interface{}) (*Theme, error)
	UpdateTheme(Theme) (*Theme, error)
	DeleteTheme(int64) error
}

// ThemeClient handles communication with the theme related methods of
// the Shopify API.
type ThemeClient struct {
	client *Client
}

// Theme represents a Shopify theme
type Theme struct {
	ID                int64      `json:"id"`
	Name              string     `json:"name"`
	AbleToPreview     bool       `json:"previewable"`
	Processing        bool       `json:"processing"`
	Role              string     `json:"role"`
	ThemeStoreID      int64      `json:"theme_store_id"`
	AdminGraphQLApiID string     `json:"admin_graphql_api_id"`
	CreatedAt         *time.Time `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at"`
	ThemeSrc          string     `json:"src"` //TODO: latest added from shopify
}

// ThemesResource is the result from the themes/X.json endpoint
type ThemeResource struct {
	Theme *Theme `json:"theme"`
}

// ThemesResource is the result from the themes.json endpoint
type ThemesResource struct {
	Themes []Theme `json:"themes"`
}

// List all themes
func (s *ThemeClient) ListTheme(options interface{}) ([]Theme, error) {
	path := fmt.Sprintf("%s.json", themesBasePath)
	resource := new(ThemesResource)
	err := s.client.Get(path, resource, options)
	return resource.Themes, err
}

// Update a theme
func (s *ThemeClient) CreateTheme(theme Theme) (*Theme, error) {
	path := fmt.Sprintf("%s.json", themesBasePath)
	wrappedData := ThemeResource{Theme: &theme}
	resource := new(ThemeResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Theme, err
}

// Get a theme
func (s *ThemeClient) GetTheme(themeID int64, options interface{}) (*Theme, error) {
	path := fmt.Sprintf("%s/%d.json", themesBasePath, themeID)
	resource := new(ThemeResource)
	err := s.client.Get(path, resource, options)
	return resource.Theme, err
}

// Update a theme
func (s *ThemeClient) UpdateTheme(theme Theme) (*Theme, error) {
	path := fmt.Sprintf("%s/%d.json", themesBasePath, theme.ID)
	wrappedData := ThemeResource{Theme: &theme}
	resource := new(ThemeResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Theme, err
}

// Delete a theme
func (s *ThemeClient) DeleteTheme(themeID int64) error {
	path := fmt.Sprintf("%s/%d.json", themesBasePath, themeID)
	return s.client.Delete(path)
}
