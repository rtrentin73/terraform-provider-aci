package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciApplicationprofile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciApplicationprofileCreate,
		Update: resourceAciApplicationprofileUpdate,
		Read:   resourceAciApplicationprofileRead,
		Delete: resourceAciApplicationprofileDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciApplicationprofileImport,
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

			"prio": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "priority class id",
			},

			"relation_fv_rs_ap_mon_pol": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to monEPGPol",
			},
		}),
	}
}

func getRemoteApplicationprofile(client *client.Client, dn string) (*models.Applicationprofile, error) {
	fvApCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	fvAp := models.ApplicationprofileFromContainer(fvApCont)

	if fvAp.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", fvAp.DistinguishedName)
	}

	return fvAp, nil
}

func setApplicationprofileAttributes(fvAp *models.Applicationprofile, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(fvAp.DistinguishedName)
	d.Set("description", fvAp.Description)
	d.Set("tenant_dn", GetParentDn(fvAp.DistinguishedName))
	fvApMap, _ := fvAp.ToMap()

	d.Set("annotation", fvApMap["annotation"])
	d.Set("name_alias", fvApMap["nameAlias"])
	d.Set("prio", fvApMap["prio"])
	return d
}

func resourceAciApplicationprofileImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	fvAp, err := getRemoteApplicationprofile(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setApplicationprofileAttributes(fvAp, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciApplicationprofileCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	TenantDn := d.Get("tenant_dn").(string)

	fvApAttr := models.ApplicationprofileAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		fvApAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		fvApAttr.NameAlias = NameAlias.(string)
	}
	if Prio, ok := d.GetOk("prio"); ok {
		fvApAttr.Prio = Prio.(string)
	}
	fvAp := models.NewApplicationprofile(fmt.Sprintf("ap-%s", name), TenantDn, desc, fvApAttr)

	err := aciClient.Save(fvAp)
	if err != nil {
		return err
	}

	if relationTofvRsApMonPol, ok := d.GetOk("relation_fv_rs_ap_mon_pol"); ok {
		relationParam := relationTofvRsApMonPol.(string)
		err = aciClient.CreateRelationfvRsApMonPolFromApplicationprofile(fvAp.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}

	d.SetId(fvAp.DistinguishedName)
	return resourceAciApplicationprofileRead(d, m)
}

func resourceAciApplicationprofileUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	TenantDn := d.Get("tenant_dn").(string)

	fvApAttr := models.ApplicationprofileAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		fvApAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		fvApAttr.NameAlias = NameAlias.(string)
	}
	if Prio, ok := d.GetOk("prio"); ok {
		fvApAttr.Prio = Prio.(string)
	}
	fvAp := models.NewApplicationprofile(fmt.Sprintf("ap-%s", name), TenantDn, desc, fvApAttr)

	fvAp.Status = "modified"

	err := aciClient.Save(fvAp)

	if err != nil {
		return err
	}

	if d.HasChange("relation_fv_rs_ap_mon_pol") {
		_, newRelParam := d.GetChange("relation_fv_rs_ap_mon_pol")
		err = aciClient.DeleteRelationfvRsApMonPolFromApplicationprofile(fvAp.DistinguishedName)
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationfvRsApMonPolFromApplicationprofile(fvAp.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}

	d.SetId(fvAp.DistinguishedName)
	return resourceAciApplicationprofileRead(d, m)

}

func resourceAciApplicationprofileRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	fvAp, err := getRemoteApplicationprofile(aciClient, dn)

	if err != nil {
		return err
	}
	setApplicationprofileAttributes(fvAp, d)
	return nil
}

func resourceAciApplicationprofileDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "fvAp")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
