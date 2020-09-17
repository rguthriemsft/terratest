resource "random_string" "defaultvault" {
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

resource "azurerm_resource_group" "example" {
  name     =  format("%s-%s-%s", "terratest", lower(random_string.default.result), "recoveryservices")
  location = var.location
}

resource "azurerm_recovery_services_vault" "example" {
  name                = lower(random_string.defaultvault.result)
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  sku                 = "Standard"
}

// Recovery Services Vault - Backup Policy 
resource "azurerm_backup_policy_vm" "example" {
  name                = format("%s-%s", lower(random_string.default.result), "backupvmpolicy")
  resource_group_name = azurerm_resource_group.example.name
  recovery_vault_name = azurerm_recovery_services_vault.example.name

  timezone = "UTC"

  backup {
    frequency = "Daily"
    time      = "23:00"
  }

  retention_daily {
    count = 10
  }

  retention_weekly {
    count    = 42
    weekdays = ["Sunday", "Wednesday", "Friday", "Saturday"]
  }

  retention_monthly {
    count    = 7
    weekdays = ["Sunday", "Wednesday"]
    weeks    = ["First", "Last"]
  }

  retention_yearly {
    count    = 77
    weekdays = ["Sunday"]
    weeks    = ["Last"]
    months   = ["January"]
  }
}