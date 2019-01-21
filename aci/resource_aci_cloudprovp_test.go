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

func TestAccAciCloudproviderprofile_Basic(t *testing.T) {
	var cloudproviderprofile models.Cloudproviderprofile
	cloud_prov_p_name := acctest.RandString(5)
	description := "cloudproviderprofile created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudproviderprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudproviderprofileConfig_basic(cloud_prov_p_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudproviderprofileExists("aci_cloudproviderprofile.foocloudproviderprofile", &cloudproviderprofile),
					testAccCheckAciCloudproviderprofileAttributes(cloud_prov_p_name, description, &cloudproviderprofile),
				),
			},
		},
	})
}

func testAccCheckAciCloudproviderprofileConfig_basic(cloud_prov_p_name string) string {
	return fmt.Sprintf(`

	resource "aci_cloudproviderprofile" "foocloudproviderprofile" {
		name 		= "%s"
		description = "cloudproviderprofile created while acceptance testing"

	}

	`, cloud_prov_p_name)
}

func testAccCheckAciCloudproviderprofileExists(name string, cloudproviderprofile *models.Cloudproviderprofile) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud provider profile %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud provider profile dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloudproviderprofileFound := models.CloudproviderprofileFromContainer(cont)
		if cloudproviderprofileFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud provider profile %s not found", rs.Primary.ID)
		}
		*cloudproviderprofile = *cloudproviderprofileFound
		return nil
	}
}

func testAccCheckAciCloudproviderprofileDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloudproviderprofile" {
			cont, err := client.Get(rs.Primary.ID)
			cloudproviderprofile := models.CloudproviderprofileFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud provider profile %s Still exists", cloudproviderprofile.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudproviderprofileAttributes(cloud_prov_p_name, description string, cloudproviderprofile *models.Cloudproviderprofile) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_prov_p_name != GetMOName(cloudproviderprofile.DistinguishedName) {
			return fmt.Errorf("Bad cloud_prov_p %s", GetMOName(cloudproviderprofile.DistinguishedName))
		}

		if description != cloudproviderprofile.Description {
			return fmt.Errorf("Bad cloudproviderprofile Description %s", cloudproviderprofile.Description)
		}

		return nil
	}
}
