package main

import (
	"github.com/adamdecaf/terraform-provider-nearlyfreespeech/nearlyfreespeech"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: nearlyfreespeech.Provider})
}
