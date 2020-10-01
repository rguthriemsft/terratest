# ---------------------------------------------------------------------------------------------------------------------
# ENVIRONMENT VARIABLES
# Define these secrets as environment variables
# ---------------------------------------------------------------------------------------------------------------------

# ARM_CLIENT_ID
# ARM_CLIENT_SECRET
# ARM_SUBSCRIPTION_ID
# ARM_TENANT_ID

# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ---------------------------------------------------------------------------------------------------------------------

# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------

variable "location" {
  description = "The location to set for the storage account."
  type        = string
  default     = "East US"
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

variable "postfix" {
  description = "A postfix string to centrally mitigate resource name collisions"
  type        = string
  default     = "resource"
}


