output "resource_group_name" {
  value = azurerm_resource_group.example.name
}

output "loganalytics_workspace_name" {
    value = azurerm_log_analytics_workspace.example.name
}

output "loganalytics_workspace_sku" {
    value = azurerm_log_analytics_workspace.example.sku
}

output "loganalytics_workspace_retention" {
    value = azurerm_log_analytics_workspace.example.retention_in_days
}