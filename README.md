# Cisco ACI Provider

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) v0.11.7

- [Go](https://golang.org/doc/install) Latest Version

## Building The Provider ##
Clone this repository to: `$GOPATH/src/github.com/ciscoecosystem/terraform-provider-cisco-aci`.

```sh
$ mkdir -p $GOPATH/src/github.com/ciscoecosystem; cd $GOPATH/src/github.com/ciscoecosystem
$ git clone github.com/ciscoecosystem/terraform-provider-aci
```

Enter the provider directory and run make build.

```sh
$ cd $GOPATH/src/github.com/ciscoecosystem/terraform-provider-cisco-aci
$ make build

```

Using The Provider
------------------
If you are building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory, run `terraform init` to initialize it.

ex.
```hcl
#configure provider with your cisco aci credentials.
provider "aci" {
  # cisco-aci user name
  username = "admin"
  # cisco-aci password
  password = "password"
  # cisco-aci url
  url      = "https://my-cisco-aci.com"
  insecure = true
}

resource "aci_tenant" "test-tenant" {
  name        = "test-tenant"
  description = "This tenant is created by terraform"
}

resource "aci_app_profile" "test-app" {
  tanent_dn   = "${aci_tenant.test-tenant.id}"
  name        = "test-app"
  description = "This app profile is created by terraform"
}
```

Developing The Provider
-----------------------
If you want to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine. You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider with sanity checks present in scripts directory and put the provider binary in `$GOPATH/bin` directory.
