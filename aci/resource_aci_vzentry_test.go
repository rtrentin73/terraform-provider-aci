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

func TestAccAciFilterentry_Basic(t *testing.T) {
	var filterentry models.Filterentry
	fv_tenant_name := acctest.RandString(5)
	vz_filter_name := acctest.RandString(5)
	vz_entry_name := acctest.RandString(5)
	description := "filterentry created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciFilterentryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciFilterentryConfig_basic(fv_tenant_name, vz_filter_name, vz_entry_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciFilterentryExists("aci_filterentry.foofilterentry", &filterentry),
					testAccCheckAciFilterentryAttributes(fv_tenant_name, vz_filter_name, vz_entry_name, description, &filterentry),
				),
			},
		},
	})
}

func testAccCheckAciFilterentryConfig_basic(fv_tenant_name, vz_filter_name, vz_entry_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_filter" "foofilter" {
		name 		= "%s"
		description = "filter created while acceptance testing"
		tenant_dn = "${aci_tenant.footenant.id}"
	}

	resource "aci_filterentry" "foofilterentry" {
		name 		= "%s"
		description = "filterentry created while acceptance testing"
		filter_dn = "${aci_filter.foofilter.id}"
	}

	`, fv_tenant_name, vz_filter_name, vz_entry_name)
}

func testAccCheckAciFilterentryExists(name string, filterentry *models.Filterentry) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Filter entry %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Filter entry dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		filterentryFound := models.FilterentryFromContainer(cont)
		if filterentryFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Filter entry %s not found", rs.Primary.ID)
		}
		*filterentry = *filterentryFound
		return nil
	}
}

func testAccCheckAciFilterentryDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_filterentry" {
			cont, err := client.Get(rs.Primary.ID)
			filterentry := models.FilterentryFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Filter entry %s Still exists", filterentry.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciFilterentryAttributes(fv_tenant_name, vz_filter_name, vz_entry_name, description string, filterentry *models.Filterentry) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if vz_entry_name != GetMOName(filterentry.DistinguishedName) {
			return fmt.Errorf("Bad vz_entry %s", GetMOName(filterentry.DistinguishedName))
		}

		if vz_filter_name != GetMOName(GetParentDn(filterentry.DistinguishedName)) {
			return fmt.Errorf(" Bad vz_filter %s", GetMOName(GetParentDn(filterentry.DistinguishedName)))
		}
		if description != filterentry.Description {
			return fmt.Errorf("Bad filterentry Description %s", filterentry.Description)
		}

		return nil
	}
}
