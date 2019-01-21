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

func TestAccAciApplicationprofile_Basic(t *testing.T) {
	var applicationprofile models.Applicationprofile
	fv_tenant_name := acctest.RandString(5)
	fv_ap_name := acctest.RandString(5)
	description := "applicationprofile created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciApplicationprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciApplicationprofileConfig_basic(fv_tenant_name, fv_ap_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciApplicationprofileExists("aci_applicationprofile.fooapplicationprofile", &applicationprofile),
					testAccCheckAciApplicationprofileAttributes(fv_tenant_name, fv_ap_name, description, &applicationprofile),
				),
			},
		},
	})
}

func testAccCheckAciApplicationprofileConfig_basic(fv_tenant_name, fv_ap_name string) string {
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

	`, fv_tenant_name, fv_ap_name)
}

func testAccCheckAciApplicationprofileExists(name string, applicationprofile *models.Applicationprofile) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Application profile %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Application profile dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		applicationprofileFound := models.ApplicationprofileFromContainer(cont)
		if applicationprofileFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Application profile %s not found", rs.Primary.ID)
		}
		*applicationprofile = *applicationprofileFound
		return nil
	}
}

func testAccCheckAciApplicationprofileDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_applicationprofile" {
			cont, err := client.Get(rs.Primary.ID)
			applicationprofile := models.ApplicationprofileFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Application profile %s Still exists", applicationprofile.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciApplicationprofileAttributes(fv_tenant_name, fv_ap_name, description string, applicationprofile *models.Applicationprofile) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if fv_ap_name != GetMOName(applicationprofile.DistinguishedName) {
			return fmt.Errorf("Bad fv_ap %s", GetMOName(applicationprofile.DistinguishedName))
		}

		if fv_tenant_name != GetMOName(GetParentDn(applicationprofile.DistinguishedName)) {
			return fmt.Errorf(" Bad fv_tenant %s", GetMOName(GetParentDn(applicationprofile.DistinguishedName)))
		}
		if description != applicationprofile.Description {
			return fmt.Errorf("Bad applicationprofile Description %s", applicationprofile.Description)
		}

		return nil
	}
}
