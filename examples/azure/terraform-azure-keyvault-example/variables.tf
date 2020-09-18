
variable "location" {
  description = "The location to set for the storage account."
  type        = string
  default     = "East US"
}

variable "resource_group_name" {
  description = "The name to set for the resource group."
  type        = string
  default     = "azure-keyvault-test"
}

variable "key_vault_name" {
  description = "The name to set for the key vault."
  type        = string
  default     = "terratestkeyvault"
}

variable "secret_name" {
  description = "The name to set for the key vault secret."
  type        = string
  default     = "secret1"
}

variable "key_name" {
  description = "The name to set for the key vault key."
  type        = string
  default     = "key1"
}

variable "certificate_name" {
  description = "The name to set for the key vault certificate."
  type        = string
  default     = "certificate1"
}


