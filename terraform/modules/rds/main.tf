resource "aws_db_subnet_group" "rds_subnet_group" {
  name = "rds-subnet-group"
  subnet_ids = var.subnet_ids
}

resource "aws_security_group" "rds_sg" {
    name        = "rds-sg"
    description = "Security group for RDS instances"
    vpc_id      = var.vpc_id

    ingress {
      description = "PostgreSQL access"
      from_port   = 5432
      to_port     = 5432
      protocol    = "tcp"
      cidr_blocks = ["10.0.0.0/16"]
    }
  egress {
        from_port   = 0
        to_port     = 0
        protocol    = "-1"
        cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_db_instance" "rds_instance" {
    identifier = var.tags.Project != "" ? "${var.tags.Project}-rds" : "rds-instance"
    instance_class = "db.t4g.micro"
    engine = "postgres"
    engine_version = "17.5"
    allocated_storage = 20
    db_name = var.db_name
    username = var.username
    password = var.password
    publicly_accessible = false
    skip_final_snapshot = true
    vpc_security_group_ids = [aws_security_group.rds_sg.id]
    db_subnet_group_name = aws_db_subnet_group.rds_subnet_group.name
    tags = var.tags
}