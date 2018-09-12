package aci

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceAciApplicationEPG() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciApplicationEPGCreate,
		Update: resourceAciApplicationEPGUpdate,
		Read:   resourceAciApplicationEPGRead,
		Delete: resourceAciApplicationEPGDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciApplicationEPGImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"fv_ap_dn": &schema.Schema{
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

			"exception_tag": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"flood_on_encap": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",

				ValidateFunc: validation.StringInSlice([]string{
					"disabled",
					"enabled",
				}, false),
			},

			"fwd_ctrl": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",

				ValidateFunc: validation.StringInSlice([]string{
					"none",
					"proxy-arp",
				}, false),
			},

			"has_mcast_source": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",

				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},

			"is_attr_based_e_pg": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",

				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},

			"match_t": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "match criteria",

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

			"pc_enf_pref": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "enforcement preference",

				ValidateFunc: validation.StringInSlice([]string{
					"enforced",
					"unenforced",
				}, false),
			},

			"pref_gr_memb": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",

				ValidateFunc: validation.StringInSlice([]string{
					"exclude",
					"include",
				}, false),
			},

			"prio": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "qos priority class id",

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

			"shutdown": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Mo doc not defined in techpub!!!",

				ValidateFunc: validation.StringInSlice([]string{
					"no",
					"yes",
				}, false),
			},
		}),
	}
}

func getRemoteApplicationEPG(client *client.Client, dn string) (*models.ApplicationEPG, error) {
	fvAEPgCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	fvAEPg := models.ApplicationEPGFromContainer(fvAEPgCont)

	if fvAEPg.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", fvAEPg.DistinguishedName)
	}

	return fvAEPg, nil
}

func setApplicationEPGAttributes(fvAEPg *models.ApplicationEPG, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(fvAEPg.DistinguishedName)
	d.Set("description", fvAEPg.Description)
	d.Set("fv_ap_dn", GetParentDn(fvAEPg.DistinguishedName))

	d.Set("annotation", fvAEPg.Annotation)
	d.Set("exception_tag", fvAEPg.ExceptionTag)
	d.Set("flood_on_encap", fvAEPg.FloodOnEncap)
	d.Set("fwd_ctrl", fvAEPg.FwdCtrl)
	d.Set("has_mcast_source", fvAEPg.HasMcastSource)
	d.Set("is_attr_based_e_pg", fvAEPg.IsAttrBasedEPg)
	d.Set("match_t", fvAEPg.MatchT)
	d.Set("name_alias", fvAEPg.NameAlias)
	d.Set("pc_enf_pref", fvAEPg.PcEnfPref)
	d.Set("pref_gr_memb", fvAEPg.PrefGrMemb)
	d.Set("prio", fvAEPg.Prio)
	d.Set("shutdown", fvAEPg.Shutdown)
	return d
}

func resourceAciApplicationEPGImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	fvAEPg, err := getRemoteApplicationEPG(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setApplicationEPGAttributes(fvAEPg, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciApplicationEPGCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	name := d.Get("name").(string)
	fvApDn := d.Get("fv_ap_dn").(string)

	fvAEPgAttr := models.ApplicationEPGAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		fvAEPgAttr.Annotation = Annotation.(string)
	}
	if ExceptionTag, ok := d.GetOk("exception_tag"); ok {
		fvAEPgAttr.ExceptionTag = ExceptionTag.(string)
	}
	if FloodOnEncap, ok := d.GetOk("flood_on_encap"); ok {
		fvAEPgAttr.FloodOnEncap = FloodOnEncap.(string)
	}
	if FwdCtrl, ok := d.GetOk("fwd_ctrl"); ok {
		fvAEPgAttr.FwdCtrl = FwdCtrl.(string)
	}
	if HasMcastSource, ok := d.GetOk("has_mcast_source"); ok {
		fvAEPgAttr.HasMcastSource = HasMcastSource.(string)
	}
	if IsAttrBasedEPg, ok := d.GetOk("is_attr_based_e_pg"); ok {
		fvAEPgAttr.IsAttrBasedEPg = IsAttrBasedEPg.(string)
	}
	if MatchT, ok := d.GetOk("match_t"); ok {
		fvAEPgAttr.MatchT = MatchT.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		fvAEPgAttr.NameAlias = NameAlias.(string)
	}
	if PcEnfPref, ok := d.GetOk("pc_enf_pref"); ok {
		fvAEPgAttr.PcEnfPref = PcEnfPref.(string)
	}
	if PrefGrMemb, ok := d.GetOk("pref_gr_memb"); ok {
		fvAEPgAttr.PrefGrMemb = PrefGrMemb.(string)
	}
	if Prio, ok := d.GetOk("prio"); ok {
		fvAEPgAttr.Prio = Prio.(string)
	}
	if Shutdown, ok := d.GetOk("shutdown"); ok {
		fvAEPgAttr.Shutdown = Shutdown.(string)
	}
	fvAEPg := models.NewApplicationEPG(fmt.Sprintf("epg-%s", name), fvApDn, desc, fvAEPgAttr)

	err := aciClient.Save(fvAEPg)
	if err != nil {
		return err
	}

	d.SetId(fvAEPg.DistinguishedName)
	return resourceAciApplicationEPGRead(d, m)
}

func resourceAciApplicationEPGUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)
	fvApDn := d.Get("fv_ap_dn").(string)

	fvAEPgAttr := models.ApplicationEPGAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		fvAEPgAttr.Annotation = Annotation.(string)
	}
	if ExceptionTag, ok := d.GetOk("exception_tag"); ok {
		fvAEPgAttr.ExceptionTag = ExceptionTag.(string)
	}
	if FloodOnEncap, ok := d.GetOk("flood_on_encap"); ok {
		fvAEPgAttr.FloodOnEncap = FloodOnEncap.(string)
	}
	if FwdCtrl, ok := d.GetOk("fwd_ctrl"); ok {
		fvAEPgAttr.FwdCtrl = FwdCtrl.(string)
	}
	if HasMcastSource, ok := d.GetOk("has_mcast_source"); ok {
		fvAEPgAttr.HasMcastSource = HasMcastSource.(string)
	}
	if IsAttrBasedEPg, ok := d.GetOk("is_attr_based_e_pg"); ok {
		fvAEPgAttr.IsAttrBasedEPg = IsAttrBasedEPg.(string)
	}
	if MatchT, ok := d.GetOk("match_t"); ok {
		fvAEPgAttr.MatchT = MatchT.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		fvAEPgAttr.NameAlias = NameAlias.(string)
	}
	if PcEnfPref, ok := d.GetOk("pc_enf_pref"); ok {
		fvAEPgAttr.PcEnfPref = PcEnfPref.(string)
	}
	if PrefGrMemb, ok := d.GetOk("pref_gr_memb"); ok {
		fvAEPgAttr.PrefGrMemb = PrefGrMemb.(string)
	}
	if Prio, ok := d.GetOk("prio"); ok {
		fvAEPgAttr.Prio = Prio.(string)
	}
	if Shutdown, ok := d.GetOk("shutdown"); ok {
		fvAEPgAttr.Shutdown = Shutdown.(string)
	}
	fvAEPg := models.NewApplicationEPG(fmt.Sprintf("epg-%s", name), fvApDn, desc, fvAEPgAttr)

	fvAEPg.Status = "modified"

	err := aciClient.Save(fvAEPg)

	if err != nil {
		return err
	}

	d.SetId(fvAEPg.DistinguishedName)
	return resourceAciApplicationEPGRead(d, m)

}

func resourceAciApplicationEPGRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	fvAEPg, err := getRemoteApplicationEPG(aciClient, dn)

	if err != nil {
		return err
	}
	setApplicationEPGAttributes(fvAEPg, d)
	return nil
}

func resourceAciApplicationEPGDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "fvAEPg")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
