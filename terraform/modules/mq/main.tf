resource "aws_security_group" "rabbitmq_sg" {
    name_prefix = "${var.broker_name}-sg"
    description = "Security group for RabbitMQ broker"
    vpc_id      = var.vpc_id

    # RabbitMQ management console (HTTPS)
    ingress {
        from_port   = 443
        to_port     = 443
        protocol    = "tcp"
        security_groups = var.security_group_ids
        description = "HTTPS access for RabbitMQ management console"
    }

    # RabbitMQ AMQP port
    ingress {
        from_port   = 5672
        to_port     = 5672
        protocol    = "tcp"
        security_groups = var.security_group_ids
        description = "AMQP access for RabbitMQ"
    }

    # RabbitMQ AMQPS port (TLS)
    ingress {
        from_port   = 5671
        to_port     = 5671
        protocol    = "tcp"
        security_groups = var.security_group_ids
        description = "AMQPS access for RabbitMQ (TLS)"
    }

    # Allow all outbound traffic
    egress {
        from_port   = 0
        to_port     = 0
        protocol    = "-1"
        cidr_blocks = ["0.0.0.0/0"]
        description = "Allow all outbound traffic"
    }

    tags = merge(var.tags, {
        Name = "${var.broker_name}-sg"
    })
}

resource "aws_mq_broker" "rabbitmq" {
    broker_name       = var.broker_name
    engine_type       = "RabbitMQ"
    engine_version    = var.engine_version
    host_instance_type = var.instance_type
    publicly_accessible = false
    auto_minor_version_upgrade = true

    dynamic "user" {
        for_each = var.users
        content {
            username = user.value.username
            password = user.value.password
        }
    }

    deployment_mode  = "SINGLE_INSTANCE"
    subnet_ids       = [var.subnet_ids[0]]
    security_groups  = [aws_security_group.rabbitmq_sg.id]
    tags = var.tags
}
