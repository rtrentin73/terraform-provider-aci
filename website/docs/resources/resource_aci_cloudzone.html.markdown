---
layout: "aci"
page_title: "ACI: aci_cloud_availability_zone"
sidebar_current: "docs-aci-resource-cloud_availability_zone"
description: |-
  Manages ACI Cloud Availability Zone
---

# aci_cloud_availability_zone #
Manages ACI Cloud Availability Zone

## Example Usage ##

```hcl
resource "aci_cloud_availability_zone" "example" {

  cloud_providers_region_dn  = "${aci_cloud_providers_region.example.id}"

    name  = "example"

  annotation  = "example"
  name_alias  = "example"
}
```
## Argument Reference ##
* `cloud_providers_region_dn` - (Required) Distinguished name of parent CloudProvidersRegion object.
* `name` - (Required) name of Object cloud_availability_zone.
* `annotation` - (Optional) annotation for object cloud_availability_zone.
* `name_alias` - (Optional) name_alias for object cloud_availability_zone.



