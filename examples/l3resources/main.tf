provider "aci" {
  username = "admin"
  password = "cisco123"
  url      = "https://192.168.10.102"
  insecure = true
}

resource "aci_tenant" "test_tenant" {
  name        = "tf_test_l3out_ten"
  #description = "This tenant is created by terraform"
}

resource "aci_l3_outside" "tf_l3_out" {
    tenant_dn = "${aci_tenant.test_tenant.id}"
    name = "test_l3Out"
  
}

resource "aci_external_network_instance_profile" "test_net_prof" {
    l3_outside_dn = "${aci_l3_outside.tf_l3_out.id}"
    name = "test_ext_network"
  
}

resource "aci_l3_ext_subnet" "test_ext_subnet" {
    external_network_instance_profile_dn = "${aci_external_network_instance_profile.test_net_prof.id}"
    ip = "10.0.1.0/8"
    description = "hello"
  
}

resource "aci_logical_node_profile" "test_nodep" {
    l3_outside_dn = "${aci_l3_outside.tf_l3_out.id}"
    name = "test_log_prof"
    relation_l3ext_rs_node_l3_out_att = ["topology/pod-1/node-201"]
  
}
