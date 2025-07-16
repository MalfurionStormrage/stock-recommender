output "frontend_instance_public_ip" {
  description = "IP pública de la instancia frontend"
  value       = aws_instance.frontend_instance.public_ip
}
