---
layout: "aci"
page_title: "ACI: aci_contract"
sidebar_current: "docs-aci-resource-contract"
description: |-
  Manages ACI Contract Model Object.
---

# aci_contract #
Manages ACI Contract Model Object.

## Example Usage ##
```hcl
resource "aci_contract" "example-contract" {
  tanent_dn   = "${aci_tenant.example-tenant.id}"
  name        = "example-contract"
  description = "This contract is created by terraform"
  scope       = "context"
  dscp        = "VA"
}
```

## Argument Reference ##

* `name` - (Required) Name of the Contract created.
* `description` - (Optional) Description for Contract Model Object.
* `tenant_dn` - (Required) Distinguished Name of parent tenant object.
* `dscp` - (Optional) The target differentiated services code point (DSCP) of the path attached to the layer 3 outside profile. Default - `unspecified`. Possible Values :- `unspecified`, `VA`, `EF`, `CS0`, `CS1`, `CS2`, `CS3`, `CS4`, `CS5`, `CS6`, `CS7`, `AF11`, `AF12`, `AF13`, `AF21`, `AF22`, `AF23`, `AF31`, `AF32`, `AF33`, `AF41`, `AF42`, `AF43`.
* `scope` - (Optional) Represents the scope of current contract. If the scope is set as application-profile, the epg can only communicate with epgs in the same application-profile. Default - `context`. Possible Values :- `context`, `global`, `tenant`, `application-profile`.