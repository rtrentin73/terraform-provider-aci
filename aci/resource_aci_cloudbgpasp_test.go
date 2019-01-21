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

func TestAccAciAutonomoussystemprofile_Basic(t *testing.T) {
	var autonomoussystemprofile models.Autonomoussystemprofile
	cloud_bgp_as_p_name := acctest.RandString(5)
	description := "autonomoussystemprofile created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciAutonomoussystemprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciAutonomoussystemprofileConfig_basic(cloud_bgp_as_p_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciAutonomoussystemprofileExists("aci_autonomoussystemprofile.fooautonomoussystemprofile", &autonomoussystemprofile),
					testAccCheckAciAutonomoussystemprofileAttributes(cloud_bgp_as_p_name, description, &autonomoussystemprofile),
				),
			},
		},
	})
}

func testAccCheckAciAutonomoussystemprofileConfig_basic(cloud_bgp_as_p_name string) string {
	return fmt.Sprintf(`

	resource "aci_autonomoussystemprofile" "fooautonomoussystemprofile" {
		name 		= "%s"
		description = "autonomoussystemprofile created while acceptance testing"

	}

	`, cloud_bgp_as_p_name)
}

func testAccCheckAciAutonomoussystemprofileExists(name string, autonomoussystemprofile *models.Autonomoussystemprofile) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Autonomous system profile %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Autonomous system profile dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		autonomoussystemprofileFound := models.AutonomoussystemprofileFromContainer(cont)
		if autonomoussystemprofileFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Autonomous system profile %s not found", rs.Primary.ID)
		}
		*autonomoussystemprofile = *autonomoussystemprofileFound
		return nil
	}
}

func testAccCheckAciAutonomoussystemprofileDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_autonomoussystemprofile" {
			cont, err := client.Get(rs.Primary.ID)
			autonomoussystemprofile := models.AutonomoussystemprofileFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Autonomous system profile %s Still exists", autonomoussystemprofile.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciAutonomoussystemprofileAttributes(cloud_bgp_as_p_name, description string, autonomoussystemprofile *models.Autonomoussystemprofile) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_bgp_as_p_name != GetMOName(autonomoussystemprofile.DistinguishedName) {
			return fmt.Errorf("Bad cloud_bgp_as_p %s", GetMOName(autonomoussystemprofile.DistinguishedName))
		}

		if description != autonomoussystemprofile.Description {
			return fmt.Errorf("Bad autonomoussystemprofile Description %s", autonomoussystemprofile.Description)
		}

		return nil
	}
}
