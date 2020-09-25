output "resource_group_name" {
  value = azurerm_resource_group.resourcegroup.name
}

output "loganalytics_workspace_name" {
    value = azurerm_log_analytics_workspace.loganalyticsworkspace.name
}

output "loganalytics_workspace_sku" {
    value = azurerm_log_analytics_workspace.loganalyticsworkspace.sku
}

output "loganalytics_workspace_retention" {
    value = azurerm_log_analytics_workspace.loganalyticsworkspace.retention_in_days
}