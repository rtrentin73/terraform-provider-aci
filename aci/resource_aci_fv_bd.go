package aci

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceAciBridgeDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciBridgeDomainCreate,
		Update: resourceAciBridgeDomainUpdate,
		Read:   resourceAciBridgeDomainRead,
		Delete: resourceAciBridgeDomainDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciBridgeDomainImport,
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

			"optimize_wan_bandwidth": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",

				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},

			"annotation": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"arp_flood": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "arp flood enable",

				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},

			"ep_clear": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",

				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},

			"ep_move_detect_mode": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ep move detection garp based mode",

				ValidateFunc: validation.StringInSlice([]string{
					"garp",
				}, false),
			},

			"host_based_routing": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",

				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},

			"intersite_bum_traffic_allow": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",

				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},

			"intersite_l2_stretch": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",

				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},

			"ip_learning": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Endpoint Dataplane Learning",

				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},

			"limit_ip_learn_to_subnets": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "limits ip learning to bd subnets only",

				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},

			"ll_addr": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "override of system generated ipv6 link-local address",
			},

			"mac": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "mac address",
			},

			"mcast_allow": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",

				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},

			"multi_dst_pkt_act": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "forwarding method for multi destinations",

				ValidateFunc: validation.StringInSlice([]string{
					"bd-flood",
					"drop",
					"encap-flood",
				}, false),
			},

			"name_alias": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "component type",

				ValidateFunc: validation.StringInSlice([]string{
					"fc",
					"regular",
				}, false),
			},

			"unicast_route": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unicast routing",

				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},

			"unk_mac_ucast_act": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "forwarding method for l2 destinations",

				ValidateFunc: validation.StringInSlice([]string{
					"flood",
					"proxy",
				}, false),
			},

			"unk_mcast_act": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "parameter used by node to forward data",

				ValidateFunc: validation.StringInSlice([]string{
					"flood",
					"opt-flood",
				}, false),
			},

			"vmac": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",

				ValidateFunc: validation.StringInSlice([]string{
					"not-applicable",
				}, false),
			},
		}),
	}
}

func getRemoteBridgeDomain(client *client.Client, dn string) (*models.BridgeDomain, error) {
	fvBDCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	fvBD := models.BridgeDomainFromContainer(fvBDCont)

	if fvBD.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", fvBD.DistinguishedName)
	}

	return fvBD, nil
}

func setBridgeDomainAttributes(fvBD *models.BridgeDomain, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(fvBD.DistinguishedName)
	d.Set("description", fvBD.Description)
	d.Set("fv_tenant_dn", GetParentDn(fvBD.DistinguishedName))

	d.Set("optimize_wan_bandwidth", fvBD.OptimizeWanBandwidth)
	d.Set("annotation", fvBD.Annotation)
	d.Set("arp_flood", fvBD.ArpFlood)
	d.Set("ep_clear", fvBD.EpClear)
	d.Set("ep_move_detect_mode", fvBD.EpMoveDetectMode)
	d.Set("host_based_routing", fvBD.HostBasedRouting)
	d.Set("intersite_bum_traffic_allow", fvBD.IntersiteBumTrafficAllow)
	d.Set("intersite_l2_stretch", fvBD.IntersiteL2Stretch)
	d.Set("ip_learning", fvBD.IpLearning)
	d.Set("limit_ip_learn_to_subnets", fvBD.LimitIpLearnToSubnets)
	d.Set("ll_addr", fvBD.LlAddr)
	d.Set("mac", fvBD.Mac)
	d.Set("mcast_allow", fvBD.McastAllow)
	d.Set("multi_dst_pkt_act", fvBD.MultiDstPktAct)
	d.Set("name_alias", fvBD.NameAlias)
	d.Set("type", fvBD.Type)
	d.Set("unicast_route", fvBD.UnicastRoute)
	d.Set("unk_mac_ucast_act", fvBD.UnkMacUcastAct)
	d.Set("unk_mcast_act", fvBD.UnkMcastAct)
	d.Set("vmac", fvBD.Vmac)
	return d
}

func resourceAciBridgeDomainImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	fvBD, err := getRemoteBridgeDomain(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setBridgeDomainAttributes(fvBD, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciBridgeDomainCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	name := d.Get("name").(string)
	fvTenantDn := d.Get("fv_tenant_dn").(string)

	fvBDAttr := models.BridgeDomainAttributes{}
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
	if Vmac, ok := d.GetOk("vmac"); ok {
		fvBDAttr.Vmac = Vmac.(string)
	}
	fvBD := models.NewBridgeDomain(fmt.Sprintf("BD-%s", name), fvTenantDn, desc, fvBDAttr)

	err := aciClient.Save(fvBD)
	if err != nil {
		return err
	}

	d.SetId(fvBD.DistinguishedName)
	return resourceAciBridgeDomainRead(d, m)
}

func resourceAciBridgeDomainUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)
	fvTenantDn := d.Get("fv_tenant_dn").(string)

	fvBDAttr := models.BridgeDomainAttributes{}
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
	if Vmac, ok := d.GetOk("vmac"); ok {
		fvBDAttr.Vmac = Vmac.(string)
	}
	fvBD := models.NewBridgeDomain(fmt.Sprintf("BD-%s", name), fvTenantDn, desc, fvBDAttr)

	fvBD.Status = "modified"

	err := aciClient.Save(fvBD)

	if err != nil {
		return err
	}

	d.SetId(fvBD.DistinguishedName)
	return resourceAciBridgeDomainRead(d, m)

}

func resourceAciBridgeDomainRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	fvBD, err := getRemoteBridgeDomain(aciClient, dn)

	if err != nil {
		return err
	}
	setBridgeDomainAttributes(fvBD, d)
	return nil
}

func resourceAciBridgeDomainDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "fvBD")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
