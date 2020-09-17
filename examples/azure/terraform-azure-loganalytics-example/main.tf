resource "random_string" "default" {
  length = 3  
  lower = true
  number = false
  special = false
}

resource "azurerm_resource_group" "main" {
  name     =  format("%s-%s-%s", "terratest", lower(random_string.default.result), "loganalytics")
  location = var.location
}

resource "random_string" "defaultkv" {
  length = 8  
  lower = true
  number = false
  special = false
}

resource "azurerm_log_analytics_workspace" "logws" {
  name                = lower(random_string.defaultkv.result)
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name
  sku                 = "PerGB2018"
  retention_in_days   = 30
}