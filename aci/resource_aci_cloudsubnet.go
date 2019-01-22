package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciCloudSubnet() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciCloudSubnetCreate,
		Update: resourceAciCloudSubnetUpdate,
		Read:   resourceAciCloudSubnetRead,
		Delete: resourceAciCloudSubnetDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudSubnetImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"cloud_cidr_pool_dn": &schema.Schema{
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
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"name_alias": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"scope": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "capability domain",
			},

			"usage": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "usage of the port",
			},

			"relation_cloud_rs_zone_attach": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to cloudZone",
			},
			"relation_cloud_rs_subnet_to_flow_log": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to cloudAwsFlowLogPol",
			},
		}),
	}
}

func getRemoteCloudSubnet(client *client.Client, dn string) (*models.CloudSubnet, error) {
	cloudSubnetCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudSubnet := models.CloudSubnetFromContainer(cloudSubnetCont)

	if cloudSubnet.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudSubnet.DistinguishedName)
	}

	return cloudSubnet, nil
}

func setCloudSubnetAttributes(cloudSubnet *models.CloudSubnet, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudSubnet.DistinguishedName)
	d.Set("description", cloudSubnet.Description)
	d.Set("cloud_cidr_pool_dn", GetParentDn(cloudSubnet.DistinguishedName))
	cloudSubnetMap, _ := cloudSubnet.ToMap()

	d.Set("annotation", cloudSubnetMap["annotation"])
	d.Set("ip", cloudSubnetMap["ip"])
	d.Set("name_alias", cloudSubnetMap["nameAlias"])
	d.Set("scope", cloudSubnetMap["scope"])
	d.Set("usage", cloudSubnetMap["usage"])
	return d
}

func resourceAciCloudSubnetImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudSubnet, err := getRemoteCloudSubnet(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCloudSubnetAttributes(cloudSubnet, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudSubnetCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	ip := d.Get("ip").(string)

	CloudCIDRPoolDn := d.Get("cloud_cidr_pool_dn").(string)

	cloudSubnetAttr := models.CloudSubnetAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudSubnetAttr.Annotation = Annotation.(string)
	}
	if Ip, ok := d.GetOk("ip"); ok {
		cloudSubnetAttr.Ip = Ip.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudSubnetAttr.NameAlias = NameAlias.(string)
	}
	if Scope, ok := d.GetOk("scope"); ok {
		cloudSubnetAttr.Scope = Scope.(string)
	}
	if Usage, ok := d.GetOk("usage"); ok {
		cloudSubnetAttr.Usage = Usage.(string)
	}
	cloudSubnet := models.NewCloudSubnet(fmt.Sprintf("subnet-[%s]", ip), CloudCIDRPoolDn, desc, cloudSubnetAttr)

	err := aciClient.Save(cloudSubnet)
	if err != nil {
		return err
	}

	if relationTocloudRsZoneAttach, ok := d.GetOk("relation_cloud_rs_zone_attach"); ok {
		relationParam := relationTocloudRsZoneAttach.(string)
		err = aciClient.CreateRelationcloudRsZoneAttachFromCloudSubnet(cloudSubnet.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTocloudRsSubnetToFlowLog, ok := d.GetOk("relation_cloud_rs_subnet_to_flow_log"); ok {
		relationParam := relationTocloudRsSubnetToFlowLog.(string)
		err = aciClient.CreateRelationcloudRsSubnetToFlowLogFromCloudSubnet(cloudSubnet.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}

	d.SetId(cloudSubnet.DistinguishedName)
	return resourceAciCloudSubnetRead(d, m)
}

func resourceAciCloudSubnetUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	ip := d.Get("ip").(string)

	CloudCIDRPoolDn := d.Get("cloud_cidr_pool_dn").(string)

	cloudSubnetAttr := models.CloudSubnetAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudSubnetAttr.Annotation = Annotation.(string)
	}
	if Ip, ok := d.GetOk("ip"); ok {
		cloudSubnetAttr.Ip = Ip.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudSubnetAttr.NameAlias = NameAlias.(string)
	}
	if Scope, ok := d.GetOk("scope"); ok {
		cloudSubnetAttr.Scope = Scope.(string)
	}
	if Usage, ok := d.GetOk("usage"); ok {
		cloudSubnetAttr.Usage = Usage.(string)
	}
	cloudSubnet := models.NewCloudSubnet(fmt.Sprintf("subnet-[%s]", ip), CloudCIDRPoolDn, desc, cloudSubnetAttr)

	cloudSubnet.Status = "modified"

	err := aciClient.Save(cloudSubnet)

	if err != nil {
		return err
	}

	if d.HasChange("relation_cloud_rs_zone_attach") {
		_, newRelParam := d.GetChange("relation_cloud_rs_zone_attach")
		err = aciClient.DeleteRelationcloudRsZoneAttachFromCloudSubnet(cloudSubnet.DistinguishedName)
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationcloudRsZoneAttachFromCloudSubnet(cloudSubnet.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_cloud_rs_subnet_to_flow_log") {
		_, newRelParam := d.GetChange("relation_cloud_rs_subnet_to_flow_log")
		err = aciClient.DeleteRelationcloudRsSubnetToFlowLogFromCloudSubnet(cloudSubnet.DistinguishedName)
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationcloudRsSubnetToFlowLogFromCloudSubnet(cloudSubnet.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}

	d.SetId(cloudSubnet.DistinguishedName)
	return resourceAciCloudSubnetRead(d, m)

}

func resourceAciCloudSubnetRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudSubnet, err := getRemoteCloudSubnet(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudSubnetAttributes(cloudSubnet, d)
	return nil
}

func resourceAciCloudSubnetDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudSubnet")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
