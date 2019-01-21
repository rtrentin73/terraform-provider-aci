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

func TestAccAciCloudapplicationcontainer_Basic(t *testing.T) {
	var cloudapplicationcontainer models.Cloudapplicationcontainer
	fv_tenant_name := acctest.RandString(5)
	cloud_app_name := acctest.RandString(5)
	description := "cloudapplicationcontainer created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudapplicationcontainerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudapplicationcontainerConfig_basic(fv_tenant_name, cloud_app_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudapplicationcontainerExists("aci_cloudapplicationcontainer.foocloudapplicationcontainer", &cloudapplicationcontainer),
					testAccCheckAciCloudapplicationcontainerAttributes(fv_tenant_name, cloud_app_name, description, &cloudapplicationcontainer),
				),
			},
		},
	})
}

func testAccCheckAciCloudapplicationcontainerConfig_basic(fv_tenant_name, cloud_app_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_cloudapplicationcontainer" "foocloudapplicationcontainer" {
		name 		= "%s"
		description = "cloudapplicationcontainer created while acceptance testing"
		tenant_dn = "${aci_tenant.footenant.id}"
	}

	`, fv_tenant_name, cloud_app_name)
}

func testAccCheckAciCloudapplicationcontainerExists(name string, cloudapplicationcontainer *models.Cloudapplicationcontainer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud application container %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud application container dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloudapplicationcontainerFound := models.CloudapplicationcontainerFromContainer(cont)
		if cloudapplicationcontainerFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud application container %s not found", rs.Primary.ID)
		}
		*cloudapplicationcontainer = *cloudapplicationcontainerFound
		return nil
	}
}

func testAccCheckAciCloudapplicationcontainerDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloudapplicationcontainer" {
			cont, err := client.Get(rs.Primary.ID)
			cloudapplicationcontainer := models.CloudapplicationcontainerFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud application container %s Still exists", cloudapplicationcontainer.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudapplicationcontainerAttributes(fv_tenant_name, cloud_app_name, description string, cloudapplicationcontainer *models.Cloudapplicationcontainer) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_app_name != GetMOName(cloudapplicationcontainer.DistinguishedName) {
			return fmt.Errorf("Bad cloud_app %s", GetMOName(cloudapplicationcontainer.DistinguishedName))
		}

		if fv_tenant_name != GetMOName(GetParentDn(cloudapplicationcontainer.DistinguishedName)) {
			return fmt.Errorf(" Bad fv_tenant %s", GetMOName(GetParentDn(cloudapplicationcontainer.DistinguishedName)))
		}
		if description != cloudapplicationcontainer.Description {
			return fmt.Errorf("Bad cloudapplicationcontainer Description %s", cloudapplicationcontainer.Description)
		}

		return nil
	}
}
