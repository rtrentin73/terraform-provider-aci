---
layout: "aci"
page_title: "ACI: aci_tenant"
sidebar_current: "docs-aci-resource-tenant"
description: |-
  It manages single tenant.
---
# aci_tenant #

It manages single tenant.

## Example Usage ##

```hcl
resource "aci_tenant" "example" {
  name        = "example"
  description = "This tenant is created by terraform"
}
```

## Argument Reference ##

* `name` - (Required) Name of tenant.  
* `description` - (Optional) Description with tenant.
* `status` - (Optional) Status of object. Possible values are `created`, `updated`. Default :- `created`.
