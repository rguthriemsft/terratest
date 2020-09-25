package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2020-03-01-preview/operationalinsights"
)

// LogAnalyticsWorkspaceExists indicates whether the operatonal insights workspaces exists; otherwise false or error
func LogAnalyticsWorkspaceExists(workspaceName, resourceGroupName, subscriptionID string) bool {
	ws, err := GetLogAnalyticsWorkspaceE(workspaceName, resourceGroupName, subscriptionID)
	if err != nil {
		return false
	}

	return (*ws.Name == workspaceName)
}

// GetLogAnalyticsWorkspaceSku return the log analytics workspace SKU as string as one of the following: Free, Standard, Premium, PerGB2018, CapacityReservation; otherwise empty string "".
func GetLogAnalyticsWorkspaceSku(workspaceName, resourceGroupName, subscriptionID string) string {
	ws, err := GetLogAnalyticsWorkspaceE(workspaceName, resourceGroupName, subscriptionID)
	if err != nil {
		return ""
	}
	return string(ws.Sku.Name)
}

// GetLogAnalyticsWorkspaceRetentionPeriodDays returns the log analytics workspace retention period in days; otherwise -1.
func GetLogAnalyticsWorkspaceRetentionPeriodDays(workspaceName, resourceGroupName, subscriptionID string) int32 {
	ws, err := GetLogAnalyticsWorkspaceE(workspaceName, resourceGroupName, subscriptionID)
	if err != nil {
		return -1
	}
	return *ws.RetentionInDays
}

// GetLogAnalyticsWorkspaceE gets an operational insights workspace if it exists in a subscription.
// This function would fail the test if there is an error.
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

// GetLogAnalyticsWorkspacesClientE return workspaces client; otherwise error.
// This function would fail the test if there is an error.
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
