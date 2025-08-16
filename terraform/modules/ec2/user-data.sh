#!/bin/bash
sudo apt update -y
sudo apt install -y nginx
sudo systemctl start nginx
sudo systemctl enable nginx

echo "<h1>Hello from Terraform EC2 Web Server with NGINX</h1>" | sudo tee /var/www/html/index.html
