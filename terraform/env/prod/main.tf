terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "6.9.0"
    }
  }
}
provider "aws" {
  region     = var.region
  access_key = var.access_key
  secret_key = var.secret_key
}

module "vpc" {
  source                = "../../modules/vpc"
  vpc_cidr              = var.vpc_cidr
  public_subnet_cidr    = var.public_subnet_cidr
  private_subnet_cidr   = var.private_subnet_cidr
  private_subnet_cidr_2 = var.private_subnet_cidr_2
  az1                   = var.az1
  az2                   = var.az2
  tags                  = var.tags
}

module "ec2" {
  source        = "../../modules/ec2"
  ami_id        = var.ami_id
  instance_type = var.instance_type
  subnet_id     = module.vpc.public_subnet_id
  ec2_key_name  = var.ec2_key_name
  vpc_id        = module.vpc.vpc_id
  tags          = var.tags
}

module "rds" {
  source     = "../../modules/rds"
  username   = var.username
  password   = var.password
  subnet_ids = module.vpc.private_subnet_ids
  vpc_id     = module.vpc.vpc_id
  tags       = var.tags
}

module "mq" {
  source              = "../../modules/mq"
  broker_name         = var.broker_name
  users               = var.users
  subnet_ids          = module.vpc.private_subnet_ids
  vpc_id              = module.vpc.vpc_id
  security_group_ids = [module.ec2.ec2_sg_id]
  tags                = var.tags
}