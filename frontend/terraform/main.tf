provider "aws" {
  region = "us-east-1"
}

#resource "aws_key_pair" "frontend_key" {
#  key_name   = var.key_name
#  public_key = file(var.public_key_path)
#}

resource "aws_security_group" "frontend_sg" {
  name        = "frontend_sg"
  description = "Allow HTTP for frontend"
  
  ingress {
    from_port   = 80
    to_port     = 80
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

resource "aws_instance" "frontend_instance" {
  ami                    = "ami-0150ccaf51ab55a51" # Ubuntu 22.04 us-east-1
  instance_type          = "t2.micro"
  key_name               = var.key_name
  vpc_security_group_ids = [aws_security_group.frontend_sg.id]

  user_data = <<-EOF
              #!/bin/bash
              apt update -y
              apt install -y docker.io
              systemctl start docker
              docker run -d -p 80:80 6abraham6/frontend
              EOF

  tags = {
    Name = "frontend-instance"
  }
}
