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

func TestAccAciCloudExternalEPg_Basic(t *testing.T) {
	var cloud_external_e_pg models.CloudExternalEPg
	fv_tenant_name := acctest.RandString(5)
	cloud_app_name := acctest.RandString(5)
	cloud_ext_e_pg_name := acctest.RandString(5)
	description := "cloud_external_e_pg created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudExternalEPgDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudExternalEPgConfig_basic(fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudExternalEPgExists("aci_cloud_external_e_pg.foocloud_external_e_pg", &cloud_external_e_pg),
					testAccCheckAciCloudExternalEPgAttributes(fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name, description, &cloud_external_e_pg),
				),
			},
		},
	})
}

func testAccCheckAciCloudExternalEPgConfig_basic(fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name string) string {
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

	resource "aci_cloud_external_e_pg" "foocloud_external_e_pg" {
		name 		= "%s"
		description = "cloud_external_e_pg created while acceptance testing"
		cloud_applicationcontainer_dn = "${aci_cloud_applicationcontainer.foocloud_applicationcontainer.id}"
	}

	`, fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name)
}

func testAccCheckAciCloudExternalEPgExists(name string, cloud_external_e_pg *models.CloudExternalEPg) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud External EPg %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud External EPg dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloud_external_e_pgFound := models.CloudExternalEPgFromContainer(cont)
		if cloud_external_e_pgFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud External EPg %s not found", rs.Primary.ID)
		}
		*cloud_external_e_pg = *cloud_external_e_pgFound
		return nil
	}
}

func testAccCheckAciCloudExternalEPgDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloud_external_e_pg" {
			cont, err := client.Get(rs.Primary.ID)
			cloud_external_e_pg := models.CloudExternalEPgFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud External EPg %s Still exists", cloud_external_e_pg.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudExternalEPgAttributes(fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name, description string, cloud_external_e_pg *models.CloudExternalEPg) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_ext_e_pg_name != GetMOName(cloud_external_e_pg.DistinguishedName) {
			return fmt.Errorf("Bad cloud_ext_e_pg %s", GetMOName(cloud_external_e_pg.DistinguishedName))
		}

		if cloud_app_name != GetMOName(GetParentDn(cloud_external_e_pg.DistinguishedName)) {
			return fmt.Errorf(" Bad cloud_app %s", GetMOName(GetParentDn(cloud_external_e_pg.DistinguishedName)))
		}
		if description != cloud_external_e_pg.Description {
			return fmt.Errorf("Bad cloud_external_e_pg Description %s", cloud_external_e_pg.Description)
		}

		return nil
	}
}
