---
layout: "aci"
page_title: "ACI: aci_cloudexternalepg"
sidebar_current: "docs-aci-resource-cloudexternalepg"
description: |-
  Manages ACI Cloud external epg
---

# aci_cloudexternalepg #
Manages ACI Cloud external epg

## Example Usage ##

```hcl
resource "aci_cloudexternalepg" "example" {

  cloudapplicationcontainer_dn  = "${aci_cloudapplicationcontainer.example.id}"

    name  = "example"

  annotation  = "example"
  exception_tag  = "example"
  flood_on_encap  = "example"
  match_t  = "example"
  name_alias  = "example"
  pref_gr_memb  = "example"
  prio  = "example"
  route_reachability  = "example"
}
```
## Argument Reference ##
* `cloudapplicationcontainer_dn` - (Required) Distinguished name of parent Cloudapplicationcontainer object.
* `name` - (Required) name of Object cloudexternalepg.
* `annotation` - (Optional) annotation for object cloudexternalepg.
* `exception_tag` - (Optional) exception_tag for object cloudexternalepg.
* `flood_on_encap` - (Optional) flood_on_encap for object cloudexternalepg.
* `match_t` - (Optional) match criteria
* `name_alias` - (Optional) name_alias for object cloudexternalepg.
* `pref_gr_memb` - (Optional) pref_gr_memb for object cloudexternalepg.
* `prio` - (Optional) qos priority class id
* `route_reachability` - (Optional) route_reachability for object cloudexternalepg.

* `relation_fv_rs_sec_inherited` - (Optional) Relation to class fvEPg. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_prov` - (Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_cons_if` - (Optional) Relation to class vzCPIf. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_cust_qos_pol` - (Optional) Relation to class qosCustomPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_cons` - (Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_cloud_rs_cloud_e_pg_ctx` - (Optional) Relation to class fvCtx. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_prot_by` - (Optional) Relation to class vzTaboo. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_intra_epg` - (Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.
                


