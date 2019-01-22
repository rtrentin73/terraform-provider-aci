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

func TestAccAciAutonomousSystemProfile_Basic(t *testing.T) {
	var autonomous_system_profile models.AutonomousSystemProfile
	cloud_bgp_as_p_name := acctest.RandString(5)
	description := "autonomous_system_profile created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciAutonomousSystemProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciAutonomousSystemProfileConfig_basic(cloud_bgp_as_p_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciAutonomousSystemProfileExists("aci_autonomous_system_profile.fooautonomous_system_profile", &autonomous_system_profile),
					testAccCheckAciAutonomousSystemProfileAttributes(cloud_bgp_as_p_name, description, &autonomous_system_profile),
				),
			},
		},
	})
}

func testAccCheckAciAutonomousSystemProfileConfig_basic(cloud_bgp_as_p_name string) string {
	return fmt.Sprintf(`

	resource "aci_autonomous_system_profile" "fooautonomous_system_profile" {
		name 		= "%s"
		description = "autonomous_system_profile created while acceptance testing"

	}

	`, cloud_bgp_as_p_name)
}

func testAccCheckAciAutonomousSystemProfileExists(name string, autonomous_system_profile *models.AutonomousSystemProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Autonomous System Profile %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Autonomous System Profile dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		autonomous_system_profileFound := models.AutonomousSystemProfileFromContainer(cont)
		if autonomous_system_profileFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Autonomous System Profile %s not found", rs.Primary.ID)
		}
		*autonomous_system_profile = *autonomous_system_profileFound
		return nil
	}
}

func testAccCheckAciAutonomousSystemProfileDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_autonomous_system_profile" {
			cont, err := client.Get(rs.Primary.ID)
			autonomous_system_profile := models.AutonomousSystemProfileFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Autonomous System Profile %s Still exists", autonomous_system_profile.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciAutonomousSystemProfileAttributes(cloud_bgp_as_p_name, description string, autonomous_system_profile *models.AutonomousSystemProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_bgp_as_p_name != GetMOName(autonomous_system_profile.DistinguishedName) {
			return fmt.Errorf("Bad cloud_bgp_as_p %s", GetMOName(autonomous_system_profile.DistinguishedName))
		}

		if description != autonomous_system_profile.Description {
			return fmt.Errorf("Bad autonomous_system_profile Description %s", autonomous_system_profile.Description)
		}

		return nil
	}
}
