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

func TestAccAciBridgedomain_Basic(t *testing.T) {
	var bridgedomain models.Bridgedomain
	fv_tenant_name := acctest.RandString(5)
	fv_bd_name := acctest.RandString(5)
	description := "bridgedomain created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciBridgedomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciBridgedomainConfig_basic(fv_tenant_name, fv_bd_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciBridgedomainExists("aci_bridgedomain.foobridgedomain", &bridgedomain),
					testAccCheckAciBridgedomainAttributes(fv_tenant_name, fv_bd_name, description, &bridgedomain),
				),
			},
		},
	})
}

func testAccCheckAciBridgedomainConfig_basic(fv_tenant_name, fv_bd_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_bridgedomain" "foobridgedomain" {
		name 		= "%s"
		description = "bridgedomain created while acceptance testing"
		tenant_dn = "${aci_tenant.footenant.id}"
	}

	`, fv_tenant_name, fv_bd_name)
}

func testAccCheckAciBridgedomainExists(name string, bridgedomain *models.Bridgedomain) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Bridge domain %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Bridge domain dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		bridgedomainFound := models.BridgedomainFromContainer(cont)
		if bridgedomainFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Bridge domain %s not found", rs.Primary.ID)
		}
		*bridgedomain = *bridgedomainFound
		return nil
	}
}

func testAccCheckAciBridgedomainDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_bridgedomain" {
			cont, err := client.Get(rs.Primary.ID)
			bridgedomain := models.BridgedomainFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Bridge domain %s Still exists", bridgedomain.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciBridgedomainAttributes(fv_tenant_name, fv_bd_name, description string, bridgedomain *models.Bridgedomain) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if fv_bd_name != GetMOName(bridgedomain.DistinguishedName) {
			return fmt.Errorf("Bad fv_bd %s", GetMOName(bridgedomain.DistinguishedName))
		}

		if fv_tenant_name != GetMOName(GetParentDn(bridgedomain.DistinguishedName)) {
			return fmt.Errorf(" Bad fv_tenant %s", GetMOName(GetParentDn(bridgedomain.DistinguishedName)))
		}
		if description != bridgedomain.Description {
			return fmt.Errorf("Bad bridgedomain Description %s", bridgedomain.Description)
		}

		return nil
	}
}
