output "resource_group_name" {
  value = azurerm_resource_group.resource_group.name
}

output "key_vault_name" {
  value = azurerm_key_vault.key_vault.name
}

output "secret_name" {
  value = azurerm_key_vault_secret.key_vault_secret.name
}

output "key_name" {
  value = azurerm_key_vault_key.key_vault_key.name
}

output "certificate_name" {
  value = azurerm_key_vault_certificate.key_vault_certificate.name
}