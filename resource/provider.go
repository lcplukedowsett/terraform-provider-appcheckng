package main

import (
	"fmt"
	"terraform-provider-appcheckng/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("APPCHECK_API_KEY", nil),
				Description: "The API key for AppCheckNG",
			},
			"endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("APPCHECK_ENDPOINT", nil),
				Description: "The Endpoint for AppCheckNG API",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"appcheck_scan_definition": resourceAppCheckScanDefinition(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"appcheck_scan_definition": dataSourceAppCheckScanDefinition(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	apiKey := d.Get("api_key").(string)
	endpoint := d.Get("endpoint").(string)

	if apiKey == "" {
		return nil, fmt.Errorf("api_key must be provided")
	}
	if endpoint == "" {
		return nil, fmt.Errorf("endpoint must be provided")
	}

	return client.NewAppCheckClient(apiKey, endpoint), nil
}
