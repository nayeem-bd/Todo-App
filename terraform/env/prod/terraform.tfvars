region                = "us-east-1"
access_key            = "your access_key"
secret_key            = "your secret_key"
vpc_cidr              = "10.0.0.0/16"
public_subnet_cidr    = "10.0.1.0/24"
private_subnet_cidr   = "10.0.2.0/24"
private_subnet_cidr_2 = "10.0.3.0/24"
az1                   = "us-east-1a"
az2                   = "us-east-1b"
ami_id                = "ami-0a7d80731ae1b2435"
instance_type         = "t2.micro"
username              = "pgmaster"
password              = "pgmaster"
bucket_name           = "my-personal-s3-bucket"
tags = {
  Environment = "prod"
  Project     = "todo-app"
  Owner       = "nayeem"
  CreatedBy   = "Terraform"
}
ec2_key_name   = "nayeem-server"
broker_name    = "todo-mq-broker"
engine_version = "3.13"
users = [
  {
    username = "admin"
    password = "adminadminadmin"
}]
