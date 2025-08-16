resource "aws_security_group" "ec2_sg" {
    name        = "ec2-sg"
    description = "Security group for EC2 instances"
    vpc_id      = var.vpc_id
    ingress {
        description = "HTTP access"
        from_port   = 80
        to_port     = 80
        protocol    = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }

    ingress {
        description = "HTTP access"
        from_port   = 8080
        to_port     = 8080
        protocol    = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }

    ingress {
      description = "ssh access"
        from_port   = 22
        to_port     = 22
        protocol    = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }
    egress {
        from_port   = 0
        to_port     = 0
        protocol    = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }
}

resource "aws_instance" "web" {
  ami = var.ami_id
  instance_type = var.instance_type
  subnet_id = var.subnet_id
  user_data = file("${path.module}/user-data.sh")
  vpc_security_group_ids = [aws_security_group.ec2_sg.id]
  key_name = var.ec2_key_name
  tags = merge(var.tags, {
    Name = "${var.tags.Project != "" ? var.tags.Project : "web"}-instance"
  })
}
