package main

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MAILGUN_API_KEY", nil),
			},
			"base_url": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"mailgunv3_domain": resourceMailgunDomain(),
			"mailgunv3_route":  resourceMailgunRoute(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		APIKey:  d.Get("api_key").(string),
		BaseUrl: d.Get("base_url").(string),
	}

	log.Println("[INFO] Initializing Mailgun client")
	return config.Client(), nil
}
