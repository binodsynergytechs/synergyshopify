package synergyshopify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	customCollectionsBasePath     = "custom_collections"
	customCollectionsResourceName = "collections"
)

// CustomCollectionService is an interface for interacting with the custom
// collection endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/customcollection
type CustomCollectionService interface {
	List(interface{}) ([]CustomCollection, error)
	Count(interface{}) (int, error)
	Get(int64, interface{}) (*CustomCollection, error)
	Create(CustomCollection) (*CustomCollection, error)
	Update(CustomCollection) (*CustomCollection, error)
	Delete(int64) error
	ListCollectionWithPagination(interface{}) ([]CustomCollection, *Pagination, error)
	ListWithPaginations(interface{}) ([]CustomCollection, *Pagination, error)

	// MetafieldsService used for CustomCollection resource to communicate with Metafields resource
	MetafieldsService
}

// CustomCollectionServiceOp handles communication with the custom collection
// related methods of the Shopify API.
type CustomCollectionServiceOp struct {
	client *Client
}

// CustomCollection represents a Shopify custom collection.
type CustomCollection struct {
	ID             int64       `json:"id"`
	Handle         string      `json:"handle"`
	Title          string      `json:"title"`
	UpdatedAt      *time.Time  `json:"updated_at"`
	BodyHTML       string      `json:"body_html"`
	SortOrder      string      `json:"sort_order"`
	TemplateSuffix string      `json:"template_suffix"`
	Image          Image       `json:"image"`
	Published      bool        `json:"published"`
	PublishedAt    *time.Time  `json:"published_at"`
	PublishedScope string      `json:"published_scope"`
	Metafields     []Metafield `json:"metafields,omitempty"` // FIXME: Field Not Available Or Deprecated In Latest Shopify Model 23/04
}

// CustomCollectionResource represents the result form the custom_collections/X.json endpoint
type CustomCollectionResource struct {
	Collection *CustomCollection `json:"custom_collection"`
}

// CustomCollectionsResource represents the result from the custom_collections.json endpoint
type CustomCollectionsResource struct {
	Collections []CustomCollection `json:"custom_collections"`
}

// List products
func (s *CustomCollectionServiceOp) ListWithPaginations(options interface{}) ([]CustomCollection, *Pagination, error) {
	products, pagination, err := s.ListCollectionWithPagination(options)
	if err != nil {
		return nil, nil, err
	}
	return products, pagination, nil
}

// createAndDoGetHeaders creates an executes a request while returning the response headers.
func (c *Client) createAndDoGetHeaders(method, relPath string, data, options, resource interface{}) (http.Header, error) {
	if strings.HasPrefix(relPath, "/") {
		// make sure it's a relative path
		relPath = strings.TrimLeft(relPath, "/")
	}

	relPath = path.Join(c.pathPrefix, relPath)
	req, err := c.NewRequest(method, relPath, data, options)
	if err != nil {
		return nil, err
	}

	return c.doGetHeaders(req, resource)
}

// doGetHeaders executes a request, decoding the response into `v` and also returns any response headers.
func (c *Client) doGetHeaders(req *http.Request, v interface{}) (http.Header, error) {
	var resp *http.Response
	var err error
	retries := c.retries
	c.attempts = 0
	c.logRequest(req)

	for {
		c.attempts++
		resp, err = c.Client.Do(req)
		c.logResponse(resp)
		if err != nil {
			return nil, err // http client errors, not api responses
		}

		respErr := CheckResponseError(resp)
		if respErr == nil {
			break // no errors, break out of the retry loop
		}

		// retry scenario, close resp and any continue will retry
		resp.Body.Close()

		if retries <= 1 {
			return nil, respErr
		}

		if rateLimitErr, isRetryErr := respErr.(RateLimitError); isRetryErr {
			// back off and retry

			wait := time.Duration(rateLimitErr.RetryAfter) * time.Second
			c.log.Debugf("rate limited waiting %s", wait.String())
			time.Sleep(wait)
			retries--
			continue
		}

		var doRetry bool
		switch resp.StatusCode {
		case http.StatusServiceUnavailable:
			c.log.Debugf("service unavailable, retrying")
			doRetry = true
			retries--
		}

		if doRetry {
			continue
		}

		// no retry attempts, just return the err
		return nil, respErr
	}

	c.logResponse(resp)
	defer resp.Body.Close()

	if c.apiVersion == defaultApiVersion && resp.Header.Get("X-Shopify-API-Version") != "" {
		// if using stable on first request set the api version
		c.apiVersion = resp.Header.Get("X-Shopify-API-Version")
		c.log.Infof("api version not set, now using %s", c.apiVersion)
	}

	if v != nil {
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&v)
		if err != nil {
			return nil, err
		}
	}

	if s := strings.Split(resp.Header.Get("X-Shopify-Shop-Api-Call-Limit"), "/"); len(s) == 2 {
		c.RateLimits.RequestCount, _ = strconv.Atoi(s[0])
		c.RateLimits.BucketSize, _ = strconv.Atoi(s[1])
	}

	c.RateLimits.RetryAfterSeconds, _ = strconv.ParseFloat(resp.Header.Get("Retry-After"), 64)

	return resp.Header, nil
}

// ListWithPagination lists products and return pagination to retrieve next/previous results.
func (s *CustomCollectionServiceOp) ListCollectionWithPagination(options interface{}) ([]CustomCollection, *Pagination, error) {
	path := fmt.Sprintf("%s.json", customCollectionsBasePath)
	resource := new(CustomCollectionsResource)
	// headers := http.Header{}

	headers, err := s.client.createAndDoGetHeaders("GET", path, nil, options, resource)
	if err != nil {
		return nil, nil, err
	}

	// Extract pagination info from header
	linkHeader := headers.Get("Link")
	// fmt.Println("links:", linkHeader)
	pagination, err := extractPagination(linkHeader)
	if err != nil {
		return nil, nil, err
	}

	return resource.Collections, pagination, nil
}

// List custom collections
func (s *CustomCollectionServiceOp) List(options interface{}) ([]CustomCollection, error) {
	path := fmt.Sprintf("%s.json", customCollectionsBasePath)
	resource := new(CustomCollectionsResource)
	err := s.client.Get(path, resource, options)
	return resource.Collections, err
}

// Count custom collections
func (s *CustomCollectionServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", customCollectionsBasePath)
	return s.client.Count(path, options)
}

// Get individual custom collection
func (s *CustomCollectionServiceOp) Get(collectionID int64, options interface{}) (*CustomCollection, error) {
	path := fmt.Sprintf("%s/%d.json", customCollectionsBasePath, collectionID)
	resource := new(CustomCollectionResource)
	err := s.client.Get(path, resource, options)
	return resource.Collection, err
}

// Create a new custom collection
// See Image for the details of the Image creation for a collection.
func (s *CustomCollectionServiceOp) Create(collection CustomCollection) (*CustomCollection, error) {
	path := fmt.Sprintf("%s.json", customCollectionsBasePath)
	wrappedData := CustomCollectionResource{Collection: &collection}
	resource := new(CustomCollectionResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Collection, err
}

// Update an existing custom collection
func (s *CustomCollectionServiceOp) Update(collection CustomCollection) (*CustomCollection, error) {
	path := fmt.Sprintf("%s/%d.json", customCollectionsBasePath, collection.ID)
	wrappedData := CustomCollectionResource{Collection: &collection}
	resource := new(CustomCollectionResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Collection, err
}

// Delete an existing custom collection.
func (s *CustomCollectionServiceOp) Delete(collectionID int64) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", customCollectionsBasePath, collectionID))
}

// List metafields for a custom collection
func (s *CustomCollectionServiceOp) ListMetafields(customCollectionID int64, options interface{}) ([]Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: customCollectionsResourceName, resourceID: customCollectionID}
	return metafieldService.List(options)
}

// Count metafields for a custom collection
func (s *CustomCollectionServiceOp) CountMetafields(customCollectionID int64, options interface{}) (int, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: customCollectionsResourceName, resourceID: customCollectionID}
	return metafieldService.Count(options)
}

// Get individual metafield for a custom collection
func (s *CustomCollectionServiceOp) GetMetafield(customCollectionID int64, metafieldID int64, options interface{}) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: customCollectionsResourceName, resourceID: customCollectionID}
	return metafieldService.Get(metafieldID, options)
}

// Create a new metafield for a custom collection
func (s *CustomCollectionServiceOp) CreateMetafield(customCollectionID int64, metafield Metafield) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: customCollectionsResourceName, resourceID: customCollectionID}
	return metafieldService.Create(metafield)
}

// Update an existing metafield for a custom collection
func (s *CustomCollectionServiceOp) UpdateMetafield(customCollectionID int64, metafield Metafield) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: customCollectionsResourceName, resourceID: customCollectionID}
	return metafieldService.Update(metafield)
}

// // Delete an existing metafield for a custom collection
func (s *CustomCollectionServiceOp) DeleteMetafield(customCollectionID int64, metafieldID int64) error {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: customCollectionsResourceName, resourceID: customCollectionID}
	return metafieldService.Delete(metafieldID)
}
