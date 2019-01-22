---
layout: "aci"
page_title: "ACI: aci_cloud_endpoint_selectorfor_external_e_pgs"
sidebar_current: "docs-aci-resource-cloud_endpoint_selectorfor_external_e_pgs"
description: |-
  Manages ACI Cloud Endpoint Selector for External EPgs
---

# aci_cloud_endpoint_selectorfor_external_e_pgs #
Manages ACI Cloud Endpoint Selector for External EPgs

## Example Usage ##

```hcl
resource "aci_cloud_endpoint_selectorfor_external_e_pgs" "example" {

  cloud_external_e_pg_dn  = "${aci_cloud_external_e_pg.example.id}"

    name  = "example"

  annotation  = "example"
  is_shared  = "example"
  match_expression  = "example"
  name_alias  = "example"
  subnet  = "example"
}
```
## Argument Reference ##
* `cloud_external_e_pg_dn` - (Required) Distinguished name of parent CloudExternalEPg object.
* `name` - (Required) name of Object cloud_endpoint_selectorfor_external_e_pgs.
* `annotation` - (Optional) annotation for object cloud_endpoint_selectorfor_external_e_pgs.
* `is_shared` - (Optional) is_shared for object cloud_endpoint_selectorfor_external_e_pgs.
* `match_expression` - (Optional) match_expression for object cloud_endpoint_selectorfor_external_e_pgs.
* `name_alias` - (Optional) name_alias for object cloud_endpoint_selectorfor_external_e_pgs.
* `subnet` - (Optional) subnet for object cloud_endpoint_selectorfor_external_e_pgs.



