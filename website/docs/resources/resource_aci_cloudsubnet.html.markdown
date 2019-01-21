---
layout: "aci"
page_title: "ACI: aci_cloudsubnet"
sidebar_current: "docs-aci-resource-cloudsubnet"
description: |-
  Manages ACI Cloud subnet
---

# aci_cloudsubnet #
Manages ACI Cloud subnet

## Example Usage ##

```hcl
resource "aci_cloudsubnet" "example" {

  cloudcidrpool_dn  = "${aci_cloudcidrpool.example.id}"

    ip  = "example"

  annotation  = "example"
  ip  = "example"
  name_alias  = "example"
  scope  = "example"
  usage  = "example"
}
```
## Argument Reference ##
* `cloudcidrpool_dn` - (Required) Distinguished name of parent Cloudcidrpool object.
* `ip` - (Required) ip of Object cloudsubnet.
* `annotation` - (Optional) annotation for object cloudsubnet.
* `ip` - (Optional) ip address
* `name_alias` - (Optional) name_alias for object cloudsubnet.
* `scope` - (Optional) capability domain
* `usage` - (Optional) usage of the port

* `relation_cloud_rs_zone_attach` - (Optional) Relation to class cloudZone. Cardinality - N_TO_ONE. Type - String.
                
* `relation_cloud_rs_subnet_to_flow_log` - (Optional) Relation to class cloudAwsFlowLogPol. Cardinality - N_TO_ONE. Type - String.
                


