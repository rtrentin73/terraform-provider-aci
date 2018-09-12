---
layout: "aci"
page_title: "ACI: aci_subject"
sidebar_current: "docs-aci-resource-subject"
description: |-
  Manages ACI Subject Model Object.
---

# aci_subject #
Manages ACI subject Model Object.

## Example Usage ##

```hcl
resource "aci_subject" "example-subject" {
  contract_dn = "${aci_contract.example-contract.id}"
  name        = "example-subject"
  description = "This subject is created by terraform"
}
```

## Argument Reference ##

* `name` - (Required) Name of the Subject model object.
* `description` - (Optional) Description for the Subject model object.
* `contract_dn` - (Required) Distinguished name of the parent contract object.
* `consumer_match` - (Optional) The subject match criteria across consumers. Default - `AtleastOne`. Possible Values :- `All`, `AtleastOne`, `AtmostOne`, `None`.
* `provider_match` - (Optional) The subject match criteria across providers. Default - `AtleastOne` . Possible Values :- `All`, `AtleastOne`, `AtmostOne`, `None`.
* `priority` - (Optional) The priority level of a sub application running behind an endpoint group, such as an Exchange server. Default - `unspecified`. Possible Values :- `unspecified`,`level1`, `level2`, `level3`.
* `dscp` - (Optional) The target differentiated services code point (DSCP) of the path attached to the layer 3 outside profile. Defualt - `unspecified` . Possible Values :- `unspecified`, `VA`, `EF`, `CS0`, `CS1`, `CS2`, `CS3`, `CS4`, `CS5`, `CS6`, `CS7`, `AF11`, `AF12`, `AF13`, `AF21`, `AF22`, `AF23`, `AF31`, `AF32`, `AF33`, `AF41`, `AF42`, `AF43`.
* `reverse_filter_ports` - (Optional) Enables the filter to apply on both ingress and egress traffic. Default - `true`.

