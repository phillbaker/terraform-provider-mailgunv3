package main

import (
	"os"
	"testing"

	// "github.com/hashicorp/terraform/config"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	mailgun "github.com/mailgun/mailgun-go/v3"
)

var testProviders map[string]terraform.ResourceProvider
var testProvider *schema.Provider

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

const (
	testDomain = "mailgun.test"
)

var server mailgun.MockServer

// Setup and shutdown the mailgun mock server for the entire test suite
func TestMain(m *testing.M) {
	server = mailgun.NewMockServer()
	defer server.Stop()
	os.Exit(m.Run())
}

func init() {
	testProvider = Provider().(*schema.Provider)
	testProviders = map[string]terraform.ResourceProvider{
		"mailgunv3": testProvider,
	}

  // Override default implementation to pull from the environment, this
  // oveerrides any default we set in the ConfigureFunc :|
	testProvider.Schema["api_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Default:  "fakekey123",
	}

	testProviderOriginalConfigureFunc := testProvider.ConfigureFunc
	testProvider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		d.Set("base_url", server.URL())
		return testProviderOriginalConfigureFunc(d)
	}

	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"mailgunv3": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("MAILGUN_API_KEY"); v == "" {
		t.Fatal("MAILGUN_API_KEY must be set for acceptance tests")
	}
}
