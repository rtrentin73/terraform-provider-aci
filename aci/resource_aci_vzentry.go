package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciFilterentry() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciFilterentryCreate,
		Update: resourceAciFilterentryUpdate,
		Read:   resourceAciFilterentryRead,
		Delete: resourceAciFilterentryDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciFilterentryImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"filter_dn": &schema.Schema{
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
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"apply_to_frag": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "fragment",
			},

			"arp_opc": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "open peripheral codes",
			},

			"d_from_port": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "end of the destination port range",
			},

			"d_to_port": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "start of the destination port range",
			},

			"ether_t": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "ethertype",
			},

			"icmpv4_t": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "",
			},

			"icmpv6_t": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "",
			},

			"match_dscp": &schema.Schema{
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

			"prot": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "level 3 ip protocol",
			},

			"s_from_port": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "start of the source port range",
			},

			"s_to_port": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "end of the source port range",
			},

			"stateful": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "stateful entry",
			},

			"tcp_rules": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "tcp flags",
			},
		}),
	}
}

func getRemoteFilterentry(client *client.Client, dn string) (*models.Filterentry, error) {
	vzEntryCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vzEntry := models.FilterentryFromContainer(vzEntryCont)

	if vzEntry.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", vzEntry.DistinguishedName)
	}

	return vzEntry, nil
}

func setFilterentryAttributes(vzEntry *models.Filterentry, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vzEntry.DistinguishedName)
	d.Set("description", vzEntry.Description)
	d.Set("filter_dn", GetParentDn(vzEntry.DistinguishedName))
	vzEntryMap, _ := vzEntry.ToMap()

	d.Set("annotation", vzEntryMap["annotation"])
	d.Set("apply_to_frag", vzEntryMap["applyToFrag"])
	d.Set("arp_opc", vzEntryMap["arpOpc"])
	d.Set("d_from_port", vzEntryMap["dFromPort"])
	d.Set("d_to_port", vzEntryMap["dToPort"])
	d.Set("ether_t", vzEntryMap["etherT"])
	d.Set("icmpv4_t", vzEntryMap["icmpv4T"])
	d.Set("icmpv6_t", vzEntryMap["icmpv6T"])
	d.Set("match_dscp", vzEntryMap["matchDscp"])
	d.Set("name_alias", vzEntryMap["nameAlias"])
	d.Set("prot", vzEntryMap["prot"])
	d.Set("s_from_port", vzEntryMap["sFromPort"])
	d.Set("s_to_port", vzEntryMap["sToPort"])
	d.Set("stateful", vzEntryMap["stateful"])
	d.Set("tcp_rules", vzEntryMap["tcpRules"])
	return d
}

func resourceAciFilterentryImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	vzEntry, err := getRemoteFilterentry(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setFilterentryAttributes(vzEntry, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciFilterentryCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	FilterDn := d.Get("filter_dn").(string)

	vzEntryAttr := models.FilterentryAttributes{}
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
	vzEntry := models.NewFilterentry(fmt.Sprintf("e-%s", name), FilterDn, desc, vzEntryAttr)

	err := aciClient.Save(vzEntry)
	if err != nil {
		return err
	}

	d.SetId(vzEntry.DistinguishedName)
	return resourceAciFilterentryRead(d, m)
}

func resourceAciFilterentryUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	FilterDn := d.Get("filter_dn").(string)

	vzEntryAttr := models.FilterentryAttributes{}
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
	vzEntry := models.NewFilterentry(fmt.Sprintf("e-%s", name), FilterDn, desc, vzEntryAttr)

	vzEntry.Status = "modified"

	err := aciClient.Save(vzEntry)

	if err != nil {
		return err
	}

	d.SetId(vzEntry.DistinguishedName)
	return resourceAciFilterentryRead(d, m)

}

func resourceAciFilterentryRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	vzEntry, err := getRemoteFilterentry(aciClient, dn)

	if err != nil {
		return err
	}
	setFilterentryAttributes(vzEntry, d)
	return nil
}

func resourceAciFilterentryDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "vzEntry")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
