package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciCloudavailabilityzone() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciCloudavailabilityzoneCreate,
		Update: resourceAciCloudavailabilityzoneUpdate,
		Read:   resourceAciCloudavailabilityzoneRead,
		Delete: resourceAciCloudavailabilityzoneDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudavailabilityzoneImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"cloudprovidersregion_dn": &schema.Schema{
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

func getRemoteCloudavailabilityzone(client *client.Client, dn string) (*models.Cloudavailabilityzone, error) {
	cloudZoneCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudZone := models.CloudavailabilityzoneFromContainer(cloudZoneCont)

	if cloudZone.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudZone.DistinguishedName)
	}

	return cloudZone, nil
}

func setCloudavailabilityzoneAttributes(cloudZone *models.Cloudavailabilityzone, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudZone.DistinguishedName)
	d.Set("description", cloudZone.Description)
	d.Set("cloudprovidersregion_dn", GetParentDn(cloudZone.DistinguishedName))
	cloudZoneMap, _ := cloudZone.ToMap()

	d.Set("annotation", cloudZoneMap["annotation"])
	d.Set("name_alias", cloudZoneMap["nameAlias"])
	return d
}

func resourceAciCloudavailabilityzoneImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudZone, err := getRemoteCloudavailabilityzone(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCloudavailabilityzoneAttributes(cloudZone, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudavailabilityzoneCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudprovidersregionDn := d.Get("cloudprovidersregion_dn").(string)

	cloudZoneAttr := models.CloudavailabilityzoneAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudZoneAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudZoneAttr.NameAlias = NameAlias.(string)
	}
	cloudZone := models.NewCloudavailabilityzone(fmt.Sprintf("zone-%s", name), CloudprovidersregionDn, desc, cloudZoneAttr)

	err := aciClient.Save(cloudZone)
	if err != nil {
		return err
	}

	d.SetId(cloudZone.DistinguishedName)
	return resourceAciCloudavailabilityzoneRead(d, m)
}

func resourceAciCloudavailabilityzoneUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudprovidersregionDn := d.Get("cloudprovidersregion_dn").(string)

	cloudZoneAttr := models.CloudavailabilityzoneAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudZoneAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudZoneAttr.NameAlias = NameAlias.(string)
	}
	cloudZone := models.NewCloudavailabilityzone(fmt.Sprintf("zone-%s", name), CloudprovidersregionDn, desc, cloudZoneAttr)

	cloudZone.Status = "modified"

	err := aciClient.Save(cloudZone)

	if err != nil {
		return err
	}

	d.SetId(cloudZone.DistinguishedName)
	return resourceAciCloudavailabilityzoneRead(d, m)

}

func resourceAciCloudavailabilityzoneRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudZone, err := getRemoteCloudavailabilityzone(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudavailabilityzoneAttributes(cloudZone, d)
	return nil
}

func resourceAciCloudavailabilityzoneDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudZone")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
