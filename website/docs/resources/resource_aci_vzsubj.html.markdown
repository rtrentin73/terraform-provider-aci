---
layout: "aci"
page_title: "ACI: aci_contract_subject"
sidebar_current: "docs-aci-resource-contract_subject"
description: |-
  Manages ACI ContractSubject
---

# aci_contract_subject #
Manages ACI ContractSubject

## Example Usage ##

```hcl
resource "aci_contract_subject" "example" {

  vz_br_cp_dn  = "${aci_vz_br_cp.example.id}"
  name  = "example"
  cons_match_t  = "example"
  name_alias  = "example"
  prio  = "example"
  prov_match_t  = "example"
  rev_flt_ports  = "example"
  target_dscp  = "example"
}
```
## Argument Reference ##
* `vz_br_cp_dn` - (Required) Distinguished name of parent vzBrCP object.
* `name` - (Required) name of Object contract_subject.
* `cons_match_t` - (Optional) consumer subject match criteria
* `name_alias` - (Optional) name_alias for object contract_subject.
* `prio` - (Optional) priority level specifier
* `prov_match_t` - (Optional) consumer subject match criteria
* `rev_flt_ports` - (Optional) enables filter to apply on ingress and egress traffic
* `target_dscp` - (Optional) target dscp
