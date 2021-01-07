# Terraform Azure Automation Account Example

This folder contains a simple Terraform module that deploys resources in [Azure](https://azure.microsoft.com/) to demonstrate how you can use Terratest to write automated tests for your Azure Terraform code. This module deploys the following:

- An [Automation Account](https://azure.microsoft.com/en-us/services/automation/) that provides the module the following:
  - `Automation Account` with the name specified in the `automation_account_name` variable.
    - `Automation Account Connection Run As Account` with the name specified in the `automation_run_as_connection_name` variable.  See the section titled *_Example Service Principal and Certificate Setup_* below on how to configure the Automation Account RunAs Account credentials.
    - `Automation Account Connection Run As Certificate` with the name specified in the `automation_run_as_certificate_name` variable and the thumbprint in the `TF_VAR_AUTOMATION_RUN_AS_CERTIFICATE_THUMBPRINT` variable.
    - `Automation Account Connection Type` with type specified in the `automation_run_as_connection_type` variable.
    - [Desired State Configuration](https://docs.microsoft.com/en-us/powershell/scripting/dsc/getting-started/winGettingStarted?view=powershell-7#:~:text=Get%20started%20with%20Desired%20State%20Configuration%20%28DSC%29%20for,Windows%20PowerShell%20Desired%20State%20Configuration%20log%20files.%20) with the name specified in the `sample_dsc_name` variable and the path to the DSC specified in the `sample_dsc_path` variable.
  - [Virtual Machine](https://docs.microsoft.com/en-us/azure/virtual-machines/) with the name specified in the `vm_name` output variable.
    - [Virtual Machine Extension](https://docs.microsoft.com/en-us/azure/virtual-machines/extensions/overview#:~:text=Troubleshoot%20extensions%20%20%20%20Namespace%20%20,Encryption%20for%20Windows%20%2012%20more%20rows%20), configured for DSC, with the `NodeConigurationName` set in the `sample_dsc_configuration_name` variable.
    - The VM includes a virtual network, a subnet, and a network interface with hard-coded configuration as it is not pertinent to this example.

Check out [test/azure/terraform_azure_automationaccount_example_test.go](/test/azure/terraform_azure_automationaccount_example_test.go) to see how you can write
automated tests for this module.

Note that the resources deployed in this module don't actually do anything; it just runs the resources for demonstration purposes.

**WARNING**: This module and the automated tests for it deploy real resources into your Azure account which can cost you
money. The resources are all part of the [Azure Free Account](https://azure.microsoft.com/en-us/free/), so if you haven't used that up,
it should be free, but you are completely responsible for all Azure charges.

## Example Service Principal and Certificate Setup

To run this example, you must create a service principal in Azure Active Directory as well as create a non-password protected self-signed certificate in .pfx format that you will need to upload into the Automation Account Run As service principal as a secret for test purposes.

The same certificate file will need to be placed in the `/examples/azure/terraform-azure-automationaccount-example/certificate/` folder with the name `runascert.pfx` so that the certificate can be uploaded into the Automation Account in order to successfully configure the Automation Account RunAs account and connection.

The documentation link [Manage an Azure Automation Run As account](https://docs.microsoft.com/en-us/azure/automation/manage-runas-account#:~:text=1%20Go%20to%20your%20Automation%20account%20and%20select,locate%20the%20role%20definition%20that%20is%20being%20used.) has additional background on the configuration requirements.  

For the example, note that the `TF_VAR_AUTOMATION_ACCOUNT_CLIENT_ID`, and the `TF_VAR_AUTOMATION_RUN_AS_CERTIFICATE_THUMBPRINT`, environment variables must be configured with the corresponding service principal values. For the Automation Account Run As conneciton certificate, place the self-signed .pfx certificate  into the `certificate` folder per above.  Also set the certificate thumbprint in the `TF_VAR_RUNAS_CERTIFICATE_THUMBPRINT` variable.

For the example to suceed, and in general when uploading a DSC to an Automation Account, you will need to kick off compilation of the DSC in the Automation Account prior to applying the DSC to a VM node, else it will fail to apply.  You can use PowerShell Core to compile the DSC in Terraform via the `null_resource` resource but you first must sign-in to Azure from PowerShell Core. The `TF_VAR_POWERSHELL_CLIENT_ID` and `TF_VAR_POWERSHELL_CLIENT_SECRET` environment variables must be configured with a service principal to sign-in to Azure from PowerShell core and compile the DSC.

*_Note: In a production system, you would store the service principal configuraiton in and create the certificate using Azure KeyVault and then configure a Terraform `azurerm_key_vault_secret` data source on the KeyVault instance to access the data securely directly from Terraform._*

## Running this module manually

1. Sign up for [Azure](https://azure.microsoft.com/)
1. Configure your Azure credentials using one of the [supported methods for Azure CLI
   tools](https://docs.microsoft.com/en-us/cli/azure/azure-cli-configuration?view=azure-cli-latest)
1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`
1. Ensure [environment variables](../README.md#review-environment-variables) are available
1. Run `terraform init`
1. Run `terraform apply`
1. When you're done, run `terraform destroy`

## Running automated tests against this module

1. Sign up for [Azure](https://azure.microsoft.com/)
1. Configure your Azure credentials using one of the [supported methods for Azure CLI
   tools](https://docs.microsoft.com/en-us/cli/azure/azure-cli-configuration?view=azure-cli-latest)
1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`
1. Configure your Terratest [Go test environment](../README.md)
1. `cd test/azure`
1. `go build terraform_azure_automationaccount_example_test.go`
1. `go test -v -run TestTerraformAzureAutomationAccountExample`
