package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"terraform-provider-appcheckng/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAppCheckScanDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppCheckScanDefinitionCreate,
		Read:   resourceAppCheckScanDefinitionRead,
		Update: resourceAppCheckScanDefinitionUpdate,
		Delete: resourceAppCheckScanDefinitionDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"targets": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"profile_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scan_hub": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAppCheckScanDefinitionCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	name := d.Get("name").(string)
	targets := d.Get("targets").([]interface{})
	profileID := d.Get("profile_id").(string)
	scanHub := d.Get("scan_hub").(string)
	tags := d.Get("tags").(string)

	data := url.Values{}
	data.Set("name", name)
	for _, target := range targets {
		data.Add("targets", target.(string))
	}
	if profileID != "" {
		data.Set("profile_id", profileID)
	}
	if scanHub != "" {
		data.Set("scan_hub", scanHub)
	}
	if tags != "" {
		data.Set("tags", tags)
	}

	resp, err := client.postRequest("scan/new", data.Encode())
	if err != nil {
		return err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp, &result); err != nil {
		return err
	}

	if !result["success"].(bool) {
		return fmt.Errorf("failed to create scan definition: %s", result["message"].(string))
	}

	d.SetId(result["scan_id"].(string))
	return resourceAppCheckScanDefinitionRead(d, m)
}

func resourceAppCheckScanDefinitionRead(d *schema.ResourceData, m interface{}) error {
	// Implement the read function
	return nil
}

func resourceAppCheckScanDefinitionUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	scanID := d.Id()
	name := d.Get("name").(string)
	targets := d.Get("targets").([]interface{})
	profileID := d.Get("profile_id").(string)
	scanHub := d.Get("scan_hub").(string)
	tags := d.Get("tags").(string)

	data := url.Values{}
	if name != "" {
		data.Set("name", name)
	}
	for _, target := range targets {
		data.Add("targets", target.(string))
	}
	if profileID != "" {
		data.Set("profile_id", profileID)
	}
	if scanHub != "" {
		data.Set("scan_hub", scanHub)
	}
	if tags != "" {
		data.Set("tags", tags)
	}

	resp, err := client.postRequest(fmt.Sprintf("scan/%s/update", scanID), data.Encode())
	if err != nil {
		return err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp, &result); err != nil {
		return err
	}

	if !result["success"].(bool) {
		return fmt.Errorf("failed to update scan definition: %s", result["message"].(string))
	}

	return resourceAppCheckScanDefinitionRead(d, m)
}

func resourceAppCheckScanDefinitionDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.Client)

	scanID := d.Id()
	_, err := client.postRequest(fmt.Sprintf("scan/%s/delete", scanID), "")
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
