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

func parseArgs() {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "resource_group_name=") {
			resourceGroupName = strings.TrimPrefix(arg, "resource_group_name=")
		} else if strings.HasPrefix(arg, "acr_name=") {
			acrName = strings.TrimPrefix(arg, "acr_name=")
		} else if strings.HasPrefix(arg, "location=") {
			location = strings.TrimPrefix(arg, "location=")
		}
	}
}

func TestTerraformInfrastructure(t *testing.T) {
	// Parse explicitement les arguments
	parseArgs()

	t.Logf("Parsed Arguments: resource_group_name=%s, acr_name=%s, location=%s", resourceGroupName, acrName, location)

	// Vérifiez que les arguments sont bien passés
	if resourceGroupName == "" || acrName == "" || location == "" {
		t.Fatalf("Missing required arguments: resource_group_name=%s, acr_name=%s, location=%s", resourceGroupName, acrName, location)
	}

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

	outputResourceGroup := terraform.Output(t, terraformOptions, "resource_group_name")
	assert.Equal(t, resourceGroupName, outputResourceGroup)

	outputAcrName := terraform.Output(t, terraformOptions, "acr_name")
	assert.Equal(t, acrName, outputAcrName)
}
