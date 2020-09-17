resource "random_string" "storedefault" {
  length = 8  
  lower = true
  number = false
  special = false
}

resource "random_string" "default" {
  length = 3  
  lower = true
  number = false
  special = false
}

resource "azurerm_resource_group" "main" {
  name     =  format("%s-%s-%s", "terratest", lower(random_string.default.result), "storage")
  location = var.location
}


resource "azurerm_storage_account" "storageaccount" {
  name                     = lower(random_string.storedefault.result)
  resource_group_name      = azurerm_resource_group.main.name
  location                 = azurerm_resource_group.main.location
  account_kind             = "StorageV2"
  account_tier             = "Standard"
  account_replication_type = "GRS"
}

resource "azurerm_storage_container" "container" {
  name = "container1"
  storage_account_name  = azurerm_storage_account.storageaccount.name
  container_access_type = "private"
}