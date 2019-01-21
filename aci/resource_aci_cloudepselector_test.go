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

func TestAccAciCloudendpointselector_Basic(t *testing.T) {
	var cloudendpointselector models.Cloudendpointselector
	fv_tenant_name := acctest.RandString(5)
	cloud_app_name := acctest.RandString(5)
	cloud_e_pg_name := acctest.RandString(5)
	cloud_ep_selector_name := acctest.RandString(5)
	description := "cloudendpointselector created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudendpointselectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudendpointselectorConfig_basic(fv_tenant_name, cloud_app_name, cloud_e_pg_name, cloud_ep_selector_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudendpointselectorExists("aci_cloudendpointselector.foocloudendpointselector", &cloudendpointselector),
					testAccCheckAciCloudendpointselectorAttributes(fv_tenant_name, cloud_app_name, cloud_e_pg_name, cloud_ep_selector_name, description, &cloudendpointselector),
				),
			},
		},
	})
}

func testAccCheckAciCloudendpointselectorConfig_basic(fv_tenant_name, cloud_app_name, cloud_e_pg_name, cloud_ep_selector_name string) string {
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

	resource "aci_cloudepg" "foocloudepg" {
		name 		= "%s"
		description = "cloudepg created while acceptance testing"
		cloudapplicationcontainer_dn = "${aci_cloudapplicationcontainer.foocloudapplicationcontainer.id}"
	}

	resource "aci_cloudendpointselector" "foocloudendpointselector" {
		name 		= "%s"
		description = "cloudendpointselector created while acceptance testing"
		cloudepg_dn = "${aci_cloudepg.foocloudepg.id}"
	}

	`, fv_tenant_name, cloud_app_name, cloud_e_pg_name, cloud_ep_selector_name)
}

func testAccCheckAciCloudendpointselectorExists(name string, cloudendpointselector *models.Cloudendpointselector) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud endpoint selector %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud endpoint selector dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloudendpointselectorFound := models.CloudendpointselectorFromContainer(cont)
		if cloudendpointselectorFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud endpoint selector %s not found", rs.Primary.ID)
		}
		*cloudendpointselector = *cloudendpointselectorFound
		return nil
	}
}

func testAccCheckAciCloudendpointselectorDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloudendpointselector" {
			cont, err := client.Get(rs.Primary.ID)
			cloudendpointselector := models.CloudendpointselectorFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud endpoint selector %s Still exists", cloudendpointselector.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudendpointselectorAttributes(fv_tenant_name, cloud_app_name, cloud_e_pg_name, cloud_ep_selector_name, description string, cloudendpointselector *models.Cloudendpointselector) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_ep_selector_name != GetMOName(cloudendpointselector.DistinguishedName) {
			return fmt.Errorf("Bad cloud_ep_selector %s", GetMOName(cloudendpointselector.DistinguishedName))
		}

		if cloud_e_pg_name != GetMOName(GetParentDn(cloudendpointselector.DistinguishedName)) {
			return fmt.Errorf(" Bad cloud_e_pg %s", GetMOName(GetParentDn(cloudendpointselector.DistinguishedName)))
		}
		if description != cloudendpointselector.Description {
			return fmt.Errorf("Bad cloudendpointselector Description %s", cloudendpointselector.Description)
		}

		return nil
	}
}
