# terraform_azcli_provider

- [terraform_azcli_provider](#terraformazcliprovider)
  - [required go modules](#required-go-modules)
  - [useful commands](#useful-commands)
  - [example configuration](#example-configuration)

This provider implements a lightweight wrapper around az cli.
Reasoning behind this provider is due to the native azurerm provider being well behind in terms or feature parity with az cli.

## required go modules

go get -u github.com/tidwall/gjson
go get -u github.com/hashicorp/terraform/helper/schema

## useful commands

go env

## example configuration

```json
provider "azcli" {
  subscription_name = "azure-abu-test"
  version = "~> 0.0"
}


locals {
  resource_group_name = "abxrg0008-test"
  cosmos_account_name = "abxtf"

}

# the official azurerm provider now supports creating database, you may chose to use that instead but including here for reference
resource "azcli_cosmos_database" "default" {
  cosmos_account_name = "${local.cosmos_account_name}"
  resource_group_name = "${local.resource_group_name}"
  name                = "testdatabase"
}

 resource "azcli_cosmos_collection" "default" {
   cosmos_account_name = "${local.cosmos_account_name}"
   resource_group_name = "${local.resource_group_name}"
   database_name       = "${azcli_cosmos_database.default.id}"
   name                = "mycollection"
   throughput = "400"
 }

  resource "azcli_cosmos_collection" "partition" {
   cosmos_account_name = "${local.cosmos_account_name}"
   resource_group_name = "${local.resource_group_name}"
   database_name       = "${azcli_cosmos_database.default.id}"
   name                = "partition"
   partition_key       = "/abu/belal/test"
   throughput = "400"
 }
```