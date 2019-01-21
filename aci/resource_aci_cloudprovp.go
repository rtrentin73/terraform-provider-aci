package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciCloudproviderprofile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciCloudproviderprofileCreate,
		Update: resourceAciCloudproviderprofileUpdate,
		Read:   resourceAciCloudproviderprofileRead,
		Delete: resourceAciCloudproviderprofileDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudproviderprofileImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"vendor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"annotation": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},
		}),
	}
}

func getRemoteCloudproviderprofile(client *client.Client, dn string) (*models.Cloudproviderprofile, error) {
	cloudProvPCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudProvP := models.CloudproviderprofileFromContainer(cloudProvPCont)

	if cloudProvP.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudProvP.DistinguishedName)
	}

	return cloudProvP, nil
}

func setCloudproviderprofileAttributes(cloudProvP *models.Cloudproviderprofile, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudProvP.DistinguishedName)
	d.Set("description", cloudProvP.Description)
	cloudProvPMap, _ := cloudProvP.ToMap()

	d.Set("annotation", cloudProvPMap["annotation"])
	d.Set("vendor", cloudProvPMap["vendor"])
	return d
}

func resourceAciCloudproviderprofileImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudProvP, err := getRemoteCloudproviderprofile(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCloudproviderprofileAttributes(cloudProvP, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudproviderprofileCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	vendor := d.Get("vendor").(string)

	cloudProvPAttr := models.CloudproviderprofileAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudProvPAttr.Annotation = Annotation.(string)
	}
	if Vendor, ok := d.GetOk("vendor"); ok {
		cloudProvPAttr.Vendor = Vendor.(string)
	}
	cloudProvP := models.NewCloudproviderprofile(fmt.Sprintf("clouddomp/provp-%s", vendor), "uni", desc, cloudProvPAttr)

	err := aciClient.Save(cloudProvP)
	if err != nil {
		return err
	}

	d.SetId(cloudProvP.DistinguishedName)
	return resourceAciCloudproviderprofileRead(d, m)
}

func resourceAciCloudproviderprofileUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	vendor := d.Get("vendor").(string)

	cloudProvPAttr := models.CloudproviderprofileAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudProvPAttr.Annotation = Annotation.(string)
	}
	if Vendor, ok := d.GetOk("vendor"); ok {
		cloudProvPAttr.Vendor = Vendor.(string)
	}
	cloudProvP := models.NewCloudproviderprofile(fmt.Sprintf("clouddomp/provp-%s", vendor), "uni", desc, cloudProvPAttr)

	cloudProvP.Status = "modified"

	err := aciClient.Save(cloudProvP)

	if err != nil {
		return err
	}

	d.SetId(cloudProvP.DistinguishedName)
	return resourceAciCloudproviderprofileRead(d, m)

}

func resourceAciCloudproviderprofileRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudProvP, err := getRemoteCloudproviderprofile(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudproviderprofileAttributes(cloudProvP, d)
	return nil
}

func resourceAciCloudproviderprofileDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudProvP")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
