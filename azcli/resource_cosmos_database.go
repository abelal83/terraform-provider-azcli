package azcli

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tidwall/gjson"
)

func resourceCosmosDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceCosmosDatabaseCreate,
		Read:   resourceCosmosDatabaseRead,
		Update: resourceCosmosDatabaseUpdate,
		Delete: resourceCosmosDatabaseDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cosmos_account_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_group_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCosmosDatabaseCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	cosmosAccountName := d.Get("cosmos_account_name").(string)

	cmd := []string{"cosmosdb", "database", "create", "--db-name", name, "-g", resourceGroupName, "-n", cosmosAccountName, "-o", "json"}
	output := c.AZCommand(cmd)

	if gjson.Valid(output) {
		log.Print("az cli output is valid json")
	} else {
		panic("az cli output not valid json")
	}

	id := gjson.Get(output, "id")
	d.SetId(id.Str)

	return nil
}

func resourceCosmosDatabaseRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	cosmosAccountName := d.Get("cosmos_account_name").(string)

	cmd := []string{"cosmosdb", "database", "show", "--db-name", name, "-g", resourceGroupName, "-n", cosmosAccountName, "-o", "json"}
	output := c.AZCommand(cmd)

	switch gjson.Valid(output) {
	case true:
		log.Print("az cli output is valid json")
	case false:
		if strings.Contains(output, "Operation Failed: Resource Not Found") {
			log.Printf("az cli - %s", output)
			//return fmt.Errorf("Couldn't find database: %s", output)
			d.SetId("")
		}
	}

	return nil
}

func resourceCosmosDatabaseUpdate(d *schema.ResourceData, m interface{}) error {
	//c := m.(*Client)
	return resourceCosmosDatabaseRead(d, m)
}

func resourceCosmosDatabaseDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	cosmosAccountName := d.Get("cosmos_account_name").(string)

	cmd := []string{"cosmosdb", "database", "delete", "--db-name", name, "-g", resourceGroupName, "-n", cosmosAccountName, "-o", "json"}
	output := c.AZCommand(cmd)

	switch gjson.Valid(output) {
	case true:
		log.Print("az cli output is valid json")
	case false:
		if strings.Contains(output, "Operation Failed: Resource Not Found") {
			log.Printf("az cli - %s", output)
		}
	}

	return nil
}
