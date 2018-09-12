---
layout: "aci"
page_title: "ACI: aci_epg"
sidebar_current: "docs-aci-resource-epg"
description: |-
  Manages ACI Endpoint Groups.
---
# aci_epg #
Manages ACI Endpoint Groups.

## Example Usage ##

```hcl
resource "aci_epg" "example" {
  application_profile_dn = "${aci_app_profile.example_ap.id}"
  name                   = "example"
  description            = "This epg is created by terraform"
}
```

## Argument Reference ##

* `name` - (Required) Name of the Endpoint Group created in ACI.
* `description` - (Optional) Description for Endpoint group.
* `application_profile_dn` - (Required) Distinguished name of parent application profile.
* `is_attr_based` - (Optional) Is EPG attribute based or not? Default :- `false`.
* `preferred_policy_control` - (Optional) Preferred Policy control scheme with EPG. Default :- `unenforced`.
* `label_match_criteria` - (Optional) Label Matching criteria for EPG. Default :- `AteastOne`.