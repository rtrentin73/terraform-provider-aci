package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciBridgedomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciBridgedomainCreate,
		Update: resourceAciBridgedomainUpdate,
		Read:   resourceAciBridgedomainRead,
		Delete: resourceAciBridgedomainDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciBridgedomainImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"tenant_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"optimize_wan_bandwidth": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"annotation": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"arp_flood": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "arp flood enable",
			},

			"ep_clear": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"ep_move_detect_mode": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "ep move detection garp based mode",
			},

			"host_based_routing": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "enables advertising host routes out of l3outs of this BD",
			},

			"intersite_bum_traffic_allow": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "",
			},

			"intersite_l2_stretch": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "",
			},

			"ip_learning": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Endpoint Dataplane Learning",
			},

			"ipv6_mcast_allow": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"limit_ip_learn_to_subnets": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "limits ip learning to bd subnets only",
			},

			"ll_addr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "override of system generated ipv6 link-local address",
			},

			"mac": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "mac address",
			},

			"mcast_allow": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"multi_dst_pkt_act": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "forwarding method for multi destinations",
			},

			"name_alias": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "component type",
			},

			"unicast_route": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Unicast routing",
			},

			"unk_mac_ucast_act": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "forwarding method for l2 destinations",
			},

			"unk_mcast_act": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "parameter used by node to forward data",
			},

			"v6unk_mcast_act": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"vmac": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"relation_fv_rs_bd_to_profile": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to rtctrlProfile",
			},
			"relation_fv_rs_mldsn": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to mldSnoopPol",
			},
			"relation_fv_rs_abd_pol_mon_pol": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to monEPGPol",
			},
			"relation_fv_rs_bd_to_nd_p": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to ndIfPol",
			},
			"relation_fv_rs_bd_flood_to": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Create relation to vzFilter",
				Set:         schema.HashString,
			},
			"relation_fv_rs_bd_to_fhs": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to fhsBDPol",
			},
			"relation_fv_rs_bd_to_relay_p": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to dhcpRelayP",
			},
			"relation_fv_rs_ctx": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to fvCtx",
			},
			"relation_fv_rs_bd_to_netflow_monitor_pol": &schema.Schema{
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Create relation to netflowMonitorPol",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tn_netflow_monitor_pol_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"flt_type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"relation_fv_rs_igmpsn": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to igmpSnoopPol",
			},
			"relation_fv_rs_bd_to_ep_ret": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to fvEpRetPol",
			},
			"relation_fv_rs_bd_to_out": &schema.Schema{
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Create relation to l3extOut",
				Set:         schema.HashString,
			},
		}),
	}
}

func getRemoteBridgedomain(client *client.Client, dn string) (*models.Bridgedomain, error) {
	fvBDCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	fvBD := models.BridgedomainFromContainer(fvBDCont)

	if fvBD.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", fvBD.DistinguishedName)
	}

	return fvBD, nil
}

func setBridgedomainAttributes(fvBD *models.Bridgedomain, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(fvBD.DistinguishedName)
	d.Set("description", fvBD.Description)
	d.Set("tenant_dn", GetParentDn(fvBD.DistinguishedName))
	fvBDMap, _ := fvBD.ToMap()

	d.Set("optimize_wan_bandwidth", fvBDMap["OptimizeWanBandwidth"])
	d.Set("annotation", fvBDMap["annotation"])
	d.Set("arp_flood", fvBDMap["arpFlood"])
	d.Set("ep_clear", fvBDMap["epClear"])
	d.Set("ep_move_detect_mode", fvBDMap["epMoveDetectMode"])
	d.Set("host_based_routing", fvBDMap["hostBasedRouting"])
	d.Set("intersite_bum_traffic_allow", fvBDMap["intersiteBumTrafficAllow"])
	d.Set("intersite_l2_stretch", fvBDMap["intersiteL2Stretch"])
	d.Set("ip_learning", fvBDMap["ipLearning"])
	d.Set("ipv6_mcast_allow", fvBDMap["ipv6McastAllow"])
	d.Set("limit_ip_learn_to_subnets", fvBDMap["limitIpLearnToSubnets"])
	d.Set("ll_addr", fvBDMap["llAddr"])
	d.Set("mac", fvBDMap["mac"])
	d.Set("mcast_allow", fvBDMap["mcastAllow"])
	d.Set("multi_dst_pkt_act", fvBDMap["multiDstPktAct"])
	d.Set("name_alias", fvBDMap["nameAlias"])
	d.Set("type", fvBDMap["type"])
	d.Set("unicast_route", fvBDMap["unicastRoute"])
	d.Set("unk_mac_ucast_act", fvBDMap["unkMacUcastAct"])
	d.Set("unk_mcast_act", fvBDMap["unkMcastAct"])
	d.Set("v6unk_mcast_act", fvBDMap["v6unkMcastAct"])
	d.Set("vmac", fvBDMap["vmac"])
	return d
}

func resourceAciBridgedomainImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	fvBD, err := getRemoteBridgedomain(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setBridgedomainAttributes(fvBD, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciBridgedomainCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	TenantDn := d.Get("tenant_dn").(string)

	fvBDAttr := models.BridgedomainAttributes{}
	if OptimizeWanBandwidth, ok := d.GetOk("optimize_wan_bandwidth"); ok {
		fvBDAttr.OptimizeWanBandwidth = OptimizeWanBandwidth.(string)
	}
	if Annotation, ok := d.GetOk("annotation"); ok {
		fvBDAttr.Annotation = Annotation.(string)
	}
	if ArpFlood, ok := d.GetOk("arp_flood"); ok {
		fvBDAttr.ArpFlood = ArpFlood.(string)
	}
	if EpClear, ok := d.GetOk("ep_clear"); ok {
		fvBDAttr.EpClear = EpClear.(string)
	}
	if EpMoveDetectMode, ok := d.GetOk("ep_move_detect_mode"); ok {
		fvBDAttr.EpMoveDetectMode = EpMoveDetectMode.(string)
	}
	if HostBasedRouting, ok := d.GetOk("host_based_routing"); ok {
		fvBDAttr.HostBasedRouting = HostBasedRouting.(string)
	}
	if IntersiteBumTrafficAllow, ok := d.GetOk("intersite_bum_traffic_allow"); ok {
		fvBDAttr.IntersiteBumTrafficAllow = IntersiteBumTrafficAllow.(string)
	}
	if IntersiteL2Stretch, ok := d.GetOk("intersite_l2_stretch"); ok {
		fvBDAttr.IntersiteL2Stretch = IntersiteL2Stretch.(string)
	}
	if IpLearning, ok := d.GetOk("ip_learning"); ok {
		fvBDAttr.IpLearning = IpLearning.(string)
	}
	if Ipv6McastAllow, ok := d.GetOk("ipv6_mcast_allow"); ok {
		fvBDAttr.Ipv6McastAllow = Ipv6McastAllow.(string)
	}
	if LimitIpLearnToSubnets, ok := d.GetOk("limit_ip_learn_to_subnets"); ok {
		fvBDAttr.LimitIpLearnToSubnets = LimitIpLearnToSubnets.(string)
	}
	if LlAddr, ok := d.GetOk("ll_addr"); ok {
		fvBDAttr.LlAddr = LlAddr.(string)
	}
	if Mac, ok := d.GetOk("mac"); ok {
		fvBDAttr.Mac = Mac.(string)
	}
	if McastAllow, ok := d.GetOk("mcast_allow"); ok {
		fvBDAttr.McastAllow = McastAllow.(string)
	}
	if MultiDstPktAct, ok := d.GetOk("multi_dst_pkt_act"); ok {
		fvBDAttr.MultiDstPktAct = MultiDstPktAct.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		fvBDAttr.NameAlias = NameAlias.(string)
	}
	if Type, ok := d.GetOk("type"); ok {
		fvBDAttr.Type = Type.(string)
	}
	if UnicastRoute, ok := d.GetOk("unicast_route"); ok {
		fvBDAttr.UnicastRoute = UnicastRoute.(string)
	}
	if UnkMacUcastAct, ok := d.GetOk("unk_mac_ucast_act"); ok {
		fvBDAttr.UnkMacUcastAct = UnkMacUcastAct.(string)
	}
	if UnkMcastAct, ok := d.GetOk("unk_mcast_act"); ok {
		fvBDAttr.UnkMcastAct = UnkMcastAct.(string)
	}
	if V6unkMcastAct, ok := d.GetOk("v6unk_mcast_act"); ok {
		fvBDAttr.V6unkMcastAct = V6unkMcastAct.(string)
	}
	if Vmac, ok := d.GetOk("vmac"); ok {
		fvBDAttr.Vmac = Vmac.(string)
	}
	fvBD := models.NewBridgedomain(fmt.Sprintf("BD-%s", name), TenantDn, desc, fvBDAttr)

	err := aciClient.Save(fvBD)
	if err != nil {
		return err
	}

	if relationTofvRsBDToProfile, ok := d.GetOk("relation_fv_rs_bd_to_profile"); ok {
		relationParam := relationTofvRsBDToProfile.(string)
		err = aciClient.CreateRelationfvRsBDToProfileFromBridgedomain(fvBD.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsMldsn, ok := d.GetOk("relation_fv_rs_mldsn"); ok {
		relationParam := relationTofvRsMldsn.(string)
		err = aciClient.CreateRelationfvRsMldsnFromBridgedomain(fvBD.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsABDPolMonPol, ok := d.GetOk("relation_fv_rs_abd_pol_mon_pol"); ok {
		relationParam := relationTofvRsABDPolMonPol.(string)
		err = aciClient.CreateRelationfvRsABDPolMonPolFromBridgedomain(fvBD.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsBDToNdP, ok := d.GetOk("relation_fv_rs_bd_to_nd_p"); ok {
		relationParam := relationTofvRsBDToNdP.(string)
		err = aciClient.CreateRelationfvRsBDToNdPFromBridgedomain(fvBD.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsBdFloodTo, ok := d.GetOk("relation_fv_rs_bd_flood_to"); ok {
		relationParamList := toStringList(relationTofvRsBdFloodTo.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsBdFloodToFromBridgedomain(fvBD.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}
	if relationTofvRsBDToFhs, ok := d.GetOk("relation_fv_rs_bd_to_fhs"); ok {
		relationParam := relationTofvRsBDToFhs.(string)
		err = aciClient.CreateRelationfvRsBDToFhsFromBridgedomain(fvBD.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsBDToRelayP, ok := d.GetOk("relation_fv_rs_bd_to_relay_p"); ok {
		relationParam := relationTofvRsBDToRelayP.(string)
		err = aciClient.CreateRelationfvRsBDToRelayPFromBridgedomain(fvBD.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsCtx, ok := d.GetOk("relation_fv_rs_ctx"); ok {
		relationParam := relationTofvRsCtx.(string)
		err = aciClient.CreateRelationfvRsCtxFromBridgedomain(fvBD.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsBDToNetflowMonitorPol, ok := d.GetOk("relation_fv_rs_bd_to_netflow_monitor_pol"); ok {

		relationParamList := relationTofvRsBDToNetflowMonitorPol.(*schema.Set).List()
		for _, relationParam := range relationParamList {
			paramMap := relationParam.(map[string]interface{})
			err = aciClient.CreateRelationfvRsBDToNetflowMonitorPolFromBridgedomain(fvBD.DistinguishedName, paramMap["tn_netflow_monitor_pol_name"].(string), paramMap["flt_type"].(string))
			if err != nil {
				return err
			}
		}

	}
	if relationTofvRsIgmpsn, ok := d.GetOk("relation_fv_rs_igmpsn"); ok {
		relationParam := relationTofvRsIgmpsn.(string)
		err = aciClient.CreateRelationfvRsIgmpsnFromBridgedomain(fvBD.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsBdToEpRet, ok := d.GetOk("relation_fv_rs_bd_to_ep_ret"); ok {
		relationParam := relationTofvRsBdToEpRet.(string)
		err = aciClient.CreateRelationfvRsBdToEpRetFromBridgedomain(fvBD.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTofvRsBDToOut, ok := d.GetOk("relation_fv_rs_bd_to_out"); ok {
		relationParamList := toStringList(relationTofvRsBDToOut.(*schema.Set).List())
		for _, relationParam := range relationParamList {
			err = aciClient.CreateRelationfvRsBDToOutFromBridgedomain(fvBD.DistinguishedName, relationParam)

			if err != nil {
				return err
			}
		}
	}

	d.SetId(fvBD.DistinguishedName)
	return resourceAciBridgedomainRead(d, m)
}

func resourceAciBridgedomainUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	TenantDn := d.Get("tenant_dn").(string)

	fvBDAttr := models.BridgedomainAttributes{}
	if OptimizeWanBandwidth, ok := d.GetOk("optimize_wan_bandwidth"); ok {
		fvBDAttr.OptimizeWanBandwidth = OptimizeWanBandwidth.(string)
	}
	if Annotation, ok := d.GetOk("annotation"); ok {
		fvBDAttr.Annotation = Annotation.(string)
	}
	if ArpFlood, ok := d.GetOk("arp_flood"); ok {
		fvBDAttr.ArpFlood = ArpFlood.(string)
	}
	if EpClear, ok := d.GetOk("ep_clear"); ok {
		fvBDAttr.EpClear = EpClear.(string)
	}
	if EpMoveDetectMode, ok := d.GetOk("ep_move_detect_mode"); ok {
		fvBDAttr.EpMoveDetectMode = EpMoveDetectMode.(string)
	}
	if HostBasedRouting, ok := d.GetOk("host_based_routing"); ok {
		fvBDAttr.HostBasedRouting = HostBasedRouting.(string)
	}
	if IntersiteBumTrafficAllow, ok := d.GetOk("intersite_bum_traffic_allow"); ok {
		fvBDAttr.IntersiteBumTrafficAllow = IntersiteBumTrafficAllow.(string)
	}
	if IntersiteL2Stretch, ok := d.GetOk("intersite_l2_stretch"); ok {
		fvBDAttr.IntersiteL2Stretch = IntersiteL2Stretch.(string)
	}
	if IpLearning, ok := d.GetOk("ip_learning"); ok {
		fvBDAttr.IpLearning = IpLearning.(string)
	}
	if Ipv6McastAllow, ok := d.GetOk("ipv6_mcast_allow"); ok {
		fvBDAttr.Ipv6McastAllow = Ipv6McastAllow.(string)
	}
	if LimitIpLearnToSubnets, ok := d.GetOk("limit_ip_learn_to_subnets"); ok {
		fvBDAttr.LimitIpLearnToSubnets = LimitIpLearnToSubnets.(string)
	}
	if LlAddr, ok := d.GetOk("ll_addr"); ok {
		fvBDAttr.LlAddr = LlAddr.(string)
	}
	if Mac, ok := d.GetOk("mac"); ok {
		fvBDAttr.Mac = Mac.(string)
	}
	if McastAllow, ok := d.GetOk("mcast_allow"); ok {
		fvBDAttr.McastAllow = McastAllow.(string)
	}
	if MultiDstPktAct, ok := d.GetOk("multi_dst_pkt_act"); ok {
		fvBDAttr.MultiDstPktAct = MultiDstPktAct.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		fvBDAttr.NameAlias = NameAlias.(string)
	}
	if Type, ok := d.GetOk("type"); ok {
		fvBDAttr.Type = Type.(string)
	}
	if UnicastRoute, ok := d.GetOk("unicast_route"); ok {
		fvBDAttr.UnicastRoute = UnicastRoute.(string)
	}
	if UnkMacUcastAct, ok := d.GetOk("unk_mac_ucast_act"); ok {
		fvBDAttr.UnkMacUcastAct = UnkMacUcastAct.(string)
	}
	if UnkMcastAct, ok := d.GetOk("unk_mcast_act"); ok {
		fvBDAttr.UnkMcastAct = UnkMcastAct.(string)
	}
	if V6unkMcastAct, ok := d.GetOk("v6unk_mcast_act"); ok {
		fvBDAttr.V6unkMcastAct = V6unkMcastAct.(string)
	}
	if Vmac, ok := d.GetOk("vmac"); ok {
		fvBDAttr.Vmac = Vmac.(string)
	}
	fvBD := models.NewBridgedomain(fmt.Sprintf("BD-%s", name), TenantDn, desc, fvBDAttr)

	fvBD.Status = "modified"

	err := aciClient.Save(fvBD)

	if err != nil {
		return err
	}

	if d.HasChange("relation_fv_rs_bd_to_profile") {
		_, newRelParam := d.GetChange("relation_fv_rs_bd_to_profile")
		err = aciClient.DeleteRelationfvRsBDToProfileFromBridgedomain(fvBD.DistinguishedName)
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationfvRsBDToProfileFromBridgedomain(fvBD.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_fv_rs_mldsn") {
		_, newRelParam := d.GetChange("relation_fv_rs_mldsn")
		err = aciClient.CreateRelationfvRsMldsnFromBridgedomain(fvBD.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_fv_rs_abd_pol_mon_pol") {
		_, newRelParam := d.GetChange("relation_fv_rs_abd_pol_mon_pol")
		err = aciClient.DeleteRelationfvRsABDPolMonPolFromBridgedomain(fvBD.DistinguishedName)
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationfvRsABDPolMonPolFromBridgedomain(fvBD.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_fv_rs_bd_to_nd_p") {
		_, newRelParam := d.GetChange("relation_fv_rs_bd_to_nd_p")
		err = aciClient.CreateRelationfvRsBDToNdPFromBridgedomain(fvBD.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_fv_rs_bd_flood_to") {
		oldRel, newRel := d.GetChange("relation_fv_rs_bd_flood_to")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsBdFloodToFromBridgedomain(fvBD.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsBdFloodToFromBridgedomain(fvBD.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}
	if d.HasChange("relation_fv_rs_bd_to_fhs") {
		_, newRelParam := d.GetChange("relation_fv_rs_bd_to_fhs")
		err = aciClient.DeleteRelationfvRsBDToFhsFromBridgedomain(fvBD.DistinguishedName)
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationfvRsBDToFhsFromBridgedomain(fvBD.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_fv_rs_bd_to_relay_p") {
		_, newRelParam := d.GetChange("relation_fv_rs_bd_to_relay_p")
		err = aciClient.DeleteRelationfvRsBDToRelayPFromBridgedomain(fvBD.DistinguishedName)
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationfvRsBDToRelayPFromBridgedomain(fvBD.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_fv_rs_ctx") {
		_, newRelParam := d.GetChange("relation_fv_rs_ctx")
		err = aciClient.CreateRelationfvRsCtxFromBridgedomain(fvBD.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_fv_rs_bd_to_netflow_monitor_pol") {
		oldRel, newRel := d.GetChange("relation_fv_rs_bd_to_netflow_monitor_pol")
		oldRelList := oldRel.(*schema.Set).List()
		newRelList := newRel.(*schema.Set).List()
		for _, relationParam := range oldRelList {
			paramMap := relationParam.(map[string]interface{})
			err = aciClient.DeleteRelationfvRsBDToNetflowMonitorPolFromBridgedomain(fvBD.DistinguishedName, paramMap["tn_netflow_monitor_pol_name"].(string), paramMap["flt_type"].(string))
			if err != nil {
				return err
			}
		}
		for _, relationParam := range newRelList {
			paramMap := relationParam.(map[string]interface{})
			err = aciClient.CreateRelationfvRsBDToNetflowMonitorPolFromBridgedomain(fvBD.DistinguishedName, paramMap["tn_netflow_monitor_pol_name"].(string), paramMap["flt_type"].(string))
			if err != nil {
				return err
			}
		}

	}
	if d.HasChange("relation_fv_rs_igmpsn") {
		_, newRelParam := d.GetChange("relation_fv_rs_igmpsn")
		err = aciClient.CreateRelationfvRsIgmpsnFromBridgedomain(fvBD.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_fv_rs_bd_to_ep_ret") {
		_, newRelParam := d.GetChange("relation_fv_rs_bd_to_ep_ret")
		err = aciClient.CreateRelationfvRsBdToEpRetFromBridgedomain(fvBD.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_fv_rs_bd_to_out") {
		oldRel, newRel := d.GetChange("relation_fv_rs_bd_to_out")
		oldRelSet := oldRel.(*schema.Set)
		newRelSet := newRel.(*schema.Set)
		relToDelete := toStringList(oldRelSet.Difference(newRelSet).List())
		relToCreate := toStringList(newRelSet.Difference(oldRelSet).List())

		for _, relDn := range relToDelete {
			err = aciClient.DeleteRelationfvRsBDToOutFromBridgedomain(fvBD.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

		for _, relDn := range relToCreate {
			err = aciClient.CreateRelationfvRsBDToOutFromBridgedomain(fvBD.DistinguishedName, relDn)
			if err != nil {
				return err
			}

		}

	}

	d.SetId(fvBD.DistinguishedName)
	return resourceAciBridgedomainRead(d, m)

}

func resourceAciBridgedomainRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	fvBD, err := getRemoteBridgedomain(aciClient, dn)

	if err != nil {
		return err
	}
	setBridgedomainAttributes(fvBD, d)
	return nil
}

func resourceAciBridgedomainDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "fvBD")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
