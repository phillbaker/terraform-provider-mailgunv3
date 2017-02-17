package main

import (
    "log"

    mailgun "github.com/mailgun/mailgun-go"
)

type Config struct {
    APIKey string
}

// Client() returns a new client for accessing mailgun.
//
func (c *Config) Client() *mailgun.Mailgun {
    domain := ""       // We don't set a domain right away
    publicApiKey := "" // We don't support email validation
    client := mailgun.NewMailgun(domain, c.APIKey, publicApiKey)

    // if err != nil {
    //   return nil, err
    // }

    log.Printf("[INFO] Mailgun Client configured ")

    return &client
}
