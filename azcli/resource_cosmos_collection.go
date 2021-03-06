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
				ForceNew: true,
			},
			"resource_group_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"throughput": {
				Type:     schema.TypeString,
				Default:  "",
				Optional: true,
			},
			"partition_key": {
				Type:     schema.TypeString,
				Default:  "",
				Optional: true,
				ForceNew: true,
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
	}

	if throughput != "" {
		cmd = append(cmd, "--throughput", throughput)
		log.Printf("[INFO] throughput value %s supplied", throughput)
	}

	if partitionKey != "" {
		cmd = append(cmd, "--partition-key-path", partitionKey)
		log.Printf("[INFO] Partition key %s supplied", partitionKey)
	}

	if indexingPolicy != "" {
		cmd = append(cmd, "--indexing-policy", indexingPolicy)
		log.Printf("[INFO] Indexing policy %s supplied", indexingPolicy)
	}

	cmd = append(cmd, "-o", "json")

	output := c.AZCommand(cmd)

	r, err := ParseAzCliOutput(output)
	if err != nil {
		return err
	}

	if r.AlreadyExists {
		log.Printf("[WARN] Collection %s already exists, will start to manage state", name)
		cmd := []string{
			"cosmosdb", "collection", "show",
			"--collection-name", name, "--db-name", dbName,
			"-g", resourceGroupName, "-n", cosmosAccountName,
			"-o", "json",
		}

		output = c.AZCommand(cmd)
		id := gjson.Get(output, "collection.id")

		resultThroughput := gjson.Get(output, "offer.content.offerThroughput")
		if resultThroughput.String() == "" {
			resultThroughput := gjson.Get(output, "offer")
			if resultThroughput.Value() == nil {
				log.Printf("[INFO] Collection does not specify throughput %s, this is using database scaling", output)
			} else {
				return fmt.Errorf("Unable to get offerthroughput from %s, we are expecting a value if set or null if scale at db level", output)
			}
		} else {
			d.Set("throughput", resultThroughput.String())
		}

		d.SetId(id.Str)
		return nil
	}

	log.Printf("[INFO] Collection %s created", name)

	resultThroughput := gjson.Get(output, "offer.content.offerThroughput")
	if resultThroughput.String() == "" {
		resultThroughput := gjson.Get(output, "offer")
		if resultThroughput.Value() == nil {
			log.Printf("[INFO] Collection does not specify throughput %s, this is using database scaling", output)
		} else {
			return fmt.Errorf("Unable to get offerthroughput from %s, we are expecting a value if set or null if scale at db level", output)
		}
	} else {
		d.Set("throughput", resultThroughput.String())
	}

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

	cmd := []string{
		"cosmosdb", "collection", "show", "--collection-name", name,
		"--db-name", dbName, "-g", resourceGroupName, "-n", cosmosAccountName,
		"-o", "json",
	}
	output := c.AZCommand(cmd)

	r, err := ParseAzCliOutput(output)
	if err != nil {
		return err
	}

	if !r.Found {
		log.Printf("[INFO] collection not found")
		d.SetId("")

		return nil
	}

	resultThroughput := gjson.Get(output, "offer.content.offerThroughput")
	if resultThroughput.String() == "" {
		resultThroughput := gjson.Get(output, "offer")
		if resultThroughput.Value() == nil {
			log.Printf("[INFO] Collection does not specify throughput %s, this is using database scaling", output)
		} else {
			return fmt.Errorf("Unable to get offerthroughput from %s, we are expecting a value if set or null if scale at db level", output)
		}
	} else {
		d.Set("throughput", resultThroughput.String())
	}

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
	}

	if throughput != "" {
		cmd = append(cmd, "--throughput", throughput)
		log.Printf("[INFO] throughput value %s supplied", throughput)
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

	resultThroughput := gjson.Get(output, "offer.content.offerThroughput")
	if resultThroughput.String() == "" {
		resultThroughput := gjson.Get(output, "offer")
		if resultThroughput.Value() == nil {
			log.Printf("[INFO] Collection does not specify throughput %s, this is using database scaling", output)
		} else {
			return fmt.Errorf("Unable to get offerthroughput from %s, we are expecting a value if set or null if scale at db level", output)
		}
	} else {
		d.Set("throughput", resultThroughput.String())
	}

	return nil
}

func resourceCosmosCollectionDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	cosmosAccountName := d.Get("cosmos_account_name").(string)
	dbName := d.Get("database_name").(string)

	cmd := []string{
		"cosmosdb", "collection", "delete",
		"--collection-name", name, "--db-name", dbName,
		"-g", resourceGroupName, "-n", cosmosAccountName,
		"-o", "json",
	}

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
