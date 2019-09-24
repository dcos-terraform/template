provider "aws" {
  alias = "master"

  assume_role {
    external_id  = "EXTERNAL_ID"
    role_arn     = "arn:aws:iam::ACCOUNT_ID:role/ROLE_NAME"
    session_name = "SESSION_NAME"
  }
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
  ssh_public_key_file = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDmnmRihQ1dLiFE8uGkvZOZzWneryuP1zDC9BM3T0i1iy0AVKsAe94adNG8DxmsE/KzFX1I8QuB8ALvAyC98oLqOa5xErjx0r+2eI/Xuj76H2cy/iAKboJh1uQnGRmN3cFuip4J+Uf9BEAQPRFaRinQT0zN+icQrgsmbqzbqLkT5F8B9cWV5zE9sydn/04tu0p4r6N2JTh97NU8/eZRfI2qhR3NQrDEwloWDw5Y/p9tizfRXwy43GWKjO1so5EjNzB/dNMckR1n7ZJ/hhttTEmKuNEO++9eixXohyKgtt5IUm48mWVzUnYmTsPca67e28VHvTu3bgDm/DqPqO7JQOetVHh6+90ljhN8V15+UbBGwSVlMhogaUnO8kdCVSBw160XeB1rkc0tDxdfV+086VqTVJGOj+9Trw+jGHQP8rY/jOQEuVESEXyquRy97JCBgYePP48fsBDA0U50VFx1MlxdUSRpW4ksF9/a+hdhZ4yW/s3+7219epp7q15EAAr+ICtaa9Gw+HXXZ+X1rnnP0+xfEGbxP218LrDvv+pJ5nFVhwSlu/EjoK3KFXGaTX5+TFfLvPOQ6uBH4qVKWropsNEngOWzlIG/Nve/zj0Bpipfj120aOSi2ufvr+JaH2dmOV+2bA9Th1O/d2zH8A5rz0mtFd76ROHyAwlBiirb454UWw== sbrandt@mesosphere.com"
  admin_ips           = "127.0.0.1/32"
  dcos_instance_os    = "centos_7.6"
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
