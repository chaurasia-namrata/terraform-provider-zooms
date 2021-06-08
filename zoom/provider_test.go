package zoom

import(
	"os"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"testing"
	"log"
	
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	os.Setenv("ZOOM_TOKEN", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MjM0NzQ4MDksImlhdCI6MTYyMjg3MDAwOX0.m3I7dQYBh96qoiKBfG7Vvw1v5G1_4-RUZOWVIs5MrU4")
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"zoom": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		log.Println("[ERROR]: ",err)
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T)  {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("ZOOM_TOKEN"); v == "" {
		
		t.Fatal("ZOOM_TOKEN must be set for acceptance tests")
	}
}
