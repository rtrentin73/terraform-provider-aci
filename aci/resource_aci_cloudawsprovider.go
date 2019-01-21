package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciCloudawsprovider() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciCloudawsproviderCreate,
		Update: resourceAciCloudawsproviderUpdate,
		Read:   resourceAciCloudawsproviderRead,
		Delete: resourceAciCloudawsproviderDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudawsproviderImport,
		},

		SchemaVersion: 1,

		Schema: AppendBaseAttrSchema(map[string]*schema.Schema{
			"tenant_dn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"access_key_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"account_id": &schema.Schema{
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

			"email": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "email address of the local user",
			},

			"http_proxy": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"is_account_in_org": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"is_trusted": &schema.Schema{
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

			"provider_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"region": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},

			"secret_access_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Mo doc not defined in techpub!!!",
			},
		}),
	}
}

func getRemoteCloudawsprovider(client *client.Client, dn string) (*models.Cloudawsprovider, error) {
	cloudAwsProviderCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudAwsProvider := models.CloudawsproviderFromContainer(cloudAwsProviderCont)

	if cloudAwsProvider.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudAwsProvider.DistinguishedName)
	}

	return cloudAwsProvider, nil
}

func setCloudawsproviderAttributes(cloudAwsProvider *models.Cloudawsprovider, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudAwsProvider.DistinguishedName)
	d.Set("description", cloudAwsProvider.Description)
	d.Set("tenant_dn", GetParentDn(cloudAwsProvider.DistinguishedName))
	cloudAwsProviderMap, _ := cloudAwsProvider.ToMap()

	d.Set("access_key_id", cloudAwsProviderMap["accessKeyId"])
	d.Set("account_id", cloudAwsProviderMap["accountId"])
	d.Set("annotation", cloudAwsProviderMap["annotation"])
	d.Set("email", cloudAwsProviderMap["email"])
	d.Set("http_proxy", cloudAwsProviderMap["httpProxy"])
	d.Set("is_account_in_org", cloudAwsProviderMap["isAccountInOrg"])
	d.Set("is_trusted", cloudAwsProviderMap["isTrusted"])
	d.Set("name_alias", cloudAwsProviderMap["nameAlias"])
	d.Set("provider_id", cloudAwsProviderMap["providerId"])
	d.Set("region", cloudAwsProviderMap["region"])
	d.Set("secret_access_key", cloudAwsProviderMap["secretAccessKey"])
	return d
}

func resourceAciCloudawsproviderImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudAwsProvider, err := getRemoteCloudawsprovider(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCloudawsproviderAttributes(cloudAwsProvider, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudawsproviderCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	TenantDn := d.Get("tenant_dn").(string)

	cloudAwsProviderAttr := models.CloudawsproviderAttributes{}
	if AccessKeyId, ok := d.GetOk("access_key_id"); ok {
		cloudAwsProviderAttr.AccessKeyId = AccessKeyId.(string)
	}
	if AccountId, ok := d.GetOk("account_id"); ok {
		cloudAwsProviderAttr.AccountId = AccountId.(string)
	}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudAwsProviderAttr.Annotation = Annotation.(string)
	}
	if Email, ok := d.GetOk("email"); ok {
		cloudAwsProviderAttr.Email = Email.(string)
	}
	if HttpProxy, ok := d.GetOk("http_proxy"); ok {
		cloudAwsProviderAttr.HttpProxy = HttpProxy.(string)
	}
	if IsAccountInOrg, ok := d.GetOk("is_account_in_org"); ok {
		cloudAwsProviderAttr.IsAccountInOrg = IsAccountInOrg.(string)
	}
	if IsTrusted, ok := d.GetOk("is_trusted"); ok {
		cloudAwsProviderAttr.IsTrusted = IsTrusted.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudAwsProviderAttr.NameAlias = NameAlias.(string)
	}
	if ProviderId, ok := d.GetOk("provider_id"); ok {
		cloudAwsProviderAttr.ProviderId = ProviderId.(string)
	}
	if Region, ok := d.GetOk("region"); ok {
		cloudAwsProviderAttr.Region = Region.(string)
	}
	if SecretAccessKey, ok := d.GetOk("secret_access_key"); ok {
		cloudAwsProviderAttr.SecretAccessKey = SecretAccessKey.(string)
	}
	cloudAwsProvider := models.NewCloudawsprovider(fmt.Sprintf("awsprovider"), TenantDn, desc, cloudAwsProviderAttr)

	err := aciClient.Save(cloudAwsProvider)
	if err != nil {
		return err
	}

	d.SetId(cloudAwsProvider.DistinguishedName)
	return resourceAciCloudawsproviderRead(d, m)
}

func resourceAciCloudawsproviderUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	TenantDn := d.Get("tenant_dn").(string)

	cloudAwsProviderAttr := models.CloudawsproviderAttributes{}
	if AccessKeyId, ok := d.GetOk("access_key_id"); ok {
		cloudAwsProviderAttr.AccessKeyId = AccessKeyId.(string)
	}
	if AccountId, ok := d.GetOk("account_id"); ok {
		cloudAwsProviderAttr.AccountId = AccountId.(string)
	}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudAwsProviderAttr.Annotation = Annotation.(string)
	}
	if Email, ok := d.GetOk("email"); ok {
		cloudAwsProviderAttr.Email = Email.(string)
	}
	if HttpProxy, ok := d.GetOk("http_proxy"); ok {
		cloudAwsProviderAttr.HttpProxy = HttpProxy.(string)
	}
	if IsAccountInOrg, ok := d.GetOk("is_account_in_org"); ok {
		cloudAwsProviderAttr.IsAccountInOrg = IsAccountInOrg.(string)
	}
	if IsTrusted, ok := d.GetOk("is_trusted"); ok {
		cloudAwsProviderAttr.IsTrusted = IsTrusted.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudAwsProviderAttr.NameAlias = NameAlias.(string)
	}
	if ProviderId, ok := d.GetOk("provider_id"); ok {
		cloudAwsProviderAttr.ProviderId = ProviderId.(string)
	}
	if Region, ok := d.GetOk("region"); ok {
		cloudAwsProviderAttr.Region = Region.(string)
	}
	if SecretAccessKey, ok := d.GetOk("secret_access_key"); ok {
		cloudAwsProviderAttr.SecretAccessKey = SecretAccessKey.(string)
	}
	cloudAwsProvider := models.NewCloudawsprovider(fmt.Sprintf("awsprovider"), TenantDn, desc, cloudAwsProviderAttr)

	cloudAwsProvider.Status = "modified"

	err := aciClient.Save(cloudAwsProvider)

	if err != nil {
		return err
	}

	d.SetId(cloudAwsProvider.DistinguishedName)
	return resourceAciCloudawsproviderRead(d, m)

}

func resourceAciCloudawsproviderRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudAwsProvider, err := getRemoteCloudawsprovider(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudawsproviderAttributes(cloudAwsProvider, d)
	return nil
}

func resourceAciCloudawsproviderDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudAwsProvider")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
