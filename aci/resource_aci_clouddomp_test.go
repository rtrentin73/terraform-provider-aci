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

func TestAccAciClouddomainprofile_Basic(t *testing.T) {
	var clouddomainprofile models.Clouddomainprofile
	cloud_dom_p_name := acctest.RandString(5)
	description := "clouddomainprofile created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciClouddomainprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciClouddomainprofileConfig_basic(cloud_dom_p_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciClouddomainprofileExists("aci_clouddomainprofile.fooclouddomainprofile", &clouddomainprofile),
					testAccCheckAciClouddomainprofileAttributes(cloud_dom_p_name, description, &clouddomainprofile),
				),
			},
		},
	})
}

func testAccCheckAciClouddomainprofileConfig_basic(cloud_dom_p_name string) string {
	return fmt.Sprintf(`

	resource "aci_clouddomainprofile" "fooclouddomainprofile" {
		name 		= "%s"
		description = "clouddomainprofile created while acceptance testing"

	}

	`, cloud_dom_p_name)
}

func testAccCheckAciClouddomainprofileExists(name string, clouddomainprofile *models.Clouddomainprofile) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud domain profile %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud domain profile dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		clouddomainprofileFound := models.ClouddomainprofileFromContainer(cont)
		if clouddomainprofileFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud domain profile %s not found", rs.Primary.ID)
		}
		*clouddomainprofile = *clouddomainprofileFound
		return nil
	}
}

func testAccCheckAciClouddomainprofileDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_clouddomainprofile" {
			cont, err := client.Get(rs.Primary.ID)
			clouddomainprofile := models.ClouddomainprofileFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud domain profile %s Still exists", clouddomainprofile.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciClouddomainprofileAttributes(cloud_dom_p_name, description string, clouddomainprofile *models.Clouddomainprofile) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_dom_p_name != GetMOName(clouddomainprofile.DistinguishedName) {
			return fmt.Errorf("Bad cloud_dom_p %s", GetMOName(clouddomainprofile.DistinguishedName))
		}

		if description != clouddomainprofile.Description {
			return fmt.Errorf("Bad clouddomainprofile Description %s", clouddomainprofile.Description)
		}

		return nil
	}
}
