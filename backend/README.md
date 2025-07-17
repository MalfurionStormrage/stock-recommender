# Backend - Stock API

API REST desarrollada en **Go (Gin)**, con **PostgreSQL** y **Docker**. Expone endpoints CRUD para gestionar acciones bursátiles (stocks).

## 🚀 Iniciar con Docker

```bash
docker-compose up --build
```

La API estará disponible en:  
`http://localhost:8080`

## ⚙️ Configuración

- Edita el archivo `config/config.json` o utiliza variables de entorno.
- La base de datos se inicializa automáticamente.

> Este backend está preparado para ser desplegado fácilmente usando Terraform sobre AWS EC2.
