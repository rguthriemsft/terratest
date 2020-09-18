output "resource_group_name" {
  value = azurerm_resource_group.example.name
}

output "storage_account_name" {
  value = azurerm_storage_account.example.name
}

output "storage_account_account_tier" {
  value = azurerm_storage_account.example.account_tier
}

output "storage_account_account_kind" {
  value = azurerm_storage_account.example.account_kind
}

output "storage_container_name" {
  value = azurerm_storage_container.example.name
}
