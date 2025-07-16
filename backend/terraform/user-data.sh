#!/bin/bash
yum update -y
amazon-linux-extras install docker -y
systemctl start docker
systemctl enable docker

# Reemplaza esto con tu nombre de imagen en Docker Hub
docker pull 6abraham6/backend:latest
docker run -d -p 8080:8080 6abraham6/backend:latest
