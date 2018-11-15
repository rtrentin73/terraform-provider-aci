provider "aci" {
  username = "admin"
  password = "cisco123"
  url      = "https://192.168.10.102"
  insecure = true
}

resource "aci_tenant" "demotenant" {
  name                        = "tf_test_tenant"
  description                 = "This tenant is created by terraform"
  relation_fv_rs_tn_deny_rule = ["uni/tn-1/flt-test_update", "uni/tn-1/flt-test_update2"] # Relation to vzFilter class. Cardinality - N_TO_M.
}
