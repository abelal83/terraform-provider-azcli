
provider "cosmos" {
  cosmos_account_name = "hello"
}

resource "cosmos_database" "default" {
  name = "testDatabase"
}
