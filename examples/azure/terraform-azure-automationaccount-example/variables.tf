# ---------------------------------------------------------------------------------------------------------------------
# ENVIRONMENT VARIABLES
# Define these secrets as environment variables
# Note that the TF_VAR_AUTOMATION_ACCOUNT_CLIENT_ID and TF_VAR_AUTOMATION_ACCOUNT_CLIENT_PASSWORD 
# variables represent a service principal in AAD
# In addition to created a service prinicpal for the automation run as connection, a certificate 
# must be uploaded to the service principal as a secret.
# The same certificate must also be uplaoded into the Azure Automation Account via Terraform 
# ---------------------------------------------------------------------------------------------------------------------

# ARM_CLIENT_ID
# ARM_CLIENT_SECRET
# ARM_SUBSCRIPTION_ID
# ARM_TENANT_ID
# TF_VAR_ARM_SUBSCRIPTION_ID
# TF_VAR_ARM_TENANT_ID
# TF_VAR_AUTOMATION_ACCOUNT_CLIENT_ID
# TF_VAR_AUTOMATION_RUN_AS_CERTIFICATE_THUMBPRINT

variable ARM_SUBSCRIPTION_ID {}
variable ARM_CLIENT_ID {}
variable ARM_CLIENT_SECRET {}
variable ARM_TENANT_ID {}
variable AUTOMATION_ACCOUNT_CLIENT_ID {}
variable AUTOMATION_RUN_AS_CERTIFICATE_THUMBPRINT {}
# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ---------------------------------------------------------------------------------------------------------------------

# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------

variable "postfix" {
  description = "Random postfix string used for each test run; set from the test file at runtime."
  type        = string
  default     = "qwefgt"
}

variable "resource_group_name" {
  description = "Name for the resource group holding resources for this example"
  type        = string
  default     = "terratest-automationaccount-rg"
}

variable "location" {
  description = "The Azure region in which to deploy this sample"
  type        = string
  default     = "East US"
}

variable "cloud_environment" {
  description = "The Azure cloud where the command is executed"
  type        = string
  default     = "AzureCloud"
}

variable "automation_account_name" {
  description = "The name of the automation account that will be created in the resource group"
  type        = string
  default     = "terratest-AutomationAccount"
}

variable "automation_run_as_connection_name" {
  description = "The name of the automation run as connection that will be created in the resource group"
  type        = string
  default     = "terratest-AutomationRunAsConnectionName"
}

variable "automation_run_as_certificate_name" {
  description = "The name of the automation account run as connection certificate name"
  type        = string
  default     = "terratest-AutomationConnectionCertificateName"
}

variable "automation_run_as_certificate_path" {
  description = "The path to the automation Run As certificate .pfx file"
  type        = string
  default     = "./certificate/SPRunAsCert.pfx"
}

variable "automation_run_as_connection_type" {
  description = "The name of the automation account run as connection type"
  type        = string
  default     = "AzureServicePrincipal"
}

variable "sample_dsc_name" {
  description = "The name of the sample DSC configuration that contains the configuraitons that can be applied"
  type        = string
  default     = "SampleDSC"
}

variable "sample_dsc_path" {
  description = "The path to the sample dsc file in the repo"
  type        = string
  default     = "./dsc/SampleDSC.ps1"
}

variable "sample_dsc_configuration_name" {
  description = "The name of the DSC configuraiton to apply to the VM"
  type        = string
  default     = "SampleDSC.NotWebServer"
}

variable "vm_name" {
  description = "The name of the test VM where the DSC configuraiton will be applied"
  type        = string
  default     = "vm01"
}

