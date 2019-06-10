package aci

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciFilterCreate,
		Update: resourceAciFilterUpdate,
		Read:   resourceAciFilterRead,
		Delete: resourceAciFilterDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciFilterImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"tenant_dn": &schema.Schema{
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

			"name_alias": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"relation_vz_rs_filt_graph_att": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to vnsInTerm",
			},
			"relation_vz_rs_fwd_r_flt_p_att": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to vzAFilterableUnit",
			},
			"relation_vz_rs_rev_r_flt_p_att": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to vzAFilterableUnit",
			},
		}),
	}
}

func getRemoteFilter(client *client.Client, dn string) (*models.Filter, error) {
	vzFilterCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vzFilter := models.FilterFromContainer(vzFilterCont)

	if vzFilter.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", vzFilter.DistinguishedName)
	}

	return vzFilter, nil
}

func setFilterAttributes(vzFilter *models.Filter, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vzFilter.DistinguishedName)
	d.Set("description", vzFilter.Description)
	d.Set("tenant_dn", GetParentDn(vzFilter.DistinguishedName))
	vzFilterMap, _ := vzFilter.ToMap()
	d.Set("name", GetMOName(vzFilter.DistinguishedName))

	d.Set("annotation", vzFilterMap["annotation"])
	d.Set("name_alias", vzFilterMap["nameAlias"])
	return d
}

func resourceAciFilterImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	vzFilter, err := getRemoteFilter(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setFilterAttributes(vzFilter, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciFilterCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	TenantDn := d.Get("tenant_dn").(string)

	vzFilterAttr := models.FilterAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		vzFilterAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		vzFilterAttr.NameAlias = NameAlias.(string)
	}
	vzFilter := models.NewFilter(fmt.Sprintf("flt-%s", name), TenantDn, desc, vzFilterAttr)

	err := aciClient.Save(vzFilter)
	if err != nil {
		return err
	}

	if relationTovzRsFiltGraphAtt, ok := d.GetOk("relation_vz_rs_filt_graph_att"); ok {
		relationParam := relationTovzRsFiltGraphAtt.(string)
		err = aciClient.CreateRelationvzRsFiltGraphAttFromFilter(vzFilter.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTovzRsFwdRFltPAtt, ok := d.GetOk("relation_vz_rs_fwd_r_flt_p_att"); ok {
		relationParam := relationTovzRsFwdRFltPAtt.(string)
		err = aciClient.CreateRelationvzRsFwdRFltPAttFromFilter(vzFilter.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTovzRsRevRFltPAtt, ok := d.GetOk("relation_vz_rs_rev_r_flt_p_att"); ok {
		relationParam := relationTovzRsRevRFltPAtt.(string)
		err = aciClient.CreateRelationvzRsRevRFltPAttFromFilter(vzFilter.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}

	d.SetId(vzFilter.DistinguishedName)
	return resourceAciFilterRead(d, m)
}

func resourceAciFilterUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	TenantDn := d.Get("tenant_dn").(string)

	vzFilterAttr := models.FilterAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		vzFilterAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		vzFilterAttr.NameAlias = NameAlias.(string)
	}
	vzFilter := models.NewFilter(fmt.Sprintf("flt-%s", name), TenantDn, desc, vzFilterAttr)

	vzFilter.Status = "modified"

	err := aciClient.Save(vzFilter)

	if err != nil {
		return err
	}

	if d.HasChange("relation_vz_rs_filt_graph_att") {
		_, newRelParam := d.GetChange("relation_vz_rs_filt_graph_att")
		err = aciClient.CreateRelationvzRsFiltGraphAttFromFilter(vzFilter.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_vz_rs_fwd_r_flt_p_att") {
		_, newRelParam := d.GetChange("relation_vz_rs_fwd_r_flt_p_att")
		err = aciClient.CreateRelationvzRsFwdRFltPAttFromFilter(vzFilter.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_vz_rs_rev_r_flt_p_att") {
		_, newRelParam := d.GetChange("relation_vz_rs_rev_r_flt_p_att")
		err = aciClient.CreateRelationvzRsRevRFltPAttFromFilter(vzFilter.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}

	d.SetId(vzFilter.DistinguishedName)
	return resourceAciFilterRead(d, m)

}

func resourceAciFilterRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	vzFilter, err := getRemoteFilter(aciClient, dn)

	if err != nil {
		return err
	}
	setFilterAttributes(vzFilter, d)
	return nil
}

func resourceAciFilterDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "vzFilter")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
