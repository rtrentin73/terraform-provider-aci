package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciCloudAvailabilityZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciCloudAvailabilityZoneCreate,
		Update: resourceAciCloudAvailabilityZoneUpdate,
		Read:   resourceAciCloudAvailabilityZoneRead,
		Delete: resourceAciCloudAvailabilityZoneDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudAvailabilityZoneImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"cloud_providers_region_dn": &schema.Schema{
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
		}),
	}
}

func getRemoteCloudAvailabilityZone(client *client.Client, dn string) (*models.CloudAvailabilityZone, error) {
	cloudZoneCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudZone := models.CloudAvailabilityZoneFromContainer(cloudZoneCont)

	if cloudZone.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudZone.DistinguishedName)
	}

	return cloudZone, nil
}

func setCloudAvailabilityZoneAttributes(cloudZone *models.CloudAvailabilityZone, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudZone.DistinguishedName)
	d.Set("description", cloudZone.Description)
	d.Set("cloud_providers_region_dn", GetParentDn(cloudZone.DistinguishedName))
	cloudZoneMap, _ := cloudZone.ToMap()

	d.Set("annotation", cloudZoneMap["annotation"])
	d.Set("name_alias", cloudZoneMap["nameAlias"])
	return d
}

func resourceAciCloudAvailabilityZoneImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudZone, err := getRemoteCloudAvailabilityZone(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCloudAvailabilityZoneAttributes(cloudZone, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudAvailabilityZoneCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudProvidersRegionDn := d.Get("cloud_providers_region_dn").(string)

	cloudZoneAttr := models.CloudAvailabilityZoneAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudZoneAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudZoneAttr.NameAlias = NameAlias.(string)
	}
	cloudZone := models.NewCloudAvailabilityZone(fmt.Sprintf("zone-%s", name), CloudProvidersRegionDn, desc, cloudZoneAttr)

	err := aciClient.Save(cloudZone)
	if err != nil {
		return err
	}

	d.SetId(cloudZone.DistinguishedName)
	return resourceAciCloudAvailabilityZoneRead(d, m)
}

func resourceAciCloudAvailabilityZoneUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudProvidersRegionDn := d.Get("cloud_providers_region_dn").(string)

	cloudZoneAttr := models.CloudAvailabilityZoneAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudZoneAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudZoneAttr.NameAlias = NameAlias.(string)
	}
	cloudZone := models.NewCloudAvailabilityZone(fmt.Sprintf("zone-%s", name), CloudProvidersRegionDn, desc, cloudZoneAttr)

	cloudZone.Status = "modified"

	err := aciClient.Save(cloudZone)

	if err != nil {
		return err
	}

	d.SetId(cloudZone.DistinguishedName)
	return resourceAciCloudAvailabilityZoneRead(d, m)

}

func resourceAciCloudAvailabilityZoneRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudZone, err := getRemoteCloudAvailabilityZone(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudAvailabilityZoneAttributes(cloudZone, d)
	return nil
}

func resourceAciCloudAvailabilityZoneDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudZone")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
