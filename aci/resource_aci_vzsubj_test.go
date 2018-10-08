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

func TestAccAciContractSubject_Basic(t *testing.T) {
	var contract_subject models.ContractSubject
	fv_tenant_name := acctest.RandString(5)
	vz_br_cp_name := acctest.RandString(5)
	vz_subj_name := acctest.RandString(5)
	description := "vz_subj created while acceptance testing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciContractSubjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciContractSubjectConfig_basic(fv_tenant_name, vz_br_cp_name, vz_subj_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciContractSubjectExists("aci_vz_subj.foovz_subj", &contract_subject),
					testAccCheckAciContractSubjectAttributes(fv_tenant_name, vz_br_cp_name, vz_subj_name, description, &contract_subject),
				),
			},
		},
	})
}

func testAccCheckAciContractSubjectConfig_basic(fv_tenant_name, vz_br_cp_name, vz_subj_name string) string {
	return fmt.Sprintf(`

	resource "aci_fv_tenant" "foofv_tenant" {
		name 		= "%s"
		description = "fv_tenant created while acceptance testing"

	}

	resource "aci_vz_br_cp" "foovz_br_cp" {
		name 		= "%s"
		description = "vz_br_cp created while acceptance testing"
		fv_tenant_dn = "${aci_fv_tenant.foofv_tenant.id}"
	}

	resource "aci_vz_subj" "foovz_subj" {
		name 		= "%s"
		description = "vz_subj created while acceptance testing"
		vz_br_cp_dn = "${aci_vz_br_cp.foovz_br_cp.id}"
	}

	`, fv_tenant_name, vz_br_cp_name, vz_subj_name)
}

func testAccCheckAciContractSubjectExists(name string, contract_subject *models.ContractSubject) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Contract Subject %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Contract Subject dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		contract_subjectFound := models.ContractSubjectFromContainer(cont)
		if contract_subjectFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Contract Subject %s not found", rs.Primary.ID)
		}
		*contract_subject = *contract_subjectFound
		return nil
	}
}

func testAccCheckAciContractSubjectDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_vz_subj" {
			cont, err := client.Get(rs.Primary.ID)
			contract_subject := models.ContractSubjectFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Contract Subject %s Still exists", contract_subject.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciContractSubjectAttributes(fv_tenant_name, vz_br_cp_name, vz_subj_name, description string, contract_subject *models.ContractSubject) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if fv_tenant_name != GetMOName(contract_subject.DistinguishedName) {
			return fmt.Errorf("Bad fv_tenant %s", GetMOName(contract_subject.DistinguishedName))
		}

		if vz_br_cp_name != GetMOName(contract_subject.DistinguishedName) {
			return fmt.Errorf("Bad vz_br_cp %s", GetMOName(contract_subject.DistinguishedName))
		}

		if vz_subj_name != GetMOName(contract_subject.DistinguishedName) {
			return fmt.Errorf("Bad vz_subj %s", GetMOName(contract_subject.DistinguishedName))
		}

		if description != contract_subject.Description {
			return fmt.Errorf("Bad contract_subject Description %s", contract_subject.Description)
		}

		return nil
	}
}
