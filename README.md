# 📊 Stock Recommender

Este proyecto es una aplicación web full stack para la gestión y recomendación de acciones (stocks), compuesta por un backend en Go con PostgreSQL y un frontend en Vue 3 con TypeScript.

---

## 🚀 Tecnologías

### Backend (Go + PostgreSQL)
- Go (Gin Framework)
- PostgreSQL
- Docker
- Terraform (para despliegue en AWS EC2)
- Swagger (documentación API)

### Frontend (Vue 3 + Vite + TailwindCSS)
- Vue 3 (Composition API)
- TypeScript
- Vue Router
- Pinia (manejo de estado)
- Axios (HTTP client)
- TailwindCSS (estilos)

---

## 📁 Estructura del Proyecto

```
stock-recommender/
├── backend/
│   ├── cmd/                   # Entry point (Main.go)
│   ├── config/                # Configuración y carga de variables .env
│   ├── controllers/           # Controladores HTTP
│   ├── database/              # Conexión, migraciones y consultas SQL
│   ├── docs/                  # Documentación Swagger
│   ├── middleware/            # Middleware personalizados
│   ├── models/                # Modelos de dominio y respuestas
│   ├── repository/            # Lógica de acceso a datos
│   ├── routes/                # Ruteo de endpoints
│   ├── services/              # Lógica de negocio
│   ├── tests/                 # Tests unitarios
│   └── terraform/             # Infraestructura como código
│
├── frontend/
│   ├── src/
│   │   ├── assets/            # Imágenes e íconos
│   │   ├── components/        # Componentes reutilizables
│   │   ├── pages/             # Vistas principales (crear, listar, info)
│   │   ├── router/            # Vue Router
│   │   ├── services/          # Servicios HTTP (Axios)
│   │   ├── store/             # Pinia store
│   │   ├── styles/            # Estilos globales
│   │   └── types/             # Tipado TypeScript
│   └── public/                # Archivos estáticos
```

---

## ⚙️ Instalación y ejecución local

### 🔧 Backend

1. Clona el proyecto:

```bash
git clone https://github.com/MalfurionStormrage/stock-recommender.git
cd stock-recommender/backend
```

2. Configura el archivo `.env` con las variables necesarias:

```env
DB_PORT=8080
DB_USER=...
DB_PASSWORD=...
DB_NAME=...
DB_SSLMODE=...
AUTHORIZATION=... #KEY DE API EXTERNA PARA MIGRAR DATOS A BASE DE DATOS DEL PROYECTO
```

3. Ejecuta con Docker:

```bash
docker build -t stock-backend .
docker run -p 8080:8080 --env-file .env stock-backend
```

4. Alternativamente, ejecuta con Go:

```bash
go run ./cmd/server/main/Main.go
```

---

### 🌐 Frontend

1. Ir al directorio del frontend:

```bash
cd ../frontend
```

2. Instalar dependencias:

```bash
npm install
```

3. Configurar variables en `.env`:

```env
EXTERNALKEY=....  #KEY PARA EL USO DE LA API 'https://finnhub.io/api/v1'
```

4. Iniciar servidor de desarrollo:

```bash
npm run dev
```

---

## ☁️ Despliegue

El backend se despliega en una instancia EC2 de AWS usando Terraform y Docker. La configuración se encuentra en `backend/terraform/`.

---

## 🧪 Tests

Los tests del backend están en `backend/tests/` y se pueden ejecutar con:

```bash
go test ./...
```

---

## 🛠️ Mantenimiento

- Ignorar archivos pesados en `.gitignore`
- Evitar subir `.terraform`, `node_modules`, archivos `.env` y ejecutables `.exe`

---

## 📄 Licencia

MIT © 2025 — [MalfurionStormrage](https://github.com/MalfurionStormrage)

---
