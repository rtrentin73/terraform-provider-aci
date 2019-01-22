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

func TestAccAciCloudAvailabilityZone_Basic(t *testing.T) {
	var cloud_availability_zone models.CloudAvailabilityZone
	cloud_prov_p_name := acctest.RandString(5)
	cloud_region_name := acctest.RandString(5)
	cloud_zone_name := acctest.RandString(5)
	description := "cloud_availability_zone created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudAvailabilityZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudAvailabilityZoneConfig_basic(cloud_prov_p_name, cloud_region_name, cloud_zone_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudAvailabilityZoneExists("aci_cloud_availability_zone.foocloud_availability_zone", &cloud_availability_zone),
					testAccCheckAciCloudAvailabilityZoneAttributes(cloud_prov_p_name, cloud_region_name, cloud_zone_name, description, &cloud_availability_zone),
				),
			},
		},
	})
}

func testAccCheckAciCloudAvailabilityZoneConfig_basic(cloud_prov_p_name, cloud_region_name, cloud_zone_name string) string {
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

	resource "aci_cloud_availability_zone" "foocloud_availability_zone" {
		name 		= "%s"
		description = "cloud_availability_zone created while acceptance testing"
		cloud_providers_region_dn = "${aci_cloud_providers_region.foocloud_providers_region.id}"
	}

	`, cloud_prov_p_name, cloud_region_name, cloud_zone_name)
}

func testAccCheckAciCloudAvailabilityZoneExists(name string, cloud_availability_zone *models.CloudAvailabilityZone) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud Availability Zone %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud Availability Zone dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloud_availability_zoneFound := models.CloudAvailabilityZoneFromContainer(cont)
		if cloud_availability_zoneFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud Availability Zone %s not found", rs.Primary.ID)
		}
		*cloud_availability_zone = *cloud_availability_zoneFound
		return nil
	}
}

func testAccCheckAciCloudAvailabilityZoneDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloud_availability_zone" {
			cont, err := client.Get(rs.Primary.ID)
			cloud_availability_zone := models.CloudAvailabilityZoneFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud Availability Zone %s Still exists", cloud_availability_zone.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudAvailabilityZoneAttributes(cloud_prov_p_name, cloud_region_name, cloud_zone_name, description string, cloud_availability_zone *models.CloudAvailabilityZone) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_zone_name != GetMOName(cloud_availability_zone.DistinguishedName) {
			return fmt.Errorf("Bad cloud_zone %s", GetMOName(cloud_availability_zone.DistinguishedName))
		}

		if cloud_region_name != GetMOName(GetParentDn(cloud_availability_zone.DistinguishedName)) {
			return fmt.Errorf(" Bad cloud_region %s", GetMOName(GetParentDn(cloud_availability_zone.DistinguishedName)))
		}
		if description != cloud_availability_zone.Description {
			return fmt.Errorf("Bad cloud_availability_zone Description %s", cloud_availability_zone.Description)
		}

		return nil
	}
}
