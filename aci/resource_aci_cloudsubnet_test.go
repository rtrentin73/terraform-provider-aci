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

func TestAccAciCloudsubnet_Basic(t *testing.T) {
	var cloudsubnet models.Cloudsubnet
	fv_tenant_name := acctest.RandString(5)
	cloud_ctx_profile_name := acctest.RandString(5)
	cloud_cidr_name := acctest.RandString(5)
	cloud_subnet_name := acctest.RandString(5)
	description := "cloudsubnet created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudsubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudsubnetConfig_basic(fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name, cloud_subnet_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudsubnetExists("aci_cloudsubnet.foocloudsubnet", &cloudsubnet),
					testAccCheckAciCloudsubnetAttributes(fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name, cloud_subnet_name, description, &cloudsubnet),
				),
			},
		},
	})
}

func testAccCheckAciCloudsubnetConfig_basic(fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name, cloud_subnet_name string) string {
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

	resource "aci_cloudsubnet" "foocloudsubnet" {
		name 		= "%s"
		description = "cloudsubnet created while acceptance testing"
		cloudcidrpool_dn = "${aci_cloudcidrpool.foocloudcidrpool.id}"
	}

	`, fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name, cloud_subnet_name)
}

func testAccCheckAciCloudsubnetExists(name string, cloudsubnet *models.Cloudsubnet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud subnet %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud subnet dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloudsubnetFound := models.CloudsubnetFromContainer(cont)
		if cloudsubnetFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud subnet %s not found", rs.Primary.ID)
		}
		*cloudsubnet = *cloudsubnetFound
		return nil
	}
}

func testAccCheckAciCloudsubnetDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloudsubnet" {
			cont, err := client.Get(rs.Primary.ID)
			cloudsubnet := models.CloudsubnetFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud subnet %s Still exists", cloudsubnet.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudsubnetAttributes(fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name, cloud_subnet_name, description string, cloudsubnet *models.Cloudsubnet) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_subnet_name != GetMOName(cloudsubnet.DistinguishedName) {
			return fmt.Errorf("Bad cloud_subnet %s", GetMOName(cloudsubnet.DistinguishedName))
		}

		if cloud_cidr_name != GetMOName(GetParentDn(cloudsubnet.DistinguishedName)) {
			return fmt.Errorf(" Bad cloud_cidr %s", GetMOName(GetParentDn(cloudsubnet.DistinguishedName)))
		}
		if description != cloudsubnet.Description {
			return fmt.Errorf("Bad cloudsubnet Description %s", cloudsubnet.Description)
		}

		return nil
	}
}
