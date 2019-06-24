package azcli

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider doc
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"subscription_name": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SUBSCRIPTION_NAME", nil),
				Description: "Azure Subscription Name",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"azcli_cosmos_database":   resourceCosmosDatabase(),
			"azcli_cosmos_collection": resourceCosmosCollection(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {

	c, err := NewClient(data.Get("subscription_name").(string))
	if err != nil {
		return nil, err
	}
	return c, nil
}
