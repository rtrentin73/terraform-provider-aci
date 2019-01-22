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

func TestAccAciCloudApplicationcontainer_Basic(t *testing.T) {
	var cloud_applicationcontainer models.CloudApplicationcontainer
	fv_tenant_name := acctest.RandString(5)
	cloud_app_name := acctest.RandString(5)
	description := "cloud_applicationcontainer created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudApplicationcontainerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudApplicationcontainerConfig_basic(fv_tenant_name, cloud_app_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudApplicationcontainerExists("aci_cloud_applicationcontainer.foocloud_applicationcontainer", &cloud_applicationcontainer),
					testAccCheckAciCloudApplicationcontainerAttributes(fv_tenant_name, cloud_app_name, description, &cloud_applicationcontainer),
				),
			},
		},
	})
}

func testAccCheckAciCloudApplicationcontainerConfig_basic(fv_tenant_name, cloud_app_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_cloud_applicationcontainer" "foocloud_applicationcontainer" {
		name 		= "%s"
		description = "cloud_applicationcontainer created while acceptance testing"
		tenant_dn = "${aci_tenant.footenant.id}"
	}

	`, fv_tenant_name, cloud_app_name)
}

func testAccCheckAciCloudApplicationcontainerExists(name string, cloud_applicationcontainer *models.CloudApplicationcontainer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud Application container %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud Application container dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloud_applicationcontainerFound := models.CloudApplicationcontainerFromContainer(cont)
		if cloud_applicationcontainerFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud Application container %s not found", rs.Primary.ID)
		}
		*cloud_applicationcontainer = *cloud_applicationcontainerFound
		return nil
	}
}

func testAccCheckAciCloudApplicationcontainerDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloud_applicationcontainer" {
			cont, err := client.Get(rs.Primary.ID)
			cloud_applicationcontainer := models.CloudApplicationcontainerFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud Application container %s Still exists", cloud_applicationcontainer.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudApplicationcontainerAttributes(fv_tenant_name, cloud_app_name, description string, cloud_applicationcontainer *models.CloudApplicationcontainer) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_app_name != GetMOName(cloud_applicationcontainer.DistinguishedName) {
			return fmt.Errorf("Bad cloud_app %s", GetMOName(cloud_applicationcontainer.DistinguishedName))
		}

		if fv_tenant_name != GetMOName(GetParentDn(cloud_applicationcontainer.DistinguishedName)) {
			return fmt.Errorf(" Bad fv_tenant %s", GetMOName(GetParentDn(cloud_applicationcontainer.DistinguishedName)))
		}
		if description != cloud_applicationcontainer.Description {
			return fmt.Errorf("Bad cloud_applicationcontainer Description %s", cloud_applicationcontainer.Description)
		}

		return nil
	}
}
