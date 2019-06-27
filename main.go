package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/nikhilsbhat/terraform-provider-decode/decode"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: decode.Provider})
}