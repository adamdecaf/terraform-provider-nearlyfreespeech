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
	status, err := nfs.GetAccountStatus(meta.(*config).client)
	if err != nil {
		return err
	}

	d.SetId(d.Get("number").(string)) // TODO(adam): better
	// d.SetId(hash(rendered))
	d.Set("status", status)

	return nil
}
