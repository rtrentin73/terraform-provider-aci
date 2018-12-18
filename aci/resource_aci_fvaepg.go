package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
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
				Type:     schema.TypeString,
				Required: true,
			},

			"flood_on_encap": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"fwd_ctrl": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"is_attr_based_e_pg": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"match_t": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "match criteria",
			},

			"name_alias": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"pc_enf_pref": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "enforcement preference",
			},

			"pref_gr_memb": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"prio": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "qos priority class id",
			},

			"relation_fv_rs_bd": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to fvBD",
			},
			"relation_fv_rs_cust_qos_pol": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to qosCustomPol",
			},
			"relation_fv_rs_dom_att": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Create relation to infraDomP",
				Set:         schema.HashString,
			},
			"relation_fv_rs_fc_path_att": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Create relation to fabricPathEp",
				Set:         schema.HashString,
			},
			"relation_fv_rs_prov": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Create relation to vzBrCP",
				Set:         schema.HashString,
			},
			"relation_fv_rs_graph_def": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Create relation to vzGraphCont",
				Set:         schema.HashString,
			},
			"relation_fv_rs_cons_if": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Create relation to vzCPIf",
				Set:         schema.HashString,
			},
			"relation_fv_rs_sec_inherited": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Create relation to fvEPg",
				Set:         schema.HashString,
			},
			"relation_fv_rs_node_att": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Create relation to fabricNode",
				Set:         schema.HashString,
			},
			"relation_fv_rs_dpp_pol": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to qosDppPol",
			},
			"relation_fv_rs_cons": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Create relation to vzBrCP",
				Set:         schema.HashString,
			},
			"relation_fv_rs_prov_def": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Create relation to vzCtrctEPgCont",
				Set:         schema.HashString,
			},
			"relation_fv_rs_trust_ctrl": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to fhsTrustCtrlPol",
			},
			"relation_fv_rs_path_att": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Create relation to fabricPathEp",
				Set:         schema.HashString,
			},
			"relation_fv_rs_prot_by": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Create relation to vzTaboo",
				Set:         schema.HashString,
			},
			"relation_fv_rs_ae_pg_mon_pol": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to monEPGPol",
			},
			"relation_fv_rs_intra_epg": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
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
	fvAEPgMap, _ := fvAEPg.ToMap()

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
	fvAEPg := models.NewApplicationEPG(fmt.Sprintf("epg-%s", name), ApplicationProfileDn, desc, fvAEPgAttr)

	err := aciClient.Save(fvAEPg)
	if err != nil {
		return err
	}

	if relationTofvRsBd, ok := d.GetOk("relation_fv_rs_bd"); ok {
		relationParam := relationTofvRsBd.(string)
		err = aciClient.CreateRelationfvRsBdFromApplicationEPG(fvAEPg.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsCustQosPol, ok := d.GetOk("relation_fv_rs_cust_qos_pol"); ok {
		relationParam := relationTofvRsCustQosPol.(string)
		err = aciClient.CreateRelationfvRsCustQosPolFromApplicationEPG(fvAEPg.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsDomAtt, ok := d.GetOk("relation_fv_rs_dom_att"); ok {
		relationParamList := toStringList(relationTofvRsDomAtt.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsDomAttFromApplicationEPG(fvAEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsFcPathAtt, ok := d.GetOk("relation_fv_rs_fc_path_att"); ok {
		relationParamList := toStringList(relationTofvRsFcPathAtt.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsFcPathAttFromApplicationEPG(fvAEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsProv, ok := d.GetOk("relation_fv_rs_prov"); ok {
		relationParamList := toStringList(relationTofvRsProv.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsProvFromApplicationEPG(fvAEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsGraphDef, ok := d.GetOk("relation_fv_rs_graph_def"); ok {
		relationParamList := toStringList(relationTofvRsGraphDef.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsGraphDefFromApplicationEPG(fvAEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsConsIf, ok := d.GetOk("relation_fv_rs_cons_if"); ok {
		relationParamList := toStringList(relationTofvRsConsIf.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsConsIfFromApplicationEPG(fvAEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsSecInherited, ok := d.GetOk("relation_fv_rs_sec_inherited"); ok {
		relationParamList := toStringList(relationTofvRsSecInherited.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsSecInheritedFromApplicationEPG(fvAEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsNodeAtt, ok := d.GetOk("relation_fv_rs_node_att"); ok {
		relationParamList := toStringList(relationTofvRsNodeAtt.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsNodeAttFromApplicationEPG(fvAEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsDppPol, ok := d.GetOk("relation_fv_rs_dpp_pol"); ok {
		relationParam := relationTofvRsDppPol.(string)
		err = aciClient.CreateRelationfvRsDppPolFromApplicationEPG(fvAEPg.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsCons, ok := d.GetOk("relation_fv_rs_cons"); ok {
		relationParamList := toStringList(relationTofvRsCons.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsConsFromApplicationEPG(fvAEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsProvDef, ok := d.GetOk("relation_fv_rs_prov_def"); ok {
		relationParamList := toStringList(relationTofvRsProvDef.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsProvDefFromApplicationEPG(fvAEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsTrustCtrl, ok := d.GetOk("relation_fv_rs_trust_ctrl"); ok {
		relationParam := relationTofvRsTrustCtrl.(string)
		err = aciClient.CreateRelationfvRsTrustCtrlFromApplicationEPG(fvAEPg.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsPathAtt, ok := d.GetOk("relation_fv_rs_path_att"); ok {
		relationParamList := toStringList(relationTofvRsPathAtt.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsPathAttFromApplicationEPG(fvAEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsProtBy, ok := d.GetOk("relation_fv_rs_prot_by"); ok {
		relationParamList := toStringList(relationTofvRsProtBy.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsProtByFromApplicationEPG(fvAEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsAEPgMonPol, ok := d.GetOk("relation_fv_rs_ae_pg_mon_pol"); ok {
		relationParam := relationTofvRsAEPgMonPol.(string)
		err = aciClient.CreateRelationfvRsAEPgMonPolFromApplicationEPG(fvAEPg.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsIntraEpg, ok := d.GetOk("relation_fv_rs_intra_epg"); ok {
		relationParamList := toStringList(relationTofvRsIntraEpg.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsIntraEpgFromApplicationEPG(fvAEPg.DistinguishedName, relationParam)

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
	fvAEPg := models.NewApplicationEPG(fmt.Sprintf("epg-%s", name), ApplicationProfileDn, desc, fvAEPgAttr)

	fvAEPg.Status = "modified"

	err := aciClient.Save(fvAEPg)

	if err != nil {
		return err
	}

	if d.HasChange("relation_fv_rs_bd") {
		_, newRelParam := d.GetChange("relation_fv_rs_bd")
		err = aciClient.CreateRelationfvRsBdFromApplicationEPG(fvAEPg.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_fv_rs_cust_qos_pol") {
		_, newRelParam := d.GetChange("relation_fv_rs_cust_qos_pol")
		err = aciClient.CreateRelationfvRsCustQosPolFromApplicationEPG(fvAEPg.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_fv_rs_dom_att") {
		oldRel, newRel := d.GetChange("relation_fv_rs_dom_att")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsDomAttFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsDomAttFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_fv_rs_fc_path_att") {
		oldRel, newRel := d.GetChange("relation_fv_rs_fc_path_att")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsFcPathAttFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsFcPathAttFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_fv_rs_prov") {
		oldRel, newRel := d.GetChange("relation_fv_rs_prov")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsProvFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsProvFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_fv_rs_graph_def") {
		oldRel, newRel := d.GetChange("relation_fv_rs_graph_def")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsGraphDefFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_fv_rs_cons_if") {
		oldRel, newRel := d.GetChange("relation_fv_rs_cons_if")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsConsIfFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsConsIfFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_fv_rs_sec_inherited") {
		oldRel, newRel := d.GetChange("relation_fv_rs_sec_inherited")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsSecInheritedFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsSecInheritedFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_fv_rs_node_att") {
		oldRel, newRel := d.GetChange("relation_fv_rs_node_att")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsNodeAttFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsNodeAttFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_fv_rs_dpp_pol") {
		_, newRelParam := d.GetChange("relation_fv_rs_dpp_pol")
		err = aciClient.DeleteRelationfvRsDppPolFromApplicationEPG(fvAEPg.DistinguishedName)
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationfvRsDppPolFromApplicationEPG(fvAEPg.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_fv_rs_cons") {
		oldRel, newRel := d.GetChange("relation_fv_rs_cons")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsConsFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsConsFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_fv_rs_prov_def") {
		oldRel, newRel := d.GetChange("relation_fv_rs_prov_def")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsProvDefFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_fv_rs_trust_ctrl") {
		_, newRelParam := d.GetChange("relation_fv_rs_trust_ctrl")
		err = aciClient.DeleteRelationfvRsTrustCtrlFromApplicationEPG(fvAEPg.DistinguishedName)
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationfvRsTrustCtrlFromApplicationEPG(fvAEPg.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_fv_rs_path_att") {
		oldRel, newRel := d.GetChange("relation_fv_rs_path_att")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsPathAttFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsPathAttFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_fv_rs_prot_by") {
		oldRel, newRel := d.GetChange("relation_fv_rs_prot_by")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsProtByFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsProtByFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_fv_rs_ae_pg_mon_pol") {
		_, newRelParam := d.GetChange("relation_fv_rs_ae_pg_mon_pol")
		err = aciClient.DeleteRelationfvRsAEPgMonPolFromApplicationEPG(fvAEPg.DistinguishedName)
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationfvRsAEPgMonPolFromApplicationEPG(fvAEPg.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_fv_rs_intra_epg") {
		oldRel, newRel := d.GetChange("relation_fv_rs_intra_epg")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsIntraEpgFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsIntraEpgFromApplicationEPG(fvAEPg.DistinguishedName, relDn)
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
