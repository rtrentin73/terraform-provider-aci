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

func TestAccAciCloudcidrpool_Basic(t *testing.T) {
	var cloudcidrpool models.Cloudcidrpool
	fv_tenant_name := acctest.RandString(5)
	cloud_ctx_profile_name := acctest.RandString(5)
	cloud_cidr_name := acctest.RandString(5)
	description := "cloudcidrpool created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudcidrpoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudcidrpoolConfig_basic(fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudcidrpoolExists("aci_cloudcidrpool.foocloudcidrpool", &cloudcidrpool),
					testAccCheckAciCloudcidrpoolAttributes(fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name, description, &cloudcidrpool),
				),
			},
		},
	})
}

func testAccCheckAciCloudcidrpoolConfig_basic(fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name string) string {
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

	resource "aci_cloudcidrpool" "foocloudcidrpool" {
		name 		= "%s"
		description = "cloudcidrpool created while acceptance testing"
		cloudcontextprofile_dn = "${aci_cloudcontextprofile.foocloudcontextprofile.id}"
	}

	`, fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name)
}

func testAccCheckAciCloudcidrpoolExists(name string, cloudcidrpool *models.Cloudcidrpool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud cidr pool %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud cidr pool dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloudcidrpoolFound := models.CloudcidrpoolFromContainer(cont)
		if cloudcidrpoolFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud cidr pool %s not found", rs.Primary.ID)
		}
		*cloudcidrpool = *cloudcidrpoolFound
		return nil
	}
}

func testAccCheckAciCloudcidrpoolDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloudcidrpool" {
			cont, err := client.Get(rs.Primary.ID)
			cloudcidrpool := models.CloudcidrpoolFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud cidr pool %s Still exists", cloudcidrpool.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudcidrpoolAttributes(fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name, description string, cloudcidrpool *models.Cloudcidrpool) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_cidr_name != GetMOName(cloudcidrpool.DistinguishedName) {
			return fmt.Errorf("Bad cloud_cidr %s", GetMOName(cloudcidrpool.DistinguishedName))
		}

		if cloud_ctx_profile_name != GetMOName(GetParentDn(cloudcidrpool.DistinguishedName)) {
			return fmt.Errorf(" Bad cloud_ctx_profile %s", GetMOName(GetParentDn(cloudcidrpool.DistinguishedName)))
		}
		if description != cloudcidrpool.Description {
			return fmt.Errorf("Bad cloudcidrpool Description %s", cloudcidrpool.Description)
		}

		return nil
	}
}
