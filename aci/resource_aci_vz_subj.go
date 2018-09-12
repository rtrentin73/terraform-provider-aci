package aci

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceAciContractSubject() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciContractSubjectCreate,
		Update: resourceAciContractSubjectUpdate,
		Read:   resourceAciContractSubjectRead,
		Delete: resourceAciContractSubjectDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciContractSubjectImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"vz_br_cp_dn": &schema.Schema{
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

			"cons_match_t": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "consumer subject match criteria",

				ValidateFunc: validation.StringInSlice([]string{
					"All",
					"AtleastOne",
					"AtmostOne",
					"None",
				}, false),
			},

			"name_alias": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"prio": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "priority level specifier",

				ValidateFunc: validation.StringInSlice([]string{
					"level1",
					"level2",
					"level3",
					"level4",
					"level5",
					"level6",
					"unspecified",
				}, false),
			},

			"prov_match_t": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "consumer subject match criteria",

				ValidateFunc: validation.StringInSlice([]string{
					"All",
					"AtleastOne",
					"AtmostOne",
					"None",
				}, false),
			},

			"rev_flt_ports": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "enables filter to apply on ingress and egress traffic",

				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},

			"target_dscp": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "target dscp",

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
		}),
	}
}

func getRemoteContractSubject(client *client.Client, dn string) (*models.ContractSubject, error) {
	vzSubjCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	vzSubj := models.ContractSubjectFromContainer(vzSubjCont)

	if vzSubj.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", vzSubj.DistinguishedName)
	}

	return vzSubj, nil
}

func setContractSubjectAttributes(vzSubj *models.ContractSubject, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(vzSubj.DistinguishedName)
	d.Set("description", vzSubj.Description)
	d.Set("vz_br_cp_dn", GetParentDn(vzSubj.DistinguishedName))

	d.Set("annotation", vzSubj.Annotation)
	d.Set("cons_match_t", vzSubj.ConsMatchT)
	d.Set("name_alias", vzSubj.NameAlias)
	d.Set("prio", vzSubj.Prio)
	d.Set("prov_match_t", vzSubj.ProvMatchT)
	d.Set("rev_flt_ports", vzSubj.RevFltPorts)
	d.Set("target_dscp", vzSubj.TargetDscp)
	return d
}

func resourceAciContractSubjectImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	vzSubj, err := getRemoteContractSubject(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setContractSubjectAttributes(vzSubj, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciContractSubjectCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	name := d.Get("name").(string)
	vzBrCPDn := d.Get("vz_br_cp_dn").(string)

	vzSubjAttr := models.ContractSubjectAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		vzSubjAttr.Annotation = Annotation.(string)
	}
	if ConsMatchT, ok := d.GetOk("cons_match_t"); ok {
		vzSubjAttr.ConsMatchT = ConsMatchT.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		vzSubjAttr.NameAlias = NameAlias.(string)
	}
	if Prio, ok := d.GetOk("prio"); ok {
		vzSubjAttr.Prio = Prio.(string)
	}
	if ProvMatchT, ok := d.GetOk("prov_match_t"); ok {
		vzSubjAttr.ProvMatchT = ProvMatchT.(string)
	}
	if RevFltPorts, ok := d.GetOk("rev_flt_ports"); ok {
		vzSubjAttr.RevFltPorts = RevFltPorts.(string)
	}
	if TargetDscp, ok := d.GetOk("target_dscp"); ok {
		vzSubjAttr.TargetDscp = TargetDscp.(string)
	}
	vzSubj := models.NewContractSubject(fmt.Sprintf("subj-%s", name), vzBrCPDn, desc, vzSubjAttr)

	err := aciClient.Save(vzSubj)
	if err != nil {
		return err
	}

	d.SetId(vzSubj.DistinguishedName)
	return resourceAciContractSubjectRead(d, m)
}

func resourceAciContractSubjectUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)
	vzBrCPDn := d.Get("vz_br_cp_dn").(string)

	vzSubjAttr := models.ContractSubjectAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		vzSubjAttr.Annotation = Annotation.(string)
	}
	if ConsMatchT, ok := d.GetOk("cons_match_t"); ok {
		vzSubjAttr.ConsMatchT = ConsMatchT.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		vzSubjAttr.NameAlias = NameAlias.(string)
	}
	if Prio, ok := d.GetOk("prio"); ok {
		vzSubjAttr.Prio = Prio.(string)
	}
	if ProvMatchT, ok := d.GetOk("prov_match_t"); ok {
		vzSubjAttr.ProvMatchT = ProvMatchT.(string)
	}
	if RevFltPorts, ok := d.GetOk("rev_flt_ports"); ok {
		vzSubjAttr.RevFltPorts = RevFltPorts.(string)
	}
	if TargetDscp, ok := d.GetOk("target_dscp"); ok {
		vzSubjAttr.TargetDscp = TargetDscp.(string)
	}
	vzSubj := models.NewContractSubject(fmt.Sprintf("subj-%s", name), vzBrCPDn, desc, vzSubjAttr)

	vzSubj.Status = "modified"

	err := aciClient.Save(vzSubj)

	if err != nil {
		return err
	}

	d.SetId(vzSubj.DistinguishedName)
	return resourceAciContractSubjectRead(d, m)

}

func resourceAciContractSubjectRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	vzSubj, err := getRemoteContractSubject(aciClient, dn)

	if err != nil {
		return err
	}
	setContractSubjectAttributes(vzSubj, d)
	return nil
}

func resourceAciContractSubjectDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "vzSubj")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
