---
layout: "aci"
page_title: "ACI: aci_cloudapplicationcontainer"
sidebar_current: "docs-aci-resource-cloudapplicationcontainer"
description: |-
  Manages ACI Cloud application container
---

# aci_cloudapplicationcontainer #
Manages ACI Cloud application container

## Example Usage ##

```hcl
resource "aci_cloudapplicationcontainer" "example" {

  tenant_dn  = "${aci_tenant.example.id}"

    name  = "example"

  annotation  = "example"
  name_alias  = "example"
}
```
## Argument Reference ##
* `tenant_dn` - (Required) Distinguished name of parent Tenant object.
* `name` - (Required) name of Object cloudapplicationcontainer.
* `annotation` - (Optional) annotation for object cloudapplicationcontainer.
* `name_alias` - (Optional) name_alias for object cloudapplicationcontainer.



