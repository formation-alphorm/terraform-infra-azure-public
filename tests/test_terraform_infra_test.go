package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformInfrastructureVerification(t *testing.T) {
	// Configuration pour lire l'état existant
	terraformOptions := &terraform.Options{
		TerraformDir: "../terraform", // Chemin du répertoire Terraform utilisé par le pipeline
	}

	// Lire les sorties Terraform sans appliquer
	outputResourceGroup := terraform.Output(t, terraformOptions, "resource_group_name")
	assert.Equal(t, "terraform-iac-test-rg", outputResourceGroup, "Resource group name mismatch")

	outputAcrName := terraform.Output(t, terraformOptions, "acr_name")
	assert.Equal(t, "iacterraformprojecttestacr", outputAcrName, "ACR name mismatch")

	outputAppServicePlanName := terraform.Output(t, terraformOptions, "app_service_plan_name")
	assert.Equal(t, "iac-test-app-service-plan", outputAppServicePlanName, "App Service Plan name mismatch")

	outputAppServiceName := terraform.Output(t, terraformOptions, "app_service_name")
	assert.Equal(t, "iacprojecttestapp", outputAppServiceName, "App Service name mismatch")
}
