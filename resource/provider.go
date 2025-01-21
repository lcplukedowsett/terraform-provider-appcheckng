package main

import (
	"terraform-provider-appcheckng/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("APPCHECK_API_KEY", nil),
				Description: "The API key for AppCheckNG",
			},
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
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

	return client.NewAppCheckClient(apiKey, endpoint), nil
}
