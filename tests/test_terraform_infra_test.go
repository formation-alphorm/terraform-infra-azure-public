package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformInfrastructureVerification(t *testing.T) {
	// Configuration pour lire l'état Terraform existant
	terraformOptions := &terraform.Options{
		TerraformDir: "../terraform", // Chemin vers le répertoire Terraform
	}

	// Récupérer les outputs
	outputResourceGroup, err := terraform.OutputE(t, terraformOptions, "resource_group_name")
	if err != nil {
		t.Fatalf("Failed to get output 'resource_group_name': %v", err)
	}
	t.Logf("Resource Group Name: %s", outputResourceGroup)
	assert.Equal(t, "terraform-iac-test-rg", outputResourceGroup)

	outputAcrName, err := terraform.OutputE(t, terraformOptions, "acr_name")
	if err != nil {
		t.Fatalf("Failed to get output 'acr_name': %v", err)
	}
	t.Logf("ACR Name: %s", outputAcrName)
	assert.Equal(t, "iacterraformprojecttestacr", outputAcrName)

	outputAppServicePlanName, err := terraform.OutputE(t, terraformOptions, "app_service_plan_name")
	if err != nil {
		t.Fatalf("Failed to get output 'app_service_plan_name': %v", err)
	}
	t.Logf("App Service Plan Name: %s", outputAppServicePlanName)
	assert.Equal(t, "iac-test-app-service-plan", outputAppServicePlanName)

	outputAppServiceName, err := terraform.OutputE(t, terraformOptions, "app_service_name")
	if err != nil {
		t.Fatalf("Failed to get output 'app_service_name': %v", err)
	}
	t.Logf("App Service Name: %s", outputAppServiceName)
	assert.Equal(t, "iacprojecttestapp", outputAppServiceName)
}
