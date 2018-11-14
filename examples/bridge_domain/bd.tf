resource "aci_tenant" "tenant_for_bridge_domain" {
  name        = "tenant_for_bd"
  description = "This tenant is created by terraform ACI provider"
}

resource "aci_bridge_domain" "demobd" {
  tenant_dn   = "${aci_tenant.tenant_for_bridge_domain.id}"
  name        = "test_tf_bd"
  description = "This bridge domain is created by terraform ACI provider"
  mac         = "00:22:BD:F8:19:FF"
  relation_fv_rs_bd_to_profile = "testprofile"
  relation_fv_rs_bd_to_relay_p = "testrelay"
  relation_fv_rs_abd_pol_mon_pol = "testabdpol"
  relation_fv_rs_bd_flood_to = ["uni/tn-1/flt-test_update"]
  relation_fv_rs_bd_to_fhs = "testfhs"
  relation_fv_rs_bd_to_netflow_monitor_pol{
    tn_netflow_monitor_pol_name = "testmonpolname"
    flt_type = "ipv4"
  }
  relation_fv_rs_bd_to_out = ["testbdout"]
}
