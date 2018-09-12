package aci

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/ciscoecosystem/aci-go-client/client"
// 	"github.com/ciscoecosystem/aci-go-client/models"
// 	"github.com/hashicorp/terraform/helper/acctest"
// 	"github.com/hashicorp/terraform/helper/resource"
// 	"github.com/hashicorp/terraform/terraform"
// )

// func TestAccAciTenant_Basic(t *testing.T) {
// 	var tenant models.Tenant
// 	tenant_name := acctest.RandString(5)
// 	description := "Tenant created while acceptance testing"

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckAciTenantDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccCheckAciConfig_basic(tenant_name, description),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckAciTenantExists("aci_tenant.footenant", &tenant),
// 					testAccCheckAciTenantAttributes(tenant_name, description, &tenant),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccAciTenant_Update(t *testing.T) {
// 	var tenant models.Tenant
// 	tenant_name := acctest.RandString(5)
// 	description := "Tenant created while acceptance testing"

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckAciTenantDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccCheckAciConfig_basic(tenant_name, description),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckAciTenantExists("aci_tenant.footenant", &tenant),
// 					testAccCheckAciTenantAttributes(tenant_name, description, &tenant),
// 				),
// 			},
// 			{
// 				Config: testAccCheckAciConfig_basic(tenant_name, "description updated"),
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckAciTenantExists("aci_tenant.footenant", &tenant),
// 					testAccCheckAciTenantAttributes(tenant_name, "description updated", &tenant),
// 				),
// 			},
// 		},
// 	})
// }

// func testAccCheckAciTenantAttributes(tenant_name, description string, tenant *models.Tenant) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {

// 		if tenant_name != GetMOName(tenant.DistinguishedName) {
// 			return fmt.Errorf("Bad Tenant Name %s", GetMOName(tenant.DistinguishedName))
// 		}

// 		if description != tenant.Description {
// 			return fmt.Errorf("Bad Tenant Description %s", tenant.Description)
// 		}

// 		return nil
// 	}
// }
// func testAccCheckAciTenantExists(name string, tenant *models.Tenant) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		rs, ok := s.RootModule().Resources[name]

// 		if !ok {
// 			return fmt.Errorf("Tenant %s not found", name)
// 		}

// 		if rs.Primary.ID == "" {
// 			return fmt.Errorf("No tenant dn was set")
// 		}

// 		client := testAccProvider.Meta().(*client.Client)

// 		cont, err := client.SM.Get(rs.Primary.ID)
// 		if err != nil {
// 			return err
// 		}

// 		tenantFound := models.TenantFromContainer(cont)
// 		if tenantFound.DistinguishedName != rs.Primary.ID {
// 			return fmt.Errorf("Tenant %s not found", rs.Primary.ID)
// 		}
// 		*tenant = *tenantFound
// 		return nil
// 	}
// }

// func testAccCheckAciConfig_basic(tenant_name, description string) string {
// 	return fmt.Sprintf(`
// 	resource "aci_tenant" "footenant" {
// 		name 		= "%s"
// 		description = "%s"
// 	}
// 	`, tenant_name, description)
// }

// func testAccCheckAciTenantDestroy(s *terraform.State) error {
// 	client := testAccProvider.Meta().(*client.Client)

// 	for _, rs := range s.RootModule().Resources {
// 		if rs.Type != "aci_tenant" {
// 			continue
// 		}
// 		cont, err := client.SM.Get(rs.Primary.ID)
// 		fmt.Printf("\n\n Container %+v Error %+v", cont, err)

// 		te := models.TenantFromContainer(cont)
// 		fmt.Printf("\n\nTenant %+v", te)
// 		if err == nil {
// 			fmt.Println(err)
// 			return fmt.Errorf("Tenant %s Still exists", rs.Primary.ID)
// 		}
// 	}

// 	return nil
// }
