---
layout: "aci"
page_title: "ACI: aci_applicationprofile"
sidebar_current: "docs-aci-resource-applicationprofile"
description: |-
  Manages ACI Application profile
---

# aci_applicationprofile #
Manages ACI Application profile

## Example Usage ##

```hcl
resource "aci_applicationprofile" "example" {

  tenant_dn  = "${aci_tenant.example.id}"

    name  = "example"

  annotation  = "example"
  name_alias  = "example"
  prio  = "example"
}
```
## Argument Reference ##
* `tenant_dn` - (Required) Distinguished name of parent Tenant object.
* `name` - (Required) name of Object applicationprofile.
* `annotation` - (Optional) annotation for object applicationprofile.
* `name_alias` - (Optional) name_alias for object applicationprofile.
* `prio` - (Optional) priority class id

* `relation_fv_rs_ap_mon_pol` - (Optional) Relation to class monEPGPol. Cardinality - N_TO_ONE. Type - String.
                


