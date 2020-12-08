output "lb_default_name" {
  value = azurerm_lb.default.name
}

output "lb_private_name" {
  value = azurerm_lb.private.name
}

output "lb_private_fe_config_static_name" {
  value = azurerm_lb.private.frontend_ip_configuration[0].name
}

output "lb_private_fe_config_dynamic_name" {
  value = azurerm_lb.private.frontend_ip_configuration[1].name
}

output "lb_private_ip_static" {
  value = azurerm_lb.private.frontend_ip_configuration[0].private_ip_address
}

output "lb_private_ip_dynamic" {
  value = azurerm_lb.private.frontend_ip_configuration[1].private_ip_address
}

output "lb_public_name" {
  value = azurerm_lb.public.name
}

output "lb_public_fe_config_name" {
  value = azurerm_lb.public.frontend_ip_configuration[0].name
}

output "public_address_name" {
  value = azurerm_public_ip.pip.name
}

output "resource_group_name" {
  value = azurerm_resource_group.lb_rg.name
}
