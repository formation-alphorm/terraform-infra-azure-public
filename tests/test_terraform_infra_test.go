package test

import (
	"os"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var resourceGroupName string
var acrName string
var location string
var appServicePlanName string
var appServiceName string

// Fonction pour analyser les arguments passés via -args
func parseArgs() {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "resource_group_name=") {
			resourceGroupName = strings.TrimPrefix(arg, "resource_group_name=")
		} else if strings.HasPrefix(arg, "acr_name=") {
			acrName = strings.TrimPrefix(arg, "acr_name=")
		} else if strings.HasPrefix(arg, "location=") {
			location = strings.TrimPrefix(arg, "location=")
		} else if strings.HasPrefix(arg, "app_service_plan_name=") {
			appServicePlanName = strings.TrimPrefix(arg, "app_service_plan_name=")
		} else if strings.HasPrefix(arg, "app_service_name=") {
			appServiceName = strings.TrimPrefix(arg, "app_service_name=")
		}
	}
}

func TestTerraformInfrastructure(t *testing.T) {
	// Analyse des arguments
	parseArgs()

	// Logs pour déboguer les arguments
	t.Logf("Parsed Arguments: resource_group_name=%s, acr_name=%s, location=%s, app_service_plan_name=%s, app_service_name=%s",
		resourceGroupName, acrName, location, appServicePlanName, appServiceName)

	// Vérifiez que tous les arguments requis sont présents
	if resourceGroupName == "" || acrName == "" || location == "" || appServicePlanName == "" || appServiceName == "" {
		t.Fatalf("Missing required arguments: resource_group_name=%s, acr_name=%s, location=%s, app_service_plan_name=%s, app_service_name=%s",
			resourceGroupName, acrName, location, appServicePlanName, appServiceName)
	}

	// Configuration Terraform
	terraformOptions := &terraform.Options{
		TerraformDir: "../terraform",
		Vars: map[string]interface{}{
			"resource_group_name":    resourceGroupName,
			"acr_name":               acrName,
			"location":               location,
			"app_service_plan_name":  appServicePlanName,
			"app_service_name":       appServiceName,
		},
	}

	// Détruire les ressources après le test
	defer terraform.Destroy(t, terraformOptions)

	// Initialiser et appliquer Terraform
	t.Log("Initializing and applying Terraform")
	terraform.InitAndApply(t, terraformOptions)

	// Valider les résultats
	outputResourceGroup := terraform.Output(t, terraformOptions, "resource_group_name")
	assert.Equal(t, resourceGroupName, outputResourceGroup)

	outputAcrName := terraform.Output(t, terraformOptions, "acr_name")
	assert.Equal(t, acrName, outputAcrName)

	outputAppServicePlanName := terraform.Output(t, terraformOptions, "app_service_plan_name")
	assert.Equal(t, appServicePlanName, outputAppServicePlanName)

	outputAppServiceName := terraform.Output(t, terraformOptions, "app_service_name")
	assert.Equal(t, appServiceName, outputAppServiceName)
}
