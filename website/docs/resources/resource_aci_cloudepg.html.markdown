---
layout: "aci"
page_title: "ACI: aci_cloudepg"
sidebar_current: "docs-aci-resource-cloudepg"
description: |-
  Manages ACI Cloud epg
---

# aci_cloudepg #
Manages ACI Cloud epg

## Example Usage ##

```hcl
resource "aci_cloudepg" "example" {

  cloudapplicationcontainer_dn  = "${aci_cloudapplicationcontainer.example.id}"

    name  = "example"

  annotation  = "example"
  exception_tag  = "example"
  flood_on_encap  = "example"
  match_t  = "example"
  name_alias  = "example"
  pref_gr_memb  = "example"
  prio  = "example"
}
```
## Argument Reference ##
* `cloudapplicationcontainer_dn` - (Required) Distinguished name of parent Cloudapplicationcontainer object.
* `name` - (Required) name of Object cloudepg.
* `annotation` - (Optional) annotation for object cloudepg.
* `exception_tag` - (Optional) exception_tag for object cloudepg.
* `flood_on_encap` - (Optional) flood_on_encap for object cloudepg.
* `match_t` - (Optional) match criteria
* `name_alias` - (Optional) name_alias for object cloudepg.
* `pref_gr_memb` - (Optional) pref_gr_memb for object cloudepg.
* `prio` - (Optional) qos priority class id

* `relation_fv_rs_sec_inherited` - (Optional) Relation to class fvEPg. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_prov` - (Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_cons_if` - (Optional) Relation to class vzCPIf. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_cust_qos_pol` - (Optional) Relation to class qosCustomPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_cons` - (Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_cloud_rs_cloud_e_pg_ctx` - (Optional) Relation to class fvCtx. Cardinality - N_TO_ONE. Type - String.
                
* `relation_fv_rs_prot_by` - (Optional) Relation to class vzTaboo. Cardinality - N_TO_M. Type - Set of String.
                
* `relation_fv_rs_intra_epg` - (Optional) Relation to class vzBrCP. Cardinality - N_TO_M. Type - Set of String.
                


