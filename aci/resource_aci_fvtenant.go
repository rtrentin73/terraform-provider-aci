package aci

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciTenant() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciTenantCreate,
		Update: resourceAciTenantUpdate,
		Read:   resourceAciTenantRead,
		Delete: resourceAciTenantDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciTenantImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func getRemoteTenant(client *client.Client, dn string) (*models.Tenant, error) {
	fvTenantCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	fvTenant := models.TenantFromContainer(fvTenantCont)

	if fvTenant.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", fvTenant.DistinguishedName)
	}

	return fvTenant, nil
}

func setTenantAttributes(fvTenant *models.Tenant, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(fvTenant.DistinguishedName)
	d.Set("description", fvTenant.Description)
	fvTenant_map, _ := fvTenant.ToMap()

	d.Set("name_alias", fvTenant_map["nameAlias"])
	return d
}

func resourceAciTenantImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	fvTenant, err := getRemoteTenant(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setTenantAttributes(fvTenant, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciTenantCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	name := d.Get("name").(string)
	fvTenantAttr := models.TenantAttributes{}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		fvTenantAttr.NameAlias = NameAlias.(string)
	}
	fvTenant := models.NewTenant(fmt.Sprintf("tn-%s", name), "uni", desc, fvTenantAttr)

	err := aciClient.Save(fvTenant)
	if err != nil {
		return err
	}

	d.SetId(fvTenant.DistinguishedName)
	return resourceAciTenantRead(d, m)
}

func resourceAciTenantUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	fvTenantAttr := models.TenantAttributes{}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		fvTenantAttr.NameAlias = NameAlias.(string)
	}
	fvTenant := models.NewTenant(fmt.Sprintf("tn-%s", name), "uni", desc, fvTenantAttr)

	fvTenant.Status = "modified"

	err := aciClient.Save(fvTenant)

	if err != nil {
		return err
	}

	d.SetId(fvTenant.DistinguishedName)
	return resourceAciTenantRead(d, m)

}

func resourceAciTenantRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	fvTenant, err := getRemoteTenant(aciClient, dn)

	if err != nil {
		return err
	}
	setTenantAttributes(fvTenant, d)
	return nil
}

func resourceAciTenantDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "fvTenant")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
