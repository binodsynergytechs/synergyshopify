package goshopify

import (
	"fmt"
	"time"
)

const blogsBasePath = "blogs"

// BlogRepository is an interface for interfacing with the blogs endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/online_store/blog
type BlogRepository interface {
	ListBlog(interface{}) ([]Blog, error)
	CountBlog(interface{}) (int, error)
	GetBlog(int64, interface{}) (*Blog, error)
	CreateBlog(Blog) (*Blog, error)
	UpdateBlog(Blog) (*Blog, error)
	DeleteBlog(int64) error
}

// BlogClient handles communication with the blog related methods of
// the Shopify API.
type BlogClient struct {
	client *Client
}

// Blog represents a Shopify blog
type Blog struct {
	ID                 int64      `json:"id"`
	Title              string     `json:"title"`
	AbleToComment      string     `json:"commentable"`
	FeedBurner         string     `json:"feedburner"`
	FeedBurnerLocation string     `json:"feedburner_location"`
	Handle             string     `json:"handle"`
	MetaField          MetaField  `json:"metafield"` // FIXME: Field Not Available Or Deprecated In Latest Shopify Model 23/04
	Tags               string     `json:"tags"`
	TemplateSuffix     string     `json:"template_suffix"`
	CreatedAt          *time.Time `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
	AdminGraphqlApiID  string     `json:"admin_graphql_api_id,omitempty"`
}

// MultipleBlogsResponse is the result from the blogs.json endpoint
type MultipleBlogsResponse struct {
	Blogs []Blog `json:"blogs"`
}

// Represents the result from the blogs/X.json endpoint
type SingleBlogResponse struct {
	Blog *Blog `json:"blog"`
}

// List all blogs
func (bc *BlogClient) ListBlog(options interface{}) ([]Blog, error) {
	path := fmt.Sprintf("%s.json", blogsBasePath)
	resource := new(MultipleBlogsResponse)
	err := bc.client.Get(path, resource, options)
	return resource.Blogs, err
}

// Count blogs
func (bc *BlogClient) CountBlog(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", blogsBasePath)
	return bc.client.Count(path, options)
}

// Get single blog
func (bc *BlogClient) GetBlog(blogId int64, options interface{}) (*Blog, error) {
	path := fmt.Sprintf("%s/%d.json", blogsBasePath, blogId)
	resource := new(SingleBlogResponse)
	err := bc.client.Get(path, resource, options)
	return resource.Blog, err
}

// Create a new blog
func (bc *BlogClient) CreateBlog(blog Blog) (*Blog, error) {
	path := fmt.Sprintf("%s.json", blogsBasePath)
	wrappedData := SingleBlogResponse{Blog: &blog}
	resource := new(SingleBlogResponse)
	err := bc.client.Post(path, wrappedData, resource)
	return resource.Blog, err
}

// Update an existing blog
func (bc *BlogClient) UpdateBlog(blog Blog) (*Blog, error) {
	path := fmt.Sprintf("%s/%d.json", blogsBasePath, blog.ID)
	wrappedData := SingleBlogResponse{Blog: &blog}
	resource := new(SingleBlogResponse)
	err := bc.client.Put(path, wrappedData, resource)
	return resource.Blog, err
}

// Delete an blog
func (bc *BlogClient) Delete(blogId int64) error {
	return bc.client.Delete(fmt.Sprintf("%s/%d.json", blogsBasePath, blogId))
}
