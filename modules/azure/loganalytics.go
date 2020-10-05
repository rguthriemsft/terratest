package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2020-03-01-preview/operationalinsights"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// LogAnalyticsWorkspaceExists indicates whether the operatonal insights workspaces exists.
// This function would fail the test if there is an error.
func LogAnalyticsWorkspaceExists(t testing.TestingT, workspaceName string, resourceGroupName string, subscriptionID string) bool {
	exists, err := LogAnalyticsWorkspaceExistsE(workspaceName, resourceGroupName, subscriptionID)
	require.NoError(t, err)
	return exists
}

// GetLogAnalyticsWorkspaceSku return the log analytics workspace SKU as string as one of the following: Free, Standard, Premium, PerGB2018, CapacityReservation; otherwise empty string "".
// This function would fail the test if there is an error.
func GetLogAnalyticsWorkspaceSku(t testing.TestingT, workspaceName string, resourceGroupName string, subscriptionID string) string {
	ws, err := GetLogAnalyticsWorkspaceE(workspaceName, resourceGroupName, subscriptionID)
	require.NoError(t, err)
	return string(ws.Sku.Name)
}

// GetLogAnalyticsWorkspaceRetentionPeriodDays returns the log analytics workspace retention period in days.
// This function would fail the test if there is an error.
func GetLogAnalyticsWorkspaceRetentionPeriodDays(t testing.TestingT, workspaceName string, resourceGroupName string, subscriptionID string) int32 {
	ws, err := GetLogAnalyticsWorkspaceE(workspaceName, resourceGroupName, subscriptionID)
	require.NoError(t, err)
	return *ws.RetentionInDays
}

// GetLogAnalyticsWorkspaceE gets an operational insights workspace if it exists in a subscription.
func GetLogAnalyticsWorkspaceE(workspaceName, resoureGroupName, subscriptionID string) (*operationalinsights.Workspace, error) {
	client, err := GetLogAnalyticsWorkspacesClientE(subscriptionID)
	if err != nil {
		return nil, err
	}
	ws, err := client.Get(context.Background(), resoureGroupName, workspaceName)
	if err != nil {
		return nil, err
	}
	return &ws, nil
}

// LogAnalyticsWorkspaceExistsE indicates whether the operatonal insights workspaces exists and may return an error.
func LogAnalyticsWorkspaceExistsE(workspaceName string, resourceGroupName string, subscriptionID string) (bool, error) {
	_, err := GetLogAnalyticsWorkspaceE(workspaceName, resourceGroupName, subscriptionID)
	if err != nil {
		if ResourceNotFoundErrorExists(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// GetLogAnalyticsWorkspacesClientE return workspaces client; otherwise error.
func GetLogAnalyticsWorkspacesClientE(subscriptionID string) (*operationalinsights.WorkspacesClient, error) {
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		fmt.Println("Workspace client error getting subscription")
		return nil, err
	}
	client := operationalinsights.NewWorkspacesClient(subscriptionID)
	authorizer, err := NewAuthorizer()
	if err != nil {
		fmt.Println("authorizer error")
		return nil, err
	}
	client.Authorizer = *authorizer
	return &client, nil
}
