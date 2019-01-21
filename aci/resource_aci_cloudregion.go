package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciCloudprovidersregion() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciCloudprovidersregionCreate,
		Update: resourceAciCloudprovidersregionUpdate,
		Read:   resourceAciCloudprovidersregionRead,
		Delete: resourceAciCloudprovidersregionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudprovidersregionImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"cloudproviderprofile_dn": &schema.Schema{
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

func getRemoteCloudprovidersregion(client *client.Client, dn string) (*models.Cloudprovidersregion, error) {
	cloudRegionCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudRegion := models.CloudprovidersregionFromContainer(cloudRegionCont)

	if cloudRegion.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudRegion.DistinguishedName)
	}

	return cloudRegion, nil
}

func setCloudprovidersregionAttributes(cloudRegion *models.Cloudprovidersregion, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudRegion.DistinguishedName)
	d.Set("description", cloudRegion.Description)
	d.Set("cloudproviderprofile_dn", GetParentDn(cloudRegion.DistinguishedName))
	cloudRegionMap, _ := cloudRegion.ToMap()

	d.Set("admin_st", cloudRegionMap["adminSt"])
	d.Set("annotation", cloudRegionMap["annotation"])
	d.Set("name_alias", cloudRegionMap["nameAlias"])
	return d
}

func resourceAciCloudprovidersregionImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudRegion, err := getRemoteCloudprovidersregion(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCloudprovidersregionAttributes(cloudRegion, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudprovidersregionCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudproviderprofileDn := d.Get("cloudproviderprofile_dn").(string)

	cloudRegionAttr := models.CloudprovidersregionAttributes{}
	if AdminSt, ok := d.GetOk("admin_st"); ok {
		cloudRegionAttr.AdminSt = AdminSt.(string)
	}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudRegionAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudRegionAttr.NameAlias = NameAlias.(string)
	}
	cloudRegion := models.NewCloudprovidersregion(fmt.Sprintf("region-%s", name), CloudproviderprofileDn, desc, cloudRegionAttr)

	err := aciClient.Save(cloudRegion)
	if err != nil {
		return err
	}

	d.SetId(cloudRegion.DistinguishedName)
	return resourceAciCloudprovidersregionRead(d, m)
}

func resourceAciCloudprovidersregionUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudproviderprofileDn := d.Get("cloudproviderprofile_dn").(string)

	cloudRegionAttr := models.CloudprovidersregionAttributes{}
	if AdminSt, ok := d.GetOk("admin_st"); ok {
		cloudRegionAttr.AdminSt = AdminSt.(string)
	}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudRegionAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudRegionAttr.NameAlias = NameAlias.(string)
	}
	cloudRegion := models.NewCloudprovidersregion(fmt.Sprintf("region-%s", name), CloudproviderprofileDn, desc, cloudRegionAttr)

	cloudRegion.Status = "modified"

	err := aciClient.Save(cloudRegion)

	if err != nil {
		return err
	}

	d.SetId(cloudRegion.DistinguishedName)
	return resourceAciCloudprovidersregionRead(d, m)

}

func resourceAciCloudprovidersregionRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudRegion, err := getRemoteCloudprovidersregion(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudprovidersregionAttributes(cloudRegion, d)
	return nil
}

func resourceAciCloudprovidersregionDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudRegion")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
