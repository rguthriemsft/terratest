
variable "location" {
  description = "The location to set for the storage account."
  type        = string
  default     = "East US"
}

variable "resource_group_name" {
  description = "The name to set for the resource group."
  type        = string
  default     = "loganalytics"
}

variable "loganalytics_workspace_name" {
  description = "The name to set for the log analytics workspace."
  type        = string
  default     = "loganalyticsws"
}
