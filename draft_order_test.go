package goshopify

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/shopspring/decimal"
)

func draftOrderTests(t *testing.T, draftOrder DraftOrder) {
	// Check that dates are parsed
	d := time.Date(2019, time.April, 9, 10, 02, 43, 0, time.UTC)
	if !d.Equal(*draftOrder.CreatedAt) {
		t.Errorf("Order.CreatedAt returned %+v, expected %+v", draftOrder.CreatedAt, d)
	}

	// Check null dates
	if draftOrder.UpdatedAt == nil {
		t.Errorf("DraftOrder.UpdatedAt returned %+v, expected %+v", draftOrder.UpdatedAt, nil)
	}

	// Check prices
	p := "206.25"
	if !(p == draftOrder.TotalPrice) {
		t.Errorf("draftOrder.TotalPrice returned %+v, expected %+v", draftOrder.TotalPrice, p)
	}

	// Check null prices, notice that prices are usually not empty.
	if draftOrder.TotalTax != "0.00" {
		t.Errorf("draftOrder.TotalTax returned %+v, expected %+v", draftOrder.TotalTax, nil)
	}

	//

	// Check customer
	if draftOrder.Customer == nil {
		t.Error("Expected Customer to not be nil")
	}
	if draftOrder.Customer.Email != "bob.norman@hostmail.com" {
		t.Errorf("Customer.Email, expected %v, actual %v", "bob.norman@hostmail.com", draftOrder.Customer.Email)
	}

	ptp := decimal.NewFromFloat(199)
	lineItem := draftOrder.LineItems[0]
	if !ptp.Equals(*lineItem.Price) {
		t.Errorf("DraftOrder.LineItems[0].Price, expected %v, actual %v", "199.00", lineItem.Price)
	}
}

func TestDraftOrderGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/draft_orders/994118539.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("draft_order.json")))

	draftOrder, err := client.DraftOrder.GetDraftOrder(994118539, nil)
	if err != nil {
		t.Errorf("DraftOrder.Get returned error: %v", err)
	}
	draftOrderTests(t, *draftOrder)
}

func TestDraftOrderCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://fooshop.myshopify.com/%s/draft_orders.json", client.pathPrefix),
		httpmock.NewStringResponder(201, `{"draft_order":{"id": 1}}`))

	draftOrder := DraftOrder{
		LineItems: []LineItem{
			{
				VariantID: 1,
				Quantity:  1,
			},
		},
	}

	d, err := client.DraftOrder.CreateDraftOrder(draftOrder)
	if err != nil {
		t.Errorf("DraftOrder.Create returned error: %v", err)
	}

	expected := DraftOrder{ID: 1}
	if d.ID != expected.ID {
		t.Errorf("DraftOrder.Create returned id %d, expected %d", d.ID, expected.ID)
	}
}

func TestDraftOrderUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", fmt.Sprintf("https://fooshop.myshopify.com/%s/draft_orders/1.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"draft_order":{"id": 1}}`))

	draftOrder := DraftOrder{
		ID:            1,
		Note:          "slow order",
		TaxesIncluded: true,
	}

	d, err := client.DraftOrder.UpdateDraftOrder(draftOrder)
	if err != nil {
		t.Errorf("DraftOrder.Create returned an error %v", err)
	}

	expected := DraftOrder{ID: 1}
	if d.ID != expected.ID {
		t.Errorf("DraftOrder.Update returned id %d, expected %d", d.ID, expected.ID)
	}

}

func TestDraftOrderCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/draft_orders/count.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"count": 7}`))

	params := map[string]string{"status": "open"}
	httpmock.RegisterResponderWithQuery(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/draft_orders/count.json", client.pathPrefix),
		params,
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.DraftOrder.CountDraftOrder(nil)
	if err != nil {
		t.Errorf("DraftOrder.Count returned an error: %v", err)
	}
	expected := 7
	if cnt != expected {
		t.Errorf("DraftOrder.Count returned %d, expected %d", cnt, expected)
	}

	status := "open"
	cnt, err = client.DraftOrder.CountDraftOrder(DraftOrderCountOptions{Status: status})
	if err != nil {
		t.Errorf("DraftOrder.Count returned an error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("DraftOrder.Count returned %d, expected %d", cnt, expected)
	}
}

func TestDraftOrderList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/draft_orders.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("draft_orders.json")))

	draftOrders, err := client.DraftOrder.ListDraftOrder(nil)
	if err != nil {
		t.Errorf("DraftOrder.List returned error: %v", err)
	}

	if len(draftOrders) != 1 {
		t.Errorf("DraftOrder.List got %d orders, expected: 1", len(draftOrders))
	}
	draftOrder := draftOrders[0]
	draftOrderTests(t, draftOrder)
}

func TestDraftOrderListOptions(t *testing.T) {
	setup()
	defer teardown()
	params := map[string]string{
		"fields": "id,name",
		"limit":  "250",
		"status": "any",
	}
	httpmock.RegisterResponderWithQuery(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/draft_orders.json", client.pathPrefix),
		params,
		httpmock.NewBytesResponder(200, loadFixture("draft_orders.json")))

	options := DraftOrderListOptions{
		Limit:  250,
		Status: "any",
		Fields: "id,name",
	}

	draftOrders, err := client.DraftOrder.ListDraftOrder(options)
	if err != nil {
		t.Errorf("DraftOrder.List returned error: %v", err)
	}

	if len(draftOrders) != 1 {
		t.Errorf("DraftOrder.List got %d orders, expected: 1", len(draftOrders))
	}

	draftOrder := draftOrders[0]
	draftOrderTests(t, draftOrder)
}

func TestDraftOrderInvoice(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/draft_orders/1/send_invoice.json", client.pathPrefix),
		httpmock.NewBytesResponder(201, loadFixture("invoice.json")))
	invoice := DraftOrderInvoice{
		To:   "first@example.com",
		From: "steve@apple.com",
		Bcc: []string{
			"steve@apple.com",
		},
		Subject:       "Apple Computer Invoice",
		CustomMessage: "Thank you for ordering!",
	}
	draftInvoice, err := client.DraftOrder.InvoiceDraftOrder(1, invoice)
	if err != nil {
		t.Errorf("DraftOrder.Invoice returned an error: %v", err)
	}

	if !reflect.DeepEqual(*draftInvoice, invoice) {
		t.Errorf("DraftOrder.Invoice returned %+v, expected %+v,", draftInvoice, invoice)
	}
}

func TestDraftOrderDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"DELETE",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/draft_orders/1.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, nil))

	err := client.DraftOrder.DeleteDraftOrder(1)
	if err != nil {
		t.Errorf("DraftOrder.Delete returned an error %v", err)
	}
}
func TestDraftOrderComplete(t *testing.T) {
	setup()
	defer teardown()
	params := map[string]string{"payment_pending": "false"}
	httpmock.RegisterResponderWithQuery(
		"PUT",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/draft_orders/1/complete.json", client.pathPrefix),
		params,
		httpmock.NewBytesResponder(200, loadFixture("draft_order.json")))

	draftOrder, err := client.DraftOrder.CompleteDraftOrder(1, false)
	if err != nil {
		t.Errorf("DraftOrder.Complete returned an error %v", err)
	}
	draftOrderTests(t, *draftOrder)
}

func TestDraftOrderCompletePending(t *testing.T) {
	setup()
	defer teardown()
	params := map[string]string{"payment_pending": "true"}
	httpmock.RegisterResponderWithQuery(
		"PUT",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/draft_orders/1/complete.json", client.pathPrefix),
		params,
		httpmock.NewBytesResponder(200, loadFixture("draft_order.json")))

	draftOrder, err := client.DraftOrder.CompleteDraftOrder(1, true)
	if err != nil {
		t.Errorf("DraftOrder.Complete returned an error %v", err)
	}
	draftOrderTests(t, *draftOrder)
}

func TestDraftOrderListMetaFields(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/draft_orders/1/metaFields.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"metaFields": [{"id":1},{"id":2}]}`))

	metaFields, err := client.DraftOrder.ListMetaFields(1, nil)
	if err != nil {
		t.Errorf("DraftOrder.ListMetaFields() returned error: %v", err)
	}

	expected := []MetaField{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(metaFields, expected) {
		t.Errorf("Order.ListMetaFields() returned %+v, expected %+v", metaFields, expected)
	}
}

func TestDraftOrderCountMetaFields(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/draft_orders/1/metaFields/count.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"count": 3}`))

	params := map[string]string{"created_at_min": "2016-01-01T00:00:00Z"}
	httpmock.RegisterResponderWithQuery(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/draft_orders/1/metaFields/count.json", client.pathPrefix),
		params,
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.DraftOrder.CountMetaFields(1, nil)
	if err != nil {
		t.Errorf("Order.CountMetaFields() returned error: %v", err)
	}

	expected := 3
	if cnt != expected {
		t.Errorf("Order.CountMetaFields() returned %d, expected %d", cnt, expected)
	}

	date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	cnt, err = client.DraftOrder.CountMetaFields(1, CountOptions{CreatedAtMin: date})
	if err != nil {
		t.Errorf("Order.CountMetaFields() returned error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("Order.CountMetaFields() returned %d, expected %d", cnt, expected)
	}
}

func TestDraftOrderGetMetaField(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/draft_orders/1/metaFields/2.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"metaField": {"id":2}}`))

	metaField, err := client.DraftOrder.GetMetaField(1, 2, nil)
	if err != nil {
		t.Errorf("Order.GetMetaField() returned error: %v", err)
	}

	expected := &MetaField{ID: 2}
	if !reflect.DeepEqual(metaField, expected) {
		t.Errorf("Order.GetMetaField() returned %+v, expected %+v", metaField, expected)
	}
}

func TestDraftOrderCreateMetaField(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://fooshop.myshopify.com/%s/draft_orders/1/metaFields.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("metaField.json")))

	metaField := MetaField{
		Key:       "app_key",
		Value:     "app_value",
		ValueType: "string",
		Namespace: "affiliates",
	}

	returnedMetaField, err := client.DraftOrder.CreateMetaField(1, metaField)
	if err != nil {
		t.Errorf("Order.CreateMetaField() returned error: %v", err)
	}

	MetaFieldTests(t, *returnedMetaField)
}

func TestDraftOrderUpdateMetaField(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", fmt.Sprintf("https://fooshop.myshopify.com/%s/draft_orders/1/metaFields/2.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("metaField.json")))

	metaField := MetaField{
		ID:        2,
		Key:       "app_key",
		Value:     "app_value",
		ValueType: "string",
		Namespace: "affiliates",
	}

	returnedMetaField, err := client.DraftOrder.UpdateMetaField(1, metaField)
	if err != nil {
		t.Errorf("Order.UpdateMetaField() returned error: %v", err)
	}

	MetaFieldTests(t, *returnedMetaField)
}

func TestDraftOrderDeleteMetaField(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://fooshop.myshopify.com/%s/draft_orders/1/metaFields/2.json", client.pathPrefix),
		httpmock.NewStringResponder(200, "{}"))

	err := client.DraftOrder.DeleteMetaField(1, 2)
	if err != nil {
		t.Errorf("Order.DeleteMetaField() returned error: %v", err)
	}
}
