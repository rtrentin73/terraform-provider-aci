package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceAciApplicationProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciApplicationProfileCreate,
		Update: resourceAciApplicationProfileUpdate,
		Read:   resourceAciApplicationProfileRead,
		Delete: resourceAciApplicationProfileDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciApplicationProfileImport,
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

				ValidateFunc: validation.StringInSlice([]string{
					"level1",
					"level2",
					"level3",
					"unspecified",
				}, false),
			},

			"relation_fv_rs_ap_mon_pol": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to monEPGPol",
			},
		}),
	}
}

func getRemoteApplicationProfile(client *client.Client, dn string) (*models.ApplicationProfile, error) {
	fvApCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	fvAp := models.ApplicationProfileFromContainer(fvApCont)

	if fvAp.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", fvAp.DistinguishedName)
	}

	return fvAp, nil
}

func setApplicationProfileAttributes(fvAp *models.ApplicationProfile, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(fvAp.DistinguishedName)
	d.Set("description", fvAp.Description)
	d.Set("tenant_dn", GetParentDn(fvAp.DistinguishedName))
	fvApMap, _ := fvAp.ToMap()

	d.Set("name_alias", fvApMap["nameAlias"])
	d.Set("prio", fvApMap["prio"])
	return d
}

func resourceAciApplicationProfileImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	fvAp, err := getRemoteApplicationProfile(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setApplicationProfileAttributes(fvAp, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciApplicationProfileCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	name := d.Get("name").(string)
	TenantDn := d.Get("tenant_dn").(string)

	fvApAttr := models.ApplicationProfileAttributes{}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		fvApAttr.NameAlias = NameAlias.(string)
	}
	if Prio, ok := d.GetOk("prio"); ok {
		fvApAttr.Prio = Prio.(string)
	}
	fvAp := models.NewApplicationProfile(fmt.Sprintf("ap-%s", name), TenantDn, desc, fvApAttr)

	err := aciClient.Save(fvAp)
	if err != nil {
		return err
	}

	if relationTofvRsApMonPol, ok := d.GetOk("relation_fv_rs_ap_mon_pol"); ok {
		relationParam := relationTofvRsApMonPol.(string)
		err = aciClient.CreateRelationfvRsApMonPol(fvAp.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}

	d.SetId(fvAp.DistinguishedName)
	return resourceAciApplicationProfileRead(d, m)
}

func resourceAciApplicationProfileUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)
	TenantDn := d.Get("tenant_dn").(string)

	fvApAttr := models.ApplicationProfileAttributes{}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		fvApAttr.NameAlias = NameAlias.(string)
	}
	if Prio, ok := d.GetOk("prio"); ok {
		fvApAttr.Prio = Prio.(string)
	}
	fvAp := models.NewApplicationProfile(fmt.Sprintf("ap-%s", name), TenantDn, desc, fvApAttr)

	fvAp.Status = "modified"

	err := aciClient.Save(fvAp)

	if err != nil {
		return err
	}
	if d.HasChange("relation_fv_rs_ap_mon_pol") {
		_, newRelParam := d.GetChange("relation_fv_rs_ap_mon_pol")
		err = aciClient.DeleteRelationfvRsApMonPol(fvAp.DistinguishedName)
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationfvRsApMonPol(fvAp.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}

	d.SetId(fvAp.DistinguishedName)
	return resourceAciApplicationProfileRead(d, m)

}

func resourceAciApplicationProfileRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	fvAp, err := getRemoteApplicationProfile(aciClient, dn)

	if err != nil {
		return err
	}
	setApplicationProfileAttributes(fvAp, d)
	return nil
}

func resourceAciApplicationProfileDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "fvAp")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
