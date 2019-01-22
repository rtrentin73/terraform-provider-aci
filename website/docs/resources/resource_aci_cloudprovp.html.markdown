---
layout: "aci"
page_title: "ACI: aci_cloud_provider_profile"
sidebar_current: "docs-aci-resource-cloud_provider_profile"
description: |-
  Manages ACI Cloud Provider Profile
---

# aci_cloud_provider_profile #
Manages ACI Cloud Provider Profile

## Example Usage ##

```hcl
resource "aci_cloud_provider_profile" "example" {


    vendor  = "example"

  annotation  = "example"
  vendor  = "example"
}
```
## Argument Reference ##
* `vendor` - (Required) vendor of Object cloud_provider_profile.
* `annotation` - (Optional) annotation for object cloud_provider_profile.
* `vendor` - (Optional) vendor of the controller



