---
layout: "aci"
page_title: "ACI: aci_autonomoussystemprofile"
sidebar_current: "docs-aci-resource-autonomoussystemprofile"
description: |-
  Manages ACI Autonomous system profile
---

# aci_autonomoussystemprofile #
Manages ACI Autonomous system profile

## Example Usage ##

```hcl
resource "aci_autonomoussystemprofile" "example" {

  annotation  = "example"
  asn  = "example"
  name_alias  = "example"
}
```
## Argument Reference ##
* `annotation` - (Optional) annotation for object autonomoussystemprofile.
* `asn` - (Optional) autonomous system number
* `name_alias` - (Optional) name_alias for object autonomoussystemprofile.



