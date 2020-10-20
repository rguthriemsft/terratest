---
layout: collection-browser-doc
title: Azure SDK Client Factory
category: community
excerpt: >-
  Overview of the client factory pattern to be used with the Azure SDK when creating test modules.
tags: ["contributing", "azure"]
order: 403
nav_title: Documentation
nav_title_link: /docs/
custom_js:
  - examples
  - prism
  - collection-browser_scroll
  - collection-browser_search
  - collection-browser_toc
---

# Azure SDK Client Factory

This documentation provides and overview of the `client_factory.go` module, targeted use cases, and behaviors.  This module is intended to provide support for and simplify working with Azure's multiple cloud environments (Azure Public, Azure Government, Azure China, Azure Germany and Azure Stack).  Developers looking to contribute to additional support for Azure to Terratest should leverage client_factory and use the patterns below to add a resource REST client from Azure Go SDK.  By doing so, it provides a consistent means for developers using Terratest to test their Azure Infrastructure to connect to the correct cloud and its associated REST apis.

## Background

The Azure REST APIs support both Public and sovereign cloud environments (at the moment this includes Public, US Government, Germany, China, and Azure Stack environments).  If you are interacting with an environment other than public cloud, you need to set the base URI for the Azure REST API you are interacting with.

### Base URI

You must use the correct base URI's for the Azure REST API's (either directly or via Azure SDK for GO) to communicate with a cloud environment other than Azure Public. The Azure Go SDK supports this by using the `WithBaseURI` suffixed calls when creating service clients. For example, when using the `VirtualMachinesClient` with the public cloud, a developer would normally write code for the public cloud like so:

```golang
import (
    "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
)

func SomeVMHelperMethod() {
    subscriptionID := "your subscription ID"

    // Create a VM client and return
    vmClient, err := compute.NewVirtualMachinesClient(subscriptionID)

    // Use client / etc
}
```

However, this code will not work in non-Public cloud environments as the REST endpoints have different URIs depending on environment.  Instead, you need to use an alternative method (provided in the Azure REST SDK for Go) to get a properly configured client (*all REST API clients should support this alternate method*):

```golang
import (
    "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
)

func SomeVMHelperMethod() {
    subscriptionID := "your subscription ID"
    baseURI := "management.azure.com"

    // Create a VM client and return
    vmClient, err := compute.NewVirtualMachinesClientWithBaseURI(baseURI, subscriptionID)

    // Use client / etc
}
```

Using code similar to above, you can communicate with any Azure cloud environment just by changing the base URI that is passed to the clients (Azure Public shown in above example).

## Lookup Environment Metadata

Developers MUST avoid hardcoding these base URI's.  Instead, they should be looked up from an authoritative source. The AutoRest-GO library (used by the Go SDK) provides such functionality. The `client_factory` module makes use of the AutoRest `EnvironmentFromName(envName string)` function to return the appropriate structure.  This method and Environment structure is documented on GoDoc [here](https://godoc.org/github.com/Azure/go-autorest/autorest/azure#EnvironmentFromName).

To configure different cloud environments, we will use the same `AZURE_ENVIRONMENT` environment variable that the Go SDK uses. This can currently be set to one of the following values:

|Value                      |Cloud Environment  |
|---------------------------|-------------------|
|"AzureChinaCloud"          |ChinaCloud         |
|"AzureGermanCloud"         |GermanCloud        |
|"AzurePublicCloud"         |PublicCloud        |
|"AzureUSGovernmentCloud"   |USGovernmentCloud  |
|"AzureStackCloud"          |Azure stack        |

When using the "AzureStackCloud" setting, you MUST also set the `AZURE_ENVIRONMENT_FILEPATH` variable to point to a JSON file containing your Azure Stack URI details.

## Putting it all together

 `client_factory` implements this pattern described above in order to instantiate and return properly configured *REST SDK for GO* clients so that test implementers don't have to consider REST API client implementation as long as they have the correct `AZURE_ENVIRONMENT` env setting.  If this environment variable is not set, the client will assume public cloud as the cloud environment to communicate with.  We strongly recommend developers creating Terratest helper methods for Azure use this pattern with client factory to create REST API clients.  This will reduce effort for Terratest users creating test for Azure resources.

>NOTE: TERRAFORM uses [ARM_ENVIRONMENT](https://www.terraform.io/docs/backends/types/azurerm.html#environment) environment variable to set the correct cloud environment.  
<!-- -->
>NOTE: The default behavior of the `client_factory` is to use the AzurePublicCloud environment. This requires no work from the developer to configure, and ensures consistent behavior with the current SDK code.

### Wait, I don't see the client in client factory for the rest api I want to interact with

 If you require a client that is not already implemented in client factory for your helper method you will need to create a corresponding method that instantiates the client and accepts base URI following the patterns discussed.  Below is a walkthrough for adding a client to client factory.

## Walkthrough, adding a client to client_factory

### Add your client namespace to client factory

In the Azure SDK for GO, each service should have a module that implements that services client.  You can find the correct module [here](https://godoc.org/github.com/Azure/azure-sdk-for-go).  Add that module to the client factory imports.  Below is an example for client imports that shows clients for compute, container service and subscriptions.

{% include examples/explorer.html example_id='client-factory' file_id='client_factory_code' class='wide quick-start-examples' skip_learn_more=true skip_view_on_github=true skip_tags=true range_id='client_factory_example.imports' %}

### Add your client method to instantiate the client

The next step is to add your method to instantiate the client.  Below is an example of adding the method to create a client for Virtual Machines, note that we lookup the environment using `getEnvironmentEndpointE` and then pass that base URI to the actual method on the Virtual Machines Module to create the client `NewVirtualMachinesClientWithBaseURI`.

{% include examples/explorer.html example_id='client-factory' file_id='client_factory_code' class='wide quick-start-examples' skip_learn_more=true skip_view_on_github=true skip_tags=true range_id='client_factory_example.CreateClient' %}

### Add a unit test to client_factory_test.go

In order to ensure that your CreateClient method works properly, add a unit test to `client_factory_test.go`.  The unit test MUST assert that the base URI is correctly set for your client.  Some key points for writing your unit test are:

- Use table-driven testing to test the various combinations of cloud environments
- Give the test case a descriptive name so it is easy to identify which test failed.
- PRs will be rejected if a client is added without a corresponding unit test.

Below is an example of the Virtual Machines client unit test:

{% include examples/explorer.html example_id='client-factory' file_id='client_factory_test' class='wide quick-start-examples' skip_learn_more=true skip_view_on_github=true skip_tags=true range_id='client_factory_example.UnitTest' %}

### Use your CreateClient method in your helper

We now can use this client creation method in our helpers to create a Virtual Machines client.  Below is an example for how to call into this create method from `client_factory`:

{% include examples/explorer.html example_id='client-factory' file_id='client_factory_helper' class='wide quick-start-examples' skip_learn_more=true skip_view_on_github=true skip_tags=true range_id='client_factory_example.helper' %}
