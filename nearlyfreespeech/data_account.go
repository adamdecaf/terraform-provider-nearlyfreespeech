package nearlyfreespeech

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAccount() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAccountRead,
		Schema: map[string]*schema.Schema{
			"number": {
				Type: schema.TypeString,
				Description: "account number (XXXX-XXXXXXXX)",
				Required: true,
				ForceNew: true,
			},
			"status": &schema.Schema{
				Type: schema.TypeString,
				Description: "Status of NFS account",
				Computed:    true,
			},
		},
	}
}

func dataSourceAccountRead(d *schema.ResourceData, meta interface{}) error {
	// d.Set("rendered", rendered)
	// d.SetId(hash(rendered))
	return nil
}
