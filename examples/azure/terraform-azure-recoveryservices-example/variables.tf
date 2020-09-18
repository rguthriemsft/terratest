variable "location" {
  description = "The location to set for the storage account."
  type        = string
  default     = "East US"
}

variable "resource_group_name" {
  description = "The name to set for the resource group."
  default     = "terratest-azure-resources"
}

variable "vault_name" {
  description = "The name to set for the recovery services vault."
  type        = string
  default     = "vault"
}

variable "policy_name" {
  description = "The name to set for the recovery services backup policy."
  type        = string
  default     = "backup-policy"
}

