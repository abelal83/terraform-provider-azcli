package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatabaseCreate,
		Read:   resourceDatabaseRead,
		Update: resourceDatabaseUpdate,
		Delete: resourceDatabaseDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"default_primary_connection_string": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDatabaseCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	name := d.Get("name").(string)
	d.Set("cosmos_account_name", c.CosmosAccountName)
	//c.AZCommand("account show")
	d.Set("default_primary_connection_string", "crap")
	d.SetId(name)
	//return fmt.Errorf("unable to create Database %T", d)
	return nil
	//return resourceDatabaseRead(d, m)
}

func resourceDatabaseRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	name := d.Get("name").(string)
	d.SetId(name)
	return fmt.Errorf("unable to read from %s", c.CosmosAccountName)
	//d.Set("cosmos_account_name", "renamed")
	//return nil
}

func resourceDatabaseUpdate(d *schema.ResourceData, m interface{}) error {
	//c := m.(*Client)
	return resourceDatabaseRead(d, m)
}

func resourceDatabaseDelete(d *schema.ResourceData, m interface{}) error {
	//c := m.(*Client)
	return nil
}
