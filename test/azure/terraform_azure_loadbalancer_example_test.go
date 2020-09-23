// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerraformAzureLoadBalancerExample(t *testing.T) {
	t.Parallel()

	// initialize resource names, with random unique suffixes
	resourceGroupName := fmt.Sprintf("terratest-loadbalancer-rg-%s", random.UniqueId())
	loadBalancer01Name := fmt.Sprintf("lb-public-%s", random.UniqueId())
	loadBalancer02Name := fmt.Sprintf("lb-private-%s", random.UniqueId())

	frontendIPConfigForLB01 := fmt.Sprintf("cfg-%s", random.UniqueId())
	publicIPAddressForLB01 := fmt.Sprintf("pip-%s", random.UniqueId())

	vnetForLB02 := fmt.Sprintf("vnet-%s", random.UniqueId())
	frontendSubnetID := fmt.Sprintf("snt-%s", random.UniqueId())

	// loadbalancer::tag::1:: Configure Terraform setting up a path to Terraform code.
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../../examples/azure/terraform-azure-loadbalancer-example",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"resource_group_name": resourceGroupName,
			"loadbalancer01_name": loadBalancer01Name,
			"loadbalancer02_name": loadBalancer02Name,
			"vnet_name":           vnetForLB02,
			"lb01_feconfig":       frontendIPConfigForLB01,
			"pip_forlb01":         publicIPAddressForLB01,
			"feSubnet_forlb02":    frontendSubnetID,
		},
	}

	// config
	FrontendIPAllocationMethod := "Dynamic"

	// loadbalancer::tag::4:: At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// loadbalancer::tag::2:: Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// loadbalancer::tag::3:: Run `terraform output` to get the values of output variables

	frontendIPConfigForLB02 := terraform.Output(t, terraformOptions, "feIPConfig_forlb02")
	frontendIPAllocForLB02 := "Static"

	// loadbalancer::tag::5 Set expected variables for test

	// happy path tests
	t.Run("Load Balancer 01", func(t *testing.T) {
		// load balancer 01 (with Public IP) exists
		lb01Exists, err := azure.LoadBalancerExistsE(loadBalancer01Name, resourceGroupName, "")
		assert.NoError(t, err, "Load Balancer error.")
		assert.True(t, lb01Exists)

	})

	t.Run("Frontend Config for LB01", func(t *testing.T) {
		// Read the LB information
		lb01, err := azure.GetLoadBalancerE(loadBalancer01Name, resourceGroupName, "")
		require.NoError(t, err)
		lb01Props := lb01.LoadBalancerPropertiesFormat
		fe01Config := (*lb01Props.FrontendIPConfigurations)[0]

		// Verify settings
		assert.Equal(t, frontendIPConfigForLB01, *fe01Config.Name, "LB01 Frontend IP config name")
	})

	t.Run("IP Checks for LB01", func(t *testing.T) {
		// Read the LB information
		lb01, err := azure.GetLoadBalancerE(loadBalancer01Name, resourceGroupName, "")
		require.NoError(t, err)
		lb01Props := lb01.LoadBalancerPropertiesFormat
		fe01Config := (*lb01Props.FrontendIPConfigurations)[0]
		fe01Props := *fe01Config.FrontendIPConfigurationPropertiesFormat

		// Ensure PrivateIPAddress is nil for LB01
		assert.Nil(t, fe01Props.PrivateIPAddress, "LB01 shouldn't have PrivateIPAddress")

		// Ensure PublicIPAddress Resource exists, no need to check PublicIPAddress value
		publicIPAddressResource, err := azure.GetPublicIPAddressE(publicIPAddressForLB01, resourceGroupName, "")
		require.NoError(t, err)
		assert.NotNil(t, publicIPAddressResource, fmt.Sprintf("Public IP Resource for LB01 Frontend: %s", publicIPAddressForLB01))

		// Verify that expected PublicIPAddressResource is assigned to Load Balancer
		pipResourceName, err := azure.GetSliceLastValueE(*fe01Props.PublicIPAddress.ID, "/")
		require.NoError(t, err)
		assert.Equal(t, publicIPAddressForLB01, pipResourceName, "LB01 Public IP Address Resource Name")

		assert.Equal(t, FrontendIPAllocationMethod, string(fe01Props.PrivateIPAllocationMethod), "LB01 Frontend IP allocation method")
		assert.Nil(t, fe01Props.Subnet, "LB01 shouldn't have Subnet")
	})

	t.Run("Load Balancer 02", func(t *testing.T) {
		// load balancer 02 (with Private IP on vnet/subnet) exists
		lb02Exists, err := azure.LoadBalancerExistsE(loadBalancer02Name, resourceGroupName, "")
		assert.NoError(t, err, "Load Balancer error.")
		assert.True(t, lb02Exists)
	})

	t.Run("IP Check for Load Balancer 02", func(t *testing.T) {
		// Read LB02 information
		lb02, err := azure.GetLoadBalancerE(loadBalancer02Name, resourceGroupName, "")
		require.NoError(t, err)
		lb02Props := lb02.LoadBalancerPropertiesFormat
		fe02Config := (*lb02Props.FrontendIPConfigurations)[0]
		fe02Props := *fe02Config.FrontendIPConfigurationPropertiesFormat

		assert.Equal(t, frontendIPConfigForLB02, *fe02Props.PrivateIPAddress, "LB02 Frontend IP address")
		assert.Equal(t, frontendIPAllocForLB02, string(fe02Props.PrivateIPAllocationMethod), "LB02 Frontend IP allocation method")
		subnetID, err := azure.GetSliceLastValueE(*fe02Props.Subnet.ID, "/")
		require.NoError(t, err, "LB02 Frontend subnet not found")
		assert.Equal(t, frontendSubnetID, subnetID, "LB02 Frontend subnet ID")
	})
}
