package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciCloudEndpointSelectorforExternalEPgs() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciCloudEndpointSelectorforExternalEPgsCreate,
		Update: resourceAciCloudEndpointSelectorforExternalEPgsUpdate,
		Read:   resourceAciCloudEndpointSelectorforExternalEPgsRead,
		Delete: resourceAciCloudEndpointSelectorforExternalEPgsDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudEndpointSelectorforExternalEPgsImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"cloud_external_e_pg_dn": &schema.Schema{
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

func getRemoteCloudEndpointSelectorforExternalEPgs(client *client.Client, dn string) (*models.CloudEndpointSelectorforExternalEPgs, error) {
	cloudExtEPSelectorCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudExtEPSelector := models.CloudEndpointSelectorforExternalEPgsFromContainer(cloudExtEPSelectorCont)

	if cloudExtEPSelector.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudExtEPSelector.DistinguishedName)
	}

	return cloudExtEPSelector, nil
}

func setCloudEndpointSelectorforExternalEPgsAttributes(cloudExtEPSelector *models.CloudEndpointSelectorforExternalEPgs, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudExtEPSelector.DistinguishedName)
	d.Set("description", cloudExtEPSelector.Description)
	d.Set("cloud_external_e_pg_dn", GetParentDn(cloudExtEPSelector.DistinguishedName))
	cloudExtEPSelectorMap, _ := cloudExtEPSelector.ToMap()

	d.Set("annotation", cloudExtEPSelectorMap["annotation"])
	d.Set("is_shared", cloudExtEPSelectorMap["isShared"])
	d.Set("match_expression", cloudExtEPSelectorMap["matchExpression"])
	d.Set("name_alias", cloudExtEPSelectorMap["nameAlias"])
	d.Set("subnet", cloudExtEPSelectorMap["subnet"])
	return d
}

func resourceAciCloudEndpointSelectorforExternalEPgsImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudExtEPSelector, err := getRemoteCloudEndpointSelectorforExternalEPgs(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCloudEndpointSelectorforExternalEPgsAttributes(cloudExtEPSelector, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudEndpointSelectorforExternalEPgsCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	//name := d.Get("name").(string)

	CloudExternalEPgDn := d.Get("cloud_external_e_pg_dn").(string)

	cloudExtEPSelectorAttr := models.CloudEndpointSelectorforExternalEPgsAttributes{}
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
	cloudExtEPSelector := models.NewCloudEndpointSelectorforExternalEPgs(fmt.Sprintf("extepselector-[%s]", cloudExtEPSelectorAttr.Subnet), CloudExternalEPgDn, desc, cloudExtEPSelectorAttr)

	err := aciClient.Save(cloudExtEPSelector)
	if err != nil {
		return err
	}

	d.SetId(cloudExtEPSelector.DistinguishedName)
	return resourceAciCloudEndpointSelectorforExternalEPgsRead(d, m)
}

func resourceAciCloudEndpointSelectorforExternalEPgsUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	CloudExternalEPgDn := d.Get("cloud_external_e_pg_dn").(string)

	cloudExtEPSelectorAttr := models.CloudEndpointSelectorforExternalEPgsAttributes{}
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
	cloudExtEPSelector := models.NewCloudEndpointSelectorforExternalEPgs(fmt.Sprintf("extepselector-%s", name), CloudExternalEPgDn, desc, cloudExtEPSelectorAttr)

	cloudExtEPSelector.Status = "modified"

	err := aciClient.Save(cloudExtEPSelector)

	if err != nil {
		return err
	}

	d.SetId(cloudExtEPSelector.DistinguishedName)
	return resourceAciCloudEndpointSelectorforExternalEPgsRead(d, m)

}

func resourceAciCloudEndpointSelectorforExternalEPgsRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudExtEPSelector, err := getRemoteCloudEndpointSelectorforExternalEPgs(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudEndpointSelectorforExternalEPgsAttributes(cloudExtEPSelector, d)
	return nil
}

func resourceAciCloudEndpointSelectorforExternalEPgsDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudExtEPSelector")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
