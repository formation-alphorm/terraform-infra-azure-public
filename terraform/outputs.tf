output "acr_login_server" {
  value = azurerm_container_registry.acr.login_server
}

output "app_service_url" {
  value = azurerm_app_service.app.default_site_hostname
}

output "resource_group_name" {
  value = var.resource_group_name
}

output "acr_name" {
  value = azurerm_container_registry.acr.name
}

output "app_service_plan_name" {
  value = azurerm_app_service_plan.asp.name
}

output "app_service_name" {
  value = azurerm_app_service.app.name
}
