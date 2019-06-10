package aci

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciCloudExternalEPg() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciCloudExternalEPgCreate,
		Update: resourceAciCloudExternalEPgUpdate,
		Read:   resourceAciCloudExternalEPgRead,
		Delete: resourceAciCloudExternalEPgDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudExternalEPgImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"cloud_applicationcontainer_dn": &schema.Schema{
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

func getRemoteCloudExternalEPg(client *client.Client, dn string) (*models.CloudExternalEPg, error) {
	cloudExtEPgCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudExtEPg := models.CloudExternalEPgFromContainer(cloudExtEPgCont)

	if cloudExtEPg.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudExtEPg.DistinguishedName)
	}

	return cloudExtEPg, nil
}

func setCloudExternalEPgAttributes(cloudExtEPg *models.CloudExternalEPg, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudExtEPg.DistinguishedName)
	d.Set("description", cloudExtEPg.Description)
	d.Set("cloud_applicationcontainer_dn", GetParentDn(cloudExtEPg.DistinguishedName))
	cloudExtEPgMap, _ := cloudExtEPg.ToMap()
	d.Set("name", GetMOName(cloudExtEPg.DistinguishedName))

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

func resourceAciCloudExternalEPgImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudExtEPg, err := getRemoteCloudExternalEPg(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCloudExternalEPgAttributes(cloudExtEPg, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudExternalEPgCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudApplicationcontainerDn := d.Get("cloud_applicationcontainer_dn").(string)

	cloudExtEPgAttr := models.CloudExternalEPgAttributes{}
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
	cloudExtEPg := models.NewCloudExternalEPg(fmt.Sprintf("cloudextepg-%s", name), CloudApplicationcontainerDn, desc, cloudExtEPgAttr)

	err := aciClient.Save(cloudExtEPg)
	if err != nil {
		return err
	}

	if relationTofvRsSecInherited, ok := d.GetOk("relation_fv_rs_sec_inherited"); ok {
		relationParamList := toStringList(relationTofvRsSecInherited.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsSecInheritedFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsProv, ok := d.GetOk("relation_fv_rs_prov"); ok {
		relationParamList := toStringList(relationTofvRsProv.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsProvFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsConsIf, ok := d.GetOk("relation_fv_rs_cons_if"); ok {
		relationParamList := toStringList(relationTofvRsConsIf.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsConsIfFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsCustQosPol, ok := d.GetOk("relation_fv_rs_cust_qos_pol"); ok {
		relationParam := relationTofvRsCustQosPol.(string)
		err = aciClient.CreateRelationfvRsCustQosPolFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsCons, ok := d.GetOk("relation_fv_rs_cons"); ok {
		relationParamList := toStringList(relationTofvRsCons.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsConsFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTocloudRsCloudEPgCtx, ok := d.GetOk("relation_cloud_rs_cloud_e_pg_ctx"); ok {
		relationParam := relationTocloudRsCloudEPgCtx.(string)
		err = aciClient.CreateRelationcloudRsCloudEPgCtxFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsProtBy, ok := d.GetOk("relation_fv_rs_prot_by"); ok {
		relationParamList := toStringList(relationTofvRsProtBy.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsProtByFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsIntraEpg, ok := d.GetOk("relation_fv_rs_intra_epg"); ok {
		relationParamList := toStringList(relationTofvRsIntraEpg.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsIntraEpgFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}

	d.SetId(cloudExtEPg.DistinguishedName)
	return resourceAciCloudExternalEPgRead(d, m)
}

func resourceAciCloudExternalEPgUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudApplicationcontainerDn := d.Get("cloud_applicationcontainer_dn").(string)

	cloudExtEPgAttr := models.CloudExternalEPgAttributes{}
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
	cloudExtEPg := models.NewCloudExternalEPg(fmt.Sprintf("cloudextepg-%s", name), CloudApplicationcontainerDn, desc, cloudExtEPgAttr)

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
			err = aciClient.DeleteRelationfvRsSecInheritedFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsSecInheritedFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relDn)
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
			err = aciClient.DeleteRelationfvRsProvFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsProvFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relDn)
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
			err = aciClient.DeleteRelationfvRsConsIfFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsConsIfFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_fv_rs_cust_qos_pol") {
		_, newRelParam := d.GetChange("relation_fv_rs_cust_qos_pol")
		err = aciClient.CreateRelationfvRsCustQosPolFromCloudExternalEPg(cloudExtEPg.DistinguishedName, newRelParam.(string))
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
			err = aciClient.DeleteRelationfvRsConsFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsConsFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_cloud_rs_cloud_e_pg_ctx") {
		_, newRelParam := d.GetChange("relation_cloud_rs_cloud_e_pg_ctx")
		err = aciClient.CreateRelationcloudRsCloudEPgCtxFromCloudExternalEPg(cloudExtEPg.DistinguishedName, newRelParam.(string))
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
			err = aciClient.DeleteRelationfvRsProtByFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsProtByFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relDn)
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
			err = aciClient.DeleteRelationfvRsIntraEpgFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsIntraEpgFromCloudExternalEPg(cloudExtEPg.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}

	d.SetId(cloudExtEPg.DistinguishedName)
	return resourceAciCloudExternalEPgRead(d, m)

}

func resourceAciCloudExternalEPgRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudExtEPg, err := getRemoteCloudExternalEPg(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudExternalEPgAttributes(cloudExtEPg, d)
	return nil
}

func resourceAciCloudExternalEPgDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudExtEPg")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
