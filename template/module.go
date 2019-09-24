package template

import (
	"fmt"
	"reflect"
	"strings"
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

	intLoad := module.AWS
	switch moduleparams.Provider.Name {
	case "aws":
		s := reflect.ValueOf(&intLoad).Elem()
		typeOfT := s.Type()
		if intLoad.ModuleSource == "" {
			f, _ := typeOfT.FieldByName("ModuleSource")
			intLoad.ModuleSource = f.Tag.Get("default")
		}
		if intLoad.ModuleVersion == "" {
			f, _ := typeOfT.FieldByName("ModuleVersion")
			intLoad.ModuleVersion = f.Tag.Get("default")
		}
		for k, v := range moduleparams.Vars {
			for i := 0; i < s.NumField(); i++ {
				if strings.Replace(typeOfT.Field(i).Tag.Get("json"), ",omitempty", "", -1) == k {
					fieldname := typeOfT.Field(i).Name
					fieldReflect := s.FieldByName(fieldname)
					if fieldReflect.IsValid() {
						fieldReflect.SetString(v)
					}
				}
			}
		}
		module.AWS = intLoad
		// case "azurerm":
		// 	intLoad := AzureRM{}
		// case "gcp":
		// 	intLoad := GCP{}
	}

	// get and set default if module name is not set
	typ := reflect.TypeOf(moduleparams)
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
