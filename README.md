# terraform_azcli_provider

- [terraform_azcli_provider](#terraformazcliprovider)
  - [Synopsis](#synopsis)
  - [Dev Stuff](#dev-stuff)
    - [Required Go Modules](#required-go-modules)
    - [Useful Go Commands](#useful-go-commands)
  - [User Stuff](#user-stuff)
    - [Authenticating to Azure](#authenticating-to-azure)
      - [Provider Example Usage](#provider-example-usage)
    - [Resources](#resources)
      - [azcli_cosmos_database](#azclicosmosdatabase)
        - [Database Example Usage](#database-example-usage)
      - [azcli_cosmos_collection](#azclicosmoscollection)
        - [Collection Example Usage](#collection-example-usage)
    - [Full Example](#full-example)

## Synopsis

`Before using this provider, consider using the native azurerm provider first`

This provider implements a lightweight wrapper around Microsoft's AZ CLI. Its primary purpose is to fill in  gaps where native terraform module lacks support. For example, creating CosmosDB collections.

One benefit this provider implements is auto import, this means if your resources already exist the provider will simply generate the state as needed without a requirement for destroying\recreating or indeed importing.

## Dev Stuff

### Required Go Modules

In order to build this provider you will require these modules

go get -u github.com/tidwall/gjson</br>
go get -u github.com/hashicorp/terraform/helper/schema

### Useful Go Commands

go env

## User Stuff

### Authenticating  to Azure

This provider expects the AZ cli to be logged in before use, the only parameter supported is subscription name

#### Provider Example Usage

```hcl
provider "azcli" {
  subscription_name = "my-azure-subscription"
  version = "~> 0.0"
}
```

### Resources

#### azcli_cosmos_database

Manages Cosmos account databases.

`Note`: Official Terraform azurerm provider now supports creating databases.

##### Database Example Usage

-----

```hcl
locals {
  resource_group_name = "abxrg0008-test"
  cosmos_account_name = "abxtf"

}

resource "azcli_cosmos_database" "default" {
  cosmos_account_name = "${local.cosmos_account_name}"
  resource_group_name = "${local.resource_group_name}"
  name                = "testdatabase"
}
```

#### azcli_cosmos_collection

Manages Cosmos database collections.

##### Collection Example Usage

-----

```hcl
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

### Full Example

```hcl
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