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

func TestAccAciCloudepg_Basic(t *testing.T) {
	var cloudepg models.Cloudepg
	fv_tenant_name := acctest.RandString(5)
	cloud_app_name := acctest.RandString(5)
	cloud_e_pg_name := acctest.RandString(5)
	description := "cloudepg created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudepgDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudepgConfig_basic(fv_tenant_name, cloud_app_name, cloud_e_pg_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudepgExists("aci_cloudepg.foocloudepg", &cloudepg),
					testAccCheckAciCloudepgAttributes(fv_tenant_name, cloud_app_name, cloud_e_pg_name, description, &cloudepg),
				),
			},
		},
	})
}

func testAccCheckAciCloudepgConfig_basic(fv_tenant_name, cloud_app_name, cloud_e_pg_name string) string {
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

	`, fv_tenant_name, cloud_app_name, cloud_e_pg_name)
}

func testAccCheckAciCloudepgExists(name string, cloudepg *models.Cloudepg) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud epg %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud epg dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloudepgFound := models.CloudepgFromContainer(cont)
		if cloudepgFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud epg %s not found", rs.Primary.ID)
		}
		*cloudepg = *cloudepgFound
		return nil
	}
}

func testAccCheckAciCloudepgDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloudepg" {
			cont, err := client.Get(rs.Primary.ID)
			cloudepg := models.CloudepgFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud epg %s Still exists", cloudepg.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudepgAttributes(fv_tenant_name, cloud_app_name, cloud_e_pg_name, description string, cloudepg *models.Cloudepg) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_e_pg_name != GetMOName(cloudepg.DistinguishedName) {
			return fmt.Errorf("Bad cloud_e_pg %s", GetMOName(cloudepg.DistinguishedName))
		}

		if cloud_app_name != GetMOName(GetParentDn(cloudepg.DistinguishedName)) {
			return fmt.Errorf(" Bad cloud_app %s", GetMOName(GetParentDn(cloudepg.DistinguishedName)))
		}
		if description != cloudepg.Description {
			return fmt.Errorf("Bad cloudepg Description %s", cloudepg.Description)
		}

		return nil
	}
}
