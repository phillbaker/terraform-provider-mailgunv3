package main

import (
	"log"

	mailgun "github.com/mailgun/mailgun-go/v3"
)

type Config struct {
	APIKey  string
	BaseUrl string
}

// Client() returns a new client for accessing mailgun.
//
func (c *Config) Client() *mailgun.MailgunImpl {
	domain := "" // We don't set a domain right away
	client := mailgun.NewMailgun(domain, c.APIKey)

	if c.BaseUrl != "" {
		client.SetAPIBase(c.BaseUrl)
	}

	log.Printf("[INFO] Mailgun Client configured ")

	return client
}
