package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider fuck this
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"cosmos_account_name": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("COSMOS_ACCOUNT_NAME", nil),
				Description: "Azure Cosmos Account Name",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cosmos_database": resourceDatabase(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	config := Config{
		CosmosAccountName: data.Get("cosmos_account_name").(string),
	}
	return config.Client()
}
