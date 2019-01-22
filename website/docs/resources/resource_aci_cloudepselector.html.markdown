---
layout: "aci"
page_title: "ACI: aci_cloud_endpoint_selector"
sidebar_current: "docs-aci-resource-cloud_endpoint_selector"
description: |-
  Manages ACI Cloud Endpoint Selector
---

# aci_cloud_endpoint_selector #
Manages ACI Cloud Endpoint Selector

## Example Usage ##

```hcl
resource "aci_cloud_endpoint_selector" "example" {

  cloud_e_pg_dn  = "${aci_cloud_e_pg.example.id}"

    name  = "example"

  annotation  = "example"
  match_expression  = "example"
  name_alias  = "example"
}
```
## Argument Reference ##
* `cloud_e_pg_dn` - (Required) Distinguished name of parent CloudEPg object.
* `name` - (Required) name of Object cloud_endpoint_selector.
* `annotation` - (Optional) annotation for object cloud_endpoint_selector.
* `match_expression` - (Optional) match_expression for object cloud_endpoint_selector.
* `name_alias` - (Optional) name_alias for object cloud_endpoint_selector.



