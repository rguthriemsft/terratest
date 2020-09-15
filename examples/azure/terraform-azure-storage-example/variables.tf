
variable "prefix" {
  description = "The prefix that will be attached to all resources deployed."
  type        = string
  default     = "testprefix"
}

variable "location" {
  description = "The location to set for the storage account."
  type        = string
  default     = "East US"
}
