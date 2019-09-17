package template

import (
	"fmt"
	"reflect"
)

type ModuleParams struct {
	Provider Provider          `json:"provider"`
	Name     string            `default:"dcos" json:"name"`
	Tags     map[string]string `json:"tags"`
	Vars     map[string]string `json:"vars"`
}

type Module struct {
	Provider  Provider          `hcle:"omit" json:"provider"`
	Providers map[string]string `hcl:"providers" json:"providers,omitempty"`
	Tags      map[string]string `hcl:"tags" hcle:"omitempty" json:"tags,omitempty"`
	Name      string            `hcl:",key" json:"name"`
	AWS       `hcl:",squash"`
}

func CreateModule(moduleparams ModuleParams) Module {
	var module Module
	// var structLoad interface{}

	// switch moduleparams.Provider.Name {
	// case "aws":
	// 	structLoad = &AWS{}
	// case "azurerm":
	// 	structLoad = &AzureRM{}
	// case "gcp":
	// 	structLoad = &GCP{}
	// }

	// set the
	typ := reflect.TypeOf(module.AWS)
	if module.AWS.ModuleSource == "" {
		f, _ := typ.FieldByName("ModuleSource")
		module.AWS.ModuleSource = f.Tag.Get("default")
	}
	if module.AWS.ModuleVersion == "" {
		f, _ := typ.FieldByName("ModuleVersion")
		module.AWS.ModuleVersion = f.Tag.Get("default")
	}

	// get and set default if module name is not set
	typ = reflect.TypeOf(moduleparams)
	if moduleparams.Name == "" {
		f, _ := typ.FieldByName("Name")
		module.Name = f.Tag.Get("default")
	} else {
		module.Name = moduleparams.Name
	}

	// Set provider
	module.Provider = moduleparams.Provider

	// Set providers in module
	module.Providers = make(map[string]string, 1)
	if moduleparams.Provider.Alias == "" {
		providerNameAlias := fmt.Sprintf("%v", module.Provider.Name)
		module.Providers[module.Provider.Name] = providerNameAlias
	} else {
		providerNameAlias := fmt.Sprintf("%v.%v", module.Provider.Name, moduleparams.Provider.Alias)
		module.Providers[module.Provider.Name] = providerNameAlias
	}

	// Set tags
	module.Tags = moduleparams.Tags

	return module
}
