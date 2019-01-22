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

func TestAccAciCloudProvidersRegion_Basic(t *testing.T) {
	var cloud_providers_region models.CloudProvidersRegion
	cloud_prov_p_name := acctest.RandString(5)
	cloud_region_name := acctest.RandString(5)
	description := "cloud_providers_region created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudProvidersRegionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudProvidersRegionConfig_basic(cloud_prov_p_name, cloud_region_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudProvidersRegionExists("aci_cloud_providers_region.foocloud_providers_region", &cloud_providers_region),
					testAccCheckAciCloudProvidersRegionAttributes(cloud_prov_p_name, cloud_region_name, description, &cloud_providers_region),
				),
			},
		},
	})
}

func testAccCheckAciCloudProvidersRegionConfig_basic(cloud_prov_p_name, cloud_region_name string) string {
	return fmt.Sprintf(`

	resource "aci_cloud_provider_profile" "foocloud_provider_profile" {
		name 		= "%s"
		description = "cloud_provider_profile created while acceptance testing"

	}

	resource "aci_cloud_providers_region" "foocloud_providers_region" {
		name 		= "%s"
		description = "cloud_providers_region created while acceptance testing"
		cloud_provider_profile_dn = "${aci_cloud_provider_profile.foocloud_provider_profile.id}"
	}

	`, cloud_prov_p_name, cloud_region_name)
}

func testAccCheckAciCloudProvidersRegionExists(name string, cloud_providers_region *models.CloudProvidersRegion) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud Providers Region %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud Providers Region dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloud_providers_regionFound := models.CloudProvidersRegionFromContainer(cont)
		if cloud_providers_regionFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud Providers Region %s not found", rs.Primary.ID)
		}
		*cloud_providers_region = *cloud_providers_regionFound
		return nil
	}
}

func testAccCheckAciCloudProvidersRegionDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloud_providers_region" {
			cont, err := client.Get(rs.Primary.ID)
			cloud_providers_region := models.CloudProvidersRegionFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud Providers Region %s Still exists", cloud_providers_region.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudProvidersRegionAttributes(cloud_prov_p_name, cloud_region_name, description string, cloud_providers_region *models.CloudProvidersRegion) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_region_name != GetMOName(cloud_providers_region.DistinguishedName) {
			return fmt.Errorf("Bad cloud_region %s", GetMOName(cloud_providers_region.DistinguishedName))
		}

		if cloud_prov_p_name != GetMOName(GetParentDn(cloud_providers_region.DistinguishedName)) {
			return fmt.Errorf(" Bad cloud_prov_p %s", GetMOName(GetParentDn(cloud_providers_region.DistinguishedName)))
		}
		if description != cloud_providers_region.Description {
			return fmt.Errorf("Bad cloud_providers_region Description %s", cloud_providers_region.Description)
		}

		return nil
	}
}
