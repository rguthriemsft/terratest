// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAzureAutomationAccountExample(t *testing.T) {
	t.Parallel()

	// subscriptionID is overridden by the environment variable "ARM_SUBSCRIPTION_ID"
	subscriptionID := ""
	uniquePostfix := random.UniqueId()
	expectedAutomationAccountName := "terratest-AutomationAccount"
	expectedSampleDSCName := "SampleDSC"

	// Construct options for TF apply
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../../examples/azure/terraform-azure-automationaccount-example",
		Vars: map[string]interface{}{
			"postfix":                 uniquePostfix,
			"automation_account_name": expectedAutomationAccountName,
			"sample_dsc_name":         expectedSampleDSCName,
		},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	automationAccountName := terraform.Output(t, terraformOptions, "automation_account_name")
	// runAsAccountName := terraform.Output(t, terraformOptions, "runas_account_name")
	// runAsCetificateName := terraform.Output(t, terraformOptions, "runas_certificate_name")

	// Check that the automation account deployed successfully
	actualAutomationAccountExists := azure.AutomationAccountExists(t, automationAccountName, resourceGroupName, subscriptionID)
	assert.True(t, actualAutomationAccountExists)
	//Check that the sample DSC was uploaded successfully into the deployed automation account
	actualDSCExists := azure.AutomationAccountDSCExists(t, automationAccountName, expectedSampleDSCName, resourceGroupName, subscriptionID)
	assert.True(t, actualDSCExists)
	// Check that the DSC in the automation account successfully compiled
	dscCompiled := azure.AutomationAccountDSCCompiled(t, automationAccountName, expectedSampleDSCName, resourceGroupName, subscriptionID)
	assert.True(t, dscCompiled)
}
