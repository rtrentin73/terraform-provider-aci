---
layout: "aci"
page_title: "ACI: aci_cloudavailabilityzone"
sidebar_current: "docs-aci-resource-cloudavailabilityzone"
description: |-
  Manages ACI Cloud availability zone
---

# aci_cloudavailabilityzone #
Manages ACI Cloud availability zone

## Example Usage ##

```hcl
resource "aci_cloudavailabilityzone" "example" {

  cloudprovidersregion_dn  = "${aci_cloudprovidersregion.example.id}"

    name  = "example"

  annotation  = "example"
  name_alias  = "example"
}
```
## Argument Reference ##
* `cloudprovidersregion_dn` - (Required) Distinguished name of parent Cloudprovidersregion object.
* `name` - (Required) name of Object cloudavailabilityzone.
* `annotation` - (Optional) annotation for object cloudavailabilityzone.
* `name_alias` - (Optional) name_alias for object cloudavailabilityzone.



