

resource "azurerm_resource_group" "example" {
  name     =  var.resource_group_name
  location = var.location
}

resource "azurerm_storage_account" "example" {
  name                     = var.storage_account_name
  resource_group_name      = azurerm_resource_group.example.name
  location                 = azurerm_resource_group.example.location
  account_kind             = var.storage_account_kind
  account_tier             = var.storage_account_tier
  account_replication_type = var.storage_replication_type
}

resource "azurerm_storage_container" "example" {
  name = "container1"
  storage_account_name  = azurerm_storage_account.example.name
  container_access_type = var.container_access_type
}