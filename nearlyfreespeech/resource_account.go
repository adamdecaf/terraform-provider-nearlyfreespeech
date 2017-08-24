package nearlyfreespeech

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountCreate,
		Read:   resourceAccountRead,
		Update: resourceAccountUpdate,
		Delete: resourceAccountDelete,

		Schema: map[string]*schema.Schema{
			"number": {
				Type: schema.TypeString,
				Description: "account number (XXXX-XXXXXXXX)",
				Required: true,
				ForceNew: true,
			},
			"balance_warnings": {
				Type: schema.TypeList,
				Description: "balance amounts to trigger warnings",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: false,
				Optional: true,
				ForceNew: false,
			},
			"friendly_name": {
				Type: schema.TypeString,
				Description: "friendly name for account",
				Required: false,
				Optional: true,
				ForceNew: false,
			},
		},
	}
}

func resourceAccountCreate(d *schema.ResourceData, meta interface{}) error {
	// sourceDir := d.Get("source_dir").(string)
	// d.SetId(hash)
	return nil
}

func resourceAccountRead(d *schema.ResourceData, meta interface{}) error {
	// sourceDir := d.Get("source_dir").(string)
	// d.setId("..")
	return nil
}

func resourceAccountUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAccountDelete(d *schema.ResourceData, _ interface{}) error {
	// d.SetId("")
	// destinationDir := d.Get("destination_dir").(string)

	return nil
}
