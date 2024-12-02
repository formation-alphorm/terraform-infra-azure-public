package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformInfrastructureVerification(t *testing.T) {
	// Chemin vers le répertoire Terraform
	terraformOptions := &terraform.Options{
		TerraformDir: "../terraform",
	}

	// Fonction pour gérer les outputs (JSON ou chaînes brutes)
	getOutput := func(outputKey string) string {
		rawOutput := terraform.Output(t, terraformOptions, outputKey)
		// Nettoyage des guillemets s'il s'agit d'une chaîne brute
		return strings.Trim(rawOutput, "\"")
	}

	// Vérifiez chaque output et affichez les logs
	outputResourceGroup := getOutput("resource_group_name")
	t.Logf("Resource Group Name: %s", outputResourceGroup)
	assert.Equal(t, "terraform-iac-test-rg", outputResourceGroup)

	outputAcrName := getOutput("acr_name")
	t.Logf("ACR Name: %s", outputAcrName)
	assert.Equal(t, "iacterraformprojecttestacr", outputAcrName)

	outputAppServicePlanName := getOutput("app_service_plan_name")
	t.Logf("App Service Plan Name: %s", outputAppServicePlanName)
	assert.Equal(t, "iac-test-app-service-plan", outputAppServicePlanName)

	outputAppServiceName := getOutput("app_service_name")
	t.Logf("App Service Name: %s", outputAppServiceName)
	assert.Equal(t, "iacprojecttestapp", outputAppServiceName)
}
