---
layout: "aci"
page_title: "ACI: aci_cloudcontextprofile"
sidebar_current: "docs-aci-resource-cloudcontextprofile"
description: |-
  Manages ACI Cloud context profile
---

# aci_cloudcontextprofile #
Manages ACI Cloud context profile

## Example Usage ##

```hcl
resource "aci_cloudcontextprofile" "example" {

  tenant_dn  = "${aci_tenant.example.id}"

    name  = "example"

  annotation  = "example"
  name_alias  = "example"
  type  = "example"
}
```
## Argument Reference ##
* `tenant_dn` - (Required) Distinguished name of parent Tenant object.
* `name` - (Required) name of Object cloudcontextprofile.
* `annotation` - (Optional) annotation for object cloudcontextprofile.
* `name_alias` - (Optional) name_alias for object cloudcontextprofile.
* `type` - (Optional) component type

* `relation_cloud_rs_ctx_to_flow_log` - (Optional) Relation to class cloudAwsFlowLogPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_cloud_rs_to_ctx` - (Optional) Relation to class fvCtx. Cardinality - N_TO_ONE. Type - String.
                
* `relation_cloud_rs_ctx_profile_to_region` - (Optional) Relation to class cloudRegion. Cardinality - N_TO_ONE. Type - String.
                


