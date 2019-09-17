package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateModuleDCOS(t *testing.T) {
	assert := assert.New(t)

	module := CreateModule(ModuleParams{Provider: Provider{Name: "aws"}})
	assert.NotNil(module)
	assert.NotNil(module.Name)
	assert.Equal("dcos", module.Name, "Module name not correct")
	module = CreateModule(ModuleParams{Provider: Provider{Name: "azurerm"}})
	assert.NotNil(module)
	assert.NotNil(module.Name)
	assert.Equal("dcos", module.Name, "Module name not correct")
	module = CreateModule(ModuleParams{Provider: Provider{Name: "gcp"}})
	assert.NotNil(module)
	assert.NotNil(module.Name)
	assert.Equal("dcos", module.Name, "Module name not correct")
}

func TestCreateModuleDCOSModuleNameChange(t *testing.T) {
	assert := assert.New(t)

	module := CreateModule(ModuleParams{Provider: Provider{Name: "aws"}, Name: "foobar"})
	assert.NotNil(module)
	assert.NotNil(module.Name)
	assert.Equal("foobar", module.Name, "Module name not correct")
	module = CreateModule(ModuleParams{Provider: Provider{Name: "azurerm"}, Name: "foobar"})
	assert.NotNil(module)
	assert.NotNil(module.Name)
	assert.Equal("foobar", module.Name, "Module name not correct")
	module = CreateModule(ModuleParams{Provider: Provider{Name: "gcp"}, Name: "foobar"})
	assert.NotNil(module)
	assert.NotNil(module.Name)
	assert.Equal("foobar", module.Name, "Module name not correct")
}

func TestCreateModuleDCOSProviderAlias(t *testing.T) {
	assert := assert.New(t)

	module := CreateModule(ModuleParams{Provider: Provider{Name: "aws", Alias: "master"}})
	assert.NotNil(module)
	assert.NotNil(module.Providers)
	assert.Equal(map[string]string{"aws": "aws.master"}, module.Providers, "Providers not correct")
	module = CreateModule(ModuleParams{Provider: Provider{Name: "azurerm", Alias: "master"}})
	assert.NotNil(module)
	assert.NotNil(module.Providers)
	assert.Equal(map[string]string{"azurerm": "azurerm.master"}, module.Providers, "Providers not correct")
	module = CreateModule(ModuleParams{Provider: Provider{Name: "gcp", Alias: "master"}})
	assert.NotNil(module)
	assert.NotNil(module.Providers)
	assert.Equal(map[string]string{"gcp": "gcp.master"}, module.Providers, "Providers not correct")
}

func TestCreateModuleDCOSTags(t *testing.T) {
	assert := assert.New(t)

	module := CreateModule(ModuleParams{Provider: Provider{Name: "aws"}, Tags: map[string]string{"owner": "tester"}})
	assert.NotNil(module)
	assert.NotNil(module.Tags)
	assert.Equal(map[string]string{"owner": "tester"}, module.Tags, "Tags not correct")
	module = CreateModule(ModuleParams{Provider: Provider{Name: "azurerm"}, Tags: map[string]string{"owner": "tester"}})
	assert.NotNil(module)
	assert.NotNil(module.Tags)
	assert.Equal(map[string]string{"owner": "tester"}, module.Tags, "Tags not correct")
	module = CreateModule(ModuleParams{Provider: Provider{Name: "gcp"}, Tags: map[string]string{"owner": "tester"}})
	assert.NotNil(module)
	assert.NotNil(module.Tags)
	assert.Equal(map[string]string{"owner": "tester"}, module.Tags, "Tags not correct")
}
