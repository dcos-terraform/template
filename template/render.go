package template

import (
	"log"

	"github.com/rodaine/hclencoder"
)

type Config struct {
	Providers []Provider        `hcl:"provider" json:"provider"`
	Locals    map[string]string `hcl:"locals" hcle:"omitempty" json:"locals,omitempty"`
	Modules   []Module          `hcl:"module" json:"module"`
	Outputs   []Output          `hcl:"output" json:"output"`
}

func Render(tfmain Config) (string, error) {
	hcl, err := hclencoder.Encode(tfmain)
	if err != nil {
		log.Fatal("unable to encode: ", err)
	}

	return string(hcl), nil
}
