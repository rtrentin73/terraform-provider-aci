package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciCloudapplicationcontainer() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciCloudapplicationcontainerCreate,
		Update: resourceAciCloudapplicationcontainerUpdate,
		Read:   resourceAciCloudapplicationcontainerRead,
		Delete: resourceAciCloudapplicationcontainerDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudapplicationcontainerImport,
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

func getRemoteCloudapplicationcontainer(client *client.Client, dn string) (*models.Cloudapplicationcontainer, error) {
	cloudAppCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudApp := models.CloudapplicationcontainerFromContainer(cloudAppCont)

	if cloudApp.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudApp.DistinguishedName)
	}

	return cloudApp, nil
}

func setCloudapplicationcontainerAttributes(cloudApp *models.Cloudapplicationcontainer, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudApp.DistinguishedName)
	d.Set("description", cloudApp.Description)
	d.Set("tenant_dn", GetParentDn(cloudApp.DistinguishedName))
	cloudAppMap, _ := cloudApp.ToMap()

	d.Set("annotation", cloudAppMap["annotation"])
	d.Set("name_alias", cloudAppMap["nameAlias"])
	return d
}

func resourceAciCloudapplicationcontainerImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudApp, err := getRemoteCloudapplicationcontainer(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCloudapplicationcontainerAttributes(cloudApp, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudapplicationcontainerCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	TenantDn := d.Get("tenant_dn").(string)

	cloudAppAttr := models.CloudapplicationcontainerAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudAppAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudAppAttr.NameAlias = NameAlias.(string)
	}
	cloudApp := models.NewCloudapplicationcontainer(fmt.Sprintf("cloudapp-%s", name), TenantDn, desc, cloudAppAttr)

	err := aciClient.Save(cloudApp)
	if err != nil {
		return err
	}

	d.SetId(cloudApp.DistinguishedName)
	return resourceAciCloudapplicationcontainerRead(d, m)
}

func resourceAciCloudapplicationcontainerUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	TenantDn := d.Get("tenant_dn").(string)

	cloudAppAttr := models.CloudapplicationcontainerAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudAppAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudAppAttr.NameAlias = NameAlias.(string)
	}
	cloudApp := models.NewCloudapplicationcontainer(fmt.Sprintf("cloudapp-%s", name), TenantDn, desc, cloudAppAttr)

	cloudApp.Status = "modified"

	err := aciClient.Save(cloudApp)

	if err != nil {
		return err
	}

	d.SetId(cloudApp.DistinguishedName)
	return resourceAciCloudapplicationcontainerRead(d, m)

}

func resourceAciCloudapplicationcontainerRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudApp, err := getRemoteCloudapplicationcontainer(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudapplicationcontainerAttributes(cloudApp, d)
	return nil
}

func resourceAciCloudapplicationcontainerDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudApp")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
