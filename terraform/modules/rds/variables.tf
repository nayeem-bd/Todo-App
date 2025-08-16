variable "username" {
    description = "The username for the RDS instance"
    type        = string
}

variable "password" {
    description = "The password for the RDS instance"
    type        = string
}

variable "subnet_ids" {
    description = "A list of subnet IDs for the RDS instance"
    type        = list(string)
}

variable "vpc_id" {
    description = "The ID of the VPC where the RDS instance will be launched"
    type        = string
}

variable "tags" {
    description = "A map of tags to assign to the RDS instance"
    type        = map(string)
    default     = {}
}