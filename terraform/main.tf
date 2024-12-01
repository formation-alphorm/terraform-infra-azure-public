data "azurerm_resource_group" "main" {
  name = var.resource_group_name
}

resource "azurerm_container_registry" "acr" {
  name                = var.acr_name
  resource_group_name = var.resource_group_name
  location            = avar.location
  sku                 = "Basic"
  admin_enabled       = true
}

resource "azurerm_app_service_plan" "asp" {
  name                = var.app_service_plan_name
  location            = avar.location
  resource_group_name = var.resource_group_name
  sku {
    tier = "Basic"
    size = "B1"
  }
}

resource "azurerm_app_service" "app" {
  name                = var.app_service_name
  location            = avar.location
  resource_group_name = var.resource_group_name
  app_service_plan_id = azurerm_app_service_plan.asp.id

  site_config {
    always_on = true
  }

  app_settings = {
    "DOCKER_REGISTRY_SERVER_URL"      = azurerm_container_registry.acr.login_server
    "DOCKER_REGISTRY_SERVER_USERNAME" = azurerm_container_registry.acr.admin_username
    "DOCKER_REGISTRY_SERVER_PASSWORD" = azurerm_container_registry.acr.admin_password
  }
}