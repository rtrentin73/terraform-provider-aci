---
layout: "aci"
page_title: "ACI: aci_clouddomainprofile"
sidebar_current: "docs-aci-resource-clouddomainprofile"
description: |-
  Manages ACI Cloud domain profile
---

# aci_clouddomainprofile #
Manages ACI Cloud domain profile

## Example Usage ##

```hcl
resource "aci_clouddomainprofile" "example" {

  annotation  = "example"
  name_alias  = "example"
  site_id  = "example"
}
```
## Argument Reference ##
* `annotation` - (Optional) annotation for object clouddomainprofile.
* `name_alias` - (Optional) name_alias for object clouddomainprofile.
* `site_id` - (Optional) site_id for object clouddomainprofile.



