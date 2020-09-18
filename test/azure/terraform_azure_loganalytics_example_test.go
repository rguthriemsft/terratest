// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAzureLogAnalyticsExample(t *testing.T) {
	t.Parallel()

	expectedResourceGroupName := fmt.Sprintf("terratest-rg-%s", random.UniqueId())
	expectedWorkspaceName := fmt.Sprintf("loganalyticsws-%s", random.UniqueId())

	// website::tag::1:: Configure Terraform setting up a path to Terraform code.
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../../examples/azure/terraform-azure-loganalytics-example",
		Vars: map[string]interface{}{
			"resource_group_name":         expectedResourceGroupName,
			"loganalytics_workspace_name": expectedWorkspaceName,
		},
	}

	// website::tag::4:: At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// website::tag::2:: Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// website::tag::3:: Run `terraform output` to get the values of output variables
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	workspaceName := terraform.Output(t, terraformOptions, "loganalytics_workspace_name")
	sku := terraform.Output(t, terraformOptions, "loganalytics_workspace_sku")
	retentionPeriodString := terraform.Output(t, terraformOptions, "loganalytics_workspace_retention")

	// website::tag::4:: Verify the Log Analytics properties and ensure it matches the output.
	workspaceExists := azure.LogAnalyticsWorkspaceExists(workspaceName, resourceGroupName, "")
	assert.True(t, workspaceExists, "log analytics workspace not found.")

	actualSku := azure.GetLogAnalyticsWorkspaceSku(workspaceName, resourceGroupName, "")
	assert.Equal(t, sku, strings.ToLower(actualSku), "log analytics sku mismatch")

	var actualRetentionPeriod = azure.GetLogAnalyticsWorkspaceRetentionPeriodDays(workspaceName, resourceGroupName, "")
	expectedPeriod, _ := strconv.ParseInt(retentionPeriodString, 10, 32)
	assert.Equal(t, int32(expectedPeriod), actualRetentionPeriod, "log analytics retention period mismatch")
}
