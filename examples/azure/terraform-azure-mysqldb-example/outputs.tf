output "resource_group_name" {
  value = azurerm_resource_group.mysql_rg.name
}

output "mysql_server_name" {
  value = azurerm_mysql_server.mysqlserver.name
}

output "sql_server_full_domain_name" {
  value = azurerm_mysql_server.mysqlserver.fqdn
}

output "sql_server_admin_login" {
  value = azurerm_mysql_server.mysqlserver.administrator_login
}

output "sql_server_admin_login_pass" {
  value     = azurerm_mysql_server.mysqlserver.administrator_login_password
  sensitive = true
}

output "mysql_database_name" {
  value = azurerm_mysql_database.mysqldb.name
}
