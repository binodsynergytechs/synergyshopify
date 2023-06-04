package goshopify

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/shopspring/decimal"
)

func TestCustomerList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/customers.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"customers": [{"id":1},{"id":2}]}`))

	customers, err := client.Customer.ListCustomer(nil)
	if err != nil {
		t.Errorf("Customer.List returned error: %v", err)
	}

	expected := []Customer{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(customers, expected) {
		t.Errorf("Customer.List returned %+v, expected %+v", customers, expected)
	}
}

func TestCustomerListWithPagination(t *testing.T) {
	setup()
	defer teardown()

	listURL := fmt.Sprintf("https://fooshop.myshopify.com/%s/customers.json", client.pathPrefix)

	// The strconv.Atoi error changed in go 1.8, 1.7 is still being tested/supported.
	limitConversionErrorMessage := `strconv.Atoi: parsing "invalid": invalid syntax`
	if runtime.Version()[2:5] == "1.7" {
		limitConversionErrorMessage = `strconv.ParseInt: parsing "invalid": invalid syntax`
	}

	cases := []struct {
		body               string
		linkHeader         string
		expectedCustomers  []Customer
		expectedPagination *Pagination
		expectedErr        error
	}{
		// Expect empty pagination when there is no link header
		{
			`{"customers": [{"id":1},{"id":2}]}`,
			"",
			[]Customer{{ID: 1}, {ID: 2}},
			new(Pagination),
			nil,
		},
		// Invalid link header responses
		{
			"{}",
			"invalid link",
			[]Customer(nil),
			nil,
			ResponseDecodingError{Message: "could not extract pagination link header"},
		},
		{
			"{}",
			`<:invalid.url>; rel="next"`,
			[]Customer(nil),
			nil,
			ResponseDecodingError{Message: "pagination does not contain a valid URL"},
		},
		{
			"{}",
			`<http://valid.url?%invalid_query>; rel="next"`,
			[]Customer(nil),
			nil,
			errors.New(`invalid URL escape "%in"`),
		},
		{
			"{}",
			`<http://valid.url>; rel="next"`,
			[]Customer(nil),
			nil,
			ResponseDecodingError{Message: "page_info is missing"},
		},
		{
			"{}",
			`<http://valid.url?page_info=foo&limit=invalid>; rel="next"`,
			[]Customer(nil),
			nil,
			errors.New(limitConversionErrorMessage),
		},
		// Valid link header responses
		{
			`{"customers": [{"id":1}]}`,
			`<http://valid.url?page_info=foo&limit=2>; rel="next"`,
			[]Customer{{ID: 1}},
			&Pagination{
				NextPageOptions: &ListOptions{PageInfo: "foo", Limit: 2},
			},
			nil,
		},
		{
			`{"customers": [{"id":2}]}`,
			`<http://valid.url?page_info=foo>; rel="next", <http://valid.url?page_info=bar>; rel="previous"`,
			[]Customer{{ID: 2}},
			&Pagination{
				NextPageOptions:     &ListOptions{PageInfo: "foo"},
				PreviousPageOptions: &ListOptions{PageInfo: "bar"},
			},
			nil,
		},
	}
	for i, c := range cases {
		response := &http.Response{
			StatusCode: 200,
			Body:       httpmock.NewRespBodyFromString(c.body),
			Header: http.Header{
				"Link": {c.linkHeader},
			},
		}

		httpmock.RegisterResponder("GET", listURL, httpmock.ResponderFromResponse(response))

		customers, pagination, err := client.Customer.ListWithPaginationCustomer(nil)
		if !reflect.DeepEqual(customers, c.expectedCustomers) {
			t.Errorf("test %d Customer.ListWithPagination customers returned %+v, expected %+v", i, customers, c.expectedCustomers)
		}

		if !reflect.DeepEqual(pagination, c.expectedPagination) {
			t.Errorf(
				"test %d Customer.ListWithPagination pagination returned %+v, expected %+v",
				i,
				pagination,
				c.expectedPagination,
			)
		}

		if (c.expectedErr != nil || err != nil) && err.Error() != c.expectedErr.Error() {
			t.Errorf(
				"test %d Customer.ListWithPagination err returned %+v, expected %+v",
				i,
				err,
				c.expectedErr,
			)
		}
	}
}

func TestCustomerCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/customers/count.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"count": 5}`))

	params := map[string]string{"created_at_min": "2016-01-01T00:00:00Z"}
	httpmock.RegisterResponderWithQuery(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/customers/count.json", client.pathPrefix),
		params,
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.Customer.CountCustomer(nil)
	if err != nil {
		t.Errorf("Customer.Count returned error: %v", err)
	}

	expected := 5
	if cnt != expected {
		t.Errorf("Customer.Count returned %d, expected %d", cnt, expected)
	}

	date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	cnt, err = client.Customer.CountCustomer(CountOptions{CreatedAtMin: date})
	if err != nil {
		t.Errorf("Customer.Count returned error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("Customer.Count returned %d, expected %d", cnt, expected)
	}
}

func TestCustomerSearch(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/customers/search.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"customers": [{"id":1},{"id":2}]}`))

	customers, err := client.Customer.SearchCustomer(nil)
	if err != nil {
		t.Errorf("Customer.Search returned error: %v", err)
	}

	expected := []Customer{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(customers, expected) {
		t.Errorf("Customer.Search returned %+v, expected %+v", customers, expected)
	}
}

func TestCustomerGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/customers/1.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("customer.json")))

	customer, err := client.Customer.GetCustomer(1, nil)
	if err != nil {
		t.Errorf("Customer.Get returned error: %v", err)
	}

	address1 := &CustomerAddress{ID: 1, CustomerID: 1, FirstName: "Test", LastName: "Citizen", Company: "",
		Address1: "1 Smith St", Address2: "", City: "BRISBANE", Province: "Queensland", Country: "Australia",
		Zip: "4000", Phone: "1111 111 111", Name: "Test Citizen", ProvinceCode: "QLD", CountryCode: "AU",
		CountryName: "Australia", Default: true}
	createdAt := time.Date(2017, time.September, 23, 18, 15, 47, 0, time.UTC)
	updatedAt := time.Date(2017, time.September, 23, 18, 15, 47, 0, time.UTC)
	totalSpent := decimal.NewFromFloat(278.60)

	expectation := &Customer{
		ID:               1,
		Email:            "test@example.com",
		FirstName:        "Test",
		LastName:         "Citizen",
		AcceptsMarketing: true,
		VerifiedEmail:    true,
		TaxExempt:        false,
		OrdersCount:      4,
		State:            "enabled",
		TotalSpent:       &totalSpent,
		LastOrderId:      123,
		Note:             "",
		Phone:            "",
		DefaultAddress:   address1,
		Addresses:        []*CustomerAddress{address1},
		CreatedAt:        &createdAt,
		UpdatedAt:        &updatedAt,
	}

	if customer.ID != expectation.ID {
		t.Errorf("Customer.ID returned %+v, expected %+v", customer.ID, expectation.ID)
	}
	if customer.Email != expectation.Email {
		t.Errorf("Customer.Email returned %+v, expected %+v", customer.Email, expectation.Email)
	}
	if customer.FirstName != expectation.FirstName {
		t.Errorf("Customer.FirstName returned %+v, expected %+v", customer.FirstName, expectation.FirstName)
	}
	if customer.LastName != expectation.LastName {
		t.Errorf("Customer.LastName returned %+v, expected %+v", customer.LastName, expectation.LastName)
	}
	if customer.AcceptsMarketing != expectation.AcceptsMarketing {
		t.Errorf("Customer.AcceptsMarketing returned %+v, expected %+v", customer.AcceptsMarketing, expectation.AcceptsMarketing)
	}
	if !customer.CreatedAt.Equal(*expectation.CreatedAt) {
		t.Errorf("Customer.CreatedAt returned %+v, expected %+v", customer.CreatedAt, expectation.CreatedAt)
	}
	if !customer.UpdatedAt.Equal(*expectation.UpdatedAt) {
		t.Errorf("Customer.UpdatedAt returned %+v, expected %+v", customer.UpdatedAt, expectation.UpdatedAt)
	}
	if customer.OrdersCount != expectation.OrdersCount {
		t.Errorf("Customer.OrdersCount returned %+v, expected %+v", customer.OrdersCount, expectation.OrdersCount)
	}
	if customer.State != expectation.State {
		t.Errorf("Customer.State returned %+v, expected %+v", customer.State, expectation.State)
	}
	if !expectation.TotalSpent.Truncate(2).Equals(customer.TotalSpent.Truncate(2)) {
		t.Errorf("Customer.TotalSpent returned %+v, expected %+v", customer.TotalSpent, expectation.TotalSpent)
	}
	if customer.LastOrderId != expectation.LastOrderId {
		t.Errorf("Customer.LastOrderId returned %+v, expected %+v", customer.LastOrderId, expectation.LastOrderId)
	}
	if customer.Note != expectation.Note {
		t.Errorf("Customer.Note returned %+v, expected %+v", customer.Note, expectation.Note)
	}
	if customer.VerifiedEmail != expectation.VerifiedEmail {
		t.Errorf("Customer.Note returned %+v, expected %+v", customer.VerifiedEmail, expectation.VerifiedEmail)
	}
	if customer.TaxExempt != expectation.TaxExempt {
		t.Errorf("Customer.TaxExempt returned %+v, expected %+v", customer.TaxExempt, expectation.TaxExempt)
	}
	if customer.Phone != expectation.Phone {
		t.Errorf("Customer.Phone returned %+v, expected %+v", customer.Phone, expectation.Phone)
	}
	if customer.DefaultAddress == nil {
		t.Errorf("Customer.Address is nil, expected not nil")
	} else {
		if customer.DefaultAddress.ID != expectation.DefaultAddress.ID {
			t.Errorf("Customer.DefaultAddress.ID returned %+v, expected %+v", customer.DefaultAddress.ID, expectation.DefaultAddress.ID)
		}
		if customer.DefaultAddress.CustomerID != expectation.DefaultAddress.CustomerID {
			t.Errorf("Customer.DefaultAddress.CustomerID returned %+v, expected %+v", customer.DefaultAddress.CustomerID, expectation.DefaultAddress.CustomerID)
		}
		if customer.DefaultAddress.FirstName != expectation.DefaultAddress.FirstName {
			t.Errorf("Customer.DefaultAddress.FirstName returned %+v, expected %+v", customer.DefaultAddress.FirstName, expectation.DefaultAddress.FirstName)
		}
		if customer.DefaultAddress.LastName != expectation.DefaultAddress.LastName {
			t.Errorf("Customer.DefaultAddress.LastName returned %+v, expected %+v", customer.DefaultAddress.LastName, expectation.DefaultAddress.LastName)
		}
		if customer.DefaultAddress.Company != expectation.DefaultAddress.Company {
			t.Errorf("Customer.DefaultAddress.Company returned %+v, expected %+v", customer.DefaultAddress.Company, expectation.DefaultAddress.Company)
		}
		if customer.DefaultAddress.Address1 != expectation.DefaultAddress.Address1 {
			t.Errorf("Customer.DefaultAddress.Address1 returned %+v, expected %+v", customer.DefaultAddress.Address1, expectation.DefaultAddress.Address1)
		}
		if customer.DefaultAddress.Address2 != expectation.DefaultAddress.Address2 {
			t.Errorf("Customer.DefaultAddress.Address2 returned %+v, expected %+v", customer.DefaultAddress.Address2, expectation.DefaultAddress.Address2)
		}
		if customer.DefaultAddress.City != expectation.DefaultAddress.City {
			t.Errorf("Customer.DefaultAddress.City returned %+v, expected %+v", customer.DefaultAddress.City, expectation.DefaultAddress.City)
		}
		if customer.DefaultAddress.Province != expectation.DefaultAddress.Province {
			t.Errorf("Customer.DefaultAddress.Province returned %+v, expected %+v", customer.DefaultAddress.Province, expectation.DefaultAddress.Province)
		}
		if customer.DefaultAddress.Country != expectation.DefaultAddress.Country {
			t.Errorf("Customer.DefaultAddress.Country returned %+v, expected %+v", customer.DefaultAddress.Country, expectation.DefaultAddress.Country)
		}
		if customer.DefaultAddress.Zip != expectation.DefaultAddress.Zip {
			t.Errorf("Customer.DefaultAddress.Zip returned %+v, expected %+v", customer.DefaultAddress.Zip, expectation.DefaultAddress.Zip)
		}
		if customer.DefaultAddress.Phone != expectation.DefaultAddress.Phone {
			t.Errorf("Customer.DefaultAddress.Phone returned %+v, expected %+v", customer.DefaultAddress.Phone, expectation.DefaultAddress.Phone)
		}
		if customer.DefaultAddress.Name != expectation.DefaultAddress.Name {
			t.Errorf("Customer.DefaultAddress.Name returned %+v, expected %+v", customer.DefaultAddress.Name, expectation.DefaultAddress.Name)
		}
		if customer.DefaultAddress.ProvinceCode != expectation.DefaultAddress.ProvinceCode {
			t.Errorf("Customer.DefaultAddress.ProvinceCode returned %+v, expected %+v", customer.DefaultAddress.ProvinceCode, expectation.DefaultAddress.ProvinceCode)
		}
		if customer.DefaultAddress.CountryCode != expectation.DefaultAddress.CountryCode {
			t.Errorf("Customer.DefaultAddress.ID returned %+v, expected %+v", customer.DefaultAddress.ID, expectation.DefaultAddress.ID)
		}
		if customer.DefaultAddress.CountryCode != expectation.DefaultAddress.CountryCode {
			t.Errorf("Customer.DefaultAddress.CountryCode returned %+v, expected %+v", customer.DefaultAddress.CountryCode, expectation.DefaultAddress.CountryCode)
		}
		if customer.DefaultAddress.CountryName != expectation.DefaultAddress.CountryName {
			t.Errorf("Customer.DefaultAddress.CountryName returned %+v, expected %+v", customer.DefaultAddress.CountryName, expectation.DefaultAddress.CountryName)
		}
		if customer.DefaultAddress.Default != expectation.DefaultAddress.Default {
			t.Errorf("Customer.DefaultAddress.Default returned %+v, expected %+v", customer.DefaultAddress.Default, expectation.DefaultAddress.Default)
		}
	}
	if len(customer.Addresses) != len(expectation.Addresses) {
		t.Errorf("Customer.Addresses count returned %d, expected %d", len(customer.Addresses), len(expectation.Addresses))
	}
}

func TestCustomerUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", fmt.Sprintf("https://fooshop.myshopify.com/%s/customers/1.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("customer.json")))

	customer := Customer{
		ID:   1,
		Tags: "new",
	}

	returnedCustomer, err := client.Customer.UpdateCustomer(customer)
	if err != nil {
		t.Errorf("Customer.Update returned error: %v", err)
	}

	expectedCustomerID := int64(1)
	if returnedCustomer.ID != expectedCustomerID {
		t.Errorf("Customer.ID returned %+v expected %+v", returnedCustomer.ID, expectedCustomerID)
	}
}

func TestCustomerCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://fooshop.myshopify.com/%s/customers.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("customer.json")))

	customer := Customer{
		ID:   1,
		Tags: "new",
	}

	returnedCustomer, err := client.Customer.CreateCustomer(customer)
	if err != nil {
		t.Errorf("Customer.Create returned error: %v", err)
	}

	expectedCustomerID := int64(1)
	if returnedCustomer.ID != expectedCustomerID {
		t.Errorf("Customer.ID returned %+v expected %+v", returnedCustomer.ID, expectedCustomerID)
	}
}

func TestCustomerDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://fooshop.myshopify.com/%s/customers/1.json", client.pathPrefix),
		httpmock.NewStringResponder(200, ""))

	err := client.Customer.DeleteCustomer(1)
	if err != nil {
		t.Errorf("Customer.Delete returned error: %v", err)
	}
}

func TestCustomerListMetaFields(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/customers/1/metafields.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"metafields": [{"id":1},{"id":2}]}`))

	metaFields, err := client.Customer.ListMetaFields(1, nil)
	if err != nil {
		t.Errorf("Customer.ListMetaFields() returned error: %v", err)
	}

	expected := []MetaField{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(metaFields, expected) {
		t.Errorf("Customer.ListMetaFields() returned %+v, expected %+v", metaFields, expected)
	}
}

func TestCustomerCountMetaFields(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/customers/1/metafields/count.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"count": 3}`))

	params := map[string]string{"created_at_min": "2016-01-01T00:00:00Z"}
	httpmock.RegisterResponderWithQuery(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/customers/1/metafields/count.json", client.pathPrefix),
		params,
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.Customer.CountMetaFields(1, nil)
	if err != nil {
		t.Errorf("Customer.CountMetaFields() returned error: %v", err)
	}

	expected := 3
	if cnt != expected {
		t.Errorf("Customer.CountMetaFields() returned %d, expected %d", cnt, expected)
	}

	date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	cnt, err = client.Customer.CountMetaFields(1, CountOptions{CreatedAtMin: date})
	if err != nil {
		t.Errorf("Customer.CountMetaFields() returned error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("Customer.CountMetaFields() returned %d, expected %d", cnt, expected)
	}
}

func TestCustomerGetMetaField(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/customers/1/metafields/2.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"metafield": {"id":2}}`))

	metaField, err := client.Customer.GetMetaField(1, 2, nil)
	if err != nil {
		t.Errorf("Customer.GetMetaField() returned error: %v", err)
	}

	expected := &MetaField{ID: 2}
	if !reflect.DeepEqual(metaField, expected) {
		t.Errorf("Customer.GetMetaField() returned %+v, expected %+v", metaField, expected)
	}
}

func TestCustomerCreateMetaField(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://fooshop.myshopify.com/%s/customers/1/metafields.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("metafield.json")))

	metaField := MetaField{
		Key:       "app_key",
		Value:     "app_value",
		ValueType: "string",
		Type:      "single_line_text_field",
		Namespace: "affiliates",
	}

	returnedMetaField, err := client.Customer.CreateMetaField(1, metaField)
	if err != nil {
		t.Errorf("Customer.CreateMetaField() returned error: %v", err)
	}

	MetaFieldTests(t, *returnedMetaField)
}

func TestCustomerUpdateMetaField(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", fmt.Sprintf("https://fooshop.myshopify.com/%s/customers/1/metafields/2.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("metafield.json")))

	metafield := MetaField{
		ID:        2,
		Key:       "app_key",
		Value:     "app_value",
		ValueType: "string",
		Type:      "single_line_text_field",
		Namespace: "affiliates",
	}

	returnedMetafield, err := client.Customer.UpdateMetaField(1, metafield)
	if err != nil {
		t.Errorf("Customer.UpdateMetaField() returned error: %v", err)
	}

	MetaFieldTests(t, *returnedMetafield)
}

func TestCustomerDeleteMetaField(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://fooshop.myshopify.com/%s/customers/1/metafields/2.json", client.pathPrefix),
		httpmock.NewStringResponder(200, "{}"))

	err := client.Customer.DeleteMetaField(1, 2)
	if err != nil {
		t.Errorf("Customer.DeleteMetaField() returned error: %v", err)
	}
}

func TestCustomerListOrders(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/customers/1/orders.json", client.pathPrefix),
		httpmock.NewStringResponder(200, "{\"orders\":[]}"),
	)
	params := map[string]string{"status": "any"}
	httpmock.RegisterResponderWithQuery(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/customers/1/orders.json", client.pathPrefix),
		params,
		httpmock.NewBytesResponder(200, loadFixture("orders.json")),
	)

	orders, err := client.Customer.ListOrders(1, nil)
	if err != nil {
		t.Errorf("Customer.ListOrders returned error: %v", err)
	}

	// Check that orders were parsed
	if len(orders) != 0 {
		t.Errorf("Customer.ListOrders got %v orders, expected: 1", len(orders))
	}

	orders, err = client.Customer.ListOrders(1, OrderListOptions{Status: "any"})
	if err != nil {
		t.Errorf("Customer.ListOrders returned error: %v", err)
	}

	// Check that orders were parsed
	if len(orders) != 1 {
		t.Errorf("Customer.ListOrders got %v orders, expected: 1", len(orders))
	}

	order := orders[0]
	orderTests(t, order)
}

func TestCustomerListTags(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/customers/tags.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("customer_tags.json")),
	)

	tags, err := client.Customer.ListTags(nil)
	if err != nil {
		t.Errorf("Customer.ListTags returned error: %v", err)
	}

	// Check that tags were parsed
	if len(tags) != 2 {
		t.Errorf("Customer.ListTags got %v tags, expected: 2", len(tags))
	}

	// Check correct tag was read
	if len(tags) > 0 && tags[0] != "tag1" {
		t.Errorf("Customer.ListTags got %v as the first tag, expected: 'tag1'", tags[0])
	}
}
