package azure

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/automation/mgmt/automation"
	"github.com/gruntwork-io/terratest/modules/testing"
)

// AutomationAccountExists indicates whether the specified Azure Automation Account exists.
// This function would fail the test if there is an error.
func AutomationAccountExists(t testing.TestingT, automationAccountName string, resourceGroupName string, subscriptionID string) bool {
	exists, err := AutomationAccountExistsE(t, automationAccountName, resourceGroupName, subscriptionID)
	if err != nil {
		panic("error")
	}
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
	if err != nil {
		panic(err)
	}
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
	if err != nil {
		panic(err)
	}
	return compiled
}

// AutomationAccountDSCCompiledE indicates whether the specified Azure Automation Account DSC compiled successfully.
// DSC Compilation requires the Automation Account to spin-up resources, taking time to complete.
func AutomationAccountDSCCompiledE(t testing.TestingT, automationAccountName string, dscConfiguraitonName string, resourceGroupName string, subscriptionID string) (bool, error) {
	seconds := 30 // 30 second initial delay
	sleep := time.Duration(seconds) * time.Second
	maxAttempts := 5 // try 5 times, doubling the delay each time.

	for i := 1; i < maxAttempts; i++ {
		time.Sleep(sleep)
		dscCompileJobStatus, err := AutomationAccountDSCCompileJobStatus(t, automationAccountName, dscConfiguraitonName, resourceGroupName, subscriptionID)
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

// AutomationAccountHasRunAsCertificate indicates whether the specified Azure Automation Account RunAs Certificate exists.
// This function would fail the test if there is an error.
func AutomationAccountHasRunAsCertificate() {

}

// AutomationAccountHasRunAsAccount indicates whether the specified Azure Automation Account RunAs Account exists.
// This function would fail the test if there is an error.
func AutomationAccountHasRunAsAccount() {

}

// DSCAppliedSuccessfullyToVM indicates whether the specified Azure Automation Account compiled DSC has successfully been applied to the target VM.
// This function would fail the test if there is an error.
func DSCAppliedSuccessfullyToVM() {

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

// AutomationAccountDSCCompileJobStatus returns the Azure Automation Account DscConfiguration by Automation Account name if it exists in the subscription
func AutomationAccountDSCCompileJobStatus(t testing.TestingT, automationAccountName string, dscConfigurationName string, resourceGroupName string, subscriptionID string) (string, error) {
	resourceGroupName, err := getTargetAzureResourceGroupName(resourceGroupName)
	if err != nil {
		return "", err
	}

	client, err := GetDscCompilationJobClientE(subscriptionID)
	if err != nil {
		return "", err
	}

	filter := fmt.Sprintf("properties/configuration/name eq '%s'", dscConfigurationName)
	dscCompilationJobListResultPage, err := client.ListByAutomationAccount(context.Background(), resourceGroupName, automationAccountName, filter)
	if err != nil {
		return "", err
	}

	var dscCompilationJob []automation.DscCompilationJob
	var mostRecentCompileJobTick int64
	var mostRecentCompileJobStatus string
	// Loop through filtered pages of DSC compilation jobs to find latest compilation job
	for dscCompilationJobListResultPage.NotDone() {
		dscCompilationJob = dscCompilationJobListResultPage.Values()
		// Loop through compilation jobs in the current page
		for _, element := range dscCompilationJob {
			if element.CreationTime.Unix() > mostRecentCompileJobTick {
				mostRecentCompileJobTick = element.CreationTime.Unix()
				mostRecentCompileJobStatus = (string)(element.Status)
			}
		}
		dscCompilationJobListResultPage.Next()
	}
	// Check to ensure  DSC compilation jobs are present (i.e. mostRecentCompileJobTick is non zero)
	if mostRecentCompileJobTick == 0 {
		panic("No compilation jobs present for this DSC configuraiton, or compilation jobs are 'suspeneded' in the Automation Account.")
	} else {
		return mostRecentCompileJobStatus, nil
	}
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
func GetAutomationAccountConnection(automationAccountName string, resourceGroupName string, subscriptionID string, automationAccountConnectionName string) *automation.Connection {
	client, err := GetAutomationAccountConnectionClientE(subscriptionID)
	if err != nil {
		panic("error")
	}

	// Get Automation Account Connection
	automationAccountConnection, err := client.Get(context.Background(), resourceGroupName, automationAccountName, automationAccountConnectionName)
	if err != nil {
		panic("error")
	}

	return &automationAccountConnection
}

// GetAutomationAccountConnectionE gets the RunAs Connection if it exists in the Azure Automation Account
func GetAutomationAccountConnectionE(automationAccountName string, resourceGroupName string, subscriptionID string, automationAccountConnectionName string) (*automation.Connection, error) {
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
func GetAutomationAccountCertificate(automationAccountName string, resourceGroupName string, subscriptionID string, automationAccountCertificateName *string) *automation.Certificate {
	client, err := GetAutomationAccountCertficateClientE(subscriptionID)
	if err != nil {
		panic("error")
	}

	// Get Automation Account Connection
	automationAccountCertificate, err := client.Get(context.Background(), resourceGroupName, automationAccountName, *automationAccountCertificateName)
	if err != nil {
		panic("error")
	}

	return &automationAccountCertificate
}

// GetAutomationAccountCertificateE gets the RunAs Connection Certificate if it exists in the Azure Automation Account
func GetAutomationAccountCertificateE(automationAccountName string, resourceGroupName string, subscriptionID string, automationAccountCertificateName string) (*automation.Certificate, error) {
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
