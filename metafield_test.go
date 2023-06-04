package goshopify

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
)

func MetaFieldTests(t *testing.T, metafield MetaField) {
	// Check that ID is assigned to the returned metafield
	expectedInt := int64(1)
	if metafield.ID != expectedInt {
		t.Errorf("Metafield.ID returned %+v, expected %+v", metafield.ID, expectedInt)
	}
}

func TestMetafieldList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/metafields.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"metafields": [{"id":1},{"id":2}]}`))

	metaFields, err := client.MetaField.ListMetaField(nil)
	if err != nil {
		t.Errorf("MetaField.List returned error: %v", err)
	}

	expected := []MetaField{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(metaFields, expected) {
		t.Errorf("MetaField.List returned %+v, expected %+v", metaFields, expected)
	}
}

func TestMetafieldCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/metafields/count.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"count": 3}`))

	params := map[string]string{"created_at_min": "2016-01-01T00:00:00Z"}
	httpmock.RegisterResponderWithQuery(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/metafields/count.json", client.pathPrefix),
		params,
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.MetaField.CountMetaField(nil)
	if err != nil {
		t.Errorf("MetaField.Count returned error: %v", err)
	}

	expected := 3
	if cnt != expected {
		t.Errorf("MetaField.Count returned %d, expected %d", cnt, expected)
	}

	date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	cnt, err = client.MetaField.CountMetaField(CountOptions{CreatedAtMin: date})
	if err != nil {
		t.Errorf("MetaField.Count returned error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("MetaField.Count returned %d, expected %d", cnt, expected)
	}
}

func TestMetafieldGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/metafields/1.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("metafield.json")))

	metafield, err := client.MetaField.GetMetaField(1, nil)
	if err != nil {
		t.Errorf("MetaField.Get returned error: %v", err)
	}

	createdAt := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC)
	expected := &MetaField{
		ID:                1,
		Key:               "app_key",
		Value:             "app_value",
		ValueType:         "string",
		Type:              "single_line_text_field",
		Namespace:         "affiliates",
		Description:       "some amaaazing app's value",
		OwnerId:           1,
		CreatedAt:         &createdAt,
		UpdatedAt:         &updatedAt,
		OwnerResource:     "shop",
		AdminGraphqlApiID: "gid://shopify/Metafield/1",
	}
	if !reflect.DeepEqual(metafield, expected) {
		t.Errorf("MetaField.Get returned %+v, expected %+v", metafield, expected)
	}
}

func TestMetafieldCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://fooshop.myshopify.com/%s/metafields.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("metafield.json")))

	metafield := MetaField{
		Namespace: "inventory",
		Key:       "warehouse",
		Value:     "25",
		ValueType: "integer",
		Type:      "single_line_text_field",
	}

	returnedMetafield, err := client.MetaField.CreateMetaField(metafield)
	if err != nil {
		t.Errorf("Metafield.Create returned error: %v", err)
	}

	MetaFieldTests(t, *returnedMetafield)
}

func TestMetafieldUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", fmt.Sprintf("https://fooshop.myshopify.com/%s/metafields/1.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("metafield.json")))

	metafield := MetaField{
		ID:        1,
		Value:     "something new",
		ValueType: "string",
		Type:      "single_line_text_field",
	}

	returnedMetafield, err := client.MetaField.UpdateMetaField(metafield)
	if err != nil {
		t.Errorf("Metafield.Update returned error: %v", err)
	}

	MetaFieldTests(t, *returnedMetafield)
}

func TestMetafieldDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://fooshop.myshopify.com/%s/metafields/1.json", client.pathPrefix),
		httpmock.NewStringResponder(200, "{}"))

	err := client.MetaField.DeleteMetaField(1)
	if err != nil {
		t.Errorf("Metafield.Delete returned error: %v", err)
	}
}
