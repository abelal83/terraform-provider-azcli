provider "azcli" {
  subscription_name = "hello"
  version = "~> 0.0"
}

provider "azurerm" {
  
}

resource "azcli_cosmos_database" "default" {
  cosmos_account_name = "abx"
  resource_group_name = "terraform-provider"
  name                = "testdatabase"
}

 resource "azcli_cosmos_collection" "default" {
   cosmos_account_name = "abx"
   resource_group_name = "terraform-provider"
   database_name       = "${azcli_cosmos_database.default.id}"
   name                = "mycollection"
   throughput = "400"
 }

  resource "azcli_cosmos_collection" "partition" {
   cosmos_account_name = "abx"
   resource_group_name = "terraform-provider"
   database_name       = "${azcli_cosmos_database.default.id}"
   name                = "partition"
   partition_key       = "/abu/belal"
   throughput = "400"
 }

 resource "azcli_cosmos_collection" "indexing_policy" {
   cosmos_account_name = "abx"
   resource_group_name = "terraform-provider"
   database_name       = "${azcli_cosmos_database.default.id}"
   name                = "indexing"
   throughput = "400"
   indexing_policy = "${file("indexing_policy.json")}"
 }