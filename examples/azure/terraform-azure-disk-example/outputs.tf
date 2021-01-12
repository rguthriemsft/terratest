output "resource_group_name" {
  value = azurerm_resource_group.disk_rg.name
}

output "disk_name" {
  value = azurerm_managed_disk.disk.name
}

output "disk_type" {
  value = azurerm_managed_disk.disk.storage_account_type
}
