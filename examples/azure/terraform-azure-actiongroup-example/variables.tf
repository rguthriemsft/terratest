# ---------------------------------------------------------------------------------------------------------------------
# ENVIRONMENT VARIABLES
# Define these secrets as environment variables
# ---------------------------------------------------------------------------------------------------------------------

# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ---------------------------------------------------------------------------------------------------------------------

variable "resource_group_name" {
  description = "Name of the resource group that exists in Azure"
  type        = string
}

variable "app_name" {
  description = "The base name of the application used in the naming convention."
  type        = string
}

variable "location" {
  description = "Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created."
  type        = string
}

variable "short_name" {
  description = "Shorthand name for SMS texts."
  type        = string
}

# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------

variable "enable_email" {
  description = "Enable email alert capabilities"
  type        = bool
  default     = false
}

variable "email_name" {
  description = "Friendly Name for email address"
  type        = string
  default     = ""
}

variable "email_address" {
  description = "email address to send alerts to"
  type        = string
  default     = ""
}

variable "enable_sms" {
  description = "Enable Texting Alerts"
  type        = bool
  default     = false
}

variable "sms_name" {
  description = "Friendly Name for phone number"
  type        = string
  default     = ""
}

variable "sms_country_code" {
  description = "Country Code for phone number"
  type        = number
  default     = 1
}

variable "sms_phone_number" {
  description = "Phone number for text alerts"
  type        = number
  default     = 0
}

variable "enable_webhook" {
  description = "Enable Web Hook Alerts"
  type        = bool
  default     = false
}

variable "webhook_name" {
  description = "Friendly Name for web hook"
  type        = string
  default     = ""
}

variable "webhook_service_uri" {
  description = "The full URI for the webhook"
  type        = string
  default     = ""
}