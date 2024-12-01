package test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformInfrastructure(t *testing.T) {
	resourceGroupName := os.Getenv("resource_group_name")
	acrName := os.Getenv("acr_name")
	location := os.Getenv("location")

	terraformOptions := &terraform.Options{
		TerraformDir: "../terraform",
		Vars: map[string]interface{}{
			"resource_group_name": resourceGroupName,
			"acr_name":            acrName,
			"location":            location,
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	outputACRName := terraform.Output(t, terraformOptions, "acr_name")
	assert.Equal(t, acrName, outputACRName)
}
