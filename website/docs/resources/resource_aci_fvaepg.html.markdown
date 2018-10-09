---
layout: "aci"
page_title: "ACI: aci_application_epg"
sidebar_current: "docs-aci-resource-application_epg"
description: |-
  Manages ACI Application EPG
---

# aci_application_epg #
Manages ACI Application EPG

## Example Usage ##

```hcl
resource "aci_application_epg" "example" {

  application_profile_dn  = "${aci_application_profile.example.id}"
  name  = "example"
  flood_on_encap  = "example"
  fwd_ctrl  = "example"
  is_attr_based_e_pg  = "example"
  match_t  = "example"
  name_alias  = "example"
  pc_enf_pref  = "example"
  pref_gr_memb  = "example"
  prio  = "example"
}
```
## Argument Reference ##
* `application_profile_dn` - (Required) Distinguished name of parent ApplicationProfile object.
* `name` - (Required) name of Object application_epg.
* `flood_on_encap` - (Optional) flood_on_encap for object application_epg.
* `fwd_ctrl` - (Optional) fwd_ctrl for object application_epg.
* `is_attr_based_e_pg` - (Optional) is_attr_based_e_pg for object application_epg.
* `match_t` - (Optional) match criteria
* `name_alias` - (Optional) name_alias for object application_epg.
* `pc_enf_pref` - (Optional) enforcement preference
* `pref_gr_memb` - (Optional) pref_gr_memb for object application_epg.
* `prio` - (Optional) qos priority class id
