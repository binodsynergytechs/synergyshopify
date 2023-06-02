package goshopify

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

// This function tests the ListAbandonedCheckout method of the AbandonedCheckoutServiceOp struct.
// It sets up the necessary dependencies, makes an HTTP request to retrieve a list of abandoned checkouts,
// and compares the response with the expected result.

func TestListAbandonedCheckouts(t *testing.T) {
	setup()
	defer teardown()
	// Register a responder for the GET request to retrieve the abandoned checkouts.
	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/checkouts.json", client.pathPrefix),
		httpmock.NewStringResponder(
			200,
			`{"checkouts": [{"id":1},{"id":2}]}`,
		),
	)

	// Call the ListAbandonedCheckout method and capture the response and error.
	abandonedCheckouts, err := client.AbandonedCheckout.ListAbandonedCheckout(nil)
	if err != nil {
		t.Errorf("AbandonedCheckout.ListAbandonedCheckout returned an error: %v", err)
	}

	// Define the expected list of abandoned checkouts.
	expected := []AbandonedCheckout{{ID: 1}, {ID: 2}}

	// Compare the actual response with the expected result.
	if !reflect.DeepEqual(abandonedCheckouts, expected) {
		t.Errorf("AbandonedCheckout.ListAbandonedCheckout returned %+v, expected %+v", abandonedCheckouts, expected)
	}
}
