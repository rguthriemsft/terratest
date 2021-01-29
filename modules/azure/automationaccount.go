package azure

import (
	"context"
	"fmt"
	"time"

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

// AutomationAccountDscExists indicates whether the specified Azure Automation Account DSC exists.
// This function would fail the test if there is an error.
func AutomationAccountDscExists(t testing.TestingT, dscConfiguraitonName string, automationAccountName string, resourceGroupName string, subscriptionID string) bool {
	exists, err := AutomationAccountDscExistsE(t, dscConfiguraitonName, automationAccountName, resourceGroupName, subscriptionID)
	require.NoError(t, err)

	return exists
}

// AutomationAccountDscExistsE indicates whether the specified Azure Automation Account exists.
func AutomationAccountDscExistsE(t testing.TestingT, dscConfiguraitonName string, automationAccountName string, resourceGroupName string, subscriptionID string) (bool, error) {
	_, err := GetAutomationAccountDscConfigurationE(t, dscConfiguraitonName, automationAccountName, resourceGroupName, subscriptionID)
	if err != nil {
		if ResourceNotFoundErrorExists(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// AutomationAccountDscCompiled indicates whether the specified Azure Automation Account DSC compiled successfully.
// This function would fail the test if there is an error.
func AutomationAccountDscCompiled(t testing.TestingT, dscConfiguraitonName string, automationAccountName string, resourceGroupName string, subscriptionID string) bool {
	compiled, err := AutomationAccountDscCompiledE(t, dscConfiguraitonName, automationAccountName, resourceGroupName, subscriptionID)
	require.NoError(t, err)

	return compiled
}

// AutomationAccountDscCompiledE indicates whether the specified Azure Automation Account DSC compiled successfully.
// DSC compilation is performed via a Terraform null_resource using PowerShell Core, executing async.
// Compilation can take a few minutes to spin up resources requiring a retry mechanism to allow for this
func AutomationAccountDscCompiledE(t testing.TestingT, dscConfiguraitonName string, automationAccountName string, resourceGroupName string, subscriptionID string) (bool, error) {
	seconds := 30 // 30 second initial delay
	sleep := time.Duration(seconds) * time.Second
	maxAttempts := 5 // try 5 times, doubling the delay each time.

	for i := 1; i < maxAttempts; i++ {
		time.Sleep(sleep)
		dscCompileJobStatus, err := AutomationAccountDscCompileJobStatusE(t, dscConfiguraitonName, automationAccountName, resourceGroupName, subscriptionID)
		if err != nil {
			return false, err
		}
		// check if status === Completed
		if dscCompileJobStatus == "Completed" {
			return true, nil
		}
		sleep = 2 * sleep
	}
	return false, nil
}

// AutomationAccountRunAsCertificateThumbprintMatches indicates whether the specified Azure Automation Account RunAs Certificate exists.
// This function would fail the test if there is an error.
func AutomationAccountRunAsCertificateThumbprintMatches(t testing.TestingT, runAsCertificateThumbprint string, runAsCertificateName string, automationAccountName string, resourceGroupName string, subscriptionID string) bool {
	exists, err := AutomationAccountRunAsCertificateThumbprintMatchesE(t, runAsCertificateThumbprint, runAsCertificateName, automationAccountName, resourceGroupName, subscriptionID)
	require.NoError(t, err)

	return exists
}

// AutomationAccountRunAsCertificateThumbprintMatchesE indicates whether the specified Azure Automation Account RunAs Certificate exists.
func AutomationAccountRunAsCertificateThumbprintMatchesE(t testing.TestingT, runAsCertificateThumbprint string, runAsCertificateName string, automationAccountName string, resourceGroupName string, subscriptionID string) (bool, error) {
	certificate, err := GetAutomationAccountCertificateE(t, runAsCertificateName, automationAccountName, resourceGroupName, subscriptionID)
	if err != nil {
		if ResourceNotFoundErrorExists(err) {
			return false, nil
		}
		return false, err
	}
	return *certificate.CertificateProperties.Thumbprint == runAsCertificateThumbprint, nil
}

// AutomationAccountRunAsConnectionValidates indicates whether the specified Azure Automation Account RunAs Account exists.
// This function would fail the test if there is an error.
func AutomationAccountRunAsConnectionValidates(t testing.TestingT, automationAccountrunAsAccountName string, runAsConnectionType string, runAsCertificateThumbprint string, automationAccountName string, resourceGroupName string, subscriptionID string) bool {
	exists, err := AutomationAccountRunAsConnectionValidatesE(t, automationAccountrunAsAccountName, runAsConnectionType, runAsCertificateThumbprint, automationAccountName, resourceGroupName, subscriptionID)
	require.NoError(t, err)

	return exists
}

// AutomationAccountRunAsConnectionValidatesE indicates whether the specified Azure Automation Account RunAs Account exists.
func AutomationAccountRunAsConnectionValidatesE(t testing.TestingT, automationAccountrunAsAccountName string, runAsConnectionType string, runAsCertificateThumbprint string, automationAccountName string, resourceGroupName string, subscriptionID string) (bool, error) {
	runAsAccountConnection, err := GetAutomationAccountRunAsConnectionE(t, automationAccountrunAsAccountName, automationAccountName, resourceGroupName, subscriptionID)
	if err != nil {
		if ResourceNotFoundErrorExists(err) {
			return false, nil
		}
		return false, err
	}
	return *runAsAccountConnection.Name == automationAccountrunAsAccountName &&
		*runAsAccountConnection.ConnectionProperties.ConnectionType.Name == runAsConnectionType &&
		*runAsAccountConnection.ConnectionProperties.FieldDefinitionValues["CertificateThumbprint"] == runAsCertificateThumbprint, nil
}

// AutomationAccountDscAppliedSuccessfullyToVM indicates whether the specified Azure Automation Account compiled DSC has successfully been applied to the target VM.
// This function would fail the test if there is an error.
func AutomationAccountDscAppliedSuccessfullyToVM(t testing.TestingT, dscConfiguraitonName string, vmName string, automationAccountName string, resourceGroupName string, subscriptionID string) bool {
	exists, err := AutomationAccountDscAppliedSuccessfullyToVME(t, dscConfiguraitonName, vmName, automationAccountName, resourceGroupName, subscriptionID)
	require.NoError(t, err)

	return exists
}

// AutomationAccountDscAppliedSuccessfullyToVME indicates whether the specified Azure Automation Account compiled DSC has successfully been applied to the target VM.
func AutomationAccountDscAppliedSuccessfullyToVME(t testing.TestingT, dscConfiguraitonName string, vmName string, automationAccountName string, resourceGroupName string, subscriptionID string) (bool, error) {
	dscNodeConfig, err := GetAutomationAccountDscNodeConfigurationE(t, dscConfiguraitonName, vmName, automationAccountName, resourceGroupName, subscriptionID)
	if err != nil {
		if ResourceNotFoundErrorExists(err) {
			return false, nil
		}
		return false, err
	}
	return *dscNodeConfig.Status == "Compliant", nil
}

/////////
// Helper Methods for above checks
/////////

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

// GetAutomationAccountDscConfigurationE returns the Azure Automation Account DscConfiguration by Automation Account name if it exists in the subscription
func GetAutomationAccountDscConfigurationE(t testing.TestingT, dscConfigurationName string, automationAccountName string, resourceGroupName string, subscriptionID string) (*automation.DscConfiguration, error) {
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

// AutomationAccountDscCompileJobStatusE returns the Azure Automation Account DscConfiguration by Automation Account name if it exists in the subscription
func AutomationAccountDscCompileJobStatusE(t testing.TestingT, dscConfigurationName string, automationAccountName string, resourceGroupName string, subscriptionID string) (string, error) {
	client, err := GetDscCompilationJobClientE(subscriptionID)
	if err != nil {
		return "", err
	}

	filter := fmt.Sprintf("properties/configuration/name eq '%s'", dscConfigurationName)
	dscCompilationJobListResultPage, err := client.ListByAutomationAccount(context.Background(), resourceGroupName, automationAccountName, filter)
	if err != nil {
		return "", err
	}

	var dscCompilationJobs []automation.DscCompilationJob
	var mostRecentCompileJobTick int64
	var mostRecentCompileJobStatus string
	// Loop through filtered pages of DSC compilation jobs to find latest compilation job
	for dscCompilationJobListResultPage.NotDone() {
		dscCompilationJobs = dscCompilationJobListResultPage.Values()
		// Loop through compilation jobs in the current page
		for _, element := range dscCompilationJobs {
			if element.CreationTime.Unix() > mostRecentCompileJobTick {
				mostRecentCompileJobTick = element.CreationTime.Unix()
				mostRecentCompileJobStatus = (string)(element.Status)
			}
		}
		err := dscCompilationJobListResultPage.Next()
		if err != nil {
			return "", err
		}
	}
	// Check to ensure  DSC compilation jobs are present (i.e. mostRecentCompileJobTick is non zero)
	if mostRecentCompileJobTick == 0 {
		panic("No compilation jobs present for this DSC configuraiton, or compilation jobs are 'suspeneded' in the Automation Account.")
	} else {
		return mostRecentCompileJobStatus, nil
	}
}

// GetAutomationAccountCertificateE gets the RunAs Connection Certificate if it exists in the Azure Automation Account
func GetAutomationAccountCertificateE(t testing.TestingT, automationAccountCertificateName string, automationAccountName string, resourceGroupName string, subscriptionID string) (*automation.Certificate, error) {
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

// GetAutomationAccountDscNodeConfigurationE gets the Node Configuration if it exists in the Azure Automation Account
func GetAutomationAccountDscNodeConfigurationE(t testing.TestingT, dscConfiguraitonName string, vmName string, automationAccountName string, resourceGroupName string, subscriptionID string) (*automation.DscNode, error) {
	client, err := GetAutomationAccountDscNodeConfigClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	filter := fmt.Sprintf("name eq '%s'", vmName)
	dscNodeListResultPage, err := client.ListByAutomationAccount(context.Background(), resourceGroupName, automationAccountName, filter)
	if err != nil {
		return nil, err
	}

	var dscNodeList []automation.DscNode
	var dscNodeID string
	for dscNodeListResultPage.NotDone() {
		dscNodeList = dscNodeListResultPage.Values()
		// Loop through compilation jobs in the current page
		for _, element := range dscNodeList {
			if *element.Name == vmName && *element.NodeConfiguration.Name == dscConfiguraitonName {
				dscNodeID = *element.NodeID
			}
		}

		err := dscNodeListResultPage.Next()
		if err != nil {
			return nil, err
		}
	}

	// Get Automation Account Connection
	dscNodeConfig, err := client.Get(context.Background(), resourceGroupName, automationAccountName, dscNodeID)
	if err != nil {
		return nil, err
	}

	return &dscNodeConfig, nil
}

// GetAutomationAccountRunAsConnectionE gets the RunAs Connection if it exists in the Azure Automation Account
func GetAutomationAccountRunAsConnectionE(t testing.TestingT, automationAccountRunAsConnectionName string, automationAccountName string, resourceGroupName string, subscriptionID string) (*automation.Connection, error) {
	client, err := GetAutomationAccountRunAsConnectionClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get Automation Account Connection
	automationAccountRunAsConnection, err := client.Get(context.Background(), resourceGroupName, automationAccountName, automationAccountRunAsConnectionName)
	if err != nil {
		return nil, err
	}

	return &automationAccountRunAsConnection, nil
}

// GetCertificateClientE returns the Azure Automation Account Certfiicate client
func GetCertificateClientE(subscriptionID string) (*automation.CertificateClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Automation Account Certificate client
	client := automation.NewCertificateClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	client.Authorizer = *authorizer

	return &client, nil
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

// GetAutomationAccountRunAsConnectionClientE gets the RunAs Connection if it exists in the Azure Automation Account
func GetAutomationAccountRunAsConnectionClientE(subscriptionID string) (*automation.ConnectionClient, error) {
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

// GetAutomationAccountDscNodeConfigClientE gets the RunAs Connection if it exists in the Azure Automation Account
func GetAutomationAccountDscNodeConfigClientE(subscriptionID string) (*automation.DscNodeClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the Automation Account Client
	client := automation.NewDscNodeClient(subscriptionID)

	// setup authorizer
	auth, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	// Setup auth and create request params
	client.Authorizer = *auth

	return &client, nil
}
