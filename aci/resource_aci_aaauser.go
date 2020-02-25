package aci

import (
	"fmt"
	"log"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAciUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciUserCreate,
		Update: resourceAciUserUpdate,
		Read:   resourceAciUserRead,
		Delete: resourceAciUserDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciUserImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"account_status": &schema.Schema{
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
			// "otp_enable": &schema.Schema{
			// 	Type:     schema.TypeBool,
			// 	Optional: true,
			// 	Computed: true,
			// },
			"otp_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"phone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pwd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// "pwd_life_time": &schema.Schema{
			// 	Type:     schema.TypeInt,
			// 	Optional: true,
			// 	Computed: true,
			// },
			// "pwd_update_required": &schema.Schema{
			// 	Type:     schema.TypeBool,
			// 	Optional: true,
			// 	Computed: true,
			// },
			"rbac_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// "unix_user_id": &schema.Schema{
			// 	Type:     schema.TypeInt,
			// 	Optional: true,
			// 	Computed: true,
			// },
			"annotation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clear_pwd_history": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expiration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expires": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"first_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		}),
	}
}

func getRemoteUser(client *client.Client, dn string) (*models.User, error) {
	aaaUserCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	aaaUser := models.UserFromContainer(aaaUserCont)

	if aaaUser.DistinguishedName == "" {
		return nil, fmt.Errorf("User %s not found", aaaUser.DistinguishedName)
	}

	return aaaUser, nil
}

func setUserAttributes(aaaUser *models.User, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(aaaUser.DistinguishedName)
	d.Set("description", aaaUser.Description)
	aaaUserMap, _ := aaaUser.ToMap()

	d.Set("name", aaaUserMap["name"])
	d.Set("account_status", aaaUserMap["account_status"])
	d.Set("name_alias", aaaUserMap["name_alias"])
	//d.Set("otp_enable", aaaUserMap["otp_enable"])
	d.Set("otp_key", aaaUserMap["otp_key"])
	d.Set("phone", aaaUserMap["phone"])
	d.Set("pwd", aaaUserMap["pwd"])
	//d.Set("pwd_life_time", aaaUserMap["pwd_life_time"])
	//d.Set("pwd_update_required", aaaUserMap["pwd_update_required"])
	d.Set("rbac_string", aaaUserMap["rbac_string"])
	//d.Set("unix_user_id", aaaUserMap["unix_user_id"])
	d.Set("annotation", aaaUserMap["annotation"])
	d.Set("cert_attribute", aaaUserMap["cert_attribute"])
	d.Set("clear_pwd_history", aaaUserMap["clear_pwd_history"])
	d.Set("email", aaaUserMap["email"])
	d.Set("expiration", aaaUserMap["expiration"])
	d.Set("expires", aaaUserMap["expires"])
	d.Set("first_name", aaaUserMap["first_name"])
	d.Set("last_name", aaaUserMap["last_name"])
	return d
}

func resourceAciUserImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] %s: Beginning Import", d.Id())
	aciClient := m.(*client.Client)

	dn := d.Id()

	aaaUser, err := getRemoteUser(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setUserAttributes(aaaUser, d)

	log.Printf("[DEBUG] %s: Import finished successfully", d.Id())

	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciUserCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] User: Beginning Creation")
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	name := d.Get("name").(string)

	userAttr := models.UserAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		userAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		userAttr.NameAlias = NameAlias.(string)
	}
	if AccountStatus, ok := d.GetOk("account_status"); ok {
		userAttr.AccountStatus = AccountStatus.(string)
	}
	// if OtpEnable, ok := d.GetOk("otp_enable"); ok {
	// 	userAttr.OtpEnable = OtpEnable.(bool)
	// }
	if OtpKey, ok := d.GetOk("otp_key"); ok {
		userAttr.OtpKey = OtpKey.(string)
	}
	if Phone, ok := d.GetOk("phone"); ok {
		userAttr.Phone = Phone.(string)
	}
	if Pwd, ok := d.GetOk("pwd"); ok {
		userAttr.Pwd = Pwd.(string)
	}
	// if Pwd_Life_Time, ok := d.GetOk("pwd_life_time"); ok {
	// 	userAttr.PwdLifetime = Pwd_Life_Time.(int)
	// }
	// if PwdUpdateRequired, ok := d.GetOk("pwd_update_required"); ok {
	// 	userAttr.PwdUpdateRequired = PwdUpdateRequired.(bool)
	// }
	if RbacString, ok := d.GetOk("rbac_string"); ok {
		userAttr.RbacString = RbacString.(string)
	}
	// if UnixUserId, ok := d.GetOk("unix_user_id"); ok {
	// 	userAttr.UnixUserId = UnixUserId.(int)
	// }
	if Cert_Attribute, ok := d.GetOk("cert_attribute"); ok {
		userAttr.CertAttribute = Cert_Attribute.(string)
	}
	if Clear_Pwd_History, ok := d.GetOk("clear_pwd_history"); ok {
		userAttr.ClearPwdHistory = Clear_Pwd_History.(string)
	}
	if Email, ok := d.GetOk("email"); ok {
		userAttr.Email = Email.(string)
	}
	if Expiration, ok := d.GetOk("expiration"); ok {
		userAttr.Expiration = Expiration.(string)
	}
	if Expires, ok := d.GetOk("expires"); ok {
		userAttr.Expires = Expires.(string)
	}
	if FirstName, ok := d.GetOk("first_name"); ok {
		userAttr.FirstName = FirstName.(string)
	}
	if LastName, ok := d.GetOk("last_name"); ok {
		userAttr.LastName = LastName.(string)
	}

	user := models.NewUser(fmt.Sprintf("tn-%s", name), "uni", desc, userAttr)

	err := aciClient.Save(user)
	if err != nil {
		return err
	}
	d.Partial(true)

	d.SetPartial("name")

	d.Partial(false)

	d.SetId(user.DistinguishedName)
	log.Printf("[DEBUG] %s: Creation finished successfully", d.Id())

	return resourceAciUserRead(d, m)
}

func resourceAciUserUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] User: Beginning Update")

	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	userAttr := models.UserAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		userAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		userAttr.NameAlias = NameAlias.(string)
	}
	if AccountStatus, ok := d.GetOk("account_status"); ok {
		userAttr.AccountStatus = AccountStatus.(string)
	}
	// if OtpEnable, ok := d.GetOk("otp_enable"); ok {
	// 	userAttr.OtpEnable = OtpEnable.(bool)
	// }
	if OtpKey, ok := d.GetOk("otp_key"); ok {
		userAttr.OtpKey = OtpKey.(string)
	}
	if Phone, ok := d.GetOk("phone"); ok {
		userAttr.Phone = Phone.(string)
	}
	if Pwd, ok := d.GetOk("pwd"); ok {
		userAttr.Pwd = Pwd.(string)
	}
	// if Pwd_Life_Time, ok := d.GetOk("pwd_life_time"); ok {
	// 	userAttr.PwdLifetime = Pwd_Life_Time.(int)
	// }
	// if PwdUpdateRequired, ok := d.GetOk("pwd_update_required"); ok {
	// 	userAttr.PwdUpdateRequired = PwdUpdateRequired.(bool)
	// }
	if RbacString, ok := d.GetOk("rbac_string"); ok {
		userAttr.RbacString = RbacString.(string)
	}
	// if UnixUserId, ok := d.GetOk("unix_user_id"); ok {
	// 	userAttr.UnixUserId = UnixUserId.(int)
	// }
	if Cert_Attribute, ok := d.GetOk("cert_attribute"); ok {
		userAttr.CertAttribute = Cert_Attribute.(string)
	}
	if Clear_Pwd_History, ok := d.GetOk("clear_pwd_history"); ok {
		userAttr.ClearPwdHistory = Clear_Pwd_History.(string)
	}
	if Email, ok := d.GetOk("email"); ok {
		userAttr.Email = Email.(string)
	}
	if Expiration, ok := d.GetOk("expiration"); ok {
		userAttr.Expiration = Expiration.(string)
	}
	if Expires, ok := d.GetOk("expires"); ok {
		userAttr.Expires = Expires.(string)
	}
	if FirstName, ok := d.GetOk("first_name"); ok {
		userAttr.FirstName = FirstName.(string)
	}
	if LastName, ok := d.GetOk("last_name"); ok {
		userAttr.LastName = LastName.(string)
	}
	aaaUser := models.NewUser(fmt.Sprintf("tn-%s", name), "uni", desc, userAttr)

	aaaUser.Status = "modified"

	err := aciClient.Save(aaaUser)

	if err != nil {
		return err
	}
	d.Partial(true)

	d.SetPartial("name")

	d.Partial(false)

	d.SetId(aaaUser.DistinguishedName)
	log.Printf("[DEBUG] %s: Update finished successfully", d.Id())

	return resourceAciUserRead(d, m)

}

func resourceAciUserRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Read", d.Id())

	aciClient := m.(*client.Client)

	dn := d.Id()
	aaaUser, err := getRemoteUser(aciClient, dn)

	if err != nil {
		d.SetId("")
		return nil
	}
	setUserAttributes(aaaUser, d)
	log.Printf("[DEBUG] %s: Read finished successfully", d.Id())

	return nil
}

func resourceAciUserDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Destroy", d.Id())

	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "aaaUser")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] %s: Destroy finished successfully", d.Id())

	d.SetId("")
	return err
}
