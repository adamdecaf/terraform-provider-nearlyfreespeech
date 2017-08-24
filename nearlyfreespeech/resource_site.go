package nearlyfreespeech

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSite() *schema.Resource {
	return &schema.Resource{
		Create: resourceSiteCreate,
		Read:   resourceSiteRead,
		Delete: resourceSiteDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "name of the site",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceSiteCreate(d *schema.ResourceData, meta interface{}) error {
	// sourceDir := d.Get("source_dir").(string)
	// d.SetId(hash)
	return nil
}

func resourceSiteRead(d *schema.ResourceData, meta interface{}) error {
	// sourceDir := d.Get("source_dir").(string)
	// d.setId("..")
	return nil
}

func resourceSiteDelete(d *schema.ResourceData, _ interface{}) error {
	// d.SetId("")
	// destinationDir := d.Get("destination_dir").(string)

	return nil
}
