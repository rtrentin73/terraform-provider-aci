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

func TestAccAciCloudCIDRPool_Basic(t *testing.T) {
	var cloud_cidr_pool models.CloudCIDRPool
	fv_tenant_name := acctest.RandString(5)
	cloud_ctx_profile_name := acctest.RandString(5)
	cloud_cidr_name := acctest.RandString(5)
	description := "cloud_cidr_pool created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudCIDRPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudCIDRPoolConfig_basic(fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudCIDRPoolExists("aci_cloud_cidr_pool.foocloud_cidr_pool", &cloud_cidr_pool),
					testAccCheckAciCloudCIDRPoolAttributes(fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name, description, &cloud_cidr_pool),
				),
			},
		},
	})
}

func testAccCheckAciCloudCIDRPoolConfig_basic(fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name string) string {
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

	`, fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name)
}

func testAccCheckAciCloudCIDRPoolExists(name string, cloud_cidr_pool *models.CloudCIDRPool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud CIDR Pool %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud CIDR Pool dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloud_cidr_poolFound := models.CloudCIDRPoolFromContainer(cont)
		if cloud_cidr_poolFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud CIDR Pool %s not found", rs.Primary.ID)
		}
		*cloud_cidr_pool = *cloud_cidr_poolFound
		return nil
	}
}

func testAccCheckAciCloudCIDRPoolDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloud_cidr_pool" {
			cont, err := client.Get(rs.Primary.ID)
			cloud_cidr_pool := models.CloudCIDRPoolFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud CIDR Pool %s Still exists", cloud_cidr_pool.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudCIDRPoolAttributes(fv_tenant_name, cloud_ctx_profile_name, cloud_cidr_name, description string, cloud_cidr_pool *models.CloudCIDRPool) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_cidr_name != GetMOName(cloud_cidr_pool.DistinguishedName) {
			return fmt.Errorf("Bad cloud_cidr %s", GetMOName(cloud_cidr_pool.DistinguishedName))
		}

		if cloud_ctx_profile_name != GetMOName(GetParentDn(cloud_cidr_pool.DistinguishedName)) {
			return fmt.Errorf(" Bad cloud_ctx_profile %s", GetMOName(GetParentDn(cloud_cidr_pool.DistinguishedName)))
		}
		if description != cloud_cidr_pool.Description {
			return fmt.Errorf("Bad cloud_cidr_pool Description %s", cloud_cidr_pool.Description)
		}

		return nil
	}
}
