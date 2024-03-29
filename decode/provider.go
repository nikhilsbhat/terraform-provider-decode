package decode

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"decode_json": decodeJSON(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"decode_json": datadecodeJSON(),
		},
	}
}
