// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package azure

import (
	"context"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/automation/mgmt/2015-10-31/automation"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// GetAutomationAccountE returns the Azure Automation Account  by name if it exists in the subscription
// This function would fail the test if there is an error.
func GetAutomationAccount(t testing.TestingT, automationAccountName *string, resourceGroupName *string) *automation.Account {
	client, err := GetAutomationAccountClient(os.Getenv(helper.SubscriptionIDEnvName))
	require.NoError(t, err)

	automationAccount, err := client.Get(context.Background(), *resourceGroupName, *automationAccountName)
	require.NoError(t, err)

	return &automationAccount
}

// GetAutomationAccountE returns the Azure Automation Account by name if it exists in the subscription
func GetAutomationAccountE(t testing.TestingT, automationAccountName *string, resourceGroupName *string, subscriptionID string) (*automation.Account, error) {
	// Validate resource group name and subscription ID
	resourceGroupName, err := getTargetAzureResourceGroupName(resourceGroupName)
	if err != nil {
		return nil, err
	}

	client, err := GetAutomationAccountClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	automationAccount, err := client.Get(context.Background(), *resourceGroupName, *automationAccountName)
	if err != nil {
		return nil, err
	}

	return &automationAccount, nil
}

// GetAutomationAccountClient returns the Azure Automation Account client
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

// GetAutomationAccountConnection gets the RunAs Connection if it exists in the Azure Automation Account
// This function would fail the test if there is an error.
func GetAutomationAccountConnection(t testing.TestingT, automationAccountName *string, resourceGroupName *string, automationAccountConnectionName *string) *automation.Connection {
	client, err := GetAutomationAccountConnectionClientE(os.Getenv(helper.SubscriptionIDEnvName))
	require.NoError(t, err)

	// Get Automation Account Connection
	automationAccountConnection, err := client.Get(context.Background(), *resourceGroupName, *automationAccountName, *automationAccountConnectionName)
	require.NoError(t, err)

	return &automationAccountConnection
}

// GetAutomationAccountConnectionE gets the RunAs Connection if it exists in the Azure Automation Account
func GetAutomationAccountConnectionE(t testing.TestingT, automationAccountName *string, resourceGroupName *string, automationAccountConnectionName *string) (*automation.Connection, error) {
	client, err := GetAutomationAccountConnectionClientE(os.Getenv(helper.SubscriptionIDEnvName))
	if err != nil {
		return nil, err
	}

	// Get Automation Account Connection
	automationAccountConnection, err := client.Get(context.Background(), *resourceGroupName, *automationAccountName, *automationAccountConnectionName)
	if err != nil {
		return nil, err
	}

	return &automationAccountConnection, nil
}

// GetAutomationAccountConnectionClient gets the RunAs Connection if it exists in the Azure Automation Account
// This function would fail the test if there is an error.
func GetAutomationAccountConnectionClient(subscriptionID string) *automation.ConnectionClient {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	require.NoError(t, err)

	client := automation.NewConnectionClient(subscriptionID)

	// setup authorizer
	auth, err := NewAuthorizer()
	require.NoError(t, err)

	// Setup auth and create request params
	client.Authorizer = *auth

	return &client
}

// GetAutomationAccountCertificate gets the RunAs Connection Certificate if it exists in the Azure Automation Account
// This function would fail the test if there is an error.
func GetAutomationAccountCertificate(t testing.TestingT, automationAccountName *string, resourceGroupName *string, automationAccountCertificateName *string) *automation.Certificate {
	client, err := GetAutomationAccountCertficateClient(os.Getenv(helper.SubscriptionIDEnvName))
	require.NoError(t, err)

	// Get Automation Account Connection
	automationAccountCertificate, err := client.Get(context.Background(), *resourceGroupName, *automationAccountName, *automationAccountCertificateName)
	require.NoError(t, err)

	return &automationAccountCertificate
}

// GetAutomationAccountCertificateE gets the RunAs Connection Certificate if it exists in the Azure Automation Account
func GetAutomationAccountCertificateE(t testing.TestingT, automationAccountName *string, resourceGroupName *string, automationAccountCertificateName *string) (*automation.Certificate, error) {
	client, err := GetAutomationAccountCertficateClient(os.Getenv(helper.SubscriptionIDEnvName))
	if err != nil {
		return nil, err
	}

	// Get Automation Account Connection
	automationAccountCertificate, err := client.Get(context.Background(), *resourceGroupName, *automationAccountName, *automationAccountCertificateName)
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
