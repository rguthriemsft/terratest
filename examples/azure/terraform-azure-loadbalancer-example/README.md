# Terraform Load Balancer Example

This folder contains a simple Terraform module that deploys resources in [Azure](https://azure.microsoft.com/) to demonstrate
how you can use Terratest to write automated tests for your Azure Terraform code. This module deploys two Load Balancers for Public and Private IP scenarios.

- A Public [Load Balancer](https://docs.microsoft.com/azure/load-balancer/) that gives the module the following:

  - `Load Balancer` with the name specified in the `lb_public_name` and configuration in the `lb_public_fe_config_name` output variables.

- A Private [Load Balancer](https://docs.microsoft.com/azure/load-balancer/) that gives the module the following:

  - `Load Balancer` with the name specified in the `lb_private_name` and static Frontend IP Configuration in the `lb_private_fe_config_static_name` output variables.

- A Default [Load Balancer](https://docs.microsoft.com/azure/load-balancer/) that gives the module the following:

  - `Load Balancer` with the name specified in the `lb_default_name` and no Frontend configuration.

- Networking that provides the following for the Load Balancer module with the following:
  - [Virtual Network](https://docs.microsoft.com/azure/virtual-network/) with the name specified in the `vnet_name` output variable.
  - [Subnet](https://docs.microsoft.com/azure/virtual-network/virtual-network-manage-subnet) with the name specified in the `subnet_name` output variable.
  - [Public IP Address](https://docs.microsoft.com/azure/virtual-network/virtual-network-public-ip-address) with the name specified in the `public_address_name` output variable.

Check out [test/azure/terraform_azure_loadbalancer_example_test.go](/test/azure/terraform_azure_loadbalancer_example_test.go) to see how you can write
automated tests for this module.

Note that the Load Balancers and their associated resources in this module don't actually do anything; they are created before running the tests, for demonstration purposes.

**WARNING**: This module and the automated tests for it deploy real resources into your Azure account which can cost you money. The resources are all part of the [Azure Free Account](https://azure.microsoft.com/free/), so if you haven't used that up, it should be free, but you are completely responsible for all Azure charges.

## Running this module manually

1. Sign up for [Azure](https://azure.microsoft.com/)
1. Configure your Azure credentials using one of the [supported methods for Azure CLI
   tools](https://docs.microsoft.com/cli/azure/azure-cli-configuration?view=azure-cli-latest)
1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`
1. Ensure [environment variables](../README.md#review-environment-variables) are available
1. Run `terraform init`
1. Run `terraform apply`
1. When you're done, run `terraform destroy`

## Running automated tests against this module

1. Sign up for [Azure](https://azure.microsoft.com/)
1. Configure your Azure credentials using one of the [supported methods for Azure CLI
   tools](https://docs.microsoft.com/cli/azure/azure-cli-configuration?view=azure-cli-latest)
1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`
1. Configure your Terratest [Go test environment](../README.md)
1. `cd test/azure`
1. `go build terraform_azure_loadbalancer_example_test.go`
1. `go test -v -run TestTerraformAzureLoadBalancerExample`
