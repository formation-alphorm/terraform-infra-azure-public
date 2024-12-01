terraform {
  backend "azurerm" {
    resource_group_name  = "tfstate-backend-rg"
    storage_account_name = "tfstatebackendanass"
    container_name       = "terraform-tfstate"
    key                  = "terraform.tfstate"
  }
}
