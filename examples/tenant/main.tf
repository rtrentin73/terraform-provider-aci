provider "aci" {
  username = "admin"
  password = "cisco123"
  url      = "https://192.168.10.102"
  insecure = true
}

resource "aci_tenant" "demotenant" {
  name        = "tf_test_tenant"
  description = "This tenant is created by terraform"
  relation_to_vz_filter = ["uni/tn-1/flt-1"]
}
