package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciClouddomainprofile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciClouddomainprofileCreate,
		Update: resourceAciClouddomainprofileUpdate,
		Read:   resourceAciClouddomainprofileRead,
		Delete: resourceAciClouddomainprofileDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciClouddomainprofileImport,
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

func getRemoteClouddomainprofile(client *client.Client, dn string) (*models.Clouddomainprofile, error) {
	cloudDomPCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudDomP := models.ClouddomainprofileFromContainer(cloudDomPCont)

	if cloudDomP.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudDomP.DistinguishedName)
	}

	return cloudDomP, nil
}

func setClouddomainprofileAttributes(cloudDomP *models.Clouddomainprofile, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudDomP.DistinguishedName)
	d.Set("description", cloudDomP.Description)
	cloudDomPMap, _ := cloudDomP.ToMap()

	d.Set("annotation", cloudDomPMap["annotation"])
	d.Set("name_alias", cloudDomPMap["nameAlias"])
	d.Set("site_id", cloudDomPMap["siteId"])
	return d
}

func resourceAciClouddomainprofileImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudDomP, err := getRemoteClouddomainprofile(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setClouddomainprofileAttributes(cloudDomP, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciClouddomainprofileCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	cloudDomPAttr := models.ClouddomainprofileAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudDomPAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudDomPAttr.NameAlias = NameAlias.(string)
	}
	if SiteId, ok := d.GetOk("site_id"); ok {
		cloudDomPAttr.SiteId = SiteId.(string)
	}
	cloudDomP := models.NewClouddomainprofile(fmt.Sprintf("clouddomp"), "uni", desc, cloudDomPAttr)

	err := aciClient.Save(cloudDomP)
	if err != nil {
		return err
	}

	d.SetId(cloudDomP.DistinguishedName)
	return resourceAciClouddomainprofileRead(d, m)
}

func resourceAciClouddomainprofileUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	cloudDomPAttr := models.ClouddomainprofileAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudDomPAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudDomPAttr.NameAlias = NameAlias.(string)
	}
	if SiteId, ok := d.GetOk("site_id"); ok {
		cloudDomPAttr.SiteId = SiteId.(string)
	}
	cloudDomP := models.NewClouddomainprofile(fmt.Sprintf("clouddomp"), "uni", desc, cloudDomPAttr)

	cloudDomP.Status = "modified"

	err := aciClient.Save(cloudDomP)

	if err != nil {
		return err
	}

	d.SetId(cloudDomP.DistinguishedName)
	return resourceAciClouddomainprofileRead(d, m)

}

func resourceAciClouddomainprofileRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudDomP, err := getRemoteClouddomainprofile(aciClient, dn)

	if err != nil {
		return err
	}
	setClouddomainprofileAttributes(cloudDomP, d)
	return nil
}

func resourceAciClouddomainprofileDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudDomP")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
