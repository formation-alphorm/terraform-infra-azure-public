package test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformInfrastructureVerification(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../terraform", // Chemin vers le répertoire Terraform
	}

	// Fonction pour gérer les outputs brut ou JSON
	getTerraformOutput := func(output string) (string, error) {
		rawOutput, err := terraform.OutputRequiredE(t, terraformOptions, output)
		if err != nil {
			return "", err
		}
		// Tente de parser comme JSON, sinon retourne tel quel
		var parsedOutput string
		err = json.Unmarshal([]byte(rawOutput), &parsedOutput)
		if err != nil {
			// Si parsing échoue, c'est probablement une chaîne brute
			parsedOutput = strings.Trim(rawOutput, "\"")
		}
		return parsedOutput, nil
	}

	// Récupérer et valider les outputs
	outputResourceGroup, err := getTerraformOutput("resource_group_name")
	if err != nil {
		t.Fatalf("Failed to get output 'resource_group_name': %v", err)
	}
	t.Logf("Resource Group Name: %s", outputResourceGroup)
	assert.Equal(t, "terraform-iac-production-rg", outputResourceGroup)

	outputAcrName, err := getTerraformOutput("acr_name")
	if err != nil {
		t.Fatalf("Failed to get output 'acr_name': %v", err)
	}
	t.Logf("ACR Name: %s", outputAcrName)
	assert.Equal(t, "iacterraformprojectproductionacr", outputAcrName)

	outputAppServicePlanName, err := getTerraformOutput("app_service_plan_name")
	if err != nil {
		t.Fatalf("Failed to get output 'app_service_plan_name': %v", err)
	}
	t.Logf("App Service Plan Name: %s", outputAppServicePlanName)
	assert.Equal(t, "iac-production-app-service-plan", outputAppServicePlanName)

	outputAppServiceName, err := getTerraformOutput("app_service_name")
	if err != nil {
		t.Fatalf("Failed to get output 'app_service_name': %v", err)
	}
	t.Logf("App Service Name: %s", outputAppServiceName)
	assert.Equal(t, "iacprojectproductionapp", outputAppServiceName)
}
