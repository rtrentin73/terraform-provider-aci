---
layout: "aci"
page_title: "ACI: aci_subnet"
sidebar_current: "docs-aci-resource-subnet"
description: |-
  Manages ACI Subnet Model Object.
---

# aci_subnet #
Manages ACI Subnet Model Object.

## Example Usage ##

```hcl
resource "aci_subnet" "example-subnet" {
  name             = "10.0.3.28/27"
  bridge_domain_dn = "${aci_bridge_domain.example-bridge-domain.id}"
  ip_address       = "10.0.3.28/27"
  scope            = ["private"]
  description      = "This subnet is created by terraform"
}
```

## Argument Reference ##
* `name` - (Required) Name of the Subnet model object. This contains the ip address with subnet mask in `xxx.xxx.xxx.xxx/xx` form.
* `description` - (Optional) Description for the Subnet Model Object.
* `bridge_domain_dn` - (Required) Distinguished name of Parent Bridge domain model object.
* `ip_address` - (Required) The IP address and mask of the default gateway.
* `scope` - (Optional) The network visibility of the subnet. Default - `private`. Possible Values :- `public`, `private`, `shared`.
* `preffered` - (Optional) Indicates if the subnet is preferred (primary) over the available alternatives. Only one preferred subnet is allowed. Default - `false`.
* `virtual` - (Optional) Treated as virtual IP address. Used in case of Bridge Domain extended to multiple sites. Default - `false`.



