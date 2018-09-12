---
layout: "aci"
page_title: "ACI: aci_ap"
sidebar_current: "docs-aci-resource-ap"
description: |-
  Manages ACI application profile.
---
# aci_app_profile #

Manages ACI application profile.

## Example Usage ##

```hcl

resource "aci_app_profile" "example" {
  tanent_dn   = "${aci_tenant.example_tenant.id}"
  name        = "example"
  description = "This app profile is created by terraform"
}
```

## Argument Reference ##

* `name` - (Required) Name of Application profile created or updated.
* `discription` - (Optional) Description for the Application Profile.
* `status` - (Optional) Status of object. Possible values are `created`, `updated`. Default :- `created`.
* `tanent_dn` - (Required) Distinguished name of parent tanent object.
