package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciCloudexternalepg() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciCloudexternalepgCreate,
		Update: resourceAciCloudexternalepgUpdate,
		Read:   resourceAciCloudexternalepgRead,
		Delete: resourceAciCloudexternalepgDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudexternalepgImport,
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

			"route_reachability": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
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

func getRemoteCloudexternalepg(client *client.Client, dn string) (*models.Cloudexternalepg, error) {
	cloudExtEPgCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudExtEPg := models.CloudexternalepgFromContainer(cloudExtEPgCont)

	if cloudExtEPg.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudExtEPg.DistinguishedName)
	}

	return cloudExtEPg, nil
}

func setCloudexternalepgAttributes(cloudExtEPg *models.Cloudexternalepg, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudExtEPg.DistinguishedName)
	d.Set("description", cloudExtEPg.Description)
	d.Set("cloudapplicationcontainer_dn", GetParentDn(cloudExtEPg.DistinguishedName))
	cloudExtEPgMap, _ := cloudExtEPg.ToMap()

	d.Set("annotation", cloudExtEPgMap["annotation"])
	d.Set("exception_tag", cloudExtEPgMap["exceptionTag"])
	d.Set("flood_on_encap", cloudExtEPgMap["floodOnEncap"])
	d.Set("match_t", cloudExtEPgMap["matchT"])
	d.Set("name_alias", cloudExtEPgMap["nameAlias"])
	d.Set("pref_gr_memb", cloudExtEPgMap["prefGrMemb"])
	d.Set("prio", cloudExtEPgMap["prio"])
	d.Set("route_reachability", cloudExtEPgMap["routeReachability"])
	return d
}

func resourceAciCloudexternalepgImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudExtEPg, err := getRemoteCloudexternalepg(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCloudexternalepgAttributes(cloudExtEPg, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudexternalepgCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudapplicationcontainerDn := d.Get("cloudapplicationcontainer_dn").(string)

	cloudExtEPgAttr := models.CloudexternalepgAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudExtEPgAttr.Annotation = Annotation.(string)
	}
	if ExceptionTag, ok := d.GetOk("exception_tag"); ok {
		cloudExtEPgAttr.ExceptionTag = ExceptionTag.(string)
	}
	if FloodOnEncap, ok := d.GetOk("flood_on_encap"); ok {
		cloudExtEPgAttr.FloodOnEncap = FloodOnEncap.(string)
	}
	if MatchT, ok := d.GetOk("match_t"); ok {
		cloudExtEPgAttr.MatchT = MatchT.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudExtEPgAttr.NameAlias = NameAlias.(string)
	}
	if PrefGrMemb, ok := d.GetOk("pref_gr_memb"); ok {
		cloudExtEPgAttr.PrefGrMemb = PrefGrMemb.(string)
	}
	if Prio, ok := d.GetOk("prio"); ok {
		cloudExtEPgAttr.Prio = Prio.(string)
	}
	if RouteReachability, ok := d.GetOk("route_reachability"); ok {
		cloudExtEPgAttr.RouteReachability = RouteReachability.(string)
	}
	cloudExtEPg := models.NewCloudexternalepg(fmt.Sprintf("cloudextepg-%s", name), CloudapplicationcontainerDn, desc, cloudExtEPgAttr)

	err := aciClient.Save(cloudExtEPg)
	if err != nil {
		return err
	}

	if relationTofvRsSecInherited, ok := d.GetOk("relation_fv_rs_sec_inherited"); ok {
		relationParamList := toStringList(relationTofvRsSecInherited.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsSecInheritedFromCloudexternalepg(cloudExtEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsProv, ok := d.GetOk("relation_fv_rs_prov"); ok {
		relationParamList := toStringList(relationTofvRsProv.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsProvFromCloudexternalepg(cloudExtEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsConsIf, ok := d.GetOk("relation_fv_rs_cons_if"); ok {
		relationParamList := toStringList(relationTofvRsConsIf.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsConsIfFromCloudexternalepg(cloudExtEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsCustQosPol, ok := d.GetOk("relation_fv_rs_cust_qos_pol"); ok {
		relationParam := relationTofvRsCustQosPol.(string)
		err = aciClient.CreateRelationfvRsCustQosPolFromCloudexternalepg(cloudExtEPg.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsCons, ok := d.GetOk("relation_fv_rs_cons"); ok {
		relationParamList := toStringList(relationTofvRsCons.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsConsFromCloudexternalepg(cloudExtEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTocloudRsCloudEPgCtx, ok := d.GetOk("relation_cloud_rs_cloud_e_pg_ctx"); ok {
		relationParam := relationTocloudRsCloudEPgCtx.(string)
		err = aciClient.CreateRelationcloudRsCloudEPgCtxFromCloudexternalepg(cloudExtEPg.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsProtBy, ok := d.GetOk("relation_fv_rs_prot_by"); ok {
		relationParamList := toStringList(relationTofvRsProtBy.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsProtByFromCloudexternalepg(cloudExtEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsIntraEpg, ok := d.GetOk("relation_fv_rs_intra_epg"); ok {
		relationParamList := toStringList(relationTofvRsIntraEpg.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsIntraEpgFromCloudexternalepg(cloudExtEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}

	d.SetId(cloudExtEPg.DistinguishedName)
	return resourceAciCloudexternalepgRead(d, m)
}

func resourceAciCloudexternalepgUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudapplicationcontainerDn := d.Get("cloudapplicationcontainer_dn").(string)

	cloudExtEPgAttr := models.CloudexternalepgAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudExtEPgAttr.Annotation = Annotation.(string)
	}
	if ExceptionTag, ok := d.GetOk("exception_tag"); ok {
		cloudExtEPgAttr.ExceptionTag = ExceptionTag.(string)
	}
	if FloodOnEncap, ok := d.GetOk("flood_on_encap"); ok {
		cloudExtEPgAttr.FloodOnEncap = FloodOnEncap.(string)
	}
	if MatchT, ok := d.GetOk("match_t"); ok {
		cloudExtEPgAttr.MatchT = MatchT.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudExtEPgAttr.NameAlias = NameAlias.(string)
	}
	if PrefGrMemb, ok := d.GetOk("pref_gr_memb"); ok {
		cloudExtEPgAttr.PrefGrMemb = PrefGrMemb.(string)
	}
	if Prio, ok := d.GetOk("prio"); ok {
		cloudExtEPgAttr.Prio = Prio.(string)
	}
	if RouteReachability, ok := d.GetOk("route_reachability"); ok {
		cloudExtEPgAttr.RouteReachability = RouteReachability.(string)
	}
	cloudExtEPg := models.NewCloudexternalepg(fmt.Sprintf("cloudextepg-%s", name), CloudapplicationcontainerDn, desc, cloudExtEPgAttr)

	cloudExtEPg.Status = "modified"

	err := aciClient.Save(cloudExtEPg)

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
			err = aciClient.DeleteRelationfvRsSecInheritedFromCloudexternalepg(cloudExtEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsSecInheritedFromCloudexternalepg(cloudExtEPg.DistinguishedName, relDn)
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
			err = aciClient.DeleteRelationfvRsProvFromCloudexternalepg(cloudExtEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsProvFromCloudexternalepg(cloudExtEPg.DistinguishedName, relDn)
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
			err = aciClient.DeleteRelationfvRsConsIfFromCloudexternalepg(cloudExtEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsConsIfFromCloudexternalepg(cloudExtEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_fv_rs_cust_qos_pol") {
		_, newRelParam := d.GetChange("relation_fv_rs_cust_qos_pol")
		err = aciClient.CreateRelationfvRsCustQosPolFromCloudexternalepg(cloudExtEPg.DistinguishedName, newRelParam.(string))
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
			err = aciClient.DeleteRelationfvRsConsFromCloudexternalepg(cloudExtEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsConsFromCloudexternalepg(cloudExtEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_cloud_rs_cloud_e_pg_ctx") {
		_, newRelParam := d.GetChange("relation_cloud_rs_cloud_e_pg_ctx")
		err = aciClient.CreateRelationcloudRsCloudEPgCtxFromCloudexternalepg(cloudExtEPg.DistinguishedName, newRelParam.(string))
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
			err = aciClient.DeleteRelationfvRsProtByFromCloudexternalepg(cloudExtEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsProtByFromCloudexternalepg(cloudExtEPg.DistinguishedName, relDn)
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
			err = aciClient.DeleteRelationfvRsIntraEpgFromCloudexternalepg(cloudExtEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsIntraEpgFromCloudexternalepg(cloudExtEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}

	d.SetId(cloudExtEPg.DistinguishedName)
	return resourceAciCloudexternalepgRead(d, m)

}

func resourceAciCloudexternalepgRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudExtEPg, err := getRemoteCloudexternalepg(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudexternalepgAttributes(cloudExtEPg, d)
	return nil
}

func resourceAciCloudexternalepgDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudExtEPg")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
