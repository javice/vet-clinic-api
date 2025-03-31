# Changelog

Todos los cambios notables en este proyecto serán documentados en este archivo.

El formato está basado en [Keep a Changelog](https://keepachangelog.com/es-ES/1.0.0/), y este proyecto adhiere a [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.1] - 2025-03-28

### Agregado V1.1.1

- Estructura base del proyecto
- Implementación de endpoints para clientes
- Implementación de endpoints para mascotas
- Tests para todos los endpoints
- Makefile para facilitar tareas comunes
- Configuración inicial de CI/CD con GitHub Actions
- Integración de Swagger para la documentación interactiva de la API.
- Middleware de Swagger configurado en `/swagger/index.html`.
- Generación automática de la documentación Swagger en el flujo de CI.

## [0.1.0] - 2025-03-22

### Agregado V0.1.0

- Configuración inicial del proyecto
- Implementación del modelo Cliente
- Implementación del modelo Mascota
- Implementación del modelo Cita
- Configuración de base de datos SQLite
- Setup básico de Gin Framework
- Implementación de repositorios base
- Estructura base de handlers
- Definición de rutas API
- Middleware CORS

### Por hacer

- Implementar autenticación y autorización
- Agregar validación más robusta de datos
- Implementar paginación en endpoints que devuelven listas
- Agregar documentación con Swagger
- Configurar contenedorización con Docker
- Agregar soporte para bases de datos adicionales
- Implementar búsqueda y filtros avanzados
- Agregar logging estructurado
- Implementar métricas y monitoreo