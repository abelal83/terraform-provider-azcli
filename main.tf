provider "azcli" {
  subscription_name = "Clarksons - Infrastructure - R&D"
  version = "~> 0.0"
}


locals {
  resource_group_name = "abxrg0008-test"
  cosmos_account_name = "abxtf"

}

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
   throughput = "500"

 }

  resource "azcli_cosmos_collection" "partition" {
   cosmos_account_name = "${local.cosmos_account_name}"
   resource_group_name = "${local.resource_group_name}"
   database_name       = "${azcli_cosmos_database.default.id}"
   name                = "partition"
   partition_key       = "/abu/belal/test"
   throughput = "400"
 }

#  resource "azcli_cosmos_collection" "indexing_policy" {
#    cosmos_account_name = "abx"
#    resource_group_name = "terraform-provider"
#    database_name       = "${azcli_cosmos_database.default.id}"
#    name                = "indexing"
#    throughput = "400"
#    indexing_policy = "${file("indexing_policy.json")}"
#  }