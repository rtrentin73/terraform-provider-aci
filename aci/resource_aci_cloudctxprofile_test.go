package aci

import (
	"fmt"
	"testing"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAciCloudcontextprofile_Basic(t *testing.T) {
	var cloudcontextprofile models.Cloudcontextprofile
	fv_tenant_name := acctest.RandString(5)
	cloud_ctx_profile_name := acctest.RandString(5)
	description := "cloudcontextprofile created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudcontextprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudcontextprofileConfig_basic(fv_tenant_name, cloud_ctx_profile_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudcontextprofileExists("aci_cloudcontextprofile.foocloudcontextprofile", &cloudcontextprofile),
					testAccCheckAciCloudcontextprofileAttributes(fv_tenant_name, cloud_ctx_profile_name, description, &cloudcontextprofile),
				),
			},
		},
	})
}

func testAccCheckAciCloudcontextprofileConfig_basic(fv_tenant_name, cloud_ctx_profile_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_cloudcontextprofile" "foocloudcontextprofile" {
		name 		= "%s"
		description = "cloudcontextprofile created while acceptance testing"
		tenant_dn = "${aci_tenant.footenant.id}"
	}

	`, fv_tenant_name, cloud_ctx_profile_name)
}

func testAccCheckAciCloudcontextprofileExists(name string, cloudcontextprofile *models.Cloudcontextprofile) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud context profile %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud context profile dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloudcontextprofileFound := models.CloudcontextprofileFromContainer(cont)
		if cloudcontextprofileFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud context profile %s not found", rs.Primary.ID)
		}
		*cloudcontextprofile = *cloudcontextprofileFound
		return nil
	}
}

func testAccCheckAciCloudcontextprofileDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloudcontextprofile" {
			cont, err := client.Get(rs.Primary.ID)
			cloudcontextprofile := models.CloudcontextprofileFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud context profile %s Still exists", cloudcontextprofile.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudcontextprofileAttributes(fv_tenant_name, cloud_ctx_profile_name, description string, cloudcontextprofile *models.Cloudcontextprofile) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_ctx_profile_name != GetMOName(cloudcontextprofile.DistinguishedName) {
			return fmt.Errorf("Bad cloud_ctx_profile %s", GetMOName(cloudcontextprofile.DistinguishedName))
		}

		if fv_tenant_name != GetMOName(GetParentDn(cloudcontextprofile.DistinguishedName)) {
			return fmt.Errorf(" Bad fv_tenant %s", GetMOName(GetParentDn(cloudcontextprofile.DistinguishedName)))
		}
		if description != cloudcontextprofile.Description {
			return fmt.Errorf("Bad cloudcontextprofile Description %s", cloudcontextprofile.Description)
		}

		return nil
	}
}
