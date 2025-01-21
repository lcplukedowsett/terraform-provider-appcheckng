package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"terraform-provider-appcheckng/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAppCheckScanDefinition() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAppCheckScanDefinitionRead,

		Schema: map[string]*schema.Schema{
			"scan_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"targets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAppCheckScanDefinitionRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	scanID := d.Get("scan_id").(string)
	resp, err := client.GetScanDetails(scanID)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		}
		return err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp, &result); err != nil {
		return err
	}

	if !result["success"].(bool) {
		return fmt.Errorf("failed to get scan definition: %s", result["message"].(string))
	}

	data := result["data"].(map[string]interface{})
	d.SetId(data["scan_id"].(string))
	d.Set("name", data["name"])
	d.Set("targets", data["targets"])
	d.Set("user_name", data["user_name"])
	d.Set("tags", data["tags"])

	return nil
}
