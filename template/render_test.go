package template

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	assert := assert.New(t)
	module := CreateModule(ModuleParams{
		Provider: Provider{Name: "aws", Alias: "master"},
		Tags:     map[string]string{"owner": "tester"},
		Name:     "",
	})

	render, _ := Render(
		Config{
			Providers: []Provider{
				module.Provider,
			},
			Locals: map[string]string{
				"cluster_name": "rendertest",
			},
			Modules: []Module{
				module,
			},
			Outputs: []Output{
				{
					Name:  "master-ips",
					Value: "${module.dcos.masters-ips}",
				},
				{
					Name:  "cluster-address",
					Value: "${module.dcos.masters-loadbalancer}",
				},
				{
					Name:  "public-agents-loadbalancer",
					Value: "${module.dcos.public-agents-loadbalancer}",
				},
			},
		},
	)

	assert.NotNil(render)
	testfile, _ := ioutil.ReadFile(fmt.Sprintf("_tests/render.tf"))
	assert.Equal(string(testfile), render, "Render error")
}
