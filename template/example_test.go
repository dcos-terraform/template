package template

import "fmt"

func Example() {
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

	fmt.Print(render)

	// Output:
	// provider "aws" {
	//   alias = "master"
	// }
	//
	// locals {
	//   cluster_name = "rendertest"
	// }
	//
	// module "dcos" {
	//   providers {
	//     aws = "aws.master"
	//   }
	//
	//   tags {
	//     owner = "tester"
	//   }
	//
	//   source              = "dcos-terraform/dcos/aws"
	//   version             = "~> 0.2.6"
	//   ssh_public_key_file = ""
	//   admin_ips           = ""
	// }
	//
	// output "master-ips" {
	//   value = "${module.dcos.masters-ips}"
	// }
	//
	// output "cluster-address" {
	//   value = "${module.dcos.masters-loadbalancer}"
	// }
	//
	// output "public-agents-loadbalancer" {
	//   value = "${module.dcos.public-agents-loadbalancer}"
	// }
	//
}
