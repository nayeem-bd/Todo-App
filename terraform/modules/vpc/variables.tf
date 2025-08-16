variable "vpc_cidr" {
    description = "The CIDR block for the VPC"
    type        = string
}

variable "tags" {
    description = "A map of tags to assign to the VPC"
    type        = map(string)
    default     = {}
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
    description = "The first availability zone for the public subnet"
    type        = string
}

variable "az2" {
    description = "The second availability zone for the private subnet"
    type        = string
}