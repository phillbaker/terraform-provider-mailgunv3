package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	mailgun "github.com/mailgun/mailgun-go/v3"
)

func TestMailgunDomainMock(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testProviders,
		CheckDestroy: testCheckMailgunDomainDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: mailgunDomaionConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckMailgunDomainExists("mailgunv3_domain.test"),
				),
			},
		},
	})
}

func testCheckMailgunDomainExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		return nil
	}
}

func testCheckMailgunDomainDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mailgunv3_domain" {
			continue
		}

		client := testProvider.Meta().(*mailgun.MailgunImpl)
		_, err := client.GetDomain(context.Background(), rs.Primary.ID)
		if err != nil {
			return nil // should be not found error
		}

		return fmt.Errorf("Snapshot repository %q still exists", rs.Primary.ID)
	}

	return nil
}

var mailgunDomaionConfig = `
resource "mailgunv3_domain" "test" {
    name = "test.terraformv3.example.com"
    spam_action = "disabled"
    smtp_password = "foobar"
}
`
