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

func TestAccAciCloudawsprovider_Basic(t *testing.T) {
	var cloudawsprovider models.Cloudawsprovider
	fv_tenant_name := acctest.RandString(5)
	cloud_aws_provider_name := acctest.RandString(5)
	description := "cloudawsprovider created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciCloudawsproviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciCloudawsproviderConfig_basic(fv_tenant_name, cloud_aws_provider_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciCloudawsproviderExists("aci_cloudawsprovider.foocloudawsprovider", &cloudawsprovider),
					testAccCheckAciCloudawsproviderAttributes(fv_tenant_name, cloud_aws_provider_name, description, &cloudawsprovider),
				),
			},
		},
	})
}

func testAccCheckAciCloudawsproviderConfig_basic(fv_tenant_name, cloud_aws_provider_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_cloudawsprovider" "foocloudawsprovider" {
		name 		= "%s"
		description = "cloudawsprovider created while acceptance testing"
		tenant_dn = "${aci_tenant.footenant.id}"
	}

	`, fv_tenant_name, cloud_aws_provider_name)
}

func testAccCheckAciCloudawsproviderExists(name string, cloudawsprovider *models.Cloudawsprovider) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Cloud aws provider %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Cloud aws provider dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloudawsproviderFound := models.CloudawsproviderFromContainer(cont)
		if cloudawsproviderFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Cloud aws provider %s not found", rs.Primary.ID)
		}
		*cloudawsprovider = *cloudawsproviderFound
		return nil
	}
}

func testAccCheckAciCloudawsproviderDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_cloudawsprovider" {
			cont, err := client.Get(rs.Primary.ID)
			cloudawsprovider := models.CloudawsproviderFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Cloud aws provider %s Still exists", cloudawsprovider.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciCloudawsproviderAttributes(fv_tenant_name, cloud_aws_provider_name, description string, cloudawsprovider *models.Cloudawsprovider) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if cloud_aws_provider_name != GetMOName(cloudawsprovider.DistinguishedName) {
			return fmt.Errorf("Bad cloud_aws_provider %s", GetMOName(cloudawsprovider.DistinguishedName))
		}

		if fv_tenant_name != GetMOName(GetParentDn(cloudawsprovider.DistinguishedName)) {
			return fmt.Errorf(" Bad fv_tenant %s", GetMOName(GetParentDn(cloudawsprovider.DistinguishedName)))
		}
		if description != cloudawsprovider.Description {
			return fmt.Errorf("Bad cloudawsprovider Description %s", cloudawsprovider.Description)
		}

		return nil
	}
}
