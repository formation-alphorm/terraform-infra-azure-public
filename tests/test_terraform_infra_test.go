package test

import (
	"encoding/json"
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
	rawOutput, err := terraform.OutputE(t, terraformOptions, "resource_group_name")
	if err != nil {
		t.Fatalf("Failed to get output 'resource_group_name': %v", err)
	}

	// Décoder la sortie brute
	var outputResourceGroup string
	if err := json.Unmarshal([]byte(rawOutput), &outputResourceGroup); err != nil {
		t.Fatalf("Failed to parse output 'resource_group_name': %v", err)
	}
	t.Logf("Resource Group Name: %s", outputResourceGroup)

	// Valider la valeur
	assert.Equal(t, "terraform-iac-production-rg", outputResourceGroup)

	// Ajoutez des assertions similaires pour les autres outputs
	rawOutput, err = terraform.OutputE(t, terraformOptions, "acr_name")
	if err != nil {
		t.Fatalf("Failed to get output 'acr_name': %v", err)
	}

	var outputAcrName string
	if err := json.Unmarshal([]byte(rawOutput), &outputAcrName); err != nil {
		t.Fatalf("Failed to parse output 'acr_name': %v", err)
	}
	t.Logf("ACR Name: %s", outputAcrName)
	assert.Equal(t, "iacterraformprojectproductionacr", outputAcrName)
}
