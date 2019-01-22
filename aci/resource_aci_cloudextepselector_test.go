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

func TestAccAciCloudEndpointSelectorforExternalEPgs_Basic(t *testing.T) {
	var cloud_endpoint_selectorfor_external_e_pgs models.CloudEndpointSelectorforExternalEPgs
	fv_tenant_name := acctest.RandString(5)
	cloud_app_name := acctest.RandString(5)
	cloud_ext_e_pg_name := acctest.RandString(5)
	cloud_ext_ep_selector_name := acctest.RandString(5)
	description := "cloud_endpoint_selectorfor_external_e_pgs created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudEndpointSelectorforExternalEPgsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudEndpointSelectorforExternalEPgsConfig_basic(fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name, cloud_ext_ep_selector_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudEndpointSelectorforExternalEPgsExists("aci_cloud_endpoint_selectorfor_external_e_pgs.foocloud_endpoint_selectorfor_external_e_pgs", &cloud_endpoint_selectorfor_external_e_pgs),
					testAccCheckAciCloudEndpointSelectorforExternalEPgsAttributes(fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name, cloud_ext_ep_selector_name, description, &cloud_endpoint_selectorfor_external_e_pgs),
				),
			},
		},
	})
}

func testAccCheckAciCloudEndpointSelectorforExternalEPgsConfig_basic(fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name, cloud_ext_ep_selector_name string) string {
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

	resource "aci_cloud_endpoint_selectorfor_external_e_pgs" "foocloud_endpoint_selectorfor_external_e_pgs" {
		name 		= "%s"
		description = "cloud_endpoint_selectorfor_external_e_pgs created while acceptance testing"
		cloud_external_e_pg_dn = "${aci_cloud_external_e_pg.foocloud_external_e_pg.id}"
	}

	`, fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name, cloud_ext_ep_selector_name)
}

func testAccCheckAciCloudEndpointSelectorforExternalEPgsExists(name string, cloud_endpoint_selectorfor_external_e_pgs *models.CloudEndpointSelectorforExternalEPgs) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud Endpoint Selector for External EPgs %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud Endpoint Selector for External EPgs dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloud_endpoint_selectorfor_external_e_pgsFound := models.CloudEndpointSelectorforExternalEPgsFromContainer(cont)
		if cloud_endpoint_selectorfor_external_e_pgsFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud Endpoint Selector for External EPgs %s not found", rs.Primary.ID)
		}
		*cloud_endpoint_selectorfor_external_e_pgs = *cloud_endpoint_selectorfor_external_e_pgsFound
		return nil
	}
}

func testAccCheckAciCloudEndpointSelectorforExternalEPgsDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloud_endpoint_selectorfor_external_e_pgs" {
			cont, err := client.Get(rs.Primary.ID)
			cloud_endpoint_selectorfor_external_e_pgs := models.CloudEndpointSelectorforExternalEPgsFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud Endpoint Selector for External EPgs %s Still exists", cloud_endpoint_selectorfor_external_e_pgs.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudEndpointSelectorforExternalEPgsAttributes(fv_tenant_name, cloud_app_name, cloud_ext_e_pg_name, cloud_ext_ep_selector_name, description string, cloud_endpoint_selectorfor_external_e_pgs *models.CloudEndpointSelectorforExternalEPgs) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_ext_ep_selector_name != GetMOName(cloud_endpoint_selectorfor_external_e_pgs.DistinguishedName) {
			return fmt.Errorf("Bad cloud_ext_ep_selector %s", GetMOName(cloud_endpoint_selectorfor_external_e_pgs.DistinguishedName))
		}

		if cloud_ext_e_pg_name != GetMOName(GetParentDn(cloud_endpoint_selectorfor_external_e_pgs.DistinguishedName)) {
			return fmt.Errorf(" Bad cloud_ext_e_pg %s", GetMOName(GetParentDn(cloud_endpoint_selectorfor_external_e_pgs.DistinguishedName)))
		}
		if description != cloud_endpoint_selectorfor_external_e_pgs.Description {
			return fmt.Errorf("Bad cloud_endpoint_selectorfor_external_e_pgs Description %s", cloud_endpoint_selectorfor_external_e_pgs.Description)
		}

		return nil
	}
}
