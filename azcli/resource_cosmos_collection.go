package azcli

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tidwall/gjson"
)

func resourceCosmosCollection() *schema.Resource {
	return &schema.Resource{
		Create: resourceCosmosCollectionCreate,
		Read:   resourceCosmosCollectionRead,
		Update: resourceCosmosCollectionUpdate,
		Delete: resourceCosmosCollectionDelete,

		Schema: map[string]*schema.Schema{
			"database_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
			"throughput": {
				Type:     schema.TypeString,
				Default:  400,
				Optional: true,
				//ForceNew: true,
			},
			"partition_key": {
				Type:     schema.TypeString,
				Default:  "",
				Optional: true,
				//ForceNew: true,
			},
			"indexing_policy": {
				Type:     schema.TypeString,
				Default:  "",
				Optional: true,
			},
		},
	}
}

func resourceCosmosCollectionCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	cosmosAccountName := d.Get("cosmos_account_name").(string)
	throughput := d.Get("throughput").(string)
	dbName := d.Get("database_name").(string)
	partitionKey := d.Get("partition_key").(string)
	indexingPolicy := d.Get("indexing_policy").(string)

	cmd := []string{
		"cosmosdb", "collection", "create",
		"--collection-name", name,
		"--db-name", dbName,
		"-g", resourceGroupName,
		"-n", cosmosAccountName,
		"--throughput", throughput,
	}

	if partitionKey != "" {
		cmd = append(cmd, "--partition-key-path", partitionKey)
	}

	if indexingPolicy != "" {
		cmd = append(cmd, "--indexing-policy", indexingPolicy)
	}

	cmd = append(cmd, "-o", "json")

	output := c.AZCommand(cmd)

	r, err := ParseAzCliOutput(output)
	if err != nil {
		return err
	}

	if r.AlreadyExists {
		// collection already exists, lets just start to manage it.
		cmd := []string{"cosmosdb", "collection", "show", "--collection-name", name, "--db-name", dbName, "-g", resourceGroupName, "-n", cosmosAccountName, "-o", "json"}
		output = c.AZCommand(cmd)
		id := gjson.Get(output, "collection.id")
		d.SetId(id.Str)
		return nil
	}

	// new resource created
	id := gjson.Get(output, "collection.id")
	d.SetId(id.Str)
	return nil

}

func resourceCosmosCollectionRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	cosmosAccountName := d.Get("cosmos_account_name").(string)
	dbName := d.Get("database_name").(string)

	cmd := []string{"cosmosdb", "collection", "show", "--collection-name", name, "--db-name", dbName, "-g", resourceGroupName, "-n", cosmosAccountName, "-o", "json"}
	output := c.AZCommand(cmd)

	r, err := ParseAzCliOutput(output)
	if err != nil {
		return err
	}

	if !r.Found {
		// collection doesn't exist
		log.Print("[INFO] collection not found")
		d.SetId("")
	}

	resultThroughput := gjson.Get(output, "offer.offerthroughput")
	d.Set("offerthroughput", resultThroughput.Str)
	return nil

}

func resourceCosmosCollectionUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	cosmosAccountName := d.Get("cosmos_account_name").(string)
	throughput := d.Get("throughput").(string)
	dbName := d.Get("database_name").(string)

	cmd := []string{
		"cosmosdb", "collection", "update",
		"--collection-name", name,
		"--db-name", dbName,
		"-g", resourceGroupName,
		"-n", cosmosAccountName,
		"--throughput", throughput,
		"-o", "json",
	}
	output := c.AZCommand(cmd)

	r, err := ParseAzCliOutput(output)
	if err != nil {
		return err
	}

	if !r.Found {
		// collection doesn't exist
		return fmt.Errorf("Collection %s could not be found. AZ CLI returned %s", name, r.CliResponse)
	}

	resultThroughput := gjson.Get(output, "offer.offerthroughput")
	d.Set("offerthroughput", resultThroughput.Str)
	return nil
}

func resourceCosmosCollectionDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	cosmosAccountName := d.Get("cosmos_account_name").(string)
	dbName := d.Get("database_name").(string)

	cmd := []string{"cosmosdb", "collection", "delete", "--collection-name", name, "--db-name", dbName, "-g", resourceGroupName, "-n", cosmosAccountName, "-o", "json"}
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
