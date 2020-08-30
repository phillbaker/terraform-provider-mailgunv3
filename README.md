# terraform-provider-mailgunv3

The original provider at https://www.terraform.io/docs/providers/mailgun/ has been upgraded to use the v3 API. Please migrate to that provider, this project is archived. Please see [#19](https://github.com/phillbaker/terraform-provider-mailgunv3/issues/19) for details.

----

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

You can also place the binary inside the plugin folder, which varies based on your operating system. Refer to the [third-party providers documentation](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins) for more information.

For general use, see [the docs for more information](https://www.terraform.io/docs/plugins/basics.html).

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

Install [Glide](https://github.com/Masterminds/glide), then:

```
# Ensure that this folder is at the following location: `${GOPATH}/src/github.com/phillbaker/terraform-provider-mailgunv3`
cd $GOPATH/src/github.com/phillbaker/terraform-provider-mailgunv3

glide install
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
