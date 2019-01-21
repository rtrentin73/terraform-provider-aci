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

func TestAccAciContractsubject_Basic(t *testing.T) {
	var contractsubject models.Contractsubject
	fv_tenant_name := acctest.RandString(5)
	vz_br_cp_name := acctest.RandString(5)
	vz_subj_name := acctest.RandString(5)
	description := "contractsubject created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciContractsubjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciContractsubjectConfig_basic(fv_tenant_name, vz_br_cp_name, vz_subj_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciContractsubjectExists("aci_contractsubject.foocontractsubject", &contractsubject),
					testAccCheckAciContractsubjectAttributes(fv_tenant_name, vz_br_cp_name, vz_subj_name, description, &contractsubject),
				),
			},
		},
	})
}

func testAccCheckAciContractsubjectConfig_basic(fv_tenant_name, vz_br_cp_name, vz_subj_name string) string {
	return fmt.Sprintf(`

	resource "aci_tenant" "footenant" {
		name 		= "%s"
		description = "tenant created while acceptance testing"

	}

	resource "aci_contract" "foocontract" {
		name 		= "%s"
		description = "contract created while acceptance testing"
		tenant_dn = "${aci_tenant.footenant.id}"
	}

	resource "aci_contractsubject" "foocontractsubject" {
		name 		= "%s"
		description = "contractsubject created while acceptance testing"
		contract_dn = "${aci_contract.foocontract.id}"
	}

	`, fv_tenant_name, vz_br_cp_name, vz_subj_name)
}

func testAccCheckAciContractsubjectExists(name string, contractsubject *models.Contractsubject) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Contract subject %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Contract subject dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		contractsubjectFound := models.ContractsubjectFromContainer(cont)
		if contractsubjectFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Contract subject %s not found", rs.Primary.ID)
		}
		*contractsubject = *contractsubjectFound
		return nil
	}
}

func testAccCheckAciContractsubjectDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_contractsubject" {
			cont, err := client.Get(rs.Primary.ID)
			contractsubject := models.ContractsubjectFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Contract subject %s Still exists", contractsubject.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciContractsubjectAttributes(fv_tenant_name, vz_br_cp_name, vz_subj_name, description string, contractsubject *models.Contractsubject) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if vz_subj_name != GetMOName(contractsubject.DistinguishedName) {
			return fmt.Errorf("Bad vz_subj %s", GetMOName(contractsubject.DistinguishedName))
		}

		if vz_br_cp_name != GetMOName(GetParentDn(contractsubject.DistinguishedName)) {
			return fmt.Errorf(" Bad vz_br_cp %s", GetMOName(GetParentDn(contractsubject.DistinguishedName)))
		}
		if description != contractsubject.Description {
			return fmt.Errorf("Bad contractsubject Description %s", contractsubject.Description)
		}

		return nil
	}
}
