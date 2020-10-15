output "resource_group_name" {
  value = azurerm_resource_group.sql_rg.name
}

output "sql_database_id" {
  value = azurerm_sql_database.sqldb.id
}

output "sql_database_name" {
  value = azurerm_sql_database.sqldb.name
}

output "sql_server_id" {
  value = azurerm_sql_server.sqlserver.id
}

output "sql_server_name" {
  value = azurerm_sql_server.sqlserver.name
}

output "sql_server_full_domain_name" {
  value = azurerm_sql_server.sqlserver.fully_qualified_domain_name
}

output "sql_server_admin_login" {
  value = azurerm_sql_server.sqlserver.administrator_login
}

output "sql_server_admin_login_pass" {
  value     = azurerm_sql_server.sqlserver.administrator_login_password
  sensitive = true
}
