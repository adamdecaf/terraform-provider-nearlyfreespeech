package nearlyfreespeech

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"nearlyfreespeech_account": dataSourceAccount(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"nearlyfreespeech_account": resourceAccount(),
			"nearlyfreespeech_site": resourceSite(),
		},
	}
}
