---
layout: "aci"
page_title: "ACI: aci_cloud_providers_region"
sidebar_current: "docs-aci-resource-cloud_providers_region"
description: |-
  Manages ACI Cloud Providers Region
---

# aci_cloud_providers_region #
Manages ACI Cloud Providers Region

## Example Usage ##

```hcl
resource "aci_cloud_providers_region" "example" {

  cloud_provider_profile_dn  = "${aci_cloud_provider_profile.example.id}"

    name  = "example"

  admin_st  = "example"
  annotation  = "example"
  name_alias  = "example"
}
```
## Argument Reference ##
* `cloud_provider_profile_dn` - (Required) Distinguished name of parent CloudProviderProfile object.
* `name` - (Required) name of Object cloud_providers_region.
* `admin_st` - (Optional) administrative state of the object or policy
* `annotation` - (Optional) annotation for object cloud_providers_region.
* `name_alias` - (Optional) name_alias for object cloud_providers_region.



