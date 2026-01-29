# Terraform Provider REMS Content

This is a provider for Terraform that manages content
of REMS instances (https://github.com/CSCfi/rems). Note that this
does not install the REMS software itself - this just manages the
content of an existing REMS instance (forms, committees etc).

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.24

## Using the provider

This will be filled in as the provider is written.

As an example (not necessarily updated, more to show the concepts) - see the generated documentation
or official example code for more details.

```terraform
terraform {
  required_providers {
    remscontent = {
      source = "registry.terraform.io/umccr/remscontent"
    }
  }
}

provider "remscontent" {
  endpoint = "rems.somewhere.com"
  api_user = "..."
  api_key  = "..."
}

resource "remscontent_form" "application_form" {
  organization_id = "Our Organisation"
  title           = "Access to XYZ data"

  fields = [
    provider::remscontent::form_field_header("xyz_applicant", { en : "Applicant" }),
    provider::remscontent::form_field_header("xyz_purpose", { en : "Purpose" }),
  ]
}
```

## Building The Provider

1. Clone the repository
2. Enter the repository directory
3. Build the provider using the Go `install` command:

```shell
go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To actually see the provider in use whilst it is not published to a registry - you will need to make a local
development override in your `~/.terraformrc` (make sure you customise the path to your Go binaries).

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/guardians/rems" = ".. path where your Go binaries get installed .."
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

To generate or update documentation, run `make generate`.

## Archive

_This template repository is built on the [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework). The template
repository built on the [Terraform Plugin SDK](https://github.com/hashicorp/terraform-plugin-sdk) can be
found at [terraform-provider-scaffolding](https://github.com/hashicorp/terraform-provider-scaffolding). See [Which SDK Should I Use?](https://developer.hashicorp.com/terraform/plugin/framework-benefits) in the Terraform documentation for additional information._

This repository is a *template* for a [Terraform](https://www.terraform.io) provider. It is intended as a starting point for creating Terraform providers, containing:

- A resource and a data source (`internal/provider/`),
- Examples (`examples/`) and generated documentation (`docs/`),
- Miscellaneous meta files.

These files contain boilerplate code that you will need to edit to
create your own Terraform provider. Tutorials for creating Terraform providers
can be found on the [HashiCorp Developer](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework) platform. _Terraform Plugin Framework specific guides are titled accordingly._

Please see the [GitHub template repository documentation](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/creating-a-repository-from-a-template) for how to create a new repository from this template on GitHub.

Once you've written your provider, you'll want to [publish it on the Terraform Registry](https://developer.hashicorp.com/terraform/registry/providers/publishing) so that others can use it.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```
