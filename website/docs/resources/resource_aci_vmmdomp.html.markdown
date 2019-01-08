---
layout: "aci"
page_title: "ACI: aci_vmm_domain"
sidebar_current: "docs-aci-resource-vmm_domain"
description: |-
  Manages ACI VMM Domain
---

# aci_vmm_domain #
Manages ACI VMM Domain

## Example Usage ##

```hcl
resource "aci_vmm_domain" "example" {

  provider_profile_dn  = "${aci_provider_profile.example.id}"
  name  = "example"
  access_mode  = "example"
  arp_learning  = "example"
  ctrl_knob  = "example"
  delimiter  = "example"
  enable_ave  = "example"
  encap_mode  = "example"
  enf_pref  = "example"
  ep_ret_time  = "example"
  mcast_addr  = "example"
  mode  = "example"
  name_alias  = "example"
  pref_encap_mode  = "example"
}
```
## Argument Reference ##
* `provider_profile_dn` - (Required) Distinguished name of parent ProviderProfile object.
* `name` - (Required) name of Object vmm_domain.
* `access_mode` - (Optional) access_mode for object vmm_domain.
* `arp_learning` - (Optional) arp_learning for object vmm_domain.
* `ctrl_knob` - (Optional) ctrl_knob for object vmm_domain.
* `delimiter` - (Optional) delimiter for object vmm_domain.
* `enable_ave` - (Optional) enable_ave for object vmm_domain.
* `encap_mode` - (Optional) encap_mode for object vmm_domain.
* `enf_pref` - (Optional) switching enforcement preference
* `ep_ret_time` - (Optional) ep_ret_time for object vmm_domain.
* `mcast_addr` - (Optional) multicast address
* `mode` - (Optional) switch used for the domain profile
* `name_alias` - (Optional) name_alias for object vmm_domain.
* `pref_encap_mode` - (Optional) pref_encap_mode for object vmm_domain.

* `relation_infra_rs_vlan_ns` - (Optional) Relation to class fvnsVlanInstP. Cardinality - N_TO_ONE. Type - String.
                
* `relation_vmm_rs_dom_mcast_addr_ns` - (Optional) Relation to class fvnsMcastAddrInstP. Cardinality - N_TO_ONE. Type - String.
                
* `relation_vmm_rs_default_cdp_if_pol` - (Optional) Relation to class cdpIfPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_vmm_rs_default_lacp_lag_pol` - (Optional) Relation to class lacpLagPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_infra_rs_vlan_ns_def` - (Optional) Relation to class fvnsAInstP. Cardinality - N_TO_ONE. Type - String.
                
* `relation_infra_rs_vip_addr_ns` - (Optional) Relation to class fvnsAddrInst. Cardinality - N_TO_ONE. Type - String.
                
* `relation_vmm_rs_default_lldp_if_pol` - (Optional) Relation to class lldpIfPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_vmm_rs_default_l2_inst_pol` - (Optional) Relation to class l2InstPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_vmm_rs_default_stp_if_pol` - (Optional) Relation to class stpIfPol. Cardinality - N_TO_ONE. Type - String.
                
* `relation_infra_rs_dom_vxlan_ns_def` - (Optional) Relation to class fvnsAInstP. Cardinality - N_TO_ONE. Type - String.
                
* `relation_vmm_rs_default_fw_pol` - (Optional) Relation to class nwsFwPol. Cardinality - N_TO_ONE. Type - String.
                


