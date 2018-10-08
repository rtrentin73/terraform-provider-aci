---
layout: "aci"
page_title: "ACI: aci_application_profile"
sidebar_current: "docs-aci-resource-application_profile"
description: |-
  Manages ACI ApplicationProfile
---

# aci_application_profile #
Manages ACI ApplicationProfile

## Example Usage ##

```hcl
resource "aci_application_profile" "example" {

  fv_tenant_dn  = "${aci_fv_tenant.example.id}"
  name  = "example"
  name_alias  = "example"
  prio  = "example"
}
```
## Argument Reference ##
* `fv_tenant_dn` - (Required) Distinguished name of parent fvTenant object.
* `name` - (Required) name of Object application_profile.
* `name_alias` - (Optional) name_alias for object application_profile.
* `prio` - (Optional) priority class id
