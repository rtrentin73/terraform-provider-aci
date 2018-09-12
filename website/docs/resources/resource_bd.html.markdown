---
layout: "aci"
page_title: "ACI: aci_bridge_domain"
sidebar_current: "docs-aci-resource-bridge-domain"
description: |-
  Manages ACI Bridge Domain.
---

# aci_bridge_domain #
Manages the aci_bridge_domain model object.

## Example Usage ##
```hcl
resource "aci_bridge_domain" "example" {
  tanent_dn   = "${aci_tenant.example-tenant.id}"
  name        = "example"
  description = "This bridge domain is created by terraform"
  mac         = "00:22:BD:F8:19:FF"
}
```

## Argument Reference ##
* `name` - (Required) Name of the bridge domain created.
* `tenant_dn` - (Required) Distinguished Name of parent tenant object.
* `description` - (Optional) Description for Bridge domain object.
* `mac` - (Required) The MAC address of the bridge domain (BD) or switched virtual interface (SVI). Every BD by default takes the fabric wide default mac address. If user wants then he can override that address and with a different one By default the BD will take a `00:22:BD:F8:19:FF` mac address.
* `arp_flood` - (Optional) A property to specify whether ARP flooding is enabled. If flooding is disabled, unicast routing will be performed on the target IP address. Default - `false`.
* `optimize_wan` - (Optional) OptimizeWan Bandwidth flag. Enabled in sites. Default - `false`.
* `move_detect_mode` - (Optional) The End Point move detection option uses the Gratuitous Address Resolution Protocol (GARP). A gratuitous ARP is an ARP broadcast-type of packet that is used to verify that no other device on the network has the same IP address as the sending device. Default - empty string.Possible Values :- `garp`.
* `allow_intersite_bum_traffic` - (Optional) Control whether BUM traffic is allowed between sites. Default - `false`.
* `intersite_l2stretch` - (Optional) Flag to enable l2stretch between sites. Default - `false`.
* `ip_learning` - (Optional) Flag to enable iplearning. Default - `false`.
* `limit_ip_to_subnets` - (Optional) Limits IP address learning to the bridge domain subnets only. Every BD can have multiple subnets associated with it. By default, all IPs are learned.
* `ll_ip_address` - (Optional)  The override of the system generated IPv6 link-local address.
* `multi_dest_forwarding` - (Optional) The multiple destination forwarding method for L2 Multicast, Broadcast, and Link Layer traffic types. Default - `bd-flood`. Possible Values :- `bd-flood`,`drop`,`encap-flood`.
* `multicast` - (Optional) Flag to enable multicast. Default - `false`.
* `unicast_route` - (Optional) The forwarding method based on predefined forwarding criteria (IP or MAC address). Default - `true`.
* `unknown_unicast_mac` - (Optional) The forwarding method for unknown layer 2 destinations. Default - `proxy`. Possible Values :- `proxy`, `flood`.

* `unknown_multicast_mac` - (Optional) The parameter used by the node (i.e. a leaf) for forwarding data for an unknown multicast destination. Default - `flood`, Possible Values :- `flood`,`opt-flood`.
* `virtual_mac` - (Optional) Virtual MAC address of the BD/SVI. This is used when the Bridge Domain is extended to multiple sites using l2 Outside. Default - `not-applicable`.



