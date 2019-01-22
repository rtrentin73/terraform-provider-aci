package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciCloudEndpointSelector() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciCloudEndpointSelectorCreate,
		Update: resourceAciCloudEndpointSelectorUpdate,
		Read:   resourceAciCloudEndpointSelectorRead,
		Delete: resourceAciCloudEndpointSelectorDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudEndpointSelectorImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"cloud_e_pg_dn": &schema.Schema{
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

func getRemoteCloudEndpointSelector(client *client.Client, dn string) (*models.CloudEndpointSelector, error) {
	cloudEPSelectorCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudEPSelector := models.CloudEndpointSelectorFromContainer(cloudEPSelectorCont)

	if cloudEPSelector.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudEPSelector.DistinguishedName)
	}

	return cloudEPSelector, nil
}

func setCloudEndpointSelectorAttributes(cloudEPSelector *models.CloudEndpointSelector, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudEPSelector.DistinguishedName)
	d.Set("description", cloudEPSelector.Description)
	d.Set("cloud_e_pg_dn", GetParentDn(cloudEPSelector.DistinguishedName))
	cloudEPSelectorMap, _ := cloudEPSelector.ToMap()

	d.Set("annotation", cloudEPSelectorMap["annotation"])
	d.Set("match_expression", cloudEPSelectorMap["matchExpression"])
	d.Set("name_alias", cloudEPSelectorMap["nameAlias"])
	return d
}

func resourceAciCloudEndpointSelectorImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudEPSelector, err := getRemoteCloudEndpointSelector(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCloudEndpointSelectorAttributes(cloudEPSelector, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudEndpointSelectorCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudEPgDn := d.Get("cloud_e_pg_dn").(string)

	cloudEPSelectorAttr := models.CloudEndpointSelectorAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudEPSelectorAttr.Annotation = Annotation.(string)
	}
	if MatchExpression, ok := d.GetOk("match_expression"); ok {
		cloudEPSelectorAttr.MatchExpression = MatchExpression.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudEPSelectorAttr.NameAlias = NameAlias.(string)
	}
	cloudEPSelector := models.NewCloudEndpointSelector(fmt.Sprintf("epselector-%s", name), CloudEPgDn, desc, cloudEPSelectorAttr)

	err := aciClient.Save(cloudEPSelector)
	if err != nil {
		return err
	}

	d.SetId(cloudEPSelector.DistinguishedName)
	return resourceAciCloudEndpointSelectorRead(d, m)
}

func resourceAciCloudEndpointSelectorUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudEPgDn := d.Get("cloud_e_pg_dn").(string)

	cloudEPSelectorAttr := models.CloudEndpointSelectorAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudEPSelectorAttr.Annotation = Annotation.(string)
	}
	if MatchExpression, ok := d.GetOk("match_expression"); ok {
		cloudEPSelectorAttr.MatchExpression = MatchExpression.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudEPSelectorAttr.NameAlias = NameAlias.(string)
	}
	cloudEPSelector := models.NewCloudEndpointSelector(fmt.Sprintf("epselector-%s", name), CloudEPgDn, desc, cloudEPSelectorAttr)

	cloudEPSelector.Status = "modified"

	err := aciClient.Save(cloudEPSelector)

	if err != nil {
		return err
	}

	d.SetId(cloudEPSelector.DistinguishedName)
	return resourceAciCloudEndpointSelectorRead(d, m)

}

func resourceAciCloudEndpointSelectorRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudEPSelector, err := getRemoteCloudEndpointSelector(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudEndpointSelectorAttributes(cloudEPSelector, d)
	return nil
}

func resourceAciCloudEndpointSelectorDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudEPSelector")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
