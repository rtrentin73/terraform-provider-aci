---
layout: "aci"
page_title: "ACI: aci_contractsubject"
sidebar_current: "docs-aci-resource-contractsubject"
description: |-
  Manages ACI Contract subject
---

# aci_contractsubject #
Manages ACI Contract subject

## Example Usage ##

```hcl
resource "aci_contractsubject" "example" {

  contract_dn  = "${aci_contract.example.id}"

    name  = "example"

  annotation  = "example"
  cons_match_t  = "example"
  name_alias  = "example"
  prio  = "example"
  prov_match_t  = "example"
  rev_flt_ports  = "example"
  target_dscp  = "example"
}
```
## Argument Reference ##
* `contract_dn` - (Required) Distinguished name of parent Contract object.
* `name` - (Required) name of Object contractsubject.
* `annotation` - (Optional) annotation for object contractsubject.
* `cons_match_t` - (Optional) consumer subject match criteria
* `name_alias` - (Optional) name_alias for object contractsubject.
* `prio` - (Optional) priority level specifier
* `prov_match_t` - (Optional) consumer subject match criteria
* `rev_flt_ports` - (Optional) enables filter to apply on ingress and egress traffic
* `target_dscp` - (Optional) target dscp

* `relation_vz_rs_subj_graph_att` - (Optional) Relation to class vnsAbsGraph. Cardinality - N_TO_ONE. Type - String.
                
* `relation_vz_rs_sdwan_pol` - (Optional) Relation to class extdevSDWanSlaPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_vz_rs_subj_filt_att` - (Optional) Relation to class vzFilter. Cardinality - N_TO_M. Type - Set of String.
                


