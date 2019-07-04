package azcli

import (
	"github.com/hashicorp/terraform/helper/schema"
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
			"location": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"app_service_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"app_service_plan_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceFunctionAppSlotCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	slotName := d.Get("slot_name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	location := d.Get("location").(string)
	appServiceNaame := d.Get("app_service_name").(string)
	appServicePlanId := d.Get("app_service_plan_id").(string)

	cmd := []string{}
}

func resourceFunctionAppSlotRead(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceFunctionAppSlotUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceFunctionAppSlotDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
