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

variable "domain_name_label" {
  description = "The Domain Name Label for the Public IP Address"
  type        = string
}

# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------

variable "dns_ip_01" {
  description = "The first DNS Server IP for the Virtual Network"
  type        = string
  default     = "10.0.0.5"
}

variable "dns_ip_02" {
  description = "The second DNS Server IP for the Virtual Network"
  type        = string
  default     = "10.0.0.6"
}

variable "location" {
  description = "The Azure Region to deploy resources too"
  type        = string
  default     = "East US"
}

variable "postfix" {
  description = "The postfix that will be attached to all resources deployed"
  type        = string
  default     = "resource"
}

variable "private_ip" {
  description = "The Static Private IP for the Internal NIC"
  type        = string
  default     = "10.0.20.5"
}

variable "subnet_prefix" {
  description = "The subnet range of IPs for the Virtual Network"
  type        = string
  default     = "10.0.20.0/24"
}
