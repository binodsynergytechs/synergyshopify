package goshopify

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/shopspring/decimal"
)

func testApplicationChargeFields(t *testing.T, charge ApplicationCharge) {
	var nilTest *bool
	cases := []struct {
		field    string
		expected interface{}
		actual   interface{}
	}{
		{"ID", int64(1017262355), charge.ID},
		{"Name", "Super Duper Expensive action", charge.Name},
		{"APIClientID", int64(755357713), charge.APIClientID},
		{"Price", decimal.NewFromFloat(100.00).String(), charge.Price.String()},
		{"Status", "pending", charge.Status},
		{"ReturnURL", "http://super-duper.shopifyapps.com/", charge.ReturnURL},
		{"Test", nilTest, charge.Test},
		{"CreatedAt", "2018-07-05T13:11:28-04:00", charge.CreatedAt.Format(time.RFC3339)},
		{"UpdatedAt", "2018-07-05T13:11:28-04:00", charge.UpdatedAt.Format(time.RFC3339)},
		{
			"DecoratedReturnURL",
			"http://super-duper.shopifyapps.com/?charge_id=1017262355",
			charge.DecoratedReturnURL,
		},
		{
			"ConfirmationURL",
			fmt.Sprintf("https://apple.myshopify.com/%s/charges/1017262355/confirm_application_charge?signature=BAhpBBMxojw%%3D--1139a82a3433b1a6771786e03f02300440e11883", client.pathPrefix),
			charge.ConfirmationURL,
		},
	}

	for _, c := range cases {
		if c.expected != c.actual {
			t.Errorf("ApplicationCharge.%s returned %v, expected %v", c.field, c.actual, c.expected)
		}
	}
}

func TestCreateApplicationCharge(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/application_charges.json", client.pathPrefix),
		httpmock.NewBytesResponder(200, loadFixture("applicationcharge.json")),
	)

	p := decimal.NewFromFloat(100.00)
	charge := ApplicationCharge{
		Name:      "Super Duper Expensive action",
		Price:     &p,
		ReturnURL: "http://super-duper.shopifyapps.com",
	}

	returnedCharge, err := client.ApplicationCharge.CreateApplicationCharge(charge)
	if err != nil {
		t.Errorf("ApplicationCharge.CreateApplicationCharge returned an error: %v", err)
	}

	testApplicationChargeFields(t, *returnedCharge)
}

func TestGetApplicationCharge(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/application_charges/1.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"application_charge": {"id":1}}`),
	)

	charge, err := client.ApplicationCharge.GetApplicationCharge(1, nil)
	if err != nil {
		t.Errorf("ApplicationCharge.GetApplicationCharge returned an error: %v", err)
	}

	expected := &ApplicationCharge{ID: 1}
	if !reflect.DeepEqual(charge, expected) {
		t.Errorf("ApplicationCharge.GetApplicationCharge returned %+v, expected %+v", charge, expected)
	}
}

func TestListApplicationCharges(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/application_charges.json", client.pathPrefix),
		httpmock.NewStringResponder(200, `{"application_charges": [{"id":1},{"id":2}]}`),
	)

	charges, err := client.ApplicationCharge.ListApplicationCharges(nil)
	if err != nil {
		t.Errorf("ApplicationCharge.ListApplicationCharges returned an error: %v", err)
	}

	expected := []ApplicationCharge{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(charges, expected) {
		t.Errorf("ApplicationCharge.ListApplicationCharges returned %+v, expected %+v", charges, expected)
	}
}

func TestActivateApplicationCharge(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/application_charges/455696195/activate.json", client.pathPrefix),
		httpmock.NewStringResponder(
			200,
			`{"application_charge":{"id":455696195,"status":"active"}}`,
		),
	)

	charge := ApplicationCharge{
		ID:     455696195,
		Status: "accepted",
	}

	returnedCharge, err := client.ApplicationCharge.ActivateApplicationCharge(charge)
	if err != nil {
		t.Errorf("ApplicationCharge.ActivateApplicationCharge returned an error: %v", err)
	}

	expected := &ApplicationCharge{ID: 455696195, Status: "active"}
	if !reflect.DeepEqual(returnedCharge, expected) {
		t.Errorf("ApplicationCharge.ActivateApplicationCharge returned %+v, expected %+v", charge, expected)
	}
}
