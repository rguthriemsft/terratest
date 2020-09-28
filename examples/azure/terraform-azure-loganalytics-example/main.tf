# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY LOG ANALYTICS
# This is an example of how to deploy a Log Analytics workspace resource.
# ---------------------------------------------------------------------------------------------------------------------
# See test/azure/terraform_azure_loganalytics_example_test.go for how to write automated tests for this code.
# ---------------------------------------------------------------------------------------------------------------------

provider "azurerm" {
  version = "~> 2.20"
  features {}
}

# ---------------------------------------------------------------------------------------------------------------------
# PIN TERRAFORM VERSION TO >= 0.12
# The examples have been upgraded to 0.12 syntax
# ---------------------------------------------------------------------------------------------------------------------

terraform {
  # This module is now only being tested with Terraform 0.13.x. However, to make upgrading easier, we are setting
  # 0.12.26 as the minimum version, as that version added support for required_providers with source URLs, making it
  # forwards compatible with 0.13.x code.
  required_version = ">= 0.12.26"
}

# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY A RESOURCE GROUP
# ---------------------------------------------------------------------------------------------------------------------

resource "azurerm_resource_group" "resourcegroup" {
  name     = "terratest-log-rg-${var.postfix}"
  location = var.location
}

resource "azurerm_log_analytics_workspace" "loganalyticsworkspace" {
  name                = "log-ws-${var.postfix}"
  location            = azurerm_resource_group.resourcegroup.location
  resource_group_name = azurerm_resource_group.resourcegroup.name
  sku                 = "PerGB2018"
  retention_in_days   = 30
}