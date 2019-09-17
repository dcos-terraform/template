provider "aws" {
  alias = "master"
}

locals {
  cluster_name = "rendertest"
}

module "dcos" {
  providers {
    aws = "aws.master"
  }

  tags {
    owner = "tester"
  }

  source              = "dcos-terraform/dcos/aws"
  version             = "~> 0.2.6"
  ssh_public_key_file = ""
  admin_ips           = ""
}

output "master-ips" {
  value = "${module.dcos.masters-ips}"
}

output "cluster-address" {
  value = "${module.dcos.masters-loadbalancer}"
}

output "public-agents-loadbalancer" {
  value = "${module.dcos.public-agents-loadbalancer}"
}
