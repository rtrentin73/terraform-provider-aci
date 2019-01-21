package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciCloudepg() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciCloudepgCreate,
		Update: resourceAciCloudepgUpdate,
		Read:   resourceAciCloudepgRead,
		Delete: resourceAciCloudepgDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudepgImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"cloudapplicationcontainer_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"annotation": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"exception_tag": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"flood_on_encap": &schema.Schema{
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

			"relation_fv_rs_sec_inherited": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Create relation to fvEPg",
				Set:         schema.HashString,
			},
			"relation_fv_rs_prov": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Create relation to vzBrCP",
				Set:         schema.HashString,
			},
			"relation_fv_rs_cons_if": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Create relation to vzCPIf",
				Set:         schema.HashString,
			},
			"relation_fv_rs_cust_qos_pol": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to qosCustomPol",
			},
			"relation_fv_rs_cons": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Create relation to vzBrCP",
				Set:         schema.HashString,
			},
			"relation_cloud_rs_cloud_e_pg_ctx": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to fvCtx",
			},
			"relation_fv_rs_prot_by": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Create relation to vzTaboo",
				Set:         schema.HashString,
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

func getRemoteCloudepg(client *client.Client, dn string) (*models.Cloudepg, error) {
	cloudEPgCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudEPg := models.CloudepgFromContainer(cloudEPgCont)

	if cloudEPg.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudEPg.DistinguishedName)
	}

	return cloudEPg, nil
}

func setCloudepgAttributes(cloudEPg *models.Cloudepg, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudEPg.DistinguishedName)
	d.Set("description", cloudEPg.Description)
	d.Set("cloudapplicationcontainer_dn", GetParentDn(cloudEPg.DistinguishedName))
	cloudEPgMap, _ := cloudEPg.ToMap()

	d.Set("annotation", cloudEPgMap["annotation"])
	d.Set("exception_tag", cloudEPgMap["exceptionTag"])
	d.Set("flood_on_encap", cloudEPgMap["floodOnEncap"])
	d.Set("match_t", cloudEPgMap["matchT"])
	d.Set("name_alias", cloudEPgMap["nameAlias"])
	d.Set("pref_gr_memb", cloudEPgMap["prefGrMemb"])
	d.Set("prio", cloudEPgMap["prio"])
	return d
}

func resourceAciCloudepgImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudEPg, err := getRemoteCloudepg(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCloudepgAttributes(cloudEPg, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudepgCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudapplicationcontainerDn := d.Get("cloudapplicationcontainer_dn").(string)

	cloudEPgAttr := models.CloudepgAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudEPgAttr.Annotation = Annotation.(string)
	}
	if ExceptionTag, ok := d.GetOk("exception_tag"); ok {
		cloudEPgAttr.ExceptionTag = ExceptionTag.(string)
	}
	if FloodOnEncap, ok := d.GetOk("flood_on_encap"); ok {
		cloudEPgAttr.FloodOnEncap = FloodOnEncap.(string)
	}
	if MatchT, ok := d.GetOk("match_t"); ok {
		cloudEPgAttr.MatchT = MatchT.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudEPgAttr.NameAlias = NameAlias.(string)
	}
	if PrefGrMemb, ok := d.GetOk("pref_gr_memb"); ok {
		cloudEPgAttr.PrefGrMemb = PrefGrMemb.(string)
	}
	if Prio, ok := d.GetOk("prio"); ok {
		cloudEPgAttr.Prio = Prio.(string)
	}
	cloudEPg := models.NewCloudepg(fmt.Sprintf("cloudepg-%s", name), CloudapplicationcontainerDn, desc, cloudEPgAttr)

	err := aciClient.Save(cloudEPg)
	if err != nil {
		return err
	}

	if relationTofvRsSecInherited, ok := d.GetOk("relation_fv_rs_sec_inherited"); ok {
		relationParamList := toStringList(relationTofvRsSecInherited.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsSecInheritedFromCloudepg(cloudEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsProv, ok := d.GetOk("relation_fv_rs_prov"); ok {
		relationParamList := toStringList(relationTofvRsProv.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsProvFromCloudepg(cloudEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsConsIf, ok := d.GetOk("relation_fv_rs_cons_if"); ok {
		relationParamList := toStringList(relationTofvRsConsIf.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsConsIfFromCloudepg(cloudEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsCustQosPol, ok := d.GetOk("relation_fv_rs_cust_qos_pol"); ok {
		relationParam := relationTofvRsCustQosPol.(string)
		err = aciClient.CreateRelationfvRsCustQosPolFromCloudepg(cloudEPg.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsCons, ok := d.GetOk("relation_fv_rs_cons"); ok {
		relationParamList := toStringList(relationTofvRsCons.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsConsFromCloudepg(cloudEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTocloudRsCloudEPgCtx, ok := d.GetOk("relation_cloud_rs_cloud_e_pg_ctx"); ok {
		relationParam := relationTocloudRsCloudEPgCtx.(string)
		err = aciClient.CreateRelationcloudRsCloudEPgCtxFromCloudepg(cloudEPg.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsProtBy, ok := d.GetOk("relation_fv_rs_prot_by"); ok {
		relationParamList := toStringList(relationTofvRsProtBy.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsProtByFromCloudepg(cloudEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsIntraEpg, ok := d.GetOk("relation_fv_rs_intra_epg"); ok {
		relationParamList := toStringList(relationTofvRsIntraEpg.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsIntraEpgFromCloudepg(cloudEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}

	d.SetId(cloudEPg.DistinguishedName)
	return resourceAciCloudepgRead(d, m)
}

func resourceAciCloudepgUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudapplicationcontainerDn := d.Get("cloudapplicationcontainer_dn").(string)

	cloudEPgAttr := models.CloudepgAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudEPgAttr.Annotation = Annotation.(string)
	}
	if ExceptionTag, ok := d.GetOk("exception_tag"); ok {
		cloudEPgAttr.ExceptionTag = ExceptionTag.(string)
	}
	if FloodOnEncap, ok := d.GetOk("flood_on_encap"); ok {
		cloudEPgAttr.FloodOnEncap = FloodOnEncap.(string)
	}
	if MatchT, ok := d.GetOk("match_t"); ok {
		cloudEPgAttr.MatchT = MatchT.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudEPgAttr.NameAlias = NameAlias.(string)
	}
	if PrefGrMemb, ok := d.GetOk("pref_gr_memb"); ok {
		cloudEPgAttr.PrefGrMemb = PrefGrMemb.(string)
	}
	if Prio, ok := d.GetOk("prio"); ok {
		cloudEPgAttr.Prio = Prio.(string)
	}
	cloudEPg := models.NewCloudepg(fmt.Sprintf("cloudepg-%s", name), CloudapplicationcontainerDn, desc, cloudEPgAttr)

	cloudEPg.Status = "modified"

	err := aciClient.Save(cloudEPg)

	if err != nil {
		return err
	}

	if d.HasChange("relation_fv_rs_sec_inherited") {
		oldRel, newRel := d.GetChange("relation_fv_rs_sec_inherited")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsSecInheritedFromCloudepg(cloudEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsSecInheritedFromCloudepg(cloudEPg.DistinguishedName, relDn)
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
			err = aciClient.DeleteRelationfvRsProvFromCloudepg(cloudEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsProvFromCloudepg(cloudEPg.DistinguishedName, relDn)
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
			err = aciClient.DeleteRelationfvRsConsIfFromCloudepg(cloudEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsConsIfFromCloudepg(cloudEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_fv_rs_cust_qos_pol") {
		_, newRelParam := d.GetChange("relation_fv_rs_cust_qos_pol")
		err = aciClient.CreateRelationfvRsCustQosPolFromCloudepg(cloudEPg.DistinguishedName, newRelParam.(string))
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
			err = aciClient.DeleteRelationfvRsConsFromCloudepg(cloudEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsConsFromCloudepg(cloudEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_cloud_rs_cloud_e_pg_ctx") {
		_, newRelParam := d.GetChange("relation_cloud_rs_cloud_e_pg_ctx")
		err = aciClient.CreateRelationcloudRsCloudEPgCtxFromCloudepg(cloudEPg.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_fv_rs_prot_by") {
		oldRel, newRel := d.GetChange("relation_fv_rs_prot_by")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsProtByFromCloudepg(cloudEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsProtByFromCloudepg(cloudEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_fv_rs_intra_epg") {
		oldRel, newRel := d.GetChange("relation_fv_rs_intra_epg")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsIntraEpgFromCloudepg(cloudEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsIntraEpgFromCloudepg(cloudEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}

	d.SetId(cloudEPg.DistinguishedName)
	return resourceAciCloudepgRead(d, m)

}

func resourceAciCloudepgRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudEPg, err := getRemoteCloudepg(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudepgAttributes(cloudEPg, d)
	return nil
}

func resourceAciCloudepgDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudEPg")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
