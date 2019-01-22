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

func TestAccAciCloudEPg_Basic(t *testing.T) {
	var cloud_e_pg models.CloudEPg
	fv_tenant_name := acctest.RandString(5)
	cloud_app_name := acctest.RandString(5)
	cloud_e_pg_name := acctest.RandString(5)
	description := "cloud_e_pg created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudEPgDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudEPgConfig_basic(fv_tenant_name, cloud_app_name, cloud_e_pg_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudEPgExists("aci_cloud_e_pg.foocloud_e_pg", &cloud_e_pg),
					testAccCheckAciCloudEPgAttributes(fv_tenant_name, cloud_app_name, cloud_e_pg_name, description, &cloud_e_pg),
				),
			},
		},
	})
}

func testAccCheckAciCloudEPgConfig_basic(fv_tenant_name, cloud_app_name, cloud_e_pg_name string) string {
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

	resource "aci_cloud_e_pg" "foocloud_e_pg" {
		name 		= "%s"
		description = "cloud_e_pg created while acceptance testing"
		cloud_applicationcontainer_dn = "${aci_cloud_applicationcontainer.foocloud_applicationcontainer.id}"
	}

	`, fv_tenant_name, cloud_app_name, cloud_e_pg_name)
}

func testAccCheckAciCloudEPgExists(name string, cloud_e_pg *models.CloudEPg) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud EPg %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud EPg dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloud_e_pgFound := models.CloudEPgFromContainer(cont)
		if cloud_e_pgFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud EPg %s not found", rs.Primary.ID)
		}
		*cloud_e_pg = *cloud_e_pgFound
		return nil
	}
}

func testAccCheckAciCloudEPgDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloud_e_pg" {
			cont, err := client.Get(rs.Primary.ID)
			cloud_e_pg := models.CloudEPgFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud EPg %s Still exists", cloud_e_pg.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudEPgAttributes(fv_tenant_name, cloud_app_name, cloud_e_pg_name, description string, cloud_e_pg *models.CloudEPg) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_e_pg_name != GetMOName(cloud_e_pg.DistinguishedName) {
			return fmt.Errorf("Bad cloud_e_pg %s", GetMOName(cloud_e_pg.DistinguishedName))
		}

		if cloud_app_name != GetMOName(GetParentDn(cloud_e_pg.DistinguishedName)) {
			return fmt.Errorf(" Bad cloud_app %s", GetMOName(GetParentDn(cloud_e_pg.DistinguishedName)))
		}
		if description != cloud_e_pg.Description {
			return fmt.Errorf("Bad cloud_e_pg Description %s", cloud_e_pg.Description)
		}

		return nil
	}
}
