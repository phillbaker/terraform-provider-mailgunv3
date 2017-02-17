# terraform-provider-mailgunv3

[![Build Status](https://travis-ci.org/phillbaker/terraform-provider-mailgunv3.svg?branch=master)](https://travis-ci.org/phillbaker/terraform-provider-mailgunv3)

This is a terraform provider that lets you provision
email related resources on [mailgun](https://mailgun.com/) host via [Terraform](https://terraform.io/). It's powered by [mailgun-go](https://github.com/mailgun/mailgun-go), using v3 of the Mailgun API.

## Installation

[Download a binary](https://github.com/phillbaker/terraform-provider-mailgunv3/releases), and put it in a good spot on your system. Then update your `~/.terraformrc` to refer to the binary:

```hcl
providers {
  mailgunv3 = "/path/to/terraform-provider-mailgunv3"
}
```

See [the docs for more information](https://www.terraform.io/docs/plugins/basics.html).

## Usage

```
provider "mailgunv3" {
    api_key = "key-5555555555555555555"
}

resource "mailgunv3_domain" "default" {
    name = "test.terraformv3.example.com"
    spam_action = "disabled"
    smtp_password = "foobar"
}

resource "mailgunv3_route" "default" {
    priority = "0"
    description = "inbound"
    expression = "match_recipient('.*@foo.example.com')"
    actions = [
      "forward('http://example.com/api/v1/foos/')",
      "stop()"
    ]
}
```

## Development

```
go get github.com/phillbaker/terraform-provider-mailgunv3
go get github.com/mailgun/mailgun-go
cd $GOPATH/src/github.com/phillbaker/terraform-provider-mailgunv3
go build -o /path/to/binary/terraform-provider-mailgunv3
```

## Licence

See LICENSE.

## Contributing

1. Fork it ( https://github.com/phillbaker/terraform-provider-mailgunv3/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
