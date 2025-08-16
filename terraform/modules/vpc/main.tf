resource "aws_vpc" "main" {
    cidr_block = var.vpc_cidr
    tags = merge(var.tags, {
        Name = var.tags.Project != "" ? "${var.tags.Project}-vpc" : "main-vpc"
    })
}

resource "aws_subnet" "public" {
    vpc_id = aws_vpc.main.id
    cidr_block = var.public_subnet_cidr
    availability_zone = var.az1
    map_public_ip_on_launch = true
    tags = var.tags
}

resource "aws_subnet" "private" {
  vpc_id = aws_vpc.main.id
    cidr_block  = var.private_subnet_cidr
    availability_zone = var.az1
    tags = var.tags
}

resource "aws_subnet" "private_2" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = var.private_subnet_cidr_2
  availability_zone = var.az2
  tags              = var.tags
}

resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.main.id
  tags   = var.tags
}

resource "aws_route_table" "public_rt" {
  vpc_id = aws_vpc.main.id
  tags   = var.tags
}

resource "aws_route" "internet_access" {
  route_table_id = aws_route_table.public_rt.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id = aws_internet_gateway.igw.id
}

resource "aws_route_table_association" "public_assoc" {
  subnet_id      = aws_subnet.public.id
  route_table_id = aws_route_table.public_rt.id
}

