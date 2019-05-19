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
   throughput = "500"
 }
