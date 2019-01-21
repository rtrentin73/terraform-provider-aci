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

func TestAccAciCloudexternalepg_Basic(t *testing.T) {
	var cloudexternalepg models.Cloudexternalepg
	fv_tenant_name := acctest.RandString(5)
	cloud_app_name := acctest.RandString(5)
	cloud_ext_e_pg_name := acctest.RandString(5)
	description := "cloudexternalepg created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudexternalepgDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudexternalepgConfig_basic(fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudexternalepgExists("aci_cloudexternalepg.foocloudexternalepg", &cloudexternalepg),
					testAccCheckAciCloudexternalepgAttributes(fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name, description, &cloudexternalepg),
				),
			},
		},
	})
}

func testAccCheckAciCloudexternalepgConfig_basic(fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name string) string {
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

	`, fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name)
}

func testAccCheckAciCloudexternalepgExists(name string, cloudexternalepg *models.Cloudexternalepg) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud external epg %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud external epg dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloudexternalepgFound := models.CloudexternalepgFromContainer(cont)
		if cloudexternalepgFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud external epg %s not found", rs.Primary.ID)
		}
		*cloudexternalepg = *cloudexternalepgFound
		return nil
	}
}

func testAccCheckAciCloudexternalepgDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloudexternalepg" {
			cont, err := client.Get(rs.Primary.ID)
			cloudexternalepg := models.CloudexternalepgFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud external epg %s Still exists", cloudexternalepg.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudexternalepgAttributes(fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name, description string, cloudexternalepg *models.Cloudexternalepg) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_ext_e_pg_name != GetMOName(cloudexternalepg.DistinguishedName) {
			return fmt.Errorf("Bad cloud_ext_e_pg %s", GetMOName(cloudexternalepg.DistinguishedName))
		}

		if cloud_app_name != GetMOName(GetParentDn(cloudexternalepg.DistinguishedName)) {
			return fmt.Errorf(" Bad cloud_app %s", GetMOName(GetParentDn(cloudexternalepg.DistinguishedName)))
		}
		if description != cloudexternalepg.Description {
			return fmt.Errorf("Bad cloudexternalepg Description %s", cloudexternalepg.Description)
		}

		return nil
	}
}
