output "resource_group_name" {
  value = azurerm_resource_group.example.name
}

output "recovery_service_vault_name" {
  value = azurerm_recovery_services_vault.example.name
}

output "backup_policy_vm_name" {
  value = azurerm_backup_policy_vm.example.name
}
