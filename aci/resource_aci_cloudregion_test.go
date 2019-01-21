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

func TestAccAciCloudprovidersregion_Basic(t *testing.T) {
	var cloudprovidersregion models.Cloudprovidersregion
	cloud_prov_p_name := acctest.RandString(5)
	cloud_region_name := acctest.RandString(5)
	description := "cloudprovidersregion created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudprovidersregionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudprovidersregionConfig_basic(cloud_prov_p_name, cloud_region_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudprovidersregionExists("aci_cloudprovidersregion.foocloudprovidersregion", &cloudprovidersregion),
					testAccCheckAciCloudprovidersregionAttributes(cloud_prov_p_name, cloud_region_name, description, &cloudprovidersregion),
				),
			},
		},
	})
}

func testAccCheckAciCloudprovidersregionConfig_basic(cloud_prov_p_name, cloud_region_name string) string {
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

	`, cloud_prov_p_name, cloud_region_name)
}

func testAccCheckAciCloudprovidersregionExists(name string, cloudprovidersregion *models.Cloudprovidersregion) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud providers region %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud providers region dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloudprovidersregionFound := models.CloudprovidersregionFromContainer(cont)
		if cloudprovidersregionFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud providers region %s not found", rs.Primary.ID)
		}
		*cloudprovidersregion = *cloudprovidersregionFound
		return nil
	}
}

func testAccCheckAciCloudprovidersregionDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloudprovidersregion" {
			cont, err := client.Get(rs.Primary.ID)
			cloudprovidersregion := models.CloudprovidersregionFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud providers region %s Still exists", cloudprovidersregion.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudprovidersregionAttributes(cloud_prov_p_name, cloud_region_name, description string, cloudprovidersregion *models.Cloudprovidersregion) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_region_name != GetMOName(cloudprovidersregion.DistinguishedName) {
			return fmt.Errorf("Bad cloud_region %s", GetMOName(cloudprovidersregion.DistinguishedName))
		}

		if cloud_prov_p_name != GetMOName(GetParentDn(cloudprovidersregion.DistinguishedName)) {
			return fmt.Errorf(" Bad cloud_prov_p %s", GetMOName(GetParentDn(cloudprovidersregion.DistinguishedName)))
		}
		if description != cloudprovidersregion.Description {
			return fmt.Errorf("Bad cloudprovidersregion Description %s", cloudprovidersregion.Description)
		}

		return nil
	}
}
