package azcli

import (
	"log"

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
				//ForceNew: true,
			},
			"resource_group_name": {
				Type:     schema.TypeString,
				Required: true,
				//ForceNew: true,
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

	r, err := ParseAzCliOutput(output)
	if err != nil {
		return err
	}

	if r.AlreadyExists {
		// database already exists, lets just start to manage it.
		cmd = []string{"cosmosdb", "database", "show", "--db-name", name, "-g", resourceGroupName, "-n", cosmosAccountName, "-o", "json"}
		output = c.AZCommand(cmd)
		id := gjson.Get(output, "id")
		d.SetId(id.Str)
		return nil
	}

	// new resource created
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

	r, err := ParseAzCliOutput(output)
	if err != nil {
		return err
	}

	if !r.Found {
		// database doesn't exist
		log.Print("[INFO] database not found")
		d.SetId("")
	}

	return nil
}

func resourceCosmosDatabaseUpdate(d *schema.ResourceData, m interface{}) error {
	//c := m.(*Client)
	// there is nothing to update for a database
	return resourceCosmosDatabaseRead(d, m)
}

func resourceCosmosDatabaseDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	cosmosAccountName := d.Get("cosmos_account_name").(string)

	cmd := []string{"cosmosdb", "database", "delete", "--db-name", name, "-g", resourceGroupName, "-n", cosmosAccountName, "-o", "json"}
	output := c.AZCommand(cmd)

	r, err := ParseAzCliOutput(output)
	if err != nil {
		return err
	}

	if !r.Found || r.CliResponse == "" {
		// database doesn't exist
		log.Print("[INFO] database not found")
		d.SetId("")
	}

	return nil
}
