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

func TestAccAciCloudDomainProfile_Basic(t *testing.T) {
	var cloud_domain_profile models.CloudDomainProfile
	cloud_dom_p_name := acctest.RandString(5)
	description := "cloud_domain_profile created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudDomainProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudDomainProfileConfig_basic(cloud_dom_p_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudDomainProfileExists("aci_cloud_domain_profile.foocloud_domain_profile", &cloud_domain_profile),
					testAccCheckAciCloudDomainProfileAttributes(cloud_dom_p_name, description, &cloud_domain_profile),
				),
			},
		},
	})
}

func testAccCheckAciCloudDomainProfileConfig_basic(cloud_dom_p_name string) string {
	return fmt.Sprintf(`

	resource "aci_cloud_domain_profile" "foocloud_domain_profile" {
		name 		= "%s"
		description = "cloud_domain_profile created while acceptance testing"

	}

	`, cloud_dom_p_name)
}

func testAccCheckAciCloudDomainProfileExists(name string, cloud_domain_profile *models.CloudDomainProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud Domain Profile %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud Domain Profile dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloud_domain_profileFound := models.CloudDomainProfileFromContainer(cont)
		if cloud_domain_profileFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud Domain Profile %s not found", rs.Primary.ID)
		}
		*cloud_domain_profile = *cloud_domain_profileFound
		return nil
	}
}

func testAccCheckAciCloudDomainProfileDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloud_domain_profile" {
			cont, err := client.Get(rs.Primary.ID)
			cloud_domain_profile := models.CloudDomainProfileFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud Domain Profile %s Still exists", cloud_domain_profile.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudDomainProfileAttributes(cloud_dom_p_name, description string, cloud_domain_profile *models.CloudDomainProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_dom_p_name != GetMOName(cloud_domain_profile.DistinguishedName) {
			return fmt.Errorf("Bad cloud_dom_p %s", GetMOName(cloud_domain_profile.DistinguishedName))
		}

		if description != cloud_domain_profile.Description {
			return fmt.Errorf("Bad cloud_domain_profile Description %s", cloud_domain_profile.Description)
		}

		return nil
	}
}
