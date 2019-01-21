---
layout: "aci"
page_title: "ACI: aci_cloudprovidersregion"
sidebar_current: "docs-aci-resource-cloudprovidersregion"
description: |-
  Manages ACI Cloud providers region
---

# aci_cloudprovidersregion #
Manages ACI Cloud providers region

## Example Usage ##

```hcl
resource "aci_cloudprovidersregion" "example" {

  cloudproviderprofile_dn  = "${aci_cloudproviderprofile.example.id}"

    name  = "example"

  admin_st  = "example"
  annotation  = "example"
  name_alias  = "example"
}
```
## Argument Reference ##
* `cloudproviderprofile_dn` - (Required) Distinguished name of parent Cloudproviderprofile object.
* `name` - (Required) name of Object cloudprovidersregion.
* `admin_st` - (Optional) administrative state of the object or policy
* `annotation` - (Optional) annotation for object cloudprovidersregion.
* `name_alias` - (Optional) name_alias for object cloudprovidersregion.



