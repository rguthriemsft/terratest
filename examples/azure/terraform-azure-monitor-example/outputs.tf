output "resource_group_name" {
  value = azurerm_resource_group.monitor.name
}

output "diagnostic_setting_name" {
  value = azurerm_monitor_diagnostic_setting.monitor.name
}

output "diagnostic_setting_id" {
  value = azurerm_monitor_diagnostic_setting.monitor.id
}

output "keyvault_id" {
  value = azurerm_key_vault.monitor.id
}
