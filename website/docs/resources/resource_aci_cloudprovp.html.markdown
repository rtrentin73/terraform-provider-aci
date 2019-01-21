---
layout: "aci"
page_title: "ACI: aci_cloudproviderprofile"
sidebar_current: "docs-aci-resource-cloudproviderprofile"
description: |-
  Manages ACI Cloud provider profile
---

# aci_cloudproviderprofile #
Manages ACI Cloud provider profile

## Example Usage ##

```hcl
resource "aci_cloudproviderprofile" "example" {


    vendor  = "example"

  annotation  = "example"
  vendor  = "example"
}
```
## Argument Reference ##
* `vendor` - (Required) vendor of Object cloudproviderprofile.
* `annotation` - (Optional) annotation for object cloudproviderprofile.
* `vendor` - (Optional) vendor of the controller



