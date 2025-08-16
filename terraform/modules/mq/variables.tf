variable "broker_name" {
    type        = string
    description = "Name of the RabbitMQ broker"
}

variable "instance_type" {
    type        = string
    default     = "mq.t3.micro"
    description = "Instance type for RabbitMQ broker"
}

variable "engine_version" {
    type        = string
    default     = "3.13"
    description = "RabbitMQ engine version"
}

variable "users" {
    type = list(object({
        username = string
        password = string
    }))
    description = "List of RabbitMQ users"
}

variable "subnet_ids" {
    type        = list(string)
    description = "Subnets for broker"
}

variable "vpc_id" {
    type        = string
    description = "VPC ID where the RabbitMQ broker will be created"
}

variable "allowed_cidr_blocks" {
    type        = list(string)
    description = "CIDR blocks allowed to access RabbitMQ"
    default     = ["10.0.0.0/8"]
}

variable "security_group_ids" {
    type        = list(string)
    description = "Security groups attached to broker"
    default     = []
}

variable "tags" {
    type        = map(string)
    default     = {}
    description = "Tags to apply to the RabbitMQ broker"
}
