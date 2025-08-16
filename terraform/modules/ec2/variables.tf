variable "ami_id" {
    description = "The AMI ID to use for the EC2 instance"
    type        = string
}

variable "instance_type" {
    description = "The type of instance to use for the EC2 instance"
    type        = string
}

variable "subnet_id" {
    description = "The ID of the subnet to launch the EC2 instance in"
    type        = string
}

variable "tags" {
    description = "A map of tags to assign to the EC2 instance"
    type        = map(string)
    default     = {}
}

variable "vpc_id" {
    description = "The ID of the VPC where the EC2 instance will be launched"
    type        = string
}

variable "ec2_key_name" {
    description = "The name of the key pair to use for the EC2 instance"
    type        = string
}