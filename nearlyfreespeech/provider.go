package nearlyfreespeech

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NFS_API_KEY", nil),
				Description: "NearlyFreeSpeech api key",
			},
			// TODO(adam): rename to account_number ?
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NFS_ACCOUNT_ID", nil),
				Description: "NearlyFreeSpeech account number",
			},
			"login": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NFS_LOGIN", nil),
				Description: "NearlyFreeSpeech login",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"nearlyfreespeech_account": dataSourceAccount(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"nearlyfreespeech_account": resourceAccount(),
			"nearlyfreespeech_site":    resourceSite(),
		},
		ConfigureFunc: setupProvider,
	}
}
