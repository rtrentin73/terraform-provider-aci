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

func TestAccAciApplicationepg_Basic(t *testing.T) {
	var applicationepg models.Applicationepg
	fv_tenant_name := acctest.RandString(5)
	fv_ap_name := acctest.RandString(5)
	fv_ae_pg_name := acctest.RandString(5)
	description := "applicationepg created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciApplicationepgDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciApplicationepgConfig_basic(fv_tenant_name, fv_ap_name, fv_ae_pg_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciApplicationepgExists("aci_applicationepg.fooapplicationepg", &applicationepg),
					testAccCheckAciApplicationepgAttributes(fv_tenant_name, fv_ap_name, fv_ae_pg_name, description, &applicationepg),
				),
			},
		},
	})
}

func testAccCheckAciApplicationepgConfig_basic(fv_tenant_name, fv_ap_name, fv_ae_pg_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_applicationprofile" "fooapplicationprofile" {
		name 		= "%s"
		description = "applicationprofile created while acceptance testing"
		tenant_dn = "${aci_tenant.footenant.id}"
	}

	resource "aci_applicationepg" "fooapplicationepg" {
		name 		= "%s"
		description = "applicationepg created while acceptance testing"
		applicationprofile_dn = "${aci_applicationprofile.fooapplicationprofile.id}"
	}

	`, fv_tenant_name, fv_ap_name, fv_ae_pg_name)
}

func testAccCheckAciApplicationepgExists(name string, applicationepg *models.Applicationepg) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Application epg %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Application epg dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		applicationepgFound := models.ApplicationepgFromContainer(cont)
		if applicationepgFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Application epg %s not found", rs.Primary.ID)
		}
		*applicationepg = *applicationepgFound
		return nil
	}
}

func testAccCheckAciApplicationepgDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_applicationepg" {
			cont, err := client.Get(rs.Primary.ID)
			applicationepg := models.ApplicationepgFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Application epg %s Still exists", applicationepg.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciApplicationepgAttributes(fv_tenant_name, fv_ap_name, fv_ae_pg_name, description string, applicationepg *models.Applicationepg) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if fv_ae_pg_name != GetMOName(applicationepg.DistinguishedName) {
			return fmt.Errorf("Bad fv_ae_pg %s", GetMOName(applicationepg.DistinguishedName))
		}

		if fv_ap_name != GetMOName(GetParentDn(applicationepg.DistinguishedName)) {
			return fmt.Errorf(" Bad fv_ap %s", GetMOName(GetParentDn(applicationepg.DistinguishedName)))
		}
		if description != applicationepg.Description {
			return fmt.Errorf("Bad applicationepg Description %s", applicationepg.Description)
		}

		return nil
	}
}
