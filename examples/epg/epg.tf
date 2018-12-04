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
  application_profile_dn       = "${aci_application_profile.app_profile_for_epg.id}"
  name                         = "tf_test_epg"
  description                  = "This epg is created by terraform ACI providers"
  flood_on_encap               = "disabled"
  fwd_ctrl                     = "none"
  is_attr_based_e_pg           = "no"
  match_t                      = "None"
  pc_enf_pref                  = "unenforced"
  pref_gr_memb                 = "exclude"
  prio                         = "unspecified"
  relation_fv_rs_bd            = "testbd_update"
  relation_fv_rs_dom_att       = ["test"]                                            # Relation to infraDomP class. Cardinality - N_TO_M.
  relation_fv_rs_fc_path_att   = ["testfabric"]                                      # Relation to fabricPathEp class. Cardinality - N_TO_M.
  relation_fv_rs_prov          = ["testcontract"]                                    # Relation to vzBrCP class. Cardinality - N_TO_M.
  relation_fv_rs_cons_if       = ["testconsif"]                                      # Relation to vzCPIf class. Cardinality - N_TO_M.
  relation_fv_rs_sec_inherited = ["testinherited"]                                   # Relation to fvEPg class. Cardinality - N_TO_M.
  relation_fv_rs_node_att      = ["testnodeatt"]                                     # Relation to fabricNode class. Cardinality - N_TO_M.
  relation_fv_rs_dpp_pol       = "testdpppol"                                        # Relation to qosDppPol class. Cardinality - N_TO_ONE.
  relation_fv_rs_cons          = ["testrscons"]                                      # Relation to vzBrCP class. Cardinality - N_TO_M.
  relation_fv_rs_trust_ctrl    = "testtrustctrl"                                     # Relation to fhsTrustCtrlPol class. Cardinality - N_TO_ONE.
  relation_fv_rs_path_att      = ["testpathatt"]                                     # Relation to fabricPathEp class. Cardinality - N_TO_M.
  relation_fv_rs_prot_by       = ["testprot"]                                        # Relation to vzTaboo class. Cardinality - N_TO_M.
  relation_fv_rs_ae_pg_mon_pol = "aepgmonpol"                                        # Relation to monEPGPol class. Cardinality - N_TO_ONE.
  relation_fv_rs_intra_epg     = ["testintraepg"]                                    # Relation to vzBrCP class. Cardinality - N_TO_M.
}
