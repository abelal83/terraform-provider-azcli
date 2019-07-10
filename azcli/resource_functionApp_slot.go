package azcli

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tidwall/gjson"
)

func resourceFunctionAppSlot() *schema.Resource {
	return &schema.Resource{
		Create: resourceFunctionAppSlotCreate,
		Read:   resourceFunctionAppSlotRead,
		Update: resourceFunctionAppSlotUpdate,
		Delete: resourceFunctionAppSlotDelete,

		Schema: map[string]*schema.Schema{
			"slot_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_group_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"function_app_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"http_20_enabled": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
				Default:  "true",
			},
			"always_on": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
				Default:  "true",
			},
			"identity": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
		},
	}
}

func resourceFunctionAppSlotCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("starting")
	c := m.(*Client)
	log.Println("client Defined")
	slotname := d.Get("slot_name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	functionAppName := d.Get("function_app_name").(string)
	http20Enabled := d.Get("http_20_enabled").(string)
	alwaysOn := d.Get("always_on").(string)
	log.Println("Defined vars")
	createCmd := []string{
		"functionapp", "deployment", "slot", "create",
		"--name", functionAppName,
		"--resource-group", resourceGroupName,
		"--slot", slotname,
		"--configuration-source", functionAppName,
		"-o", "json",
	}

	createOutput := c.AZCommand(createCmd)

	r, err := ParseAzCliOutput(createOutput)
	if r.AlreadyExists {
		log.Println("[WARN] Functionapp $s with slot $s already exist", functionAppName, slotname)
	}

	if err != nil {
		return err
	}
	id := gjson.Get(createOutput, "id")

	configCmd := []string{
		"functionapp", "config", "set",
		"--always-on", alwaysOn,
		"--http20-enabled", http20Enabled,
		"--name", functionAppName,
		"--slot", slotname,
		"--resource-group", resourceGroupName,
		"-o", "json",
	}

	configOutput := c.AZCommand(configCmd)

	configR, configErr := ParseAzCliOutput(configOutput)
	if configErr != nil {
		return err
	}

	if configR.AlreadyExists {
		log.Println("[WARN] Functionapp $s with slot $s already exist", functionAppName, slotname)
	}

	identityCmd := []string{
		"functionapp", "identity", "assign",
		"--name", functionAppName,
		"--resource-group", resourceGroupName,
		"--slot", slotname,
		"-o", "json",
	}
	identityOutput := c.AZCommand(identityCmd)

	identityR, identityErr := ParseAzCliOutput(identityOutput)
	if identityErr != nil {
		return err
	}
	if identityR.AlreadyExists {
		log.Println("[WARN] Functionapp $s with slot $s already exist", functionAppName, slotname)
	}

	principalID := gjson.Get(identityOutput, "principalId")
	log.Println("[INFO] Identity Principal $s", principalID)
	d.Set("slot_name", slotname)
	d.Set("resource_group_name", resourceGroupName)
	d.Set("function_app_name", functionAppName)
	d.Set("identity", principalID.String())
	d.Set("http20Enabled", http20Enabled)
	d.Set("always_on", alwaysOn)
	d.SetId(id.Str)
	return nil
}

func resourceFunctionAppSlotRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	slotname := d.Get("slot_name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	functionAppName := d.Get("function_app_name").(string)

	getConfigCmd := []string{
		"functionapp", "config", "show",
		"--name", functionAppName,
		"--resource-group", resourceGroupName,
		"--slot", slotname,
		"-o", "json",
	}
	configOutput := c.AZCommand(getConfigCmd)

	r, err := ParseAzCliOutput(configOutput)
	log.Println("[INFO] function app slot config $s", configOutput)
	if err != nil {
		return err
	}

	if !r.Found {
		// slot not found
		log.Print("[INFO] Slot not found")
		d.SetId("")
	}

	return nil
}
func resourceFunctionAppSlotUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceFunctionAppSlotRead(d, m)

}

func resourceFunctionAppSlotDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
