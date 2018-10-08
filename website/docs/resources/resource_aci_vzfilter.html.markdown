---
layout: "aci"
page_title: "ACI: aci_filter"
sidebar_current: "docs-aci-resource-filter"
description: |-
  Manages ACI Filter
---

# aci_filter #
Manages ACI Filter

## Example Usage ##

```hcl
resource "aci_filter" "example" {

  fv_tenant_dn  = "${aci_fv_tenant.example.id}"
  name  = "example"
  name_alias  = "example"
}
```
## Argument Reference ##
* `fv_tenant_dn` - (Required) Distinguished name of parent fvTenant object.
* `name` - (Required) name of Object filter.
* `name_alias` - (Optional) name_alias for object filter.
