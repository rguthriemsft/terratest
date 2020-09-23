output "resource_group_name" {
  value = azurerm_resource_group.resourcegroup.name
}

output "key_vault_name" {
  value = azurerm_key_vault.keyvault.name
}

output "secret_name" {
  value = azurerm_key_vault_secret.keyvaultsecret.name
}

output "key_name" {
  value = azurerm_key_vault_key.keyvaultkey.name
}

output "certificate_name" {
  value = azurerm_key_vault_certificate.keyvaultcertificate.name
}