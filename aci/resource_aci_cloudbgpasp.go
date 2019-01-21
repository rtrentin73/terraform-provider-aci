package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciAutonomoussystemprofile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciAutonomoussystemprofileCreate,
		Update: resourceAciAutonomoussystemprofileUpdate,
		Read:   resourceAciAutonomoussystemprofileRead,
		Delete: resourceAciAutonomoussystemprofileDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciAutonomoussystemprofileImport,
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

func getRemoteAutonomoussystemprofile(client *client.Client, dn string) (*models.Autonomoussystemprofile, error) {
	cloudBgpAsPCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudBgpAsP := models.AutonomoussystemprofileFromContainer(cloudBgpAsPCont)

	if cloudBgpAsP.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudBgpAsP.DistinguishedName)
	}

	return cloudBgpAsP, nil
}

func setAutonomoussystemprofileAttributes(cloudBgpAsP *models.Autonomoussystemprofile, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudBgpAsP.DistinguishedName)
	d.Set("description", cloudBgpAsP.Description)
	cloudBgpAsPMap, _ := cloudBgpAsP.ToMap()

	d.Set("annotation", cloudBgpAsPMap["annotation"])
	d.Set("asn", cloudBgpAsPMap["asn"])
	d.Set("name_alias", cloudBgpAsPMap["nameAlias"])
	return d
}

func resourceAciAutonomoussystemprofileImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudBgpAsP, err := getRemoteAutonomoussystemprofile(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setAutonomoussystemprofileAttributes(cloudBgpAsP, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciAutonomoussystemprofileCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	cloudBgpAsPAttr := models.AutonomoussystemprofileAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudBgpAsPAttr.Annotation = Annotation.(string)
	}
	if Asn, ok := d.GetOk("asn"); ok {
		cloudBgpAsPAttr.Asn = Asn.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudBgpAsPAttr.NameAlias = NameAlias.(string)
	}
	cloudBgpAsP := models.NewAutonomoussystemprofile(fmt.Sprintf("clouddomp/as"), "uni", desc, cloudBgpAsPAttr)

	err := aciClient.Save(cloudBgpAsP)
	if err != nil {
		return err
	}

	d.SetId(cloudBgpAsP.DistinguishedName)
	return resourceAciAutonomoussystemprofileRead(d, m)
}

func resourceAciAutonomoussystemprofileUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	cloudBgpAsPAttr := models.AutonomoussystemprofileAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudBgpAsPAttr.Annotation = Annotation.(string)
	}
	if Asn, ok := d.GetOk("asn"); ok {
		cloudBgpAsPAttr.Asn = Asn.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudBgpAsPAttr.NameAlias = NameAlias.(string)
	}
	cloudBgpAsP := models.NewAutonomoussystemprofile(fmt.Sprintf("clouddomp/as"), "uni", desc, cloudBgpAsPAttr)

	cloudBgpAsP.Status = "modified"

	err := aciClient.Save(cloudBgpAsP)

	if err != nil {
		return err
	}

	d.SetId(cloudBgpAsP.DistinguishedName)
	return resourceAciAutonomoussystemprofileRead(d, m)

}

func resourceAciAutonomoussystemprofileRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudBgpAsP, err := getRemoteAutonomoussystemprofile(aciClient, dn)

	if err != nil {
		return err
	}
	setAutonomoussystemprofileAttributes(cloudBgpAsP, d)
	return nil
}

func resourceAciAutonomoussystemprofileDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudBgpAsP")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
