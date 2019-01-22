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

func TestAccAciCloudProviderProfile_Basic(t *testing.T) {
	var cloud_provider_profile models.CloudProviderProfile
	cloud_prov_p_name := acctest.RandString(5)
	description := "cloud_provider_profile created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudProviderProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudProviderProfileConfig_basic(cloud_prov_p_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudProviderProfileExists("aci_cloud_provider_profile.foocloud_provider_profile", &cloud_provider_profile),
					testAccCheckAciCloudProviderProfileAttributes(cloud_prov_p_name, description, &cloud_provider_profile),
				),
			},
		},
	})
}

func testAccCheckAciCloudProviderProfileConfig_basic(cloud_prov_p_name string) string {
	return fmt.Sprintf(`

	resource "aci_cloud_provider_profile" "foocloud_provider_profile" {
		name 		= "%s"
		description = "cloud_provider_profile created while acceptance testing"

	}

	`, cloud_prov_p_name)
}

func testAccCheckAciCloudProviderProfileExists(name string, cloud_provider_profile *models.CloudProviderProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud Provider Profile %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud Provider Profile dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloud_provider_profileFound := models.CloudProviderProfileFromContainer(cont)
		if cloud_provider_profileFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud Provider Profile %s not found", rs.Primary.ID)
		}
		*cloud_provider_profile = *cloud_provider_profileFound
		return nil
	}
}

func testAccCheckAciCloudProviderProfileDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloud_provider_profile" {
			cont, err := client.Get(rs.Primary.ID)
			cloud_provider_profile := models.CloudProviderProfileFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud Provider Profile %s Still exists", cloud_provider_profile.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudProviderProfileAttributes(cloud_prov_p_name, description string, cloud_provider_profile *models.CloudProviderProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_prov_p_name != GetMOName(cloud_provider_profile.DistinguishedName) {
			return fmt.Errorf("Bad cloud_prov_p %s", GetMOName(cloud_provider_profile.DistinguishedName))
		}

		if description != cloud_provider_profile.Description {
			return fmt.Errorf("Bad cloud_provider_profile Description %s", cloud_provider_profile.Description)
		}

		return nil
	}
}
