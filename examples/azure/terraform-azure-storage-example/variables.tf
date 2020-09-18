
variable "location" {
  description = "The location to set for the storage account."
  type        = string
  default     = "East US"
}

variable "resource_group_name" {
  description = "The name to set for the resource group."
  type        = string
  default     = "azure-k8stest"
}

variable "storage_account_name" {
  description = "The name to set for the storage account."
  type        = string
  default     = "terrateststore"
}

variable "storage_account_kind" {
  description = "The kind of storage account to set"
  type        = string
  default     = "StorageV2"
}

variable "storage_account_tier" {
  description = "The tier of storage account to set"
  type        = string
  default     = "Standard"
}

variable "storage_replication_type" {
  description = "The replication type of storage account to set"
  type        = string
  default     = "GRS"
}

variable "container_access_type" {
  description = "The replication type of storage account to set"
  type        = string
  default     = "private"
}
