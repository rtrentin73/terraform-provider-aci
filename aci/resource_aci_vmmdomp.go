package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciVMMDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciVMMDomainCreate,
		Update: resourceAciVMMDomainUpdate,
		Read:   resourceAciVMMDomainRead,
		Delete: resourceAciVMMDomainDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciVMMDomainImport,
		},

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

			"access_mode": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"arp_learning": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"ctrl_knob": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"delimiter": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"enable_ave": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"encap_mode": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"enf_pref": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "switching enforcement preference",
			},

			"ep_ret_time": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"mcast_addr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "multicast address",
			},

			"mode": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "switch used for the domain profile",
			},

			"name_alias": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"pref_encap_mode": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"relation_infra_rs_vlan_ns": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to fvnsVlanInstP",
			},
			"relation_vmm_rs_dom_mcast_addr_ns": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to fvnsMcastAddrInstP",
			},
			"relation_vmm_rs_default_cdp_if_pol": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to cdpIfPol",
			},
			"relation_vmm_rs_default_lacp_lag_pol": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to lacpLagPol",
			},
			"relation_infra_rs_vlan_ns_def": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to fvnsAInstP",
			},
			"relation_infra_rs_vip_addr_ns": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to fvnsAddrInst",
			},
			"relation_vmm_rs_default_lldp_if_pol": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to lldpIfPol",
			},
			"relation_vmm_rs_default_l2_inst_pol": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to l2InstPol",
			},
			"relation_vmm_rs_default_stp_if_pol": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to stpIfPol",
			},
			"relation_infra_rs_dom_vxlan_ns_def": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to fvnsAInstP",
			},
			"relation_vmm_rs_default_fw_pol": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to nwsFwPol",
			},
		}),
	}
}

func getRemoteVMMDomain(client *client.Client, dn string) (*models.VMMDomain, error) {
	vmmDomPCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vmmDomP := models.VMMDomainFromContainer(vmmDomPCont)

	if vmmDomP.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", vmmDomP.DistinguishedName)
	}

	return vmmDomP, nil
}

func setVMMDomainAttributes(vmmDomP *models.VMMDomain, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vmmDomP.DistinguishedName)
	d.Set("description", vmmDomP.Description)
	d.Set("provider_profile_dn", GetParentDn(vmmDomP.DistinguishedName))
	vmmDomPMap, _ := vmmDomP.ToMap()

	d.Set("access_mode", vmmDomPMap["accessMode"])
	d.Set("arp_learning", vmmDomPMap["arpLearning"])
	d.Set("ctrl_knob", vmmDomPMap["ctrlKnob"])
	d.Set("delimiter", vmmDomPMap["delimiter"])
	d.Set("enable_ave", vmmDomPMap["enableAVE"])
	d.Set("encap_mode", vmmDomPMap["encapMode"])
	d.Set("enf_pref", vmmDomPMap["enfPref"])
	d.Set("ep_ret_time", vmmDomPMap["epRetTime"])
	d.Set("mcast_addr", vmmDomPMap["mcastAddr"])
	d.Set("mode", vmmDomPMap["mode"])
	d.Set("name_alias", vmmDomPMap["nameAlias"])
	d.Set("pref_encap_mode", vmmDomPMap["prefEncapMode"])
	return d
}

func resourceAciVMMDomainImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	vmmDomP, err := getRemoteVMMDomain(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setVMMDomainAttributes(vmmDomP, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciVMMDomainCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	name := d.Get("name").(string)
	ProviderProfileDn := d.Get("provider_profile_dn").(string)

	vmmDomPAttr := models.VMMDomainAttributes{}
	if AccessMode, ok := d.GetOk("access_mode"); ok {
		vmmDomPAttr.AccessMode = AccessMode.(string)
	}
	if ArpLearning, ok := d.GetOk("arp_learning"); ok {
		vmmDomPAttr.ArpLearning = ArpLearning.(string)
	}
	if CtrlKnob, ok := d.GetOk("ctrl_knob"); ok {
		vmmDomPAttr.CtrlKnob = CtrlKnob.(string)
	}
	if Delimiter, ok := d.GetOk("delimiter"); ok {
		vmmDomPAttr.Delimiter = Delimiter.(string)
	}
	if EnableAVE, ok := d.GetOk("enable_ave"); ok {
		vmmDomPAttr.EnableAVE = EnableAVE.(string)
	}
	if EncapMode, ok := d.GetOk("encap_mode"); ok {
		vmmDomPAttr.EncapMode = EncapMode.(string)
	}
	if EnfPref, ok := d.GetOk("enf_pref"); ok {
		vmmDomPAttr.EnfPref = EnfPref.(string)
	}
	if EpRetTime, ok := d.GetOk("ep_ret_time"); ok {
		vmmDomPAttr.EpRetTime = EpRetTime.(string)
	}
	if McastAddr, ok := d.GetOk("mcast_addr"); ok {
		vmmDomPAttr.McastAddr = McastAddr.(string)
	}
	if Mode, ok := d.GetOk("mode"); ok {
		vmmDomPAttr.Mode = Mode.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		vmmDomPAttr.NameAlias = NameAlias.(string)
	}
	if PrefEncapMode, ok := d.GetOk("pref_encap_mode"); ok {
		vmmDomPAttr.PrefEncapMode = PrefEncapMode.(string)
	}
	vmmDomP := models.NewVMMDomain(fmt.Sprintf("dom-%s", name), ProviderProfileDn, desc, vmmDomPAttr)

	err := aciClient.Save(vmmDomP)
	if err != nil {
		return err
	}

	if relationToinfraRsVlanNs, ok := d.GetOk("relation_infra_rs_vlan_ns"); ok {
		relationParam := relationToinfraRsVlanNs.(string)
		err = aciClient.CreateRelationinfraRsVlanNsFromVMMDomain(vmmDomP.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTovmmRsDomMcastAddrNs, ok := d.GetOk("relation_vmm_rs_dom_mcast_addr_ns"); ok {
		relationParam := relationTovmmRsDomMcastAddrNs.(string)
		err = aciClient.CreateRelationvmmRsDomMcastAddrNsFromVMMDomain(vmmDomP.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTovmmRsDefaultCdpIfPol, ok := d.GetOk("relation_vmm_rs_default_cdp_if_pol"); ok {
		relationParam := relationTovmmRsDefaultCdpIfPol.(string)
		err = aciClient.CreateRelationvmmRsDefaultCdpIfPolFromVMMDomain(vmmDomP.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTovmmRsDefaultLacpLagPol, ok := d.GetOk("relation_vmm_rs_default_lacp_lag_pol"); ok {
		relationParam := relationTovmmRsDefaultLacpLagPol.(string)
		err = aciClient.CreateRelationvmmRsDefaultLacpLagPolFromVMMDomain(vmmDomP.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationToinfraRsVlanNsDef, ok := d.GetOk("relation_infra_rs_vlan_ns_def"); ok {
		relationParam := relationToinfraRsVlanNsDef.(string)
		err = aciClient.CreateRelationinfraRsVlanNsDefFromVMMDomain(vmmDomP.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationToinfraRsVipAddrNs, ok := d.GetOk("relation_infra_rs_vip_addr_ns"); ok {
		relationParam := relationToinfraRsVipAddrNs.(string)
		err = aciClient.CreateRelationinfraRsVipAddrNsFromVMMDomain(vmmDomP.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTovmmRsDefaultLldpIfPol, ok := d.GetOk("relation_vmm_rs_default_lldp_if_pol"); ok {
		relationParam := relationTovmmRsDefaultLldpIfPol.(string)
		err = aciClient.CreateRelationvmmRsDefaultLldpIfPolFromVMMDomain(vmmDomP.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTovmmRsDefaultL2InstPol, ok := d.GetOk("relation_vmm_rs_default_l2_inst_pol"); ok {
		relationParam := relationTovmmRsDefaultL2InstPol.(string)
		err = aciClient.CreateRelationvmmRsDefaultL2InstPolFromVMMDomain(vmmDomP.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTovmmRsDefaultStpIfPol, ok := d.GetOk("relation_vmm_rs_default_stp_if_pol"); ok {
		relationParam := relationTovmmRsDefaultStpIfPol.(string)
		err = aciClient.CreateRelationvmmRsDefaultStpIfPolFromVMMDomain(vmmDomP.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationToinfraRsDomVxlanNsDef, ok := d.GetOk("relation_infra_rs_dom_vxlan_ns_def"); ok {
		relationParam := relationToinfraRsDomVxlanNsDef.(string)
		err = aciClient.CreateRelationinfraRsDomVxlanNsDefFromVMMDomain(vmmDomP.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTovmmRsDefaultFwPol, ok := d.GetOk("relation_vmm_rs_default_fw_pol"); ok {
		relationParam := relationTovmmRsDefaultFwPol.(string)
		err = aciClient.CreateRelationvmmRsDefaultFwPolFromVMMDomain(vmmDomP.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}

	d.SetId(vmmDomP.DistinguishedName)
	return resourceAciVMMDomainRead(d, m)
}

func resourceAciVMMDomainUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)
	ProviderProfileDn := d.Get("provider_profile_dn").(string)

	vmmDomPAttr := models.VMMDomainAttributes{}
	if AccessMode, ok := d.GetOk("access_mode"); ok {
		vmmDomPAttr.AccessMode = AccessMode.(string)
	}
	if ArpLearning, ok := d.GetOk("arp_learning"); ok {
		vmmDomPAttr.ArpLearning = ArpLearning.(string)
	}
	if CtrlKnob, ok := d.GetOk("ctrl_knob"); ok {
		vmmDomPAttr.CtrlKnob = CtrlKnob.(string)
	}
	if Delimiter, ok := d.GetOk("delimiter"); ok {
		vmmDomPAttr.Delimiter = Delimiter.(string)
	}
	if EnableAVE, ok := d.GetOk("enable_ave"); ok {
		vmmDomPAttr.EnableAVE = EnableAVE.(string)
	}
	if EncapMode, ok := d.GetOk("encap_mode"); ok {
		vmmDomPAttr.EncapMode = EncapMode.(string)
	}
	if EnfPref, ok := d.GetOk("enf_pref"); ok {
		vmmDomPAttr.EnfPref = EnfPref.(string)
	}
	if EpRetTime, ok := d.GetOk("ep_ret_time"); ok {
		vmmDomPAttr.EpRetTime = EpRetTime.(string)
	}
	if McastAddr, ok := d.GetOk("mcast_addr"); ok {
		vmmDomPAttr.McastAddr = McastAddr.(string)
	}
	if Mode, ok := d.GetOk("mode"); ok {
		vmmDomPAttr.Mode = Mode.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		vmmDomPAttr.NameAlias = NameAlias.(string)
	}
	if PrefEncapMode, ok := d.GetOk("pref_encap_mode"); ok {
		vmmDomPAttr.PrefEncapMode = PrefEncapMode.(string)
	}
	vmmDomP := models.NewVMMDomain(fmt.Sprintf("dom-%s", name), ProviderProfileDn, desc, vmmDomPAttr)

	vmmDomP.Status = "modified"

	err := aciClient.Save(vmmDomP)

	if err != nil {
		return err
	}

	if d.HasChange("relation_infra_rs_vlan_ns") {
		_, newRelParam := d.GetChange("relation_infra_rs_vlan_ns")
		err = aciClient.DeleteRelationinfraRsVlanNsFromVMMDomain(vmmDomP.DistinguishedName)
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationinfraRsVlanNsFromVMMDomain(vmmDomP.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_vmm_rs_dom_mcast_addr_ns") {
		_, newRelParam := d.GetChange("relation_vmm_rs_dom_mcast_addr_ns")
		err = aciClient.DeleteRelationvmmRsDomMcastAddrNsFromVMMDomain(vmmDomP.DistinguishedName)
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationvmmRsDomMcastAddrNsFromVMMDomain(vmmDomP.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_vmm_rs_default_cdp_if_pol") {
		_, newRelParam := d.GetChange("relation_vmm_rs_default_cdp_if_pol")
		err = aciClient.CreateRelationvmmRsDefaultCdpIfPolFromVMMDomain(vmmDomP.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_vmm_rs_default_lacp_lag_pol") {
		_, newRelParam := d.GetChange("relation_vmm_rs_default_lacp_lag_pol")
		err = aciClient.CreateRelationvmmRsDefaultLacpLagPolFromVMMDomain(vmmDomP.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_infra_rs_vlan_ns_def") {
		_, newRelParam := d.GetChange("relation_infra_rs_vlan_ns_def")
		err = aciClient.CreateRelationinfraRsVlanNsDefFromVMMDomain(vmmDomP.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_infra_rs_vip_addr_ns") {
		_, newRelParam := d.GetChange("relation_infra_rs_vip_addr_ns")
		err = aciClient.DeleteRelationinfraRsVipAddrNsFromVMMDomain(vmmDomP.DistinguishedName)
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationinfraRsVipAddrNsFromVMMDomain(vmmDomP.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_vmm_rs_default_lldp_if_pol") {
		_, newRelParam := d.GetChange("relation_vmm_rs_default_lldp_if_pol")
		err = aciClient.CreateRelationvmmRsDefaultLldpIfPolFromVMMDomain(vmmDomP.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_vmm_rs_default_l2_inst_pol") {
		_, newRelParam := d.GetChange("relation_vmm_rs_default_l2_inst_pol")
		err = aciClient.CreateRelationvmmRsDefaultL2InstPolFromVMMDomain(vmmDomP.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_vmm_rs_default_stp_if_pol") {
		_, newRelParam := d.GetChange("relation_vmm_rs_default_stp_if_pol")
		err = aciClient.CreateRelationvmmRsDefaultStpIfPolFromVMMDomain(vmmDomP.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_infra_rs_dom_vxlan_ns_def") {
		_, newRelParam := d.GetChange("relation_infra_rs_dom_vxlan_ns_def")
		err = aciClient.CreateRelationinfraRsDomVxlanNsDefFromVMMDomain(vmmDomP.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_vmm_rs_default_fw_pol") {
		_, newRelParam := d.GetChange("relation_vmm_rs_default_fw_pol")
		err = aciClient.CreateRelationvmmRsDefaultFwPolFromVMMDomain(vmmDomP.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}

	d.SetId(vmmDomP.DistinguishedName)
	return resourceAciVMMDomainRead(d, m)

}

func resourceAciVMMDomainRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	vmmDomP, err := getRemoteVMMDomain(aciClient, dn)

	if err != nil {
		return err
	}
	setVMMDomainAttributes(vmmDomP, d)
	return nil
}

func resourceAciVMMDomainDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "vmmDomP")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
