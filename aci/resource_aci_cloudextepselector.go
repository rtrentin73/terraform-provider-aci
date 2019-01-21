package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciCloudendpointselectorforexternalepgs() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciCloudendpointselectorforexternalepgsCreate,
		Update: resourceAciCloudendpointselectorforexternalepgsUpdate,
		Read:   resourceAciCloudendpointselectorforexternalepgsRead,
		Delete: resourceAciCloudendpointselectorforexternalepgsDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudendpointselectorforexternalepgsImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"cloudexternalepg_dn": &schema.Schema{
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

			"is_shared": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"match_expression": &schema.Schema{
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

			"subnet": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},
		}),
	}
}

func getRemoteCloudendpointselectorforexternalepgs(client *client.Client, dn string) (*models.Cloudendpointselectorforexternalepgs, error) {
	cloudExtEPSelectorCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudExtEPSelector := models.CloudendpointselectorforexternalepgsFromContainer(cloudExtEPSelectorCont)

	if cloudExtEPSelector.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudExtEPSelector.DistinguishedName)
	}

	return cloudExtEPSelector, nil
}

func setCloudendpointselectorforexternalepgsAttributes(cloudExtEPSelector *models.Cloudendpointselectorforexternalepgs, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudExtEPSelector.DistinguishedName)
	d.Set("description", cloudExtEPSelector.Description)
	d.Set("cloudexternalepg_dn", GetParentDn(cloudExtEPSelector.DistinguishedName))
	cloudExtEPSelectorMap, _ := cloudExtEPSelector.ToMap()

	d.Set("annotation", cloudExtEPSelectorMap["annotation"])
	d.Set("is_shared", cloudExtEPSelectorMap["isShared"])
	d.Set("match_expression", cloudExtEPSelectorMap["matchExpression"])
	d.Set("name_alias", cloudExtEPSelectorMap["nameAlias"])
	d.Set("subnet", cloudExtEPSelectorMap["subnet"])
	return d
}

func resourceAciCloudendpointselectorforexternalepgsImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudExtEPSelector, err := getRemoteCloudendpointselectorforexternalepgs(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCloudendpointselectorforexternalepgsAttributes(cloudExtEPSelector, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudendpointselectorforexternalepgsCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudexternalepgDn := d.Get("cloudexternalepg_dn").(string)

	cloudExtEPSelectorAttr := models.CloudendpointselectorforexternalepgsAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudExtEPSelectorAttr.Annotation = Annotation.(string)
	}
	if IsShared, ok := d.GetOk("is_shared"); ok {
		cloudExtEPSelectorAttr.IsShared = IsShared.(string)
	}
	if MatchExpression, ok := d.GetOk("match_expression"); ok {
		cloudExtEPSelectorAttr.MatchExpression = MatchExpression.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudExtEPSelectorAttr.NameAlias = NameAlias.(string)
	}
	if Subnet, ok := d.GetOk("subnet"); ok {
		cloudExtEPSelectorAttr.Subnet = Subnet.(string)
	}
	cloudExtEPSelector := models.NewCloudendpointselectorforexternalepgs(fmt.Sprintf("extepselector-%s", name), CloudexternalepgDn, desc, cloudExtEPSelectorAttr)

	err := aciClient.Save(cloudExtEPSelector)
	if err != nil {
		return err
	}

	d.SetId(cloudExtEPSelector.DistinguishedName)
	return resourceAciCloudendpointselectorforexternalepgsRead(d, m)
}

func resourceAciCloudendpointselectorforexternalepgsUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudexternalepgDn := d.Get("cloudexternalepg_dn").(string)

	cloudExtEPSelectorAttr := models.CloudendpointselectorforexternalepgsAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudExtEPSelectorAttr.Annotation = Annotation.(string)
	}
	if IsShared, ok := d.GetOk("is_shared"); ok {
		cloudExtEPSelectorAttr.IsShared = IsShared.(string)
	}
	if MatchExpression, ok := d.GetOk("match_expression"); ok {
		cloudExtEPSelectorAttr.MatchExpression = MatchExpression.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudExtEPSelectorAttr.NameAlias = NameAlias.(string)
	}
	if Subnet, ok := d.GetOk("subnet"); ok {
		cloudExtEPSelectorAttr.Subnet = Subnet.(string)
	}
	cloudExtEPSelector := models.NewCloudendpointselectorforexternalepgs(fmt.Sprintf("extepselector-%s", name), CloudexternalepgDn, desc, cloudExtEPSelectorAttr)

	cloudExtEPSelector.Status = "modified"

	err := aciClient.Save(cloudExtEPSelector)

	if err != nil {
		return err
	}

	d.SetId(cloudExtEPSelector.DistinguishedName)
	return resourceAciCloudendpointselectorforexternalepgsRead(d, m)

}

func resourceAciCloudendpointselectorforexternalepgsRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudExtEPSelector, err := getRemoteCloudendpointselectorforexternalepgs(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudendpointselectorforexternalepgsAttributes(cloudExtEPSelector, d)
	return nil
}

func resourceAciCloudendpointselectorforexternalepgsDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudExtEPSelector")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
