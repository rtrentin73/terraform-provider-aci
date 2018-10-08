---
layout: "aci"
page_title: "ACI: aci_bridge_domain"
sidebar_current: "docs-aci-resource-bridge_domain"
description: |-
  Manages ACI BridgeDomain
---

# aci_bridge_domain #
Manages ACI BridgeDomain

## Example Usage ##

```hcl
resource "aci_bridge_domain" "example" {

  fv_tenant_dn  = "${aci_fv_tenant.example.id}"
  name  = "example"
  optimize_wan_bandwidth  = "example"
  arp_flood  = "example"
  ep_clear  = "example"
  ep_move_detect_mode  = "example"
  intersite_bum_traffic_allow  = "example"
  intersite_l2_stretch  = "example"
  ip_learning  = "example"
  limit_ip_learn_to_subnets  = "example"
  ll_addr  = "example"
  mac  = "example"
  mcast_allow  = "example"
  multi_dst_pkt_act  = "example"
  name_alias  = "example"
  type  = "example"
  unicast_route  = "example"
  unk_mac_ucast_act  = "example"
  unk_mcast_act  = "example"
  vmac  = "example"
}
```
## Argument Reference ##
* `fv_tenant_dn` - (Required) Distinguished name of parent fvTenant object.
* `name` - (Required) name of Object bridge_domain.
* `optimize_wan_bandwidth` - (Optional) optimize_wan_bandwidth for object bridge_domain.
* `arp_flood` - (Optional) arp flood enable
* `ep_clear` - (Optional) ep_clear for object bridge_domain.
* `ep_move_detect_mode` - (Optional) ep move detection garp based mode
* `intersite_bum_traffic_allow` - (Optional) intersite_bum_traffic_allow for object bridge_domain.
* `intersite_l2_stretch` - (Optional) intersite_l2_stretch for object bridge_domain.
* `ip_learning` - (Optional) Endpoint Dataplane Learning
* `limit_ip_learn_to_subnets` - (Optional) limits ip learning to bd subnets only
* `ll_addr` - (Optional) override of system generated ipv6 link-local address
* `mac` - (Optional) mac address
* `mcast_allow` - (Optional) mcast_allow for object bridge_domain.
* `multi_dst_pkt_act` - (Optional) forwarding method for multi destinations
* `name_alias` - (Optional) name_alias for object bridge_domain.
* `type` - (Optional) component type
* `unicast_route` - (Optional) Unicast routing
* `unk_mac_ucast_act` - (Optional) forwarding method for l2 destinations
* `unk_mcast_act` - (Optional) parameter used by node to forward data
* `vmac` - (Optional) vmac for object bridge_domain.
