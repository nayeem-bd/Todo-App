output "broker_id" {
  value = aws_mq_broker.rabbitmq.id
}

output "broker_arn" {
  value = aws_mq_broker.rabbitmq.arn
}

output "broker_endpoints" {
  value = aws_mq_broker.rabbitmq.instances
}

output "security_group_id" {
  value = aws_security_group.rabbitmq_sg.id
}

output "security_group_arn" {
  value = aws_security_group.rabbitmq_sg.arn
}