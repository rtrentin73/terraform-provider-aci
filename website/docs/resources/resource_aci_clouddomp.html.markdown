---
layout: "aci"
page_title: "ACI: aci_cloud_domain_profile"
sidebar_current: "docs-aci-resource-cloud_domain_profile"
description: |-
  Manages ACI Cloud Domain Profile
---

# aci_cloud_domain_profile #
Manages ACI Cloud Domain Profile

## Example Usage ##

```hcl
resource "aci_cloud_domain_profile" "example" {

  annotation  = "example"
  name_alias  = "example"
  site_id  = "example"
}
```
## Argument Reference ##
* `annotation` - (Optional) annotation for object cloud_domain_profile.
* `name_alias` - (Optional) name_alias for object cloud_domain_profile.
* `site_id` - (Optional) site_id for object cloud_domain_profile.



