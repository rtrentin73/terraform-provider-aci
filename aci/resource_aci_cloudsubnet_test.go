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

func TestAccAciCloudSubnet_Basic(t *testing.T) {
	var cloud_subnet models.CloudSubnet
	fv_tenant_name := acctest.RandString(5)
	cloud_ctx_profile_name := acctest.RandString(5)
	cloud_cidr_name := acctest.RandString(5)
	cloud_subnet_name := acctest.RandString(5)
	description := "cloud_subnet created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudSubnetConfig_basic(fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name, cloud_subnet_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudSubnetExists("aci_cloud_subnet.foocloud_subnet", &cloud_subnet),
					testAccCheckAciCloudSubnetAttributes(fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name, cloud_subnet_name, description, &cloud_subnet),
				),
			},
		},
	})
}

func testAccCheckAciCloudSubnetConfig_basic(fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name, cloud_subnet_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_cloud_context_profile" "foocloud_context_profile" {
		name 		= "%s"
		description = "cloud_context_profile created while acceptance testing"
		tenant_dn = "${aci_tenant.footenant.id}"
	}

	resource "aci_cloud_cidr_pool" "foocloud_cidr_pool" {
		name 		= "%s"
		description = "cloud_cidr_pool created while acceptance testing"
		cloud_context_profile_dn = "${aci_cloud_context_profile.foocloud_context_profile.id}"
	}

	resource "aci_cloud_subnet" "foocloud_subnet" {
		name 		= "%s"
		description = "cloud_subnet created while acceptance testing"
		cloud_cidr_pool_dn = "${aci_cloud_cidr_pool.foocloud_cidr_pool.id}"
	}

	`, fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name, cloud_subnet_name)
}

func testAccCheckAciCloudSubnetExists(name string, cloud_subnet *models.CloudSubnet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud Subnet %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud Subnet dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloud_subnetFound := models.CloudSubnetFromContainer(cont)
		if cloud_subnetFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud Subnet %s not found", rs.Primary.ID)
		}
		*cloud_subnet = *cloud_subnetFound
		return nil
	}
}

func testAccCheckAciCloudSubnetDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloud_subnet" {
			cont, err := client.Get(rs.Primary.ID)
			cloud_subnet := models.CloudSubnetFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud Subnet %s Still exists", cloud_subnet.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudSubnetAttributes(fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name, cloud_subnet_name, description string, cloud_subnet *models.CloudSubnet) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_subnet_name != GetMOName(cloud_subnet.DistinguishedName) {
			return fmt.Errorf("Bad cloud_subnet %s", GetMOName(cloud_subnet.DistinguishedName))
		}

		if cloud_cidr_name != GetMOName(GetParentDn(cloud_subnet.DistinguishedName)) {
			return fmt.Errorf(" Bad cloud_cidr %s", GetMOName(GetParentDn(cloud_subnet.DistinguishedName)))
		}
		if description != cloud_subnet.Description {
			return fmt.Errorf("Bad cloud_subnet Description %s", cloud_subnet.Description)
		}

		return nil
	}
}
