---
layout: "aci"
page_title: "ACI: aci_subnet"
sidebar_current: "docs-aci-resource-subnet"
description: |-
  Manages ACI Subnet
---

# aci_subnet #
Manages ACI Subnet

## Example Usage ##

```hcl
resource "aci_subnet" "example" {

  fv_bd_dn  = "${aci_fv_bd.example.id}"
  ip  = "example"
  ctrl  = "example"
  ip  = "example"
  name_alias  = "example"
  preferred  = "example"
  scope  = "example"
  virtual  = "example"
}
```
## Argument Reference ##
* `fv_bd_dn` - (Required) Distinguished name of parent fvBD object.
* `ip` - (Required) ip of Object subnet.
* `ctrl` - (Optional) subnet control state
* `ip` - (Optional) default gateway IP address and mask
* `name_alias` - (Optional) name_alias for object subnet.
* `preferred` - (Optional) subnet preferred status
* `scope` - (Optional) subnet visibility
* `virtual` - (Optional) virtual for object subnet.
