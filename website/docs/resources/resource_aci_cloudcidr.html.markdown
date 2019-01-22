---
layout: "aci"
page_title: "ACI: aci_cloud_cidr_pool"
sidebar_current: "docs-aci-resource-cloud_cidr_pool"
description: |-
  Manages ACI Cloud CIDR Pool
---

# aci_cloud_cidr_pool #
Manages ACI Cloud CIDR Pool

## Example Usage ##

```hcl
resource "aci_cloud_cidr_pool" "example" {

  cloud_context_profile_dn  = "${aci_cloud_context_profile.example.id}"

    addr  = "example"

  addr  = "example"
  annotation  = "example"
  name_alias  = "example"
  primary  = "example"
}
```
## Argument Reference ##
* `cloud_context_profile_dn` - (Required) Distinguished name of parent CloudContextProfile object.
* `addr` - (Required) addr of Object cloud_cidr_pool.
* `addr` - (Optional) peer address
* `annotation` - (Optional) annotation for object cloud_cidr_pool.
* `name_alias` - (Optional) name_alias for object cloud_cidr_pool.
* `primary` - (Optional) primary for object cloud_cidr_pool.



