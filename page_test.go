package goshopify

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
)

func pageTests(t *testing.T, page Page) {
	// Check that ID is assigned to the returned page
	expectedInt := int64(1)
	if page.ID != expectedInt {
		t.Errorf("Page.ID returned %+v, expected %+v", page.ID, expectedInt)
	}
}

func TestPageList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"pages": [{"id":1},{"id":2}]}`))

	pages, err := client.Page.ListPage(nil)
	if err != nil {
		t.Errorf("Page.List returned error: %v", err)
	}

	expected := []Page{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(pages, expected) {
		t.Errorf("Page.List returned %+v, expected %+v", pages, expected)
	}
}

func TestPageCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/count.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"count": 3}`))

	params := map[string]string{"created_at_min": "2016-01-01T00:00:00Z"}
	httpmock.RegisterResponderWithQuery(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/count.json", client.pathPrefix),
		params,
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.Page.CountPage(nil)
	if err != nil {
		t.Errorf("Page.Count returned error: %v", err)
	}

	expected := 3
	if cnt != expected {
		t.Errorf("Page.Count returned %d, expected %d", cnt, expected)
	}

	date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	cnt, err = client.Page.CountPage(CountOptions{CreatedAtMin: date})
	if err != nil {
		t.Errorf("Page.Count returned error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("Page.Count returned %d, expected %d", cnt, expected)
	}
}

func TestPageGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"page": {"id":1}}`))

	page, err := client.Page.GetPage(1, nil)
	if err != nil {
		t.Errorf("Page.Get returned error: %v", err)
	}

	expected := &Page{ID: 1}
	if !reflect.DeepEqual(page, expected) {
		t.Errorf("Page.Get returned %+v, expected %+v", page, expected)
	}
}

func TestPageCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("page.json")))

	page := Page{
		Title:    "404",
		BodyHTML: "<strong>NOT FOUND!<\\/strong>",
	}

	returnedPage, err := client.Page.CreatePage(page)
	if err != nil {
		t.Errorf("Page.Create returned error: %v", err)
	}

	pageTests(t, *returnedPage)
}

func TestPageUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("page.json")))

	page := Page{
		ID: 1,
	}

	returnedPage, err := client.Page.UpdatePage(page)
	if err != nil {
		t.Errorf("Page.Update returned error: %v", err)
	}

	pageTests(t, *returnedPage)
}

func TestPageDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1.json", client.pathPrefix),
		httpmock.NewStringResponder(200, "{}"))

	err := client.Page.DeletePage(1)
	if err != nil {
		t.Errorf("Page.Delete returned error: %v", err)
	}
}

func TestPageListMetaFields(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1/metaFields.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"metaFields": [{"id":1},{"id":2}]}`))

	metaFields, err := client.Page.ListMetaFields(1, nil)
	if err != nil {
		t.Errorf("Page.ListMetaFields() returned error: %v", err)
	}

	expected := []MetaField{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(metaFields, expected) {
		t.Errorf("Page.ListMetaFields() returned %+v, expected %+v", metaFields, expected)
	}
}

func TestPageCountMetaFields(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1/metaFields/count.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"count": 3}`))

	params := map[string]string{"created_at_min": "2016-01-01T00:00:00Z"}
	httpmock.RegisterResponderWithQuery(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1/metaFields/count.json", client.pathPrefix),
		params,
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.Page.CountMetaFields(1, nil)
	if err != nil {
		t.Errorf("Page.CountMetaFields() returned error: %v", err)
	}

	expected := 3
	if cnt != expected {
		t.Errorf("Page.CountMetaFields() returned %d, expected %d", cnt, expected)
	}

	date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	cnt, err = client.Page.CountMetaFields(1, CountOptions{CreatedAtMin: date})
	if err != nil {
		t.Errorf("Page.CountMetaFields() returned error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("Page.CountMetaFields() returned %d, expected %d", cnt, expected)
	}
}

func TestPageGetMetaField(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1/metaFields/2.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"metaField": {"id":2}}`))

	metaField, err := client.Page.GetMetaField(1, 2, nil)
	if err != nil {
		t.Errorf("Page.GetMetaField() returned error: %v", err)
	}

	expected := &MetaField{ID: 2}
	if !reflect.DeepEqual(metaField, expected) {
		t.Errorf("Page.GetMetaField() returned %+v, expected %+v", metaField, expected)
	}
}

func TestPageCreateMetaField(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1/metaFields.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("metaField.json")))

	metaField := MetaField{
		Key:       "app_key",
		Value:     "app_value",
		ValueType: "string",
		Type:      "single_line_text_field",
		Namespace: "affiliates",
	}

	returnedMetaField, err := client.Page.CreateMetaField(1, metaField)
	if err != nil {
		t.Errorf("Page.CreateMetaField() returned error: %v", err)
	}

	MetaFieldTests(t, *returnedMetaField)
}

func TestPageUpdateMetaField(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1/metaFields/2.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("metaField.json")))

	metaField := MetaField{
		ID:        2,
		Key:       "app_key",
		Value:     "app_value",
		ValueType: "string",
		Type:      "single_line_text_field",
		Namespace: "affiliates",
	}

	returnedMetaField, err := client.Page.UpdateMetaField(1, metaField)
	if err != nil {
		t.Errorf("Page.UpdateMetaField() returned error: %v", err)
	}

	MetaFieldTests(t, *returnedMetaField)
}

func TestPageDeleteMetaField(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://fooshop.myshopify.com/%s/pages/1/metaFields/2.json", client.pathPrefix),
		httpmock.NewStringResponder(200, "{}"))

	err := client.Page.DeleteMetaField(1, 2)
	if err != nil {
		t.Errorf("Page.DeleteMetaField() returned error: %v", err)
	}
}
