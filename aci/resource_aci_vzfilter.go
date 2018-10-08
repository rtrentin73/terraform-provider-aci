package aci

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciFilterCreate,
		Update: resourceAciFilterUpdate,
		Read:   resourceAciFilterRead,
		Delete: resourceAciFilterDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciFilterImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"fv_tenant_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

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

func getRemoteFilter(client *client.Client, dn string) (*models.Filter, error) {
	vzFilterCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vzFilter := models.FilterFromContainer(vzFilterCont)

	if vzFilter.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", vzFilter.DistinguishedName)
	}

	return vzFilter, nil
}

func setFilterAttributes(vzFilter *models.Filter, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vzFilter.DistinguishedName)
	d.Set("description", vzFilter.Description)
	d.Set("fv_tenant_dn", GetParentDn(vzFilter.DistinguishedName))
	vzFilter_map, _ := vzFilter.ToMap()

	d.Set("name_alias", vzFilter_map["nameAlias"])
	return d
}

func resourceAciFilterImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	vzFilter, err := getRemoteFilter(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setFilterAttributes(vzFilter, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciFilterCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	name := d.Get("name").(string)
	fvTenantDn := d.Get("fv_tenant_dn").(string)

	vzFilterAttr := models.FilterAttributes{}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		vzFilterAttr.NameAlias = NameAlias.(string)
	}
	vzFilter := models.NewFilter(fmt.Sprintf("flt-%s", name), fvTenantDn, desc, vzFilterAttr)

	err := aciClient.Save(vzFilter)
	if err != nil {
		return err
	}

	d.SetId(vzFilter.DistinguishedName)
	return resourceAciFilterRead(d, m)
}

func resourceAciFilterUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)
	fvTenantDn := d.Get("fv_tenant_dn").(string)

	vzFilterAttr := models.FilterAttributes{}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		vzFilterAttr.NameAlias = NameAlias.(string)
	}
	vzFilter := models.NewFilter(fmt.Sprintf("flt-%s", name), fvTenantDn, desc, vzFilterAttr)

	vzFilter.Status = "modified"

	err := aciClient.Save(vzFilter)

	if err != nil {
		return err
	}

	d.SetId(vzFilter.DistinguishedName)
	return resourceAciFilterRead(d, m)

}

func resourceAciFilterRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	vzFilter, err := getRemoteFilter(aciClient, dn)

	if err != nil {
		return err
	}
	setFilterAttributes(vzFilter, d)
	return nil
}

func resourceAciFilterDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "vzFilter")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
