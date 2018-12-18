package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
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
			"bridge_domain_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"ctrl": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "subnet control state",
			},

			"name_alias": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"preferred": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "subnet preferred status",
			},

			"scope": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "subnet visibility",
			},

			"virtual": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"relation_fv_rs_bd_subnet_to_profile": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to rtctrlProfile",
			},
			"relation_fv_rs_bd_subnet_to_out": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Create relation to l3extOut",
				Set:         schema.HashString,
			},
			"relation_fv_rs_nd_pfx_pol": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to ndPfxPol",
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
	d.Set("bridge_domain_dn", GetParentDn(fvSubnet.DistinguishedName))
	fvSubnetMap, _ := fvSubnet.ToMap()

	d.Set("ctrl", fvSubnetMap["ctrl"])
	d.Set("ip", fvSubnetMap["ip"])
	d.Set("name_alias", fvSubnetMap["nameAlias"])
	d.Set("preferred", fvSubnetMap["preferred"])
	d.Set("scope", fvSubnetMap["scope"])
	d.Set("virtual", fvSubnetMap["virtual"])
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
	BridgeDomainDn := d.Get("bridge_domain_dn").(string)

	fvSubnetAttr := models.SubnetAttributes{}
	if Ctrl, ok := d.GetOk("ctrl"); ok {
		fvSubnetAttr.Ctrl = Ctrl.(string)
	}
	if Ip, ok := d.GetOk("ip"); ok {
		fvSubnetAttr.Ip = Ip.(string)
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
	fvSubnet := models.NewSubnet(fmt.Sprintf("subnet-[%s]", ip), BridgeDomainDn, desc, fvSubnetAttr)

	err := aciClient.Save(fvSubnet)
	if err != nil {
		return err
	}

	if relationTofvRsBDSubnetToProfile, ok := d.GetOk("relation_fv_rs_bd_subnet_to_profile"); ok {
		relationParam := relationTofvRsBDSubnetToProfile.(string)
		err = aciClient.CreateRelationfvRsBDSubnetToProfileFromSubnet(fvSubnet.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsBDSubnetToOut, ok := d.GetOk("relation_fv_rs_bd_subnet_to_out"); ok {
		relationParamList := toStringList(relationTofvRsBDSubnetToOut.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsBDSubnetToOutFromSubnet(fvSubnet.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsNdPfxPol, ok := d.GetOk("relation_fv_rs_nd_pfx_pol"); ok {
		relationParam := relationTofvRsNdPfxPol.(string)
		err = aciClient.CreateRelationfvRsNdPfxPolFromSubnet(fvSubnet.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}

	d.SetId(fvSubnet.DistinguishedName)
	return resourceAciSubnetRead(d, m)
}

func resourceAciSubnetUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	ip := d.Get("ip").(string)
	BridgeDomainDn := d.Get("bridge_domain_dn").(string)

	fvSubnetAttr := models.SubnetAttributes{}
	if Ctrl, ok := d.GetOk("ctrl"); ok {
		fvSubnetAttr.Ctrl = Ctrl.(string)
	}
	if Ip, ok := d.GetOk("ip"); ok {
		fvSubnetAttr.Ip = Ip.(string)
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
	fvSubnet := models.NewSubnet(fmt.Sprintf("subnet-[%s]", ip), BridgeDomainDn, desc, fvSubnetAttr)

	fvSubnet.Status = "modified"

	err := aciClient.Save(fvSubnet)

	if err != nil {
		return err
	}

	if d.HasChange("relation_fv_rs_bd_subnet_to_profile") {
		_, newRelParam := d.GetChange("relation_fv_rs_bd_subnet_to_profile")
		err = aciClient.DeleteRelationfvRsBDSubnetToProfileFromSubnet(fvSubnet.DistinguishedName)
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationfvRsBDSubnetToProfileFromSubnet(fvSubnet.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_fv_rs_bd_subnet_to_out") {
		oldRel, newRel := d.GetChange("relation_fv_rs_bd_subnet_to_out")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsBDSubnetToOutFromSubnet(fvSubnet.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsBDSubnetToOutFromSubnet(fvSubnet.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_fv_rs_nd_pfx_pol") {
		_, newRelParam := d.GetChange("relation_fv_rs_nd_pfx_pol")
		err = aciClient.DeleteRelationfvRsNdPfxPolFromSubnet(fvSubnet.DistinguishedName)
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationfvRsNdPfxPolFromSubnet(fvSubnet.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

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
