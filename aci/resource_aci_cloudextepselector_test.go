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

func TestAccAciCloudendpointselectorforexternalepgs_Basic(t *testing.T) {
	var cloudendpointselectorforexternalepgs models.Cloudendpointselectorforexternalepgs
	fv_tenant_name := acctest.RandString(5)
	cloud_app_name := acctest.RandString(5)
	cloud_ext_e_pg_name := acctest.RandString(5)
	cloud_ext_ep_selector_name := acctest.RandString(5)
	description := "cloudendpointselectorforexternalepgs created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudendpointselectorforexternalepgsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudendpointselectorforexternalepgsConfig_basic(fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name, cloud_ext_ep_selector_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudendpointselectorforexternalepgsExists("aci_cloudendpointselectorforexternalepgs.foocloudendpointselectorforexternalepgs", &cloudendpointselectorforexternalepgs),
					testAccCheckAciCloudendpointselectorforexternalepgsAttributes(fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name, cloud_ext_ep_selector_name, description, &cloudendpointselectorforexternalepgs),
				),
			},
		},
	})
}

func testAccCheckAciCloudendpointselectorforexternalepgsConfig_basic(fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name, cloud_ext_ep_selector_name string) string {
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

	resource "aci_cloudexternalepg" "foocloudexternalepg" {
		name 		= "%s"
		description = "cloudexternalepg created while acceptance testing"
		cloudapplicationcontainer_dn = "${aci_cloudapplicationcontainer.foocloudapplicationcontainer.id}"
	}

	resource "aci_cloudendpointselectorforexternalepgs" "foocloudendpointselectorforexternalepgs" {
		name 		= "%s"
		description = "cloudendpointselectorforexternalepgs created while acceptance testing"
		cloudexternalepg_dn = "${aci_cloudexternalepg.foocloudexternalepg.id}"
	}

	`, fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name, cloud_ext_ep_selector_name)
}

func testAccCheckAciCloudendpointselectorforexternalepgsExists(name string, cloudendpointselectorforexternalepgs *models.Cloudendpointselectorforexternalepgs) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud endpoint selector for external epgs %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud endpoint selector for external epgs dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloudendpointselectorforexternalepgsFound := models.CloudendpointselectorforexternalepgsFromContainer(cont)
		if cloudendpointselectorforexternalepgsFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud endpoint selector for external epgs %s not found", rs.Primary.ID)
		}
		*cloudendpointselectorforexternalepgs = *cloudendpointselectorforexternalepgsFound
		return nil
	}
}

func testAccCheckAciCloudendpointselectorforexternalepgsDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloudendpointselectorforexternalepgs" {
			cont, err := client.Get(rs.Primary.ID)
			cloudendpointselectorforexternalepgs := models.CloudendpointselectorforexternalepgsFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud endpoint selector for external epgs %s Still exists", cloudendpointselectorforexternalepgs.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudendpointselectorforexternalepgsAttributes(fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name, cloud_ext_ep_selector_name, description string, cloudendpointselectorforexternalepgs *models.Cloudendpointselectorforexternalepgs) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_ext_ep_selector_name != GetMOName(cloudendpointselectorforexternalepgs.DistinguishedName) {
			return fmt.Errorf("Bad cloud_ext_ep_selector %s", GetMOName(cloudendpointselectorforexternalepgs.DistinguishedName))
		}

		if cloud_ext_e_pg_name != GetMOName(GetParentDn(cloudendpointselectorforexternalepgs.DistinguishedName)) {
			return fmt.Errorf(" Bad cloud_ext_e_pg %s", GetMOName(GetParentDn(cloudendpointselectorforexternalepgs.DistinguishedName)))
		}
		if description != cloudendpointselectorforexternalepgs.Description {
			return fmt.Errorf("Bad cloudendpointselectorforexternalepgs Description %s", cloudendpointselectorforexternalepgs.Description)
		}

		return nil
	}
}
