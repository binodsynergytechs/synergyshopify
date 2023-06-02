package goshopify

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

// This function tests the ListAccessScopes method of the AccessScopesServiceOp struct.
// It sets up the necessary dependencies, makes an HTTP request to retrieve a list of access scopes,
// and compares the response with the expected result.

func TestListAccessScopes(t *testing.T) {
	setup()
	defer teardown() // Register a responder for the GET request to retrieve the access scopes.
	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/oauth/access_scopes.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("access_scopes.json")),
	)

	// Call the ListAccessScopes method and capture the response and error.
	scopeResponse, err := client.AccessScopes.ListAccessScopes(nil)
	if err != nil {
		t.Errorf("AccessScopes.ListAccessScopes returned an error: %v", err)
	}

	// Define the expected list of access scopes.
	expected := []AccessScope{
		{
			Handle: "scope_1",
		},
		{
			Handle: "scope_2",
		},
	}

	// Compare the actual response with the expected result.
	if !reflect.DeepEqual(scopeResponse, expected) {
		t.Errorf("AccessScopes.ListAccessScopes returned %+v, expected %+v", expected, expected)
	}
}
