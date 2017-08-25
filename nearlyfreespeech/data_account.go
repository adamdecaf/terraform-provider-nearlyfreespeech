package nearlyfreespeech

import (
	nfs "github.com/adamdecaf/go-nfs"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAccount() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAccountRead,
		Schema: map[string]*schema.Schema{
			"number": {
				Type:        schema.TypeString,
				Description: "account number (XXXX-XXXXXXXX)",
				Required:    true,
				ForceNew:    true,
			},
			"friendly_name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "friendly name associated to account",
				Computed:    true,
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Status of NFS account",
				Computed:    true,
			},
		},
	}
}

func dataSourceAccountRead(d *schema.ResourceData, meta interface{}) error {
	c, err := getClient(meta)
	if err != nil {
		return err
	}

	// set state id
	d.SetId(d.Get("number").(string)) // TODO(adam): better
	// d.SetId(hash(rendered))

	// friendly name
	fn, err := nfs.GetFriendlyName(c)
	if err != nil {
		return err
	}
	d.Set("friendly_name", fn)

	// status
	status, err := nfs.GetAccountStatus(c)
	if err != nil {
		return err
	}
	d.Set("status", status)

	return nil
}
