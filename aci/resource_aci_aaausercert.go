package aci

import (
	"fmt"
	"log"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAciUserCert() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciUserCertCreate,
		Update: resourceAciUserCertUpdate,
		Read:   resourceAciUserCertRead,
		Delete: resourceAciUserCertDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciUserCertImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"annotation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name_alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		}),
	}
}

func getRemoteUserCert(client *client.Client, dn string) (*models.UserCert, error) {
	aaaUserCertCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	aaaUserCert := models.UserCertFromContainer(aaaUserCertCont)

	if aaaUserCert.DistinguishedName == "" {
		return nil, fmt.Errorf("UserCert %s not found", aaaUserCert.DistinguishedName)
	}

	return aaaUserCert, nil
}

func setUserCertAttributes(aaaUserCert *models.UserCert, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(aaaUserCert.DistinguishedName)
	d.Set("description", aaaUserCert.Description)
	d.Set("user_name", GetParentDn(aaaUserCert.DistinguishedName))
	aaaUserCertMap, _ := aaaUserCert.ToMap()

	d.Set("name", aaaUserCertMap["name"])
	d.Set("name_alias", aaaUserCertMap["name_alias"])
	d.Set("data", aaaUserCertMap["data"])
	d.Set("annotation", aaaUserCertMap["annotation"])
	return d
}

func resourceAciUserCertImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] %s: Beginning Import", d.Id())
	aciClient := m.(*client.Client)

	dn := d.Id()

	aaaUserCert, err := getRemoteUserCert(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setUserCertAttributes(aaaUserCert, d)

	log.Printf("[DEBUG] %s: Import finished successfully", d.Id())

	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciUserCertCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] UserCert: Beginning Creation")
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	name := d.Get("name").(string)
	Username := d.Get("user_name").(string)

	userCertAttr := models.UserCertAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		userCertAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		userCertAttr.NameAlias = NameAlias.(string)
	}
	if Data, ok := d.GetOk("data"); ok {
		userCertAttr.Data = Data.(string)
	}

	UserCert := models.NewUserCert(fmt.Sprintf("usercert-%s", name), Username, desc, userCertAttr)

	err := aciClient.Save(UserCert)
	if err != nil {
		return err
	}
	d.Partial(true)

	d.SetPartial("name")

	d.Partial(false)

	d.SetId(UserCert.DistinguishedName)
	log.Printf("[DEBUG] %s: Creation finished successfully", d.Id())

	return resourceAciUserCertRead(d, m)
}

func resourceAciUserCertUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] UserCert: Beginning Update")

	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	Username := d.Get("user_name").(string)

	userCertAttr := models.UserCertAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		userCertAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		userCertAttr.NameAlias = NameAlias.(string)
	}
	if Data, ok := d.GetOk("data"); ok {
		userCertAttr.Data = Data.(string)
	}

	aaaUserCert := models.NewUserCert(fmt.Sprintf("usercert-%s", name), Username, desc, userCertAttr)

	aaaUserCert.Status = "modified"

	err := aciClient.Save(aaaUserCert)

	if err != nil {
		return err
	}
	d.Partial(true)

	d.SetPartial("name")

	d.Partial(false)

	d.SetId(aaaUserCert.DistinguishedName)
	log.Printf("[DEBUG] %s: Update finished successfully", d.Id())

	return resourceAciUserCertRead(d, m)

}

func resourceAciUserCertRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Read", d.Id())

	aciClient := m.(*client.Client)

	dn := d.Id()
	aaaUserCert, err := getRemoteUserCert(aciClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setUserCertAttributes(aaaUserCert, d)
	log.Printf("[DEBUG] %s: Read finished successfully", d.Id())

	return nil
}

func resourceAciUserCertDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Destroy", d.Id())

	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "aaaUserCert")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] %s: Destroy finished successfully", d.Id())

	d.SetId("")
	return err
}
