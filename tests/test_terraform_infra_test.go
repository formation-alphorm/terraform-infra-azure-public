package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformInfrastructureVerification(t *testing.T) {
	// Configuration pour lire l'état Terraform existant
	terraformOptions := &terraform.Options{
		TerraformDir: "../terraform", // Chemin vers le répertoire Terraform
	}

	// Récupérer les outputs avec OutputRequiredE pour les chaînes brutes
	outputResourceGroup, err := terraform.OutputRequiredE(t, terraformOptions, "resource_group_name")
	if err != nil {
		t.Fatalf("Failed to get output 'resource_group_name': %v", err)
	}

	// Supprimer les guillemets autour des chaînes
	outputResourceGroup = strings.Trim(outputResourceGroup, "\"")
	t.Logf("Resource Group Name: %s", outputResourceGroup)
	assert.Equal(t, "terraform-iac-production-rg", outputResourceGroup)

	outputAcrName, err := terraform.OutputRequiredE(t, terraformOptions, "acr_name")
	if err != nil {
		t.Fatalf("Failed to get output 'acr_name': %v", err)
	}

	outputAcrName = strings.Trim(outputAcrName, "\"")
	t.Logf("ACR Name: %s", outputAcrName)
	assert.Equal(t, "iacterraformprojectproductionacr", outputAcrName)

	// Ajoutez les autres outputs similaires
	outputAppServicePlanName, err := terraform.OutputRequiredE(t, terraformOptions, "app_service_plan_name")
	if err != nil {
		t.Fatalf("Failed to get output 'app_service_plan_name': %v", err)
	}

	outputAppServicePlanName = strings.Trim(outputAppServicePlanName, "\"")
	t.Logf("App Service Plan Name: %s", outputAppServicePlanName)
	assert.Equal(t, "iac-production-app-service-plan", outputAppServicePlanName)

	outputAppServiceName, err := terraform.OutputRequiredE(t, terraformOptions, "app_service_name")
	if err != nil {
		t.Fatalf("Failed to get output 'app_service_name': %v", err)
	}

	outputAppServiceName = strings.Trim(outputAppServiceName, "\"")
	t.Logf("App Service Name: %s", outputAppServiceName)
	assert.Equal(t, "iacprojectproductionapp", outputAppServiceName)
}
