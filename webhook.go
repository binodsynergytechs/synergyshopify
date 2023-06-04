package goshopify

import (
	"fmt"
	"time"
)

const webhooksBasePath = "webhooks"

// WebhookRepository is an interface for interfacing with the webhook endpoints of
// the Shopify API.
// See: https://help.shopify.com/api/reference/webhook
type WebhookRepository interface {
	ListWebhook(interface{}) ([]Webhook, error)
	CountWebhook(interface{}) (int, error)
	GetWebhook(int64, interface{}) (*Webhook, error)
	CreateWebhook(Webhook) (*Webhook, error)
	UpdateWebhook(Webhook) (*Webhook, error)
	DeleteWebhook(int64) error
}

// WebhookClient handles communication with the webhook-related methods of
// the Shopify API.
type WebhookClient struct {
	client *Client
}

// Webhook represents a Shopify webhook
type Webhook struct {
	ID                         int64      `json:"id"`
	Address                    string     `json:"address"`
	Topic                      string     `json:"topic"`
	Format                     string     `json:"format"`
	CreatedAt                  *time.Time `json:"created_at,omitempty"`
	UpdatedAt                  *time.Time `json:"updated_at,omitempty"`
	Fields                     []string   `json:"fields"`
	MetafieldNamespaces        []string   `json:"metafield_namespaces"`
	PrivateMetafieldNamespaces []string   `json:"private_metafield_namespaces"`
	ApiVersion                 string     `json:"api_version,omitempty"`
}

// WebhookOptions can be used for filtering webhooks on a List request.
type WebhookOptions struct {
	Address string `url:"address,omitempty"`
	Topic   string `url:"topic,omitempty"`
}

// SingleWebhookResponse represents the result from the admin/webhooks.json endpoint
type SingleWebhookResponse struct {
	Webhook *Webhook `json:"webhook"`
}

// MultipleWebhooksResponse is the root object for a webhook get request.
type MultipleWebhooksResponse struct {
	Webhooks []Webhook `json:"webhooks"`
}

// List webhooks
func (wc *WebhookClient) ListWebhook(options interface{}) ([]Webhook, error) {
	path := fmt.Sprintf("%s.json", webhooksBasePath)
	resource := new(MultipleWebhooksResponse)
	err := wc.client.Get(path, resource, options)
	return resource.Webhooks, err
}

// Count webhooks
func (wc *WebhookClient) CountWebhook(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", webhooksBasePath)
	return wc.client.Count(path, options)
}

// Get individual webhook
func (wc *WebhookClient) GetWebhook(webhookdID int64, options interface{}) (*Webhook, error) {
	path := fmt.Sprintf("%s/%d.json", webhooksBasePath, webhookdID)
	resource := new(SingleWebhookResponse)
	err := wc.client.Get(path, resource, options)
	return resource.Webhook, err
}

// Create a new webhook
func (wc *WebhookClient) CreateWebhook(webhook Webhook) (*Webhook, error) {
	path := fmt.Sprintf("%s.json", webhooksBasePath)
	wrappedData := SingleWebhookResponse{Webhook: &webhook}
	resource := new(SingleWebhookResponse)
	err := wc.client.Post(path, wrappedData, resource)
	return resource.Webhook, err
}

// Update an existing webhook.
func (wc *WebhookClient) UpdateWebhook(webhook Webhook) (*Webhook, error) {
	path := fmt.Sprintf("%s/%d.json", webhooksBasePath, webhook.ID)
	wrappedData := SingleWebhookResponse{Webhook: &webhook}
	resource := new(SingleWebhookResponse)
	err := wc.client.Put(path, wrappedData, resource)
	return resource.Webhook, err
}

// Delete an existing webhooks
func (wc *WebhookClient) DeleteWebhook(ID int64) error {
	return wc.client.Delete(fmt.Sprintf("%s/%d.json", webhooksBasePath, ID))
}
