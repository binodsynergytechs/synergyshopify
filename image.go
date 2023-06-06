package goshopify

import (
	"fmt"
	"time"
)

// ImageRepository is an interface for interacting with the image endpoints
// of the Shopify API.
// See https://help.shopify.com/api/reference/product_image
type ImageRepository interface {
	ListImage(int64, interface{}) ([]Image, error)
	CountImage(int64, interface{}) (int, error)
	GetImage(int64, int64, interface{}) (*Image, error)
	CreateImage(int64, Image) (*Image, error)
	UpdateImage(int64, Image) (*Image, error)
	DeleteImage(int64, int64) error
}

// ImageClient handles communication with the image related methods of
// the Shopify API.
type ImageClient struct {
	client *Client
}

// Image represents a Shopify product's image.
type Image struct {
	ID         int64      `json:"id,omitempty"`
	ProductID  int64      `json:"product_id,omitempty"`
	Position   int        `json:"position,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	Width      int        `json:"width,omitempty"`
	Height     int        `json:"height,omitempty"`
	Src        string     `json:"src,omitempty"`
	Attachment string     `json:"attachment,omitempty"` //FIXME: Not Available In Latest Shopify Model
	Filename   string     `json:"filename,omitempty"`   //FIXME: Not Available In Latest Shopify Model
	VariantIds []int64    `json:"variant_ids,omitempty"`
}

// SingleImageResponse represents the result form the products/X/images/Y.json endpoint
type SingleImageResponse struct {
	Image *Image `json:"image"`
}

// MultipleImagesResponse represents the result from the products/X/images.json endpoint
type MultipleImagesResponse struct {
	Images []Image `json:"images"`
}

// ListImage images
func (ic *ImageClient) ListImage(productID int64, options interface{}) ([]Image, error) {
	path := fmt.Sprintf("%s/%d/images.json", productsBasePath, productID)
	resource := new(MultipleImagesResponse)
	err := ic.client.Get(path, resource, options)
	return resource.Images, err
}

// Count images
func (ic *ImageClient) CountImage(productID int64, options interface{}) (int, error) {
	path := fmt.Sprintf("%s/%d/images/count.json", productsBasePath, productID)
	return ic.client.Count(path, options)
}

// Get individual image
func (ic *ImageClient) GetImage(productID int64, imageID int64, options interface{}) (*Image, error) {
	path := fmt.Sprintf("%s/%d/images/%d.json", productsBasePath, productID, imageID)
	resource := new(SingleImageResponse)
	err := ic.client.Get(path, resource, options)
	return resource.Image, err
}

// Create a new image
//
// There are 2 methods of creating an image in Shopify:
// 1. Src
// 2. Filename and Attachment
//
// If both Image.Filename and Image.Attachment are supplied,
// then Image.Src is not needed.  And vice versa.
//
// If both Image.Attachment and Image.Src are provided,
// Shopify will take the attachment.
//
// Shopify will accept Image.Attachment without Image.Filename.
func (ic *ImageClient) CreateImage(productID int64, image Image) (*Image, error) {
	path := fmt.Sprintf("%s/Image%d/images.json", productsBasePath, productID)
	wrappedData := SingleImageResponse{Image: &image}
	resource := new(SingleImageResponse)
	err := ic.client.Post(path, wrappedData, resource)
	return resource.Image, err
}

// Update an existing image
func (ic *ImageClient) UpdateImage(productID int64, image Image) (*Image, error) {
	path := fmt.Sprintf("%s/%d/images/%d.json", productsBasePath, productID, image.ID)
	wrappedData := SingleImageResponse{Image: &image}
	resource := new(SingleImageResponse)
	err := ic.client.Put(path, wrappedData, resource)
	return resource.Image, err
}

// Delete an existing image
func (ic *ImageClient) DeleteImage(productID int64, imageID int64) error {
	return ic.client.Delete(fmt.Sprintf("%s/%d/images/%d.json", productsBasePath, productID, imageID))
}
