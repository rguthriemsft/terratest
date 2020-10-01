output "resource_group_name" {
  value = azurerm_resource_group.net.name
}

output "virtual_network_name" {
  value = azurerm_virtual_network.net.name
}

output "subnet_name" {
  value = azurerm_subnet.net.name
}

output "public_address_name" {
  value = azurerm_public_ip.net.name
}

output "network_interface_internal" {
  value = azurerm_network_interface.net01.name
}

output "network_interface_external" {
  value = azurerm_network_interface.net02.name
}


