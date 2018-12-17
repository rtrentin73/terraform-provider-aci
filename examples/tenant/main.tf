provider "aci" {
  username = "admin"
  password = "cisco123"
  url      = "https://192.168.10.102"
  insecure = true
}

resource "aci_tenant" "demotenant" {
  name                        = "tf_test_tenant"
  description                 = "This tenant is created by terraform"
  relation_fv_rs_tn_deny_rule = ["${aci_filter.deny_rule_filter1.id}", "${aci_filter.deny_rule_filter2.id}"] # Relation to vzFilter class. Cardinality - N_TO_M.
}
