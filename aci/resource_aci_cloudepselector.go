package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciCloudendpointselector() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciCloudendpointselectorCreate,
		Update: resourceAciCloudendpointselectorUpdate,
		Read:   resourceAciCloudendpointselectorRead,
		Delete: resourceAciCloudendpointselectorDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudendpointselectorImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"cloudepg_dn": &schema.Schema{
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
		}),
	}
}

func getRemoteCloudendpointselector(client *client.Client, dn string) (*models.Cloudendpointselector, error) {
	cloudEPSelectorCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudEPSelector := models.CloudendpointselectorFromContainer(cloudEPSelectorCont)

	if cloudEPSelector.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudEPSelector.DistinguishedName)
	}

	return cloudEPSelector, nil
}

func setCloudendpointselectorAttributes(cloudEPSelector *models.Cloudendpointselector, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudEPSelector.DistinguishedName)
	d.Set("description", cloudEPSelector.Description)
	d.Set("cloudepg_dn", GetParentDn(cloudEPSelector.DistinguishedName))
	cloudEPSelectorMap, _ := cloudEPSelector.ToMap()

	d.Set("annotation", cloudEPSelectorMap["annotation"])
	d.Set("match_expression", cloudEPSelectorMap["matchExpression"])
	d.Set("name_alias", cloudEPSelectorMap["nameAlias"])
	return d
}

func resourceAciCloudendpointselectorImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudEPSelector, err := getRemoteCloudendpointselector(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCloudendpointselectorAttributes(cloudEPSelector, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudendpointselectorCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudepgDn := d.Get("cloudepg_dn").(string)

	cloudEPSelectorAttr := models.CloudendpointselectorAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudEPSelectorAttr.Annotation = Annotation.(string)
	}
	if MatchExpression, ok := d.GetOk("match_expression"); ok {
		cloudEPSelectorAttr.MatchExpression = MatchExpression.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudEPSelectorAttr.NameAlias = NameAlias.(string)
	}
	cloudEPSelector := models.NewCloudendpointselector(fmt.Sprintf("epselector-%s", name), CloudepgDn, desc, cloudEPSelectorAttr)

	err := aciClient.Save(cloudEPSelector)
	if err != nil {
		return err
	}

	d.SetId(cloudEPSelector.DistinguishedName)
	return resourceAciCloudendpointselectorRead(d, m)
}

func resourceAciCloudendpointselectorUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudepgDn := d.Get("cloudepg_dn").(string)

	cloudEPSelectorAttr := models.CloudendpointselectorAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudEPSelectorAttr.Annotation = Annotation.(string)
	}
	if MatchExpression, ok := d.GetOk("match_expression"); ok {
		cloudEPSelectorAttr.MatchExpression = MatchExpression.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudEPSelectorAttr.NameAlias = NameAlias.(string)
	}
	cloudEPSelector := models.NewCloudendpointselector(fmt.Sprintf("epselector-%s", name), CloudepgDn, desc, cloudEPSelectorAttr)

	cloudEPSelector.Status = "modified"

	err := aciClient.Save(cloudEPSelector)

	if err != nil {
		return err
	}

	d.SetId(cloudEPSelector.DistinguishedName)
	return resourceAciCloudendpointselectorRead(d, m)

}

func resourceAciCloudendpointselectorRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudEPSelector, err := getRemoteCloudendpointselector(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudendpointselectorAttributes(cloudEPSelector, d)
	return nil
}

func resourceAciCloudendpointselectorDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudEPSelector")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
