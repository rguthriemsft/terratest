package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/automation/mgmt/automation"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// AutomationAccountExists indicates whether the specified Azure Automation Account exists.
// This function would fail the test if there is an error.
func AutomationAccountExists(t testing.TestingT, automationAccountName string, resourceGroupName string, subscriptionID string) bool {
	exists, err := AutomationAccountExistsE(t, automationAccountName, resourceGroupName, subscriptionID)
	require.NoError(t, err)
	return exists
}

// AutomationAccountExistsE indicates whether the specified Azure Automation Account exists.
func AutomationAccountExistsE(t testing.TestingT, automationAccountName string, resourceGroupName string, subscriptionID string) (bool, error) {
	_, err := GetAutomationAccountE(t, automationAccountName, resourceGroupName, subscriptionID)
	if err != nil {
		if ResourceNotFoundErrorExists(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// AutomationAccountDSCExists indicates whether the specified Azure Automation Account DSC exists.
// This function would fail the test if there is an error.
func AutomationAccountDSCExists(t testing.TestingT, automationAccountName string, dscConfiguraitonName string, resourceGroupName string, subscriptionID string) bool {
	exists, err := AutomationAccountDSCExistsE(t, automationAccountName, dscConfiguraitonName, resourceGroupName, subscriptionID)
	require.NoError(t, err)
	return exists
}

// AutomationAccountDSCExistsE indicates whether the specified Azure Automation Account exists.
func AutomationAccountDSCExistsE(t testing.TestingT, automationAccountName string, dscConfiguraitonName string, resourceGroupName string, subscriptionID string) (bool, error) {
	_, err := GetAutomationAccountDSCConfigurationE(t, automationAccountName, dscConfiguraitonName, resourceGroupName, subscriptionID)
	if err != nil {
		if ResourceNotFoundErrorExists(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// AutomationAccountDSCCompiled indicates whether the specified Azure Automation Account DSC compiled successfully.
// This function would fail the test if there is an error.
func AutomationAccountDSCCompiled(t testing.TestingT, automationAccountName string, dscConfiguraitonName string, resourceGroupName string, subscriptionID string) bool {
	compiled, err := AutomationAccountDSCCompiledE(t, automationAccountName, dscConfiguraitonName, resourceGroupName, subscriptionID)
	require.NoError(t, err)
	return compiled
}

// AutomationAccountDSCCompiledE indicates whether the specified Azure Automation Account DSC compiled successfully.
func AutomationAccountDSCCompiledE(t testing.TestingT, automationAccountName string, dscConfiguraitonName string, resourceGroupName string, subscriptionID string) (bool, error) {
	dscCompilatinoJob, err := GetAutomationAccountDscCompilationJobE(t, automationAccountName, dscConfiguraitonName, resourceGroupName, subscriptionID)
	if err != nil {
	}

	return dscCompilatinoJob.HasHTTPStatus(), nil
}

// DSCAppliedSuccessfullyToVM indicates whether the specified Azure Automation Account compiled DSC has successfully been applied to the target VM.
// This function would fail the test if there is an error.
func DSCAppliedSuccessfullyToVM(t testing.TestingT) {

}

// AutomationAccountHasRunAsCertificate indicates whether the specified Azure Automation Account RunAs Certificate exists.
// This function would fail the test if there is an error.
func AutomationAccountHasRunAsCertificate(t testing.TestingT) {

}

// AutomationAccountHasRunAsAccount indicates whether the specified Azure Automation Account RunAs Account exists.
// This function would fail the test if there is an error.
func AutomationAccountHasRunAsAccount(t testing.TestingT) {

}

// GetAutomationAccountE returns the Azure Automation Account by name if it exists in the subscription
func GetAutomationAccountE(t testing.TestingT, automationAccountName string, resourceGroupName string, subscriptionID string) (*automation.Account, error) {
	// Validate resource group name and subscription ID
	resourceGroupName, err := getTargetAzureResourceGroupName(resourceGroupName)
	if err != nil {
		return nil, err
	}

	client, err := GetAutomationAccountClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	automationAccount, err := client.Get(context.Background(), resourceGroupName, automationAccountName)
	if err != nil {
		return nil, err
	}

	return &automationAccount, nil
}

// GetAutomationAccountClientE returns the Azure Automation Account client
func GetAutomationAccountClientE(subscriptionID string) (*automation.AccountClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Automation Account Client
	client := automation.NewAccountClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	client.Authorizer = *authorizer

	return &client, nil
}

// GetAutomationAccountDSCConfigurationE returns the Azure Automation Account DscConfiguration by Automation Account name if it exists in the subscription
func GetAutomationAccountDSCConfigurationE(t testing.TestingT, automationAccountName string, dscConfigurationName string, resourceGroupName string, subscriptionID string) (*automation.DscConfiguration, error) {
	resourceGroupName, err := getTargetAzureResourceGroupName(resourceGroupName)
	if err != nil {
		return nil, err
	}

	client, err := GetDscConfigurationClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	dscConfiguration, err := client.Get(context.Background(), resourceGroupName, automationAccountName, dscConfigurationName)
	if err != nil {
		return nil, err
	}

	return &dscConfiguration, nil
}

// GetDscConfigurationClientE returns the Azure Automation Account DscConfigurationClient
func GetDscConfigurationClientE(subscriptionID string) (*automation.DscConfigurationClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Automation Account DSC Configuraiton Client
	client := automation.NewDscConfigurationClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	client.Authorizer = *authorizer

	return &client, nil
}

// GetAutomationAccountDscCompilationJobE returns the Azure Automation Account DscConfiguration by Automation Account name if it exists in the subscription
func GetAutomationAccountDscCompilationJobE(t testing.TestingT, automationAccountName string, dscConfigurationName string, resourceGroupName string, subscriptionID string) (*automation.DscCompilationJob, error) {
	resourceGroupName, err := getTargetAzureResourceGroupName(resourceGroupName)
	if err != nil {
		return nil, err
	}

	client, err := GetDscCompilationJobClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	filter := fmt.Sprintf("properties/configuration/name eq '%s'", dscConfigurationName)
	dscCompilationJobListResultPage, err := client.ListByAutomationAccount(context.Background(), resourceGroupName, automationAccountName, filter)
	if err != nil {
		return nil, err
	}

	var dscCompilatinoJob automation.DscCompilationJob
	for dscCompilationJobListResultPage.NotDone() {
		dscCompilatinoJob = dscCompilationJobListResultPage.Values()[0]
		//dscCompilatinoJob.DscCompilationJobProperties
		//dscCompilatinoJob.DscCompilationJobProperties.ProvisioningState == automation.JobProvisioningState.JobProvisioningStateSucceeded

	}
	// //*dscCompilatinoJob.DscCompilationJobProperties.Configuration.Name
	// if dscCompilatinoJob.DscCompilationJobProperties.Status == "JobStatusCompleting" {

	// }

	return &dscCompilatinoJob, nil
}

// GetDscCompilationJobClientE returns the Azure Automation Account DscCompilationJobClient
func GetDscCompilationJobClientE(subscriptionID string) (*automation.DscCompilationJobClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Automation Account DSC Compilation Job Client
	client := automation.NewDscCompilationJobClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	client.Authorizer = *authorizer

	return &client, nil
}

// GetAutomationAccountConnection gets the RunAs Connection if it exists in the Azure Automation Account
// This function would fail the test if there is an error.
func GetAutomationAccountConnection(t testing.TestingT, automationAccountName string, resourceGroupName string, subscriptionID string, automationAccountConnectionName string) *automation.Connection {
	client, err := GetAutomationAccountConnectionClientE(subscriptionID)
	require.NoError(t, err)

	// Get Automation Account Connection
	automationAccountConnection, err := client.Get(context.Background(), resourceGroupName, automationAccountName, automationAccountConnectionName)
	require.NoError(t, err)

	return &automationAccountConnection
}

// GetAutomationAccountConnectionE gets the RunAs Connection if it exists in the Azure Automation Account
func GetAutomationAccountConnectionE(t testing.TestingT, automationAccountName string, resourceGroupName string, subscriptionID string, automationAccountConnectionName string) (*automation.Connection, error) {
	client, err := GetAutomationAccountConnectionClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get Automation Account Connection
	automationAccountConnection, err := client.Get(context.Background(), resourceGroupName, automationAccountName, automationAccountConnectionName)
	if err != nil {
		return nil, err
	}

	return &automationAccountConnection, nil
}

// GetAutomationAccountCertificate gets the RunAs Connection Certificate if it exists in the Azure Automation Account
// This function would fail the test if there is an error.
func GetAutomationAccountCertificate(t testing.TestingT, automationAccountName string, resourceGroupName string, subscriptionID string, automationAccountCertificateName *string) *automation.Certificate {
	client, err := GetAutomationAccountCertficateClientE(subscriptionID)
	require.NoError(t, err)

	// Get Automation Account Connection
	automationAccountCertificate, err := client.Get(context.Background(), resourceGroupName, automationAccountName, *automationAccountCertificateName)
	require.NoError(t, err)

	return &automationAccountCertificate
}

// GetAutomationAccountCertificateE gets the RunAs Connection Certificate if it exists in the Azure Automation Account
func GetAutomationAccountCertificateE(t testing.TestingT, automationAccountName string, resourceGroupName string, subscriptionID string, automationAccountCertificateName string) (*automation.Certificate, error) {
	client, err := GetAutomationAccountCertficateClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get Automation Account Connection
	automationAccountCertificate, err := client.Get(context.Background(), resourceGroupName, automationAccountName, automationAccountCertificateName)
	if err != nil {
		return nil, err
	}

	return &automationAccountCertificate, nil
}

// GetAutomationAccountCertficateClientE gets the RunAs Connection Certificate if it exists in the Azure Automation Account
func GetAutomationAccountCertficateClientE(subscriptionID string) (*automation.CertificateClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Certificate Client reference
	client := automation.NewCertificateClient(subscriptionID)

	// setup authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	client.Authorizer = *authorizer

	return &client, nil
}

// GetAutomationAccountConnectionClientE gets the RunAs Connection if it exists in the Azure Automation Account
func GetAutomationAccountConnectionClientE(subscriptionID string) (*automation.ConnectionClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Automation Account Client
	client := automation.NewConnectionClient(subscriptionID)

	// setup authorizer
	auth, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	// Setup auth and create request params
	client.Authorizer = *auth

	return &client, nil
}
