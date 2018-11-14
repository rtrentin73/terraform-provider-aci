resource "aci_tenant" "tenant_for_epg" {
  name        = "tenant_for_epg"
  description = "This tenant is created by terraform ACI provider"
}

resource "aci_application_profile" "app_profile_for_epg" {
  tenant_dn   = "${aci_tenant.tenant_for_epg.id}"
  name        = "ap_for_epg"
  description = "This app profile is created by terraform ACI providers"
}

resource "aci_application_epg" "demoepg" {
  application_profile_dn = "${aci_application_profile.app_profile_for_epg.id}"
  name                   = "tf_test_epg"
  description            = "This epg is created by terraform ACI providers"
  relation_fv_rs_dom_att = ["test"]
  relation_fv_rs_fc_path_att = ["testfabric"]
  relation_fv_rs_prov    = ["testcontract"]
  relation_fv_rs_cons_if = ["testconsif"]
  relation_fv_rs_sec_inherited = ["testinherited"]
  relation_fv_rs_node_att = ["testnodeatt"]
  relation_fv_rs_dpp_pol = "testdpppol"
  relation_fv_rs_cons = ["testrscons"]
  relation_fv_rs_trust_ctrl = "testtrustctrl"
  relation_fv_rs_path_att = ["testpathatt"]
  relation_fv_rs_prot_by = ["testprot"]
  relation_fv_rs_ae_pg_mon_pol = "aepgmonpol"
  relation_fv_rs_intra_epg = ["testintraepg"]
}
