---
layout: "aci"
page_title: "ACI: aci_cloudendpointselector"
sidebar_current: "docs-aci-resource-cloudendpointselector"
description: |-
  Manages ACI Cloud endpoint selector
---

# aci_cloudendpointselector #
Manages ACI Cloud endpoint selector

## Example Usage ##

```hcl
resource "aci_cloudendpointselector" "example" {

  cloudepg_dn  = "${aci_cloudepg.example.id}"

    name  = "example"

  annotation  = "example"
  match_expression  = "example"
  name_alias  = "example"
}
```
## Argument Reference ##
* `cloudepg_dn` - (Required) Distinguished name of parent Cloudepg object.
* `name` - (Required) name of Object cloudendpointselector.
* `annotation` - (Optional) annotation for object cloudendpointselector.
* `match_expression` - (Optional) match_expression for object cloudendpointselector.
* `name_alias` - (Optional) name_alias for object cloudendpointselector.



