package main

import (
  "fmt"
  "log"
  "time"

  "github.com/hashicorp/terraform/helper/resource"
  "github.com/hashicorp/terraform/helper/schema"
  mailgun "github.com/mailgun/mailgun-go"
)

func resourceMailgunDomain() *schema.Resource {
  return &schema.Resource{
    Create: resourceMailgunDomainCreate,
    Read:   resourceMailgunDomainRead,
    Delete: resourceMailgunDomainDelete,

    Schema: map[string]*schema.Schema{
      "name": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
        ForceNew: true,
      },

      "spam_action": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
        ForceNew: true,
        Optional: true,
      },

      "smtp_password": &schema.Schema{
        Type:     schema.TypeString,
        ForceNew: true,
        Required: true,
      },

      "smtp_login": &schema.Schema{
        Type:     schema.TypeString,
        Computed: true,
        Optional: true,
      },

      "wildcard": &schema.Schema{
        Type:     schema.TypeBool,
        Computed: true,
        ForceNew: true,
        Optional: true,
      },

      "receiving_records": &schema.Schema{
        Type:     schema.TypeList,
        Computed: true,
        Elem: &schema.Resource{
          Schema: map[string]*schema.Schema{
            "priority": &schema.Schema{
              Type:     schema.TypeString,
              Computed: true,
            },
            "record_type": &schema.Schema{
              Type:     schema.TypeString,
              Computed: true,
            },
            "valid": &schema.Schema{
              Type:     schema.TypeString,
              Computed: true,
            },
            "value": &schema.Schema{
              Type:     schema.TypeString,
              Computed: true,
            },
          },
        },
      },

      "sending_records": &schema.Schema{
        Type:     schema.TypeList,
        Computed: true,
        Elem: &schema.Resource{
          Schema: map[string]*schema.Schema{
            "name": &schema.Schema{
              Type:     schema.TypeString,
              Computed: true,
            },
            "record_type": &schema.Schema{
              Type:     schema.TypeString,
              Computed: true,
            },
            "valid": &schema.Schema{
              Type:     schema.TypeString,
              Computed: true,
            },
            "value": &schema.Schema{
              Type:     schema.TypeString,
              Computed: true,
            },
          },
        },
      },
    },
  }
}

func resourceMailgunDomainCreate(d *schema.ResourceData, meta interface{}) error {
  client := *meta.(*mailgun.Mailgun)

  name := d.Get("name").(string)
  smtpPassword := d.Get("smtp_password").(string)
  spamAction := d.Get("spam_action").(string)
  wildcard := d.Get("wildcard").(bool)

  log.Printf("[DEBUG] Domain create configuration: %s, %s, %s, %v", name, smtpPassword, spamAction, wildcard)

  err := client.CreateDomain(name, smtpPassword, spamAction, wildcard)

  if err != nil {
    return err
  }

  d.SetId(name)

  log.Printf("[INFO] Domain ID: %s", d.Id())

  // Retrieve and update state of domain
  _, err = resourceMailgunDomainRetrieve(d.Id(), &client, d)

  if err != nil {
    return err
  }

  return nil
}

func resourceMailgunDomainDelete(d *schema.ResourceData, meta interface{}) error {
  client := *meta.(*mailgun.Mailgun)

  log.Printf("[INFO] Deleting Domain: %s", d.Id())

  // Destroy the domain
  err := client.DeleteDomain(d.Id())
  if err != nil {
    return fmt.Errorf("Error deleting domain: %s", err)
  }

  // Give the destroy a chance to take effect
  return resource.Retry(1*time.Minute, func() *resource.RetryError {
    _, _, _, err = client.GetSingleDomain(d.Id())
    if err == nil {
      log.Printf("[INFO] Retrying until domain disappears...")
      return resource.RetryableError(
        fmt.Errorf("Domain seems to still exist; will check again."))
    }
    log.Printf("[INFO] Got error looking for domain, seems gone: %s", err)
    return nil
  })
}

func resourceMailgunDomainRead(d *schema.ResourceData, meta interface{}) error {
  client := *meta.(*mailgun.Mailgun)

  _, err := resourceMailgunDomainRetrieve(d.Id(), &client, d)

  if err != nil {
    return err
  }

  return nil
}

func resourceMailgunDomainRetrieve(id string, client *mailgun.Mailgun, d *schema.ResourceData) (*mailgun.Domain, error) {
  domain, receivingRecords, sendingRecords, err := (*client).GetSingleDomain(id)

  if err != nil {
    return nil, fmt.Errorf("Error retrieving domain: %s", err)
  }

  d.Set("name", domain.Name)
  d.Set("smtp_password", domain.SMTPPassword)
  d.Set("smtp_login", domain.SMTPLogin)
  d.Set("wildcard", domain.Wildcard)
  d.Set("spam_action", domain.SpamAction)
  d.Set("receiving_records", receivingRecords)
  d.Set("sending_records", sendingRecords)

  return &domain, nil
}
