---
layout: "aci"
page_title: "ACI: aci_cloudawsprovider"
sidebar_current: "docs-aci-resource-cloudawsprovider"
description: |-
  Manages ACI Cloud aws provider
---

# aci_cloudawsprovider #
Manages ACI Cloud aws provider

## Example Usage ##

```hcl
resource "aci_cloudawsprovider" "example" {

  tenant_dn  = "${aci_tenant.example.id}"
  access_key_id  = "example"
  account_id  = "example"
  annotation  = "example"
  email  = "example"
  http_proxy  = "example"
  is_account_in_org  = "example"
  is_trusted  = "example"
  name_alias  = "example"
  provider_id  = "example"
  region  = "example"
  secret_access_key  = "example"
}
```
## Argument Reference ##
* `tenant_dn` - (Required) Distinguished name of parent Tenant object.
* `access_key_id` - (Optional) access_key_id for object cloudawsprovider.
* `account_id` - (Optional) account_id for object cloudawsprovider.
* `annotation` - (Optional) annotation for object cloudawsprovider.
* `email` - (Optional) email address of the local user
* `http_proxy` - (Optional) http_proxy for object cloudawsprovider.
* `is_account_in_org` - (Optional) is_account_in_org for object cloudawsprovider.
* `is_trusted` - (Optional) is_trusted for object cloudawsprovider.
* `name_alias` - (Optional) name_alias for object cloudawsprovider.
* `provider_id` - (Optional) provider_id for object cloudawsprovider.
* `region` - (Optional) region for object cloudawsprovider.
* `secret_access_key` - (Optional) secret_access_key for object cloudawsprovider.



