package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciAutonomousSystemProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciAutonomousSystemProfileCreate,
		Update: resourceAciAutonomousSystemProfileUpdate,
		Read:   resourceAciAutonomousSystemProfileRead,
		Delete: resourceAciAutonomousSystemProfileDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciAutonomousSystemProfileImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"annotation": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"asn": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "autonomous system number",
			},

			"name_alias": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},
		}),
	}
}

func getRemoteAutonomousSystemProfile(client *client.Client, dn string) (*models.AutonomousSystemProfile, error) {
	cloudBgpAsPCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudBgpAsP := models.AutonomousSystemProfileFromContainer(cloudBgpAsPCont)

	if cloudBgpAsP.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudBgpAsP.DistinguishedName)
	}

	return cloudBgpAsP, nil
}

func setAutonomousSystemProfileAttributes(cloudBgpAsP *models.AutonomousSystemProfile, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudBgpAsP.DistinguishedName)
	d.Set("description", cloudBgpAsP.Description)
	cloudBgpAsPMap, _ := cloudBgpAsP.ToMap()

	d.Set("annotation", cloudBgpAsPMap["annotation"])
	d.Set("asn", cloudBgpAsPMap["asn"])
	d.Set("name_alias", cloudBgpAsPMap["nameAlias"])
	return d
}

func resourceAciAutonomousSystemProfileImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudBgpAsP, err := getRemoteAutonomousSystemProfile(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setAutonomousSystemProfileAttributes(cloudBgpAsP, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciAutonomousSystemProfileCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	cloudBgpAsPAttr := models.AutonomousSystemProfileAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudBgpAsPAttr.Annotation = Annotation.(string)
	}
	if Asn, ok := d.GetOk("asn"); ok {
		cloudBgpAsPAttr.Asn = Asn.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudBgpAsPAttr.NameAlias = NameAlias.(string)
	}
	cloudBgpAsP := models.NewAutonomousSystemProfile(fmt.Sprintf("clouddomp/as"), "uni", desc, cloudBgpAsPAttr)

	err := aciClient.Save(cloudBgpAsP)
	if err != nil {
		return err
	}

	d.SetId(cloudBgpAsP.DistinguishedName)
	return resourceAciAutonomousSystemProfileRead(d, m)
}

func resourceAciAutonomousSystemProfileUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	cloudBgpAsPAttr := models.AutonomousSystemProfileAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudBgpAsPAttr.Annotation = Annotation.(string)
	}
	if Asn, ok := d.GetOk("asn"); ok {
		cloudBgpAsPAttr.Asn = Asn.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudBgpAsPAttr.NameAlias = NameAlias.(string)
	}
	cloudBgpAsP := models.NewAutonomousSystemProfile(fmt.Sprintf("clouddomp/as"), "uni", desc, cloudBgpAsPAttr)

	cloudBgpAsP.Status = "modified"

	err := aciClient.Save(cloudBgpAsP)

	if err != nil {
		return err
	}

	d.SetId(cloudBgpAsP.DistinguishedName)
	return resourceAciAutonomousSystemProfileRead(d, m)

}

func resourceAciAutonomousSystemProfileRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudBgpAsP, err := getRemoteAutonomousSystemProfile(aciClient, dn)

	if err != nil {
		return err
	}
	setAutonomousSystemProfileAttributes(cloudBgpAsP, d)
	return nil
}

func resourceAciAutonomousSystemProfileDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudBgpAsP")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
