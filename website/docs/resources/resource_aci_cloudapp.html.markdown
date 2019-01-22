---
layout: "aci"
page_title: "ACI: aci_cloud_applicationcontainer"
sidebar_current: "docs-aci-resource-cloud_applicationcontainer"
description: |-
  Manages ACI Cloud Application container
---

# aci_cloud_applicationcontainer #
Manages ACI Cloud Application container

## Example Usage ##

```hcl
resource "aci_cloud_applicationcontainer" "example" {

  tenant_dn  = "${aci_tenant.example.id}"

    name  = "example"

  annotation  = "example"
  name_alias  = "example"
}
```
## Argument Reference ##
* `tenant_dn` - (Required) Distinguished name of parent Tenant object.
* `name` - (Required) name of Object cloud_applicationcontainer.
* `annotation` - (Optional) annotation for object cloud_applicationcontainer.
* `name_alias` - (Optional) name_alias for object cloud_applicationcontainer.



