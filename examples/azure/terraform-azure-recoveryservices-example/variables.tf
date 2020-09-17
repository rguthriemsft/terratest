variable "location" {
  description = "The location to set for the storage account."
  type        = string
  default     = "East US"
}

variable "backup_policy_vm" {
  description = "The name of the vm backup policy."
  type        = string
  default     = "test"
}