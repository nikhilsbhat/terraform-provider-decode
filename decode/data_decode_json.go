package decode

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func datadecodeJSON() *schema.Resource {
	return &schema.Resource{
		Read: deCODE,

		Schema: map[string]*schema.Schema{
			"jsonfile": {
				Type:        schema.TypeString,
				Description: "Path to JSON file",
				Required:    true,
				ForceNew:    true,
			},
			"json_data": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"json_map": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}
