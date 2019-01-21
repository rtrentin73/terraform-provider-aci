package aci

import (
	"fmt"
	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciCloudcontextprofile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciCloudcontextprofileCreate,
		Update: resourceAciCloudcontextprofileUpdate,
		Read:   resourceAciCloudcontextprofileRead,
		Delete: resourceAciCloudcontextprofileDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciCloudcontextprofileImport,
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

			"annotation": &schema.Schema{
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

			"type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "component type",
			},

			"relation_cloud_rs_ctx_to_flow_log": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to cloudAwsFlowLogPol",
			},
			"relation_cloud_rs_to_ctx": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to fvCtx",
			},
			"relation_cloud_rs_ctx_profile_to_region": &schema.Schema{
				Type: schema.TypeString,

				Optional:    true,
				Description: "Create relation to cloudRegion",
			},
		}),
	}
}

func getRemoteCloudcontextprofile(client *client.Client, dn string) (*models.Cloudcontextprofile, error) {
	cloudCtxProfileCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudCtxProfile := models.CloudcontextprofileFromContainer(cloudCtxProfileCont)

	if cloudCtxProfile.DistinguishedName == "" {
		return nil, fmt.Errorf("Bridge Domain %s not found", cloudCtxProfile.DistinguishedName)
	}

	return cloudCtxProfile, nil
}

func setCloudcontextprofileAttributes(cloudCtxProfile *models.Cloudcontextprofile, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(cloudCtxProfile.DistinguishedName)
	d.Set("description", cloudCtxProfile.Description)
	d.Set("tenant_dn", GetParentDn(cloudCtxProfile.DistinguishedName))
	cloudCtxProfileMap, _ := cloudCtxProfile.ToMap()

	d.Set("annotation", cloudCtxProfileMap["annotation"])
	d.Set("name_alias", cloudCtxProfileMap["nameAlias"])
	d.Set("type", cloudCtxProfileMap["type"])
	return d
}

func resourceAciCloudcontextprofileImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {

	aciClient := m.(*client.Client)

	dn := d.Id()

	cloudCtxProfile, err := getRemoteCloudcontextprofile(aciClient, dn)

	if err != nil {
		return nil, err
	}
	schemaFilled := setCloudcontextprofileAttributes(cloudCtxProfile, d)
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciCloudcontextprofileCreate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	TenantDn := d.Get("tenant_dn").(string)

	cloudCtxProfileAttr := models.CloudcontextprofileAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudCtxProfileAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudCtxProfileAttr.NameAlias = NameAlias.(string)
	}
	if Type, ok := d.GetOk("type"); ok {
		cloudCtxProfileAttr.Type = Type.(string)
	}
	cloudCtxProfile := models.NewCloudcontextprofile(fmt.Sprintf("ctxprofile-%s", name), TenantDn, desc, cloudCtxProfileAttr)

	err := aciClient.Save(cloudCtxProfile)
	if err != nil {
		return err
	}

	if relationTocloudRsCtxToFlowLog, ok := d.GetOk("relation_cloud_rs_ctx_to_flow_log"); ok {
		relationParam := relationTocloudRsCtxToFlowLog.(string)
		err = aciClient.CreateRelationcloudRsCtxToFlowLogFromCloudcontextprofile(cloudCtxProfile.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTocloudRsToCtx, ok := d.GetOk("relation_cloud_rs_to_ctx"); ok {
		relationParam := relationTocloudRsToCtx.(string)
		err = aciClient.CreateRelationcloudRsToCtxFromCloudcontextprofile(cloudCtxProfile.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}
	if relationTocloudRsCtxProfileToRegion, ok := d.GetOk("relation_cloud_rs_ctx_profile_to_region"); ok {
		relationParam := relationTocloudRsCtxProfileToRegion.(string)
		err = aciClient.CreateRelationcloudRsCtxProfileToRegionFromCloudcontextprofile(cloudCtxProfile.DistinguishedName, relationParam)
		if err != nil {
			return err
		}

	}

	d.SetId(cloudCtxProfile.DistinguishedName)
	return resourceAciCloudcontextprofileRead(d, m)
}

func resourceAciCloudcontextprofileUpdate(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)

	name := d.Get("name").(string)

	TenantDn := d.Get("tenant_dn").(string)

	cloudCtxProfileAttr := models.CloudcontextprofileAttributes{}
	if Annotation, ok := d.GetOk("annotation"); ok {
		cloudCtxProfileAttr.Annotation = Annotation.(string)
	}
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		cloudCtxProfileAttr.NameAlias = NameAlias.(string)
	}
	if Type, ok := d.GetOk("type"); ok {
		cloudCtxProfileAttr.Type = Type.(string)
	}
	cloudCtxProfile := models.NewCloudcontextprofile(fmt.Sprintf("ctxprofile-%s", name), TenantDn, desc, cloudCtxProfileAttr)

	cloudCtxProfile.Status = "modified"

	err := aciClient.Save(cloudCtxProfile)

	if err != nil {
		return err
	}

	if d.HasChange("relation_cloud_rs_ctx_to_flow_log") {
		_, newRelParam := d.GetChange("relation_cloud_rs_ctx_to_flow_log")
		err = aciClient.DeleteRelationcloudRsCtxToFlowLogFromCloudcontextprofile(cloudCtxProfile.DistinguishedName)
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationcloudRsCtxToFlowLogFromCloudcontextprofile(cloudCtxProfile.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_cloud_rs_to_ctx") {
		_, newRelParam := d.GetChange("relation_cloud_rs_to_ctx")
		err = aciClient.CreateRelationcloudRsToCtxFromCloudcontextprofile(cloudCtxProfile.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}
	if d.HasChange("relation_cloud_rs_ctx_profile_to_region") {
		_, newRelParam := d.GetChange("relation_cloud_rs_ctx_profile_to_region")
		err = aciClient.DeleteRelationcloudRsCtxProfileToRegionFromCloudcontextprofile(cloudCtxProfile.DistinguishedName)
		if err != nil {
			return err
		}
		err = aciClient.CreateRelationcloudRsCtxProfileToRegionFromCloudcontextprofile(cloudCtxProfile.DistinguishedName, newRelParam.(string))
		if err != nil {
			return err
		}

	}

	d.SetId(cloudCtxProfile.DistinguishedName)
	return resourceAciCloudcontextprofileRead(d, m)

}

func resourceAciCloudcontextprofileRead(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)

	dn := d.Id()
	cloudCtxProfile, err := getRemoteCloudcontextprofile(aciClient, dn)

	if err != nil {
		return err
	}
	setCloudcontextprofileAttributes(cloudCtxProfile, d)
	return nil
}

func resourceAciCloudcontextprofileDelete(d *schema.ResourceData, m interface{}) error {
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "cloudCtxProfile")
	if err != nil {
		return err
	}

	d.SetId("")
	return err
}
