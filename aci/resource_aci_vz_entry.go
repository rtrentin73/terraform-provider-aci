package aci

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceAciFilterEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciFilterEntryCreate,
		Update: resourceAciFilterEntryUpdate,
		Read:   resourceAciFilterEntryRead,
		Delete: resourceAciFilterEntryDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciFilterEntryImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"vz_filter_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"annotation": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"apply_to_frag": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "fragment",

				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},

			"arp_opc": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "open peripheral codes",

				ValidateFunc: validation.StringInSlice([]string{
					"reply",
					"req",
					"unspecified",
				}, false),
			},

			"d_from_port": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "end of the destination port range",

				ValidateFunc: validation.StringInSlice([]string{
					"dns",
					"ftpData",
					"http",
					"https",
					"pop3",
					"rtsp",
					"smtp",
					"unspecified",
				}, false),
			},

			"d_to_port": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "start of the destination port range",

				ValidateFunc: validation.StringInSlice([]string{
					"dns",
					"ftpData",
					"http",
					"https",
					"pop3",
					"rtsp",
					"smtp",
					"unspecified",
				}, false),
			},

			"ether_t": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ethertype",

				ValidateFunc: validation.StringInSlice([]string{
					"arp",
					"fcoe",
					"ip",
					"ipv4",
					"ipv6",
					"mac_security",
					"mpls_ucast",
					"trill",
					"unspecified",
				}, false),
			},

			"icmpv4_t": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",

				ValidateFunc: validation.StringInSlice([]string{
					"dst-unreach",
					"echo",
					"echo-rep",
					"src-quench",
					"time-exceeded",
					"unspecified",
				}, false),
			},

			"icmpv6_t": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",

				ValidateFunc: validation.StringInSlice([]string{
					"dst-unreach",
					"echo-rep",
					"echo-req",
					"nbr-advert",
					"nbr-solicit",
					"redirect",
					"time-exceeded",
					"unspecified",
				}, false),
			},

			"match_dscp": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",

				ValidateFunc: validation.StringInSlice([]string{
					"AF11",
					"AF12",
					"AF13",
					"AF21",
					"AF22",
					"AF23",
					"AF31",
					"AF32",
					"AF33",
					"AF41",
					"AF42",
					"AF43",
					"CS0",
					"CS1",
					"CS2",
					"CS3",
					"CS4",
					"CS5",
					"CS6",
					"CS7",
					"EF",
					"VA",
					"unspecified",
				}, false),
			},

			"name_alias": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"prot": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "level 3 ip protocol",

				ValidateFunc: validation.StringInSlice([]string{
					"egp",
					"eigrp",
					"icmp",
					"icmpv6",
					"igmp",
					"igp",
					"l2tp",
					"ospfigp",
					"pim",
					"tcp",
					"udp",
					"unspecified",
				}, false),
			},

			"s_from_port": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "start of the source port range",

				ValidateFunc: validation.StringInSlice([]string{
					"dns",
					"ftpData",
					"http",
					"https",
					"pop3",
					"rtsp",
					"smtp",
					"unspecified",
				}, false),
			},

			"s_to_port": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "end of the source port range",

				ValidateFunc: validation.StringInSlice([]string{
					"dns",
					"ftpData",
					"http",
					"https",
					"pop3",
					"rtsp",
					"smtp",
					"unspecified",
				}, false),
			},

			"stateful": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "stateful entry",

				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},

			"tcp_rules": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "tcp flags",

				ValidateFunc: validation.StringInSlice([]string{
					"ack",
					"est",
					"fin",
					"rst",
					"syn",
					"unspecified",
				}, false),
			},
		}),
	}
}

func getRemoteFilterEntry(client *client.Client, dn string) (*models.FilterEntry, error) {
	vzEntryCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vzEntry := models.FilterEntryFromContainer(vzEntryCont)

	if vzEntry.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", vzEntry.DistinguishedName)
	}

	return vzEntry, nil
}

func setFilterEntryAttributes(vzEntry *models.FilterEntry, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vzEntry.DistinguishedName)
	d.Set("description", vzEntry.Description)
	d.Set("vz_filter_dn", GetParentDn(vzEntry.DistinguishedName))

	d.Set("annotation", vzEntry.Annotation)
	d.Set("apply_to_frag", vzEntry.ApplyToFrag)
	d.Set("arp_opc", vzEntry.ArpOpc)
	d.Set("d_from_port", vzEntry.DFromPort)
	d.Set("d_to_port", vzEntry.DToPort)
	d.Set("ether_t", vzEntry.EtherT)
	d.Set("icmpv4_t", vzEntry.Icmpv4T)
	d.Set("icmpv6_t", vzEntry.Icmpv6T)
	d.Set("match_dscp", vzEntry.MatchDscp)
	d.Set("name_alias", vzEntry.NameAlias)
	d.Set("prot", vzEntry.Prot)
	d.Set("s_from_port", vzEntry.SFromPort)
	d.Set("s_to_port", vzEntry.SToPort)
	d.Set("stateful", vzEntry.Stateful)
	d.Set("tcp_rules", vzEntry.TcpRules)
	return d
}

func resourceAciFilterEntryImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	vzEntry, err := getRemoteFilterEntry(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setFilterEntryAttributes(vzEntry, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciFilterEntryCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	name := d.Get("name").(string)
	vzFilterDn := d.Get("vz_filter_dn").(string)

	vzEntryAttr := models.FilterEntryAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		vzEntryAttr.Annotation = Annotation.(string)
	}
	if ApplyToFrag, ok := d.GetOk("apply_to_frag"); ok {
		vzEntryAttr.ApplyToFrag = ApplyToFrag.(string)
	}
	if ArpOpc, ok := d.GetOk("arp_opc"); ok {
		vzEntryAttr.ArpOpc = ArpOpc.(string)
	}
	if DFromPort, ok := d.GetOk("d_from_port"); ok {
		vzEntryAttr.DFromPort = DFromPort.(string)
	}
	if DToPort, ok := d.GetOk("d_to_port"); ok {
		vzEntryAttr.DToPort = DToPort.(string)
	}
	if EtherT, ok := d.GetOk("ether_t"); ok {
		vzEntryAttr.EtherT = EtherT.(string)
	}
	if Icmpv4T, ok := d.GetOk("icmpv4_t"); ok {
		vzEntryAttr.Icmpv4T = Icmpv4T.(string)
	}
	if Icmpv6T, ok := d.GetOk("icmpv6_t"); ok {
		vzEntryAttr.Icmpv6T = Icmpv6T.(string)
	}
	if MatchDscp, ok := d.GetOk("match_dscp"); ok {
		vzEntryAttr.MatchDscp = MatchDscp.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		vzEntryAttr.NameAlias = NameAlias.(string)
	}
	if Prot, ok := d.GetOk("prot"); ok {
		vzEntryAttr.Prot = Prot.(string)
	}
	if SFromPort, ok := d.GetOk("s_from_port"); ok {
		vzEntryAttr.SFromPort = SFromPort.(string)
	}
	if SToPort, ok := d.GetOk("s_to_port"); ok {
		vzEntryAttr.SToPort = SToPort.(string)
	}
	if Stateful, ok := d.GetOk("stateful"); ok {
		vzEntryAttr.Stateful = Stateful.(string)
	}
	if TcpRules, ok := d.GetOk("tcp_rules"); ok {
		vzEntryAttr.TcpRules = TcpRules.(string)
	}
	vzEntry := models.NewFilterEntry(fmt.Sprintf("e-%s", name), vzFilterDn, desc, vzEntryAttr)

	err := aciClient.Save(vzEntry)
	if err != nil {
		return err
	}

	d.SetId(vzEntry.DistinguishedName)
	return resourceAciFilterEntryRead(d, m)
}

func resourceAciFilterEntryUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)
	vzFilterDn := d.Get("vz_filter_dn").(string)

	vzEntryAttr := models.FilterEntryAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		vzEntryAttr.Annotation = Annotation.(string)
	}
	if ApplyToFrag, ok := d.GetOk("apply_to_frag"); ok {
		vzEntryAttr.ApplyToFrag = ApplyToFrag.(string)
	}
	if ArpOpc, ok := d.GetOk("arp_opc"); ok {
		vzEntryAttr.ArpOpc = ArpOpc.(string)
	}
	if DFromPort, ok := d.GetOk("d_from_port"); ok {
		vzEntryAttr.DFromPort = DFromPort.(string)
	}
	if DToPort, ok := d.GetOk("d_to_port"); ok {
		vzEntryAttr.DToPort = DToPort.(string)
	}
	if EtherT, ok := d.GetOk("ether_t"); ok {
		vzEntryAttr.EtherT = EtherT.(string)
	}
	if Icmpv4T, ok := d.GetOk("icmpv4_t"); ok {
		vzEntryAttr.Icmpv4T = Icmpv4T.(string)
	}
	if Icmpv6T, ok := d.GetOk("icmpv6_t"); ok {
		vzEntryAttr.Icmpv6T = Icmpv6T.(string)
	}
	if MatchDscp, ok := d.GetOk("match_dscp"); ok {
		vzEntryAttr.MatchDscp = MatchDscp.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		vzEntryAttr.NameAlias = NameAlias.(string)
	}
	if Prot, ok := d.GetOk("prot"); ok {
		vzEntryAttr.Prot = Prot.(string)
	}
	if SFromPort, ok := d.GetOk("s_from_port"); ok {
		vzEntryAttr.SFromPort = SFromPort.(string)
	}
	if SToPort, ok := d.GetOk("s_to_port"); ok {
		vzEntryAttr.SToPort = SToPort.(string)
	}
	if Stateful, ok := d.GetOk("stateful"); ok {
		vzEntryAttr.Stateful = Stateful.(string)
	}
	if TcpRules, ok := d.GetOk("tcp_rules"); ok {
		vzEntryAttr.TcpRules = TcpRules.(string)
	}
	vzEntry := models.NewFilterEntry(fmt.Sprintf("e-%s", name), vzFilterDn, desc, vzEntryAttr)

	vzEntry.Status = "modified"

	err := aciClient.Save(vzEntry)

	if err != nil {
		return err
	}

	d.SetId(vzEntry.DistinguishedName)
	return resourceAciFilterEntryRead(d, m)

}

func resourceAciFilterEntryRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	vzEntry, err := getRemoteFilterEntry(aciClient, dn)

	if err != nil {
		return err
	}
	setFilterEntryAttributes(vzEntry, d)
	return nil
}

func resourceAciFilterEntryDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "vzEntry")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
