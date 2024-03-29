package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	_ "github.com/hashicorp/terraform/terraform"
	scalyr "github.com/ansoni/terraform-provider-scalyr/scalyr-go"
	"time"
)

func datasourceFile() *schema.Resource {
	return &schema.Resource{
		Read: datasourceFileRead,
		Schema: map[string]*schema.Schema{
			"path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"create_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mod_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceFileRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*scalyr.ScalyrConfig)
	path := d.Get("path").(string)
	res, err := client.GetFile(path)
	if err != nil {
		return err
	}
	d.Set("content", res.Content)
	d.Set("version", res.Version)
	d.Set("create_date", res.CreateDate.String())
	d.Set("mod_date", res.CreateDate.String())

	d.SetId(time.Now().UTC().String())
	return nil
}
