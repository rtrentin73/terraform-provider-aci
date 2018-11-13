package aci


import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceAciApplicationEPG() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciApplicationEPGCreate,
		Update: resourceAciApplicationEPGUpdate,
		Read:   resourceAciApplicationEPGRead,
		Delete: resourceAciApplicationEPGDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciApplicationEPGImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"application_profile_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			
			"name": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			
            
			"flood_on_encap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Mo doc not defined in techpub!!!",
                
				ValidateFunc: validation.StringInSlice([]string{
				"disabled",
                "enabled",
                }, false),
                
			},
            
			"fwd_ctrl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Mo doc not defined in techpub!!!",
                
				ValidateFunc: validation.StringInSlice([]string{
				"none",
                "proxy-arp",
                }, false),
                
			},
            
			"is_attr_based_e_pg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Mo doc not defined in techpub!!!",
                
				ValidateFunc: validation.StringInSlice([]string{
				"no",
                "yes",
                }, false),
                
			},
            
			"match_t": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "match criteria",
                
				ValidateFunc: validation.StringInSlice([]string{
				"All",
                "AtleastOne",
                "AtmostOne",
                "None",
                }, false),
                
			},
            
			"name_alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Mo doc not defined in techpub!!!",
                
			},
            
			"pc_enf_pref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "enforcement preference",
                
				ValidateFunc: validation.StringInSlice([]string{
				"enforced",
                "unenforced",
                }, false),
                
			},
            
			"pref_gr_memb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Mo doc not defined in techpub!!!",
                
				ValidateFunc: validation.StringInSlice([]string{
				"exclude",
                "include",
                }, false),
                
			},
            
			"prio": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "qos priority class id",
                
				ValidateFunc: validation.StringInSlice([]string{
				"level1",
                "level2",
                "level3",
                "unspecified",
                }, false),
                
			},
            
			
			"relation_fv_rs_dom_att": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{ Type: schema.TypeString,},
				Optional: 	 true,
				Description: "Create relation to infraDomP",
     			Set:         schema.HashString,
			},
			"relation_fv_rs_fc_path_att": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{ Type: schema.TypeString,},
				Optional: 	 true,
				Description: "Create relation to fabricPathEp",
     			Set:         schema.HashString,
			},
			"relation_fv_rs_prov": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{ Type: schema.TypeString,},
				Optional: 	 true,
				Description: "Create relation to vzBrCP",
     			Set:         schema.HashString,
			},
			"relation_fv_rs_cons_if": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{ Type: schema.TypeString,},
				Optional: 	 true,
				Description: "Create relation to vzCPIf",
     			Set:         schema.HashString,
			},
			"relation_fv_rs_sec_inherited": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{ Type: schema.TypeString,},
				Optional: 	 true,
				Description: "Create relation to fvEPg",
     			Set:         schema.HashString,
			},
			"relation_fv_rs_node_att": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{ Type: schema.TypeString,},
				Optional: 	 true,
				Description: "Create relation to fabricNode",
     			Set:         schema.HashString,
			},
			"relation_fv_rs_dpp_pol": &schema.Schema{
				Type:     schema.TypeString,

				Optional: 	 true,
				Description: "Create relation to qosDppPol",

			},
			"relation_fv_rs_cons": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{ Type: schema.TypeString,},
				Optional: 	 true,
				Description: "Create relation to vzBrCP",
     			Set:         schema.HashString,
			},
			"relation_fv_rs_trust_ctrl": &schema.Schema{
				Type:     schema.TypeString,

				Optional: 	 true,
				Description: "Create relation to fhsTrustCtrlPol",

			},
			"relation_fv_rs_path_att": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{ Type: schema.TypeString,},
				Optional: 	 true,
				Description: "Create relation to fabricPathEp",
     			Set:         schema.HashString,
			},
			"relation_fv_rs_prot_by": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{ Type: schema.TypeString,},
				Optional: 	 true,
				Description: "Create relation to vzTaboo",
     			Set:         schema.HashString,
			},
			"relation_fv_rs_ae_pg_mon_pol": &schema.Schema{
				Type:     schema.TypeString,

				Optional: 	 true,
				Description: "Create relation to monEPGPol",

			},
			"relation_fv_rs_intra_epg": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{ Type: schema.TypeString,},
				Optional: 	 true,
				Description: "Create relation to vzBrCP",
     			Set:         schema.HashString,
			},

		}),
	}
}

func getRemoteApplicationEPG(client *client.Client, dn string) (*models.ApplicationEPG, error) {
	fvAEPgCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	fvAEPg := models.ApplicationEPGFromContainer(fvAEPgCont)

	if fvAEPg.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", fvAEPg.DistinguishedName)
	}

	return fvAEPg, nil
}

func setApplicationEPGAttributes(fvAEPg *models.ApplicationEPG, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(fvAEPg.DistinguishedName)
	d.Set("description", fvAEPg.Description)
	d.Set("application_profile_dn", GetParentDn(fvAEPg.DistinguishedName))
	fvAEPgMap , _ := fvAEPg.ToMap()
     
	d.Set("flood_on_encap", fvAEPgMap["floodOnEncap"]) 
	d.Set("fwd_ctrl", fvAEPgMap["fwdCtrl"]) 
	d.Set("is_attr_based_e_pg", fvAEPgMap["isAttrBasedEPg"]) 
	d.Set("match_t", fvAEPgMap["matchT"]) 
	d.Set("name_alias", fvAEPgMap["nameAlias"]) 
	d.Set("pc_enf_pref", fvAEPgMap["pcEnfPref"]) 
	d.Set("pref_gr_memb", fvAEPgMap["prefGrMemb"]) 
	d.Set("prio", fvAEPgMap["prio"])
	return d
}

func resourceAciApplicationEPGImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	fvAEPg, err := getRemoteApplicationEPG(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setApplicationEPGAttributes(fvAEPg, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciApplicationEPGCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	name := d.Get("name").(string)
	ApplicationProfileDn := d.Get("application_profile_dn").(string)
	
	fvAEPgAttr := models.ApplicationEPGAttributes{} 
    if FloodOnEncap, ok := d.GetOk("flood_on_encap"); ok {
        fvAEPgAttr.FloodOnEncap  = FloodOnEncap.(string)
    } 
    if FwdCtrl, ok := d.GetOk("fwd_ctrl"); ok {
        fvAEPgAttr.FwdCtrl  = FwdCtrl.(string)
    } 
    if IsAttrBasedEPg, ok := d.GetOk("is_attr_based_e_pg"); ok {
        fvAEPgAttr.IsAttrBasedEPg  = IsAttrBasedEPg.(string)
    } 
    if MatchT, ok := d.GetOk("match_t"); ok {
        fvAEPgAttr.MatchT  = MatchT.(string)
    } 
    if NameAlias, ok := d.GetOk("name_alias"); ok {
        fvAEPgAttr.NameAlias  = NameAlias.(string)
    } 
    if PcEnfPref, ok := d.GetOk("pc_enf_pref"); ok {
        fvAEPgAttr.PcEnfPref  = PcEnfPref.(string)
    } 
    if PrefGrMemb, ok := d.GetOk("pref_gr_memb"); ok {
        fvAEPgAttr.PrefGrMemb  = PrefGrMemb.(string)
    } 
    if Prio, ok := d.GetOk("prio"); ok {
        fvAEPgAttr.Prio  = Prio.(string)
    }
	fvAEPg := models.NewApplicationEPG(fmt.Sprintf("epg-%s",name),ApplicationProfileDn, desc, fvAEPgAttr)  
	
	
	err := aciClient.Save(fvAEPg)
	if err != nil {
		return err
	}
	
	if  relationTofvRsDomAtt, ok := d.GetOk("relation_fv_rs_dom_att") ; ok {
		relationParamList := toStringList(relationTofvRsDomAtt.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsDomAtt(fvAEPg.DistinguishedName, relationParam)
		
			if err != nil {
				return err
			}
		}
	}
	
	if  relationTofvRsFcPathAtt, ok := d.GetOk("relation_fv_rs_fc_path_att") ; ok {
		relationParamList := toStringList(relationTofvRsFcPathAtt.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsFcPathAtt(fvAEPg.DistinguishedName, relationParam)
		
			if err != nil {
				return err
			}
		}
	}
	
	if  relationTofvRsProv, ok := d.GetOk("relation_fv_rs_prov") ; ok {
		relationParamList := toStringList(relationTofvRsProv.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsProv(fvAEPg.DistinguishedName, relationParam)
		
			if err != nil {
				return err
			}
		}
	}
	
	if  relationTofvRsConsIf, ok := d.GetOk("relation_fv_rs_cons_if") ; ok {
		relationParamList := toStringList(relationTofvRsConsIf.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsConsIf(fvAEPg.DistinguishedName, relationParam)
		
			if err != nil {
				return err
			}
		}
	}
	
	if  relationTofvRsSecInherited, ok := d.GetOk("relation_fv_rs_sec_inherited") ; ok {
		relationParamList := toStringList(relationTofvRsSecInherited.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsSecInherited(fvAEPg.DistinguishedName, relationParam)
		
			if err != nil {
				return err
			}
		}
	}
	
	if  relationTofvRsNodeAtt, ok := d.GetOk("relation_fv_rs_node_att") ; ok {
		relationParamList := toStringList(relationTofvRsNodeAtt.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsNodeAtt(fvAEPg.DistinguishedName, relationParam)
		
			if err != nil {
				return err
			}
		}
	}
	
	if  relationTofvRsDppPol, ok := d.GetOk("relation_fv_rs_dpp_pol") ; ok {
		relationParam := relationTofvRsDppPol.(string)
		err = aciClient.CreateRelationfvRsDppPol(fvAEPg.DistinguishedName, relationParam)
		if err != nil {
			return err
		}
		
	}
	
	if  relationTofvRsCons, ok := d.GetOk("relation_fv_rs_cons") ; ok {
		relationParamList := toStringList(relationTofvRsCons.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsCons(fvAEPg.DistinguishedName, relationParam)
		
			if err != nil {
				return err
			}
		}
	}
	
	if  relationTofvRsTrustCtrl, ok := d.GetOk("relation_fv_rs_trust_ctrl") ; ok {
		relationParam := relationTofvRsTrustCtrl.(string)
		err = aciClient.CreateRelationfvRsTrustCtrl(fvAEPg.DistinguishedName, relationParam)
		if err != nil {
			return err
		}
		
	}
	
	if  relationTofvRsPathAtt, ok := d.GetOk("relation_fv_rs_path_att") ; ok {
		relationParamList := toStringList(relationTofvRsPathAtt.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsPathAtt(fvAEPg.DistinguishedName, relationParam)
		
			if err != nil {
				return err
			}
		}
	}
	
	if  relationTofvRsProtBy, ok := d.GetOk("relation_fv_rs_prot_by") ; ok {
		relationParamList := toStringList(relationTofvRsProtBy.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsProtBy(fvAEPg.DistinguishedName, relationParam)
		
			if err != nil {
				return err
			}
		}
	}
	
	if  relationTofvRsAEPgMonPol, ok := d.GetOk("relation_fv_rs_ae_pg_mon_pol") ; ok {
		relationParam := relationTofvRsAEPgMonPol.(string)
		err = aciClient.CreateRelationfvRsAEPgMonPol(fvAEPg.DistinguishedName, relationParam)
		if err != nil {
			return err
		}
		
	}
	
	if  relationTofvRsIntraEpg, ok := d.GetOk("relation_fv_rs_intra_epg") ; ok {
		relationParamList := toStringList(relationTofvRsIntraEpg.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsIntraEpg(fvAEPg.DistinguishedName, relationParam)
		
			if err != nil {
				return err
			}
		}
	}
	

	d.SetId(fvAEPg.DistinguishedName)
	return resourceAciApplicationEPGRead(d, m)
}

func resourceAciApplicationEPGUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	
	name := d.Get("name").(string)
	ApplicationProfileDn := d.Get("application_profile_dn").(string)
	

    fvAEPgAttr := models.ApplicationEPGAttributes{}     
    if FloodOnEncap, ok := d.GetOk("flood_on_encap"); ok {
        fvAEPgAttr.FloodOnEncap = FloodOnEncap.(string)
    }     
    if FwdCtrl, ok := d.GetOk("fwd_ctrl"); ok {
        fvAEPgAttr.FwdCtrl = FwdCtrl.(string)
    }     
    if IsAttrBasedEPg, ok := d.GetOk("is_attr_based_e_pg"); ok {
        fvAEPgAttr.IsAttrBasedEPg = IsAttrBasedEPg.(string)
    }     
    if MatchT, ok := d.GetOk("match_t"); ok {
        fvAEPgAttr.MatchT = MatchT.(string)
    }     
    if NameAlias, ok := d.GetOk("name_alias"); ok {
        fvAEPgAttr.NameAlias = NameAlias.(string)
    }     
    if PcEnfPref, ok := d.GetOk("pc_enf_pref"); ok {
        fvAEPgAttr.PcEnfPref = PcEnfPref.(string)
    }     
    if PrefGrMemb, ok := d.GetOk("pref_gr_memb"); ok {
        fvAEPgAttr.PrefGrMemb = PrefGrMemb.(string)
    }     
    if Prio, ok := d.GetOk("prio"); ok {
        fvAEPgAttr.Prio = Prio.(string)
    }
	fvAEPg := models.NewApplicationEPG(fmt.Sprintf("epg-%s",name),ApplicationProfileDn, desc, fvAEPgAttr)  
		

	fvAEPg.Status = "modified"

	err := aciClient.Save(fvAEPg)
	
	if err != nil {
		return err
	}
	if d.HasChange("relation_fv_rs_dom_att") {
		oldc, newc := d.GetChange("relation_fv_rs_dom_att")
		oldRelSet := oldc.(*schema.Set)
		newRelSet := newc.(*schema.Set)

		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsDomAtt(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsDomAtt(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}
	}
	if d.HasChange("relation_fv_rs_fc_path_att") {
		oldc, newc := d.GetChange("relation_fv_rs_fc_path_att")
		oldRelSet := oldc.(*schema.Set)
		newRelSet := newc.(*schema.Set)

		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsFcPathAtt(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsFcPathAtt(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}
	}
	if d.HasChange("relation_fv_rs_prov") {
		oldc, newc := d.GetChange("relation_fv_rs_prov")
		oldRelSet := oldc.(*schema.Set)
		newRelSet := newc.(*schema.Set)

		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsProv(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsProv(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}
	}
	if d.HasChange("relation_fv_rs_cons_if") {
		oldc, newc := d.GetChange("relation_fv_rs_cons_if")
		oldRelSet := oldc.(*schema.Set)
		newRelSet := newc.(*schema.Set)

		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsConsIf(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsConsIf(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}
	}
	if d.HasChange("relation_fv_rs_sec_inherited") {
		oldc, newc := d.GetChange("relation_fv_rs_sec_inherited")
		oldRelSet := oldc.(*schema.Set)
		newRelSet := newc.(*schema.Set)

		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsSecInherited(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsSecInherited(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}
	}
	if d.HasChange("relation_fv_rs_node_att") {
		oldc, newc := d.GetChange("relation_fv_rs_node_att")
		oldRelSet := oldc.(*schema.Set)
		newRelSet := newc.(*schema.Set)

		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsNodeAtt(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsNodeAtt(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}
	}
	if d.HasChange("relation_fv_rs_dpp_pol") {
		oldTdn, newTdn := d.GetChange("relation_fv_rs_dpp_pol")
		err = aciClient.DeleteRelationfvRsDppPol(fvAEPg.DistinguishedName, oldTdn.(string))
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationfvRsDppPol(fvAEPg.DistinguishedName, newTdn.(string))
		if err != nil {
			return err
		}
	
	}
	if d.HasChange("relation_fv_rs_cons") {
		oldc, newc := d.GetChange("relation_fv_rs_cons")
		oldRelSet := oldc.(*schema.Set)
		newRelSet := newc.(*schema.Set)

		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsCons(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsCons(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}
	}
	if d.HasChange("relation_fv_rs_trust_ctrl") {
		oldTdn, newTdn := d.GetChange("relation_fv_rs_trust_ctrl")
		err = aciClient.DeleteRelationfvRsTrustCtrl(fvAEPg.DistinguishedName, oldTdn.(string))
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationfvRsTrustCtrl(fvAEPg.DistinguishedName, newTdn.(string))
		if err != nil {
			return err
		}
	
	}
	if d.HasChange("relation_fv_rs_path_att") {
		oldc, newc := d.GetChange("relation_fv_rs_path_att")
		oldRelSet := oldc.(*schema.Set)
		newRelSet := newc.(*schema.Set)

		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsPathAtt(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsPathAtt(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}
	}
	if d.HasChange("relation_fv_rs_prot_by") {
		oldc, newc := d.GetChange("relation_fv_rs_prot_by")
		oldRelSet := oldc.(*schema.Set)
		newRelSet := newc.(*schema.Set)

		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsProtBy(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsProtBy(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}
	}
	if d.HasChange("relation_fv_rs_ae_pg_mon_pol") {
		oldTdn, newTdn := d.GetChange("relation_fv_rs_ae_pg_mon_pol")
		err = aciClient.DeleteRelationfvRsAEPgMonPol(fvAEPg.DistinguishedName, oldTdn.(string))
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationfvRsAEPgMonPol(fvAEPg.DistinguishedName, newTdn.(string))
		if err != nil {
			return err
		}
	
	}
	if d.HasChange("relation_fv_rs_intra_epg") {
		oldc, newc := d.GetChange("relation_fv_rs_intra_epg")
		oldRelSet := oldc.(*schema.Set)
		newRelSet := newc.(*schema.Set)

		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsIntraEpg(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsIntraEpg(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}
	}
	

	d.SetId(fvAEPg.DistinguishedName)
	return resourceAciApplicationEPGRead(d, m)

}

func resourceAciApplicationEPGRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	fvAEPg, err := getRemoteApplicationEPG(aciClient, dn)

	if err != nil {
		return err
	}
	setApplicationEPGAttributes(fvAEPg, d)
	return nil
}

func resourceAciApplicationEPGDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "fvAEPg")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}