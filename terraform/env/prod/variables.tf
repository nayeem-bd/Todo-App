variable "region" {
  description = "The AWS region to deploy resources in"
  type        = string
  default     = "us-west-2"
}

variable "access_key" {
  description = "The AWS access key"
  type        = string
  sensitive   = true
}

variable "secret_key" {
  description = "The AWS secret key"
  type        = string
  sensitive   = true
}

variable "vpc_cidr" {
  description = "The CIDR block for the VPC"
  type        = string
}

variable "public_subnet_cidr" {
  description = "The CIDR block for the public subnet"
  type        = string
}

variable "private_subnet_cidr" {
  description = "The CIDR block for the private subnet"
  type        = string
}

variable "private_subnet_cidr_2" {
  description = "The CIDR block for the private subnet"
  type        = string
}

variable "az1" {
  description = "The first availability zone"
  type        = string
}

variable "az2" {
  description = "The second availability zone"
  type        = string
}

variable "tags" {
  description = "A map of tags to assign to the resources"
  type        = map(string)
  default     = {}
}

variable "ami_id" {
  description = "The AMI ID to use for the EC2 instance"
  type        = string
}

variable "instance_type" {
  description = "The instance type for the EC2 instance"
  type        = string
  default     = "t2.micro"
}

variable "username" {
  description = "The username for the RDS instance"
  type        = string
}

variable "password" {
  description = "The password for the RDS instance"
  type        = string
  sensitive   = true
}

variable "bucket_name" {
  description = "The name of the S3 bucket to create"
  type        = string
}

variable "ec2_key_name" {
  description = "The name of the key pair to use for the EC2 instance"
  type        = string
  default     = "my-key-pair"
}

variable "broker_name" {
  description = "Name of the RabbitMQ broker"
  type        = string
}

variable "engine_version" {
  description = "RabbitMQ engine version"
  type        = string
  default     = "3.13"
}

variable "users" {
  description = "List of RabbitMQ users"
  type = list(object({
    username = string
    password = string
  }))
}
