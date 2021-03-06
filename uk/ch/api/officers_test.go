package api_test

import (
	"net/url"
	"testing"

	ch "github.com/appinesshq/globiz-go/uk/ch/api"
	"github.com/appinesshq/globiz-go/uk/ch/api/tests"
)

func TestGetOfficers(t *testing.T) {
	api, err := ch.New("12345")
	if err != nil {
		t.Fatalf("expected to pass, but got: %v", err)
	}

	ts := tests.NewMockServer()
	api.URL, err = url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("expected to pass, but got: %v", err)
	}

	c, err := api.GetCompany("12345678")
	if err != nil {
		t.Fatalf("expected to pass, but got: %v", err)
	}

	if got, expected := c.Name, "TEST LTD"; got != expected {
		t.Fatalf("expected %q, but got %q", expected, got)
	}

	o, err := c.Officers()
	if err != nil {
		t.Fatalf("expected to pass, but got: %v", err)
	}

	if got, expected := o.Items[0].Name, "PERSON, Test"; got != expected {
		t.Fatalf("expected %q, but got %q", expected, got)
	}
}
