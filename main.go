package main

import (
	"github.com/abelal83/terraform_provider_cosmosdb/azcli"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return azcli.Provider()
		},
	})
}
