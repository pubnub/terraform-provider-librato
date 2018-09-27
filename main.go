package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/pubnub/terraform-provider-librato/librato"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: librato.Provider})
}
