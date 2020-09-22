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

variable "resource_group_name" {
  description = "The resource group where all resources will be deployed"
  type        = string
  default     = "terratest-keyvault-rg"
}

variable "loadbalancer01_name" {
  description = "Load balancer with a public IP"
  type        = string
  default     = "terratest-loadbalancer-lb-01"
}

variable "loadbalancer02_name" {
  description = "Load balancer with a private IP"
  type        = string
  default     = "terratest-loadbalancer-lb-02"
}

variable "vnet_name" {
  description = "Virtual Network for Load Balancer 02"
  type        = string
  default     = "terratest-loadbalancer-vnet"
}

variable "lb01_feconfig" {
  description = "Frontend Config for Load Balancer 01"
  type        = string
  default     = "terratest-loadbalancer-cfg-01"
}

variable "pip_forlb01" {
  description = "Public IP for Load Balancer 01"
  type        = string
  default     = "terratest-loadbalancer-pip-01"
}

variable "feIPConfig_forlb02" {
  description = "Frontend Config for Load Balancer 02"
  type        = string
  default     = "terratest-loadbalancer-cfg-02"
}

variable "feSubnet_forlb02" {
  description = "Frontend Subnet for Load Balancer 02"
  type        = string
  default     = "terratest-loadbalancer-snt-02"
}