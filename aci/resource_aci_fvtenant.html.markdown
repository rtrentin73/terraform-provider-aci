---
layout: "aci"
page_title: "ACI: aci_tenant"
sidebar_current: "docs-aci-resource-tenant"
description: |-
  Manages ACI Tenant
---

# aci_tenant #
Manages ACI Tenant

## Example Usage ##

```hcl
resource "aci_tenant" "example" {

  name  = "example"
  name_alias  = "example"
}
```
## Argument Reference ##
* `name` - (Required) name of Object tenant.
* `name_alias` - (Optional) name_alias for object tenant.
