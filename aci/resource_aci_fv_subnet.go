package aci

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceAciSubnet() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciSubnetCreate,
		Update: resourceAciSubnetUpdate,
		Read:   resourceAciSubnetRead,
		Delete: resourceAciSubnetDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciSubnetImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"fv_bd_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"annotation": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"ctrl": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "subnet control state",

				ValidateFunc: validation.StringInSlice([]string{
					"nd",
					"no-default-gateway",
					"querier",
					"unspecified",
				}, false),
			},

			"name_alias": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"preferred": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "subnet preferred status",

				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},

			"scope": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "subnet visibility",

				ValidateFunc: validation.StringInSlice([]string{
					"private",
					"public",
					"shared",
				}, false),
			},

			"virtual": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",

				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},
		}),
	}
}

func getRemoteSubnet(client *client.Client, dn string) (*models.Subnet, error) {
	fvSubnetCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	fvSubnet := models.SubnetFromContainer(fvSubnetCont)

	if fvSubnet.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", fvSubnet.DistinguishedName)
	}

	return fvSubnet, nil
}

func setSubnetAttributes(fvSubnet *models.Subnet, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(fvSubnet.DistinguishedName)
	d.Set("description", fvSubnet.Description)
	d.Set("fv_bd_dn", GetParentDn(fvSubnet.DistinguishedName))

	d.Set("annotation", fvSubnet.Annotation)
	d.Set("ctrl", fvSubnet.Ctrl)
	d.Set("ip", fvSubnet.Ip)
	d.Set("name_alias", fvSubnet.NameAlias)
	d.Set("preferred", fvSubnet.Preferred)
	d.Set("scope", fvSubnet.Scope)
	d.Set("virtual", fvSubnet.Virtual)
	return d
}

func resourceAciSubnetImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	fvSubnet, err := getRemoteSubnet(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setSubnetAttributes(fvSubnet, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciSubnetCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	ip := d.Get("ip").(string)
	fvBDDn := d.Get("fv_bd_dn").(string)

	fvSubnetAttr := models.SubnetAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		fvSubnetAttr.Annotation = Annotation.(string)
	}
	if Ctrl, ok := d.GetOk("ctrl"); ok {
		fvSubnetAttr.Ctrl = Ctrl.(string)
	}

	if NameAlias, ok := d.GetOk("name_alias"); ok {
		fvSubnetAttr.NameAlias = NameAlias.(string)
	}
	if Preferred, ok := d.GetOk("preferred"); ok {
		fvSubnetAttr.Preferred = Preferred.(string)
	}
	if Scope, ok := d.GetOk("scope"); ok {
		fvSubnetAttr.Scope = Scope.(string)
	}
	if Virtual, ok := d.GetOk("virtual"); ok {
		fvSubnetAttr.Virtual = Virtual.(string)
	}
	fvSubnet := models.NewSubnet(fmt.Sprintf("subnet-[%s]", ip), fvBDDn, desc, fvSubnetAttr)

	err := aciClient.Save(fvSubnet)
	if err != nil {
		return err
	}

	d.SetId(fvSubnet.DistinguishedName)
	return resourceAciSubnetRead(d, m)
}

func resourceAciSubnetUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	ip := d.Get("ip").(string)
	fvBDDn := d.Get("fv_bd_dn").(string)

	fvSubnetAttr := models.SubnetAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		fvSubnetAttr.Annotation = Annotation.(string)
	}
	if Ctrl, ok := d.GetOk("ctrl"); ok {
		fvSubnetAttr.Ctrl = Ctrl.(string)
	}

	if NameAlias, ok := d.GetOk("name_alias"); ok {
		fvSubnetAttr.NameAlias = NameAlias.(string)
	}
	if Preferred, ok := d.GetOk("preferred"); ok {
		fvSubnetAttr.Preferred = Preferred.(string)
	}
	if Scope, ok := d.GetOk("scope"); ok {
		fvSubnetAttr.Scope = Scope.(string)
	}
	if Virtual, ok := d.GetOk("virtual"); ok {
		fvSubnetAttr.Virtual = Virtual.(string)
	}
	fvSubnet := models.NewSubnet(fmt.Sprintf("subnet-[%s]", ip), fvBDDn, desc, fvSubnetAttr)

	fvSubnet.Status = "modified"

	err := aciClient.Save(fvSubnet)

	if err != nil {
		return err
	}

	d.SetId(fvSubnet.DistinguishedName)
	return resourceAciSubnetRead(d, m)

}

func resourceAciSubnetRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	fvSubnet, err := getRemoteSubnet(aciClient, dn)

	if err != nil {
		return err
	}
	setSubnetAttributes(fvSubnet, d)
	return nil
}

func resourceAciSubnetDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "fvSubnet")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
