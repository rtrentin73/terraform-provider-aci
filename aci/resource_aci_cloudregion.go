package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciCloudProvidersRegion() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciCloudProvidersRegionCreate,
		Update: resourceAciCloudProvidersRegionUpdate,
		Read:   resourceAciCloudProvidersRegionRead,
		Delete: resourceAciCloudProvidersRegionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudProvidersRegionImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"cloud_provider_profile_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"admin_st": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "administrative state of the object or policy",
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
		}),
	}
}

func getRemoteCloudProvidersRegion(client *client.Client, dn string) (*models.CloudProvidersRegion, error) {
	cloudRegionCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudRegion := models.CloudProvidersRegionFromContainer(cloudRegionCont)

	if cloudRegion.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudRegion.DistinguishedName)
	}

	return cloudRegion, nil
}

func setCloudProvidersRegionAttributes(cloudRegion *models.CloudProvidersRegion, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudRegion.DistinguishedName)
	d.Set("description", cloudRegion.Description)
	d.Set("cloud_provider_profile_dn", GetParentDn(cloudRegion.DistinguishedName))
	cloudRegionMap, _ := cloudRegion.ToMap()

	d.Set("admin_st", cloudRegionMap["adminSt"])
	d.Set("annotation", cloudRegionMap["annotation"])
	d.Set("name_alias", cloudRegionMap["nameAlias"])
	return d
}

func resourceAciCloudProvidersRegionImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudRegion, err := getRemoteCloudProvidersRegion(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCloudProvidersRegionAttributes(cloudRegion, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudProvidersRegionCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudProviderProfileDn := d.Get("cloud_provider_profile_dn").(string)

	cloudRegionAttr := models.CloudProvidersRegionAttributes{}
	if AdminSt, ok := d.GetOk("admin_st"); ok {
		cloudRegionAttr.AdminSt = AdminSt.(string)
	}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudRegionAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudRegionAttr.NameAlias = NameAlias.(string)
	}
	cloudRegion := models.NewCloudProvidersRegion(fmt.Sprintf("region-%s", name), CloudProviderProfileDn, desc, cloudRegionAttr)

	err := aciClient.Save(cloudRegion)
	if err != nil {
		return err
	}

	d.SetId(cloudRegion.DistinguishedName)
	return resourceAciCloudProvidersRegionRead(d, m)
}

func resourceAciCloudProvidersRegionUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudProviderProfileDn := d.Get("cloud_provider_profile_dn").(string)

	cloudRegionAttr := models.CloudProvidersRegionAttributes{}
	if AdminSt, ok := d.GetOk("admin_st"); ok {
		cloudRegionAttr.AdminSt = AdminSt.(string)
	}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudRegionAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudRegionAttr.NameAlias = NameAlias.(string)
	}
	cloudRegion := models.NewCloudProvidersRegion(fmt.Sprintf("region-%s", name), CloudProviderProfileDn, desc, cloudRegionAttr)

	cloudRegion.Status = "modified"

	err := aciClient.Save(cloudRegion)

	if err != nil {
		return err
	}

	d.SetId(cloudRegion.DistinguishedName)
	return resourceAciCloudProvidersRegionRead(d, m)

}

func resourceAciCloudProvidersRegionRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudRegion, err := getRemoteCloudProvidersRegion(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudProvidersRegionAttributes(cloudRegion, d)
	return nil
}

func resourceAciCloudProvidersRegionDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudRegion")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
