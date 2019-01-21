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

func TestAccAciCloudavailabilityzone_Basic(t *testing.T) {
	var cloudavailabilityzone models.Cloudavailabilityzone
	cloud_prov_p_name := acctest.RandString(5)
	cloud_region_name := acctest.RandString(5)
	cloud_zone_name := acctest.RandString(5)
	description := "cloudavailabilityzone created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudavailabilityzoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudavailabilityzoneConfig_basic(cloud_prov_p_name, cloud_region_name, cloud_zone_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudavailabilityzoneExists("aci_cloudavailabilityzone.foocloudavailabilityzone", &cloudavailabilityzone),
					testAccCheckAciCloudavailabilityzoneAttributes(cloud_prov_p_name, cloud_region_name, cloud_zone_name, description, &cloudavailabilityzone),
				),
			},
		},
	})
}

func testAccCheckAciCloudavailabilityzoneConfig_basic(cloud_prov_p_name, cloud_region_name, cloud_zone_name string) string {
	return fmt.Sprintf(`

	resource "aci_cloudproviderprofile" "foocloudproviderprofile" {
		name 		= "%s"
		description = "cloudproviderprofile created while acceptance testing"

	}

	resource "aci_cloudprovidersregion" "foocloudprovidersregion" {
		name 		= "%s"
		description = "cloudprovidersregion created while acceptance testing"
		cloudproviderprofile_dn = "${aci_cloudproviderprofile.foocloudproviderprofile.id}"
	}

	resource "aci_cloudavailabilityzone" "foocloudavailabilityzone" {
		name 		= "%s"
		description = "cloudavailabilityzone created while acceptance testing"
		cloudprovidersregion_dn = "${aci_cloudprovidersregion.foocloudprovidersregion.id}"
	}

	`, cloud_prov_p_name, cloud_region_name, cloud_zone_name)
}

func testAccCheckAciCloudavailabilityzoneExists(name string, cloudavailabilityzone *models.Cloudavailabilityzone) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud availability zone %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud availability zone dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloudavailabilityzoneFound := models.CloudavailabilityzoneFromContainer(cont)
		if cloudavailabilityzoneFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud availability zone %s not found", rs.Primary.ID)
		}
		*cloudavailabilityzone = *cloudavailabilityzoneFound
		return nil
	}
}

func testAccCheckAciCloudavailabilityzoneDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloudavailabilityzone" {
			cont, err := client.Get(rs.Primary.ID)
			cloudavailabilityzone := models.CloudavailabilityzoneFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud availability zone %s Still exists", cloudavailabilityzone.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudavailabilityzoneAttributes(cloud_prov_p_name, cloud_region_name, cloud_zone_name, description string, cloudavailabilityzone *models.Cloudavailabilityzone) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_zone_name != GetMOName(cloudavailabilityzone.DistinguishedName) {
			return fmt.Errorf("Bad cloud_zone %s", GetMOName(cloudavailabilityzone.DistinguishedName))
		}

		if cloud_region_name != GetMOName(GetParentDn(cloudavailabilityzone.DistinguishedName)) {
			return fmt.Errorf(" Bad cloud_region %s", GetMOName(GetParentDn(cloudavailabilityzone.DistinguishedName)))
		}
		if description != cloudavailabilityzone.Description {
			return fmt.Errorf("Bad cloudavailabilityzone Description %s", cloudavailabilityzone.Description)
		}

		return nil
	}
}
