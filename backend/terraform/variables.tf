variable "key_name" {
  description = "Nombre del par de claves SSH"
  type        = string
}

variable "public_key_path" {
  description = "Ruta a la clave pública SSH (.pub)"
  type        = string
}
