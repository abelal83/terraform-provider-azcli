provider "azcli" {
  subscription_name = "Clarksons - Development - R&D"
  version = "=0.0.10"
}


locals {
  resource_group_name = "ctrrg000008"

}

resource "azcli_functionapp_slot" "slot1" {
  slot_name = "slot1"
  resource_group_name = "ctrrg000008"
  function_app_name = "mrdtestapp"
  http_20_enabled = "true"
  always_on = "true"
}

resource "azcli_functionapp_slot" "slot2" {
  slot_name = "slot2"
  resource_group_name = "ctrrg000008"
  function_app_name = "mrdtestapp"
  http_20_enabled = "true"
  always_on = "true"
}

resource "azcli_functionapp_slot" "slot3" {
  slot_name = "slot3"
  resource_group_name = "ctrrg000008"
  function_app_name = "mrdtestapp"
  http_20_enabled = "true"
  always_on = "true"
}

resource "azcli_functionapp_slot" "slot4" {
  slot_name = "slot4"
  resource_group_name = "ctrrg000008"
  function_app_name = "mrdtestapp"
  http_20_enabled = "true"
  always_on = "true"
}

output "id" {
  value = azcli_functionapp_slot.slot1.id
}

output "identity" {
  value = azcli_functionapp_slot.slot1.identity
}
output "identity2" {
  value = azcli_functionapp_slot.slot3.identity
}
output "identity4" {
  value = azcli_functionapp_slot.slot4.identity
}