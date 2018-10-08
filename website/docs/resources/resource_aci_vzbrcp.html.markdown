---
layout: "aci"
page_title: "ACI: aci_contract"
sidebar_current: "docs-aci-resource-contract"
description: |-
  Manages ACI Contract
---

# aci_contract #
Manages ACI Contract

## Example Usage ##

```hcl
resource "aci_contract" "example" {

  fv_tenant_dn  = "${aci_fv_tenant.example.id}"
  name  = "example"
  name_alias  = "example"
  prio  = "example"
  scope  = "example"
  target_dscp  = "example"
}
```
## Argument Reference ##
* `fv_tenant_dn` - (Required) Distinguished name of parent fvTenant object.
* `name` - (Required) name of Object contract.
* `name_alias` - (Optional) name_alias for object contract.
* `prio` - (Optional) priority level of the service contract
* `scope` - (Optional) scope of contract
* `target_dscp` - (Optional) target dscp
