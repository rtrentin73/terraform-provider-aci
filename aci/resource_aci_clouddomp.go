package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciCloudDomainProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciCloudDomainProfileCreate,
		Update: resourceAciCloudDomainProfileUpdate,
		Read:   resourceAciCloudDomainProfileRead,
		Delete: resourceAciCloudDomainProfileDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudDomainProfileImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

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

			"site_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},
		}),
	}
}

func getRemoteCloudDomainProfile(client *client.Client, dn string) (*models.CloudDomainProfile, error) {
	cloudDomPCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudDomP := models.CloudDomainProfileFromContainer(cloudDomPCont)

	if cloudDomP.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudDomP.DistinguishedName)
	}

	return cloudDomP, nil
}

func setCloudDomainProfileAttributes(cloudDomP *models.CloudDomainProfile, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudDomP.DistinguishedName)
	d.Set("description", cloudDomP.Description)
	cloudDomPMap, _ := cloudDomP.ToMap()

	d.Set("annotation", cloudDomPMap["annotation"])
	d.Set("name_alias", cloudDomPMap["nameAlias"])
	d.Set("site_id", cloudDomPMap["siteId"])
	return d
}

func resourceAciCloudDomainProfileImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudDomP, err := getRemoteCloudDomainProfile(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCloudDomainProfileAttributes(cloudDomP, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudDomainProfileCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	cloudDomPAttr := models.CloudDomainProfileAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudDomPAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudDomPAttr.NameAlias = NameAlias.(string)
	}
	if SiteId, ok := d.GetOk("site_id"); ok {
		cloudDomPAttr.SiteId = SiteId.(string)
	}
	cloudDomP := models.NewCloudDomainProfile(fmt.Sprintf("clouddomp"), "uni", desc, cloudDomPAttr)

	err := aciClient.Save(cloudDomP)
	if err != nil {
		return err
	}

	d.SetId(cloudDomP.DistinguishedName)
	return resourceAciCloudDomainProfileRead(d, m)
}

func resourceAciCloudDomainProfileUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	cloudDomPAttr := models.CloudDomainProfileAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudDomPAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudDomPAttr.NameAlias = NameAlias.(string)
	}
	if SiteId, ok := d.GetOk("site_id"); ok {
		cloudDomPAttr.SiteId = SiteId.(string)
	}
	cloudDomP := models.NewCloudDomainProfile(fmt.Sprintf("clouddomp"), "uni", desc, cloudDomPAttr)

	cloudDomP.Status = "modified"

	err := aciClient.Save(cloudDomP)

	if err != nil {
		return err
	}

	d.SetId(cloudDomP.DistinguishedName)
	return resourceAciCloudDomainProfileRead(d, m)

}

func resourceAciCloudDomainProfileRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudDomP, err := getRemoteCloudDomainProfile(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudDomainProfileAttributes(cloudDomP, d)
	return nil
}

func resourceAciCloudDomainProfileDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudDomP")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
