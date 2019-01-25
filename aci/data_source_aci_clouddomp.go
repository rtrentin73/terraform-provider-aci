package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAciCloudDomainProfile() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceAciCloudDomainProfileRead,

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{}),
	}
}

func dataSourceAciCloudDomainProfileRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	rn := fmt.Sprintf("clouddomp")

	dn := fmt.Sprintf("uni/%s", rn)

	cloudDomP, err := getRemoteCloudDomainProfile(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudDomainProfileAttributes(cloudDomP, d)
	return nil
}
