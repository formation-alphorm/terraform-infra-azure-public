package test

import (
	"flag"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var resourceGroupName string
var acrName string
var location string

func init() {
	flag.StringVar(&resourceGroupName, "resource_group_name", "", "Name of the resource group")
	flag.StringVar(&acrName, "acr_name", "", "Name of the Azure Container Registry")
	flag.StringVar(&location, "location", "", "Location of the resources")
}

func TestTerraformInfrastructure(t *testing.T) {
	flag.Parse()

	if resourceGroupName == "" || acrName == "" || location == "" {
		t.Fatalf("Missing required arguments: resource_group_name=%s, acr_name=%s, location=%s", resourceGroupName, acrName, location)
	}

	terraformOptions := &terraform.Options{
		TerraformDir: "../terraform",orm
		Vars: map[string]interface{}{
			"resource_group_name": resourceGroupName,
			"acr_name":            acrName,
			"location":            location,
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	outputAcrName := terraform.Output(t, terraformOptions, "acr_name")
	assert.Equal(t, acrName, outputAcrName)
}
