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

func TestAccAciCloudEndpointSelector_Basic(t *testing.T) {
	var cloud_endpoint_selector models.CloudEndpointSelector
	fv_tenant_name := acctest.RandString(5)
	cloud_app_name := acctest.RandString(5)
	cloud_e_pg_name := acctest.RandString(5)
	cloud_ep_selector_name := acctest.RandString(5)
	description := "cloud_endpoint_selector created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudEndpointSelectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudEndpointSelectorConfig_basic(fv_tenant_name, cloud_app_name, cloud_e_pg_name, cloud_ep_selector_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudEndpointSelectorExists("aci_cloud_endpoint_selector.foocloud_endpoint_selector", &cloud_endpoint_selector),
					testAccCheckAciCloudEndpointSelectorAttributes(fv_tenant_name, cloud_app_name, cloud_e_pg_name, cloud_ep_selector_name, description, &cloud_endpoint_selector),
				),
			},
		},
	})
}

func testAccCheckAciCloudEndpointSelectorConfig_basic(fv_tenant_name, cloud_app_name, cloud_e_pg_name, cloud_ep_selector_name string) string {
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

	resource "aci_cloud_endpoint_selector" "foocloud_endpoint_selector" {
		name 		= "%s"
		description = "cloud_endpoint_selector created while acceptance testing"
		cloud_e_pg_dn = "${aci_cloud_e_pg.foocloud_e_pg.id}"
	}

	`, fv_tenant_name, cloud_app_name, cloud_e_pg_name, cloud_ep_selector_name)
}

func testAccCheckAciCloudEndpointSelectorExists(name string, cloud_endpoint_selector *models.CloudEndpointSelector) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud Endpoint Selector %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud Endpoint Selector dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloud_endpoint_selectorFound := models.CloudEndpointSelectorFromContainer(cont)
		if cloud_endpoint_selectorFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud Endpoint Selector %s not found", rs.Primary.ID)
		}
		*cloud_endpoint_selector = *cloud_endpoint_selectorFound
		return nil
	}
}

func testAccCheckAciCloudEndpointSelectorDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloud_endpoint_selector" {
			cont, err := client.Get(rs.Primary.ID)
			cloud_endpoint_selector := models.CloudEndpointSelectorFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud Endpoint Selector %s Still exists", cloud_endpoint_selector.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudEndpointSelectorAttributes(fv_tenant_name, cloud_app_name, cloud_e_pg_name, cloud_ep_selector_name, description string, cloud_endpoint_selector *models.CloudEndpointSelector) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_ep_selector_name != GetMOName(cloud_endpoint_selector.DistinguishedName) {
			return fmt.Errorf("Bad cloud_ep_selector %s", GetMOName(cloud_endpoint_selector.DistinguishedName))
		}

		if cloud_e_pg_name != GetMOName(GetParentDn(cloud_endpoint_selector.DistinguishedName)) {
			return fmt.Errorf(" Bad cloud_e_pg %s", GetMOName(GetParentDn(cloud_endpoint_selector.DistinguishedName)))
		}
		if description != cloud_endpoint_selector.Description {
			return fmt.Errorf("Bad cloud_endpoint_selector Description %s", cloud_endpoint_selector.Description)
		}

		return nil
	}
}
