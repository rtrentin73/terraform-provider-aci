package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAciCloudExternalEPg() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceAciCloudExternalEPgRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"cloud_applicationcontainer_dn": &schema.Schema{
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

func dataSourceAciCloudExternalEPgRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	name := d.Get("name").(string)

	rn := fmt.Sprintf("cloudextepg-%s", name)
	CloudApplicationcontainerDn := d.Get("cloud_applicationcontainer_dn").(string)

	dn := fmt.Sprintf("%s/%s", CloudApplicationcontainerDn, rn)

	cloudExtEPg, err := getRemoteCloudExternalEPg(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudExternalEPgAttributes(cloudExtEPg, d)
	return nil
}
