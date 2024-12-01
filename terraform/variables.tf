variable "subscription_id" {
  description = "Id of the subscription."
  type        = string
  default     = "cfc6c002-a925-4c65-9fd1-841e01b031c8"
}

variable "resource_group_name" {
  description = "Name of the Resource Group."
  type        = string
  default     = "terraform-iac-rg"
}

variable "location" {
  description = "The Azure region to deploy resources into."
  type        = string
  default     = "West US"
}

variable "acr_name" {
  description = "Name of the Azure Container Registry."
  type        = string
  default     = "myacr12345"
}

variable "app_service_plan_name" {
  description = "Name of the App Service Plan."
  type        = string
  default     = "my-app-service-plan"
}

variable "app_service_name" {
  description = "Name of the App Service."
  type        = string
  default     = "my-app-service"
}
