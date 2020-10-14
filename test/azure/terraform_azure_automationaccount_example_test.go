// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformAzureAutomationAccountExample(t *testing.T) {
	t.Parallel()

	randomPostfixValue := random.UniqueId()

	// Construct options for TF apply
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../../examples/azure/terraform-azure-automationaccount-example",
		Vars: map[string]interface{}{
			"postfix": randomPostfixValue,
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	automationAccountName := terraform.Output(t, terraformOptions, "automation_account_name")
	runAsAccountName := terraform.Output(t, terraformOptions, "runas_account_name")
	runAsCetificateName := terraform.Output(t, terraformOptions, "runas_certificate_name")
}
