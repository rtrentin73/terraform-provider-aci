package aci

import (
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAciVMMDomain() *schema.Resource {
	return &schema.Resource{
		Read:          dataSourceAciVMMDomainRead,
		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"provider_profile_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		}),
	}
}

func dataSourceAciVMMDomainRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	name := d.Get("name").(string)
	ProviderProfileDn := d.Get("provider_profile_dn").(string)
	vmmDomP, err := aciClient.ReadVMMDomain(name, ProviderProfileDn)

	if err != nil {
		return err
	}
	setVMMDomainAttributes(vmmDomP, d)
	return nil
}
