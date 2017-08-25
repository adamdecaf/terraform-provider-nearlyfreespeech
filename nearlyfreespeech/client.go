package nearlyfreespeech

import (
	"errors"
	nfs "github.com/adamdecaf/go-nfs"
	"github.com/hashicorp/terraform/helper/schema"
)

type config struct {
	client *nfs.Client
}

func setupProvider(r *schema.ResourceData) (interface{}, error) {
	apiKey, ok := r.Get("api_key").(string)
	if !ok {
		return nil, errors.New("api_key is not defined (or is empty)")
	}

	login, ok := r.Get("login").(string)
	if login == "" {
		return nil, errors.New("login is not defined (or is empty)")
	}

	accountId, _ := r.Get("account_id").(string)

	var client *nfs.Client
	if accountId == "" {
		client = nfs.NewClient(apiKey, login)
	} else {
		client = nfs.NewClientForAccount(accountId, apiKey, login)
	}

	return &config{
		client: client,
	}, nil
}
