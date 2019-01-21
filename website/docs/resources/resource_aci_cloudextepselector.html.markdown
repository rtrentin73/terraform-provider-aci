---
layout: "aci"
page_title: "ACI: aci_cloudendpointselectorforexternalepgs"
sidebar_current: "docs-aci-resource-cloudendpointselectorforexternalepgs"
description: |-
  Manages ACI Cloud endpoint selector for external epgs
---

# aci_cloudendpointselectorforexternalepgs #
Manages ACI Cloud endpoint selector for external epgs

## Example Usage ##

```hcl
resource "aci_cloudendpointselectorforexternalepgs" "example" {

  cloudexternalepg_dn  = "${aci_cloudexternalepg.example.id}"

    name  = "example"

  annotation  = "example"
  is_shared  = "example"
  match_expression  = "example"
  name_alias  = "example"
  subnet  = "example"
}
```
## Argument Reference ##
* `cloudexternalepg_dn` - (Required) Distinguished name of parent Cloudexternalepg object.
* `name` - (Required) name of Object cloudendpointselectorforexternalepgs.
* `annotation` - (Optional) annotation for object cloudendpointselectorforexternalepgs.
* `is_shared` - (Optional) is_shared for object cloudendpointselectorforexternalepgs.
* `match_expression` - (Optional) match_expression for object cloudendpointselectorforexternalepgs.
* `name_alias` - (Optional) name_alias for object cloudendpointselectorforexternalepgs.
* `subnet` - (Optional) subnet for object cloudendpointselectorforexternalepgs.



