---
layout: "aci"
page_title: "ACI: aci_cloudcidrpool"
sidebar_current: "docs-aci-resource-cloudcidrpool"
description: |-
  Manages ACI Cloud cidr pool
---

# aci_cloudcidrpool #
Manages ACI Cloud cidr pool

## Example Usage ##

```hcl
resource "aci_cloudcidrpool" "example" {

  cloudcontextprofile_dn  = "${aci_cloudcontextprofile.example.id}"

    addr  = "example"

  addr  = "example"
  annotation  = "example"
  name_alias  = "example"
  primary  = "example"
}
```
## Argument Reference ##
* `cloudcontextprofile_dn` - (Required) Distinguished name of parent Cloudcontextprofile object.
* `addr` - (Required) addr of Object cloudcidrpool.
* `addr` - (Optional) peer address
* `annotation` - (Optional) annotation for object cloudcidrpool.
* `name_alias` - (Optional) name_alias for object cloudcidrpool.
* `primary` - (Optional) primary for object cloudcidrpool.



