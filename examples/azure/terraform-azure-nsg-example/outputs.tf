output "vm_name" {
  value = azurerm_virtual_machine.vm_example.name
}

output "resource_group_name" {
  value = azurerm_resource_group.nsg_rg.name
}

output "nsg_name" {
  value = azurerm_network_security_group.nsg_example.name
}

output "ssh_rule_name" {
  value = azurerm_network_security_rule.allow_ssh.name
}

output "http_rule_name" {
  value = azurerm_network_security_rule.block_http.name
}
