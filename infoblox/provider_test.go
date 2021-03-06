package infoblox

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"infoblox": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("INFOBLOX_USERNAME"); v == "" {
		t.Fatal("INFOBLOX_USERNAME must be set for acceptance tests")
	}

	if v := os.Getenv("INFOBLOX_PASSWORD"); v == "" {
		t.Fatal("INFOBLOX_PASSWORD must be set for acceptance tests")
	}

	if v := os.Getenv("INFOBLOX_HOST"); v == "" {
		t.Fatal("INFOBLOX_HOST must be set for acceptance tests.")
	}

	if v := os.Getenv("INFOBLOX_SSLVERIFY"); v == "" {
		t.Fatal("INFOBLOX_SSLVERIFY must be set for acceptance tests")
	}

	if v := os.Getenv("INFOBLOX_USECOOKIES"); v == "" {
		t.Fatal("INFOBLOX_USECOOKIES must be set for acceptance tests")
	}
}
