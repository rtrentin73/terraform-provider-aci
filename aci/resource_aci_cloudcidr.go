package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciCloudcidrpool() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciCloudcidrpoolCreate,
		Update: resourceAciCloudcidrpoolUpdate,
		Read:   resourceAciCloudcidrpoolRead,
		Delete: resourceAciCloudcidrpoolDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudcidrpoolImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"cloudcontextprofile_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"addr": &schema.Schema{
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

			"primary": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},
		}),
	}
}

func getRemoteCloudcidrpool(client *client.Client, dn string) (*models.Cloudcidrpool, error) {
	cloudCidrCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudCidr := models.CloudcidrpoolFromContainer(cloudCidrCont)

	if cloudCidr.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudCidr.DistinguishedName)
	}

	return cloudCidr, nil
}

func setCloudcidrpoolAttributes(cloudCidr *models.Cloudcidrpool, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudCidr.DistinguishedName)
	d.Set("description", cloudCidr.Description)
	d.Set("cloudcontextprofile_dn", GetParentDn(cloudCidr.DistinguishedName))
	cloudCidrMap, _ := cloudCidr.ToMap()

	d.Set("addr", cloudCidrMap["addr"])
	d.Set("annotation", cloudCidrMap["annotation"])
	d.Set("name_alias", cloudCidrMap["nameAlias"])
	d.Set("primary", cloudCidrMap["primary"])
	return d
}

func resourceAciCloudcidrpoolImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudCidr, err := getRemoteCloudcidrpool(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCloudcidrpoolAttributes(cloudCidr, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudcidrpoolCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	addr := d.Get("addr").(string)

	CloudcontextprofileDn := d.Get("cloudcontextprofile_dn").(string)

	cloudCidrAttr := models.CloudcidrpoolAttributes{}
	if Addr, ok := d.GetOk("addr"); ok {
		cloudCidrAttr.Addr = Addr.(string)
	}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudCidrAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudCidrAttr.NameAlias = NameAlias.(string)
	}
	if Primary, ok := d.GetOk("primary"); ok {
		cloudCidrAttr.Primary = Primary.(string)
	}
	cloudCidr := models.NewCloudcidrpool(fmt.Sprintf("cidr-[%s]", addr), CloudcontextprofileDn, desc, cloudCidrAttr)

	err := aciClient.Save(cloudCidr)
	if err != nil {
		return err
	}

	d.SetId(cloudCidr.DistinguishedName)
	return resourceAciCloudcidrpoolRead(d, m)
}

func resourceAciCloudcidrpoolUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	addr := d.Get("addr").(string)

	CloudcontextprofileDn := d.Get("cloudcontextprofile_dn").(string)

	cloudCidrAttr := models.CloudcidrpoolAttributes{}
	if Addr, ok := d.GetOk("addr"); ok {
		cloudCidrAttr.Addr = Addr.(string)
	}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudCidrAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudCidrAttr.NameAlias = NameAlias.(string)
	}
	if Primary, ok := d.GetOk("primary"); ok {
		cloudCidrAttr.Primary = Primary.(string)
	}
	cloudCidr := models.NewCloudcidrpool(fmt.Sprintf("cidr-[%s]", addr), CloudcontextprofileDn, desc, cloudCidrAttr)

	cloudCidr.Status = "modified"

	err := aciClient.Save(cloudCidr)

	if err != nil {
		return err
	}

	d.SetId(cloudCidr.DistinguishedName)
	return resourceAciCloudcidrpoolRead(d, m)

}

func resourceAciCloudcidrpoolRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudCidr, err := getRemoteCloudcidrpool(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudcidrpoolAttributes(cloudCidr, d)
	return nil
}

func resourceAciCloudcidrpoolDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudCidr")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
