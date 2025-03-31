# Vet Clinic API

API REST de una clínica veterinaria desarrollada con Go y Gin Framework.

## Descripción

Este proyecto implementa una API REST para gestionar una clínica veterinaria, permitiendo administrar clientes, mascotas y citas. La API está construida con Go y el framework web Gin, utilizando GORM como ORM para el manejo de la base de datos.

## Características

- Gestión completa de clientes (CRUD)
- Gestión completa de mascotas (CRUD)
- Filtrado de mascotas por cliente
- Pruebas automatizadas para todos los endpoints
- Documentación interactiva con Swagger.
- Estructura modular y escalable

## Requisitos

- Go 1.18 o superior
- SQLite (para desarrollo)

## Estructura del proyecto

```
vet-clinic-api/
├── cmd/
│   └── api/
│       └── main.go            # Punto de entrada de la aplicación
├── internal/
│   ├── handlers/              # Manejadores HTTP
│   │   ├── client.go
│   │   ├── pet.go
│   │   └── handler.go
│   ├── models/                # Modelos de datos
│   │   ├── client.go
│   │   └── pet.go
│   ├── repositories/          # Capa de acceso a datos
│   │   ├── client.go
│   │   └── pet.go
│   └── routes/                # Configuración de rutas
│       └── routes.go
├── tests/                     # Tests de la API
│   ├── client_test.go
│   └── pet_test.go
├── .github/                   # Configuración de CI/CD
│   └── workflows/
│       └── ci.yml
├── Makefile                   # Comandos para compilar, ejecutar, etc.
├── go.mod                     # Dependencias del proyecto
├── go.sum
├── README.md                  # Documentación principal
└── CHANGELOG.md               # Registro de cambios
```

## Instalación

1. Clonar el repositorio:
  
  ```bash
  git clone https://github.com/tuusuario/vet-clinic-api.git
  cd vet-clinic-api
  ```
  
2. Instalar dependencias:
  
  ```bash
  make deps
  ```

3. Genera la documentación Swagger (opcional):
  
  ```bash
  swag init -g cmd/api/main.go -o docs
  ```  

4. Compilar el proyecto:
  
  ```bash
  make build
  ```

## Documentación de la API

La documentación de la API se encuentra en [docs/swagger.json](docs/swagger.json).

Puedes acceder a ella en [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
  
## Uso

### Ejecutar el servidor

```bash
make run
```

El servidor se iniciará en `http://localhost:8080` por defecto. Puedes cambiar el puerto mediante la variable de entorno `PORT`.

### Ejecutar en modo desarrollo (con hot-reload)

```bash
# Si tienes air instalado (https://github.com/cosmtrek/air)
make dev
```

### Ejecutar tests

```bash
make test
```

## Endpoints de la API

### Clientes

- `GET /api/v1/clients` - Obtener todos los clientes
- `GET /api/v1/clients/:id` - Obtener un cliente por ID
- `POST /api/v1/clients` - Crear un nuevo cliente
- `PUT /api/v1/clients/:id` - Actualizar un cliente
- `PATCH /api/v1/clients/:id` - Actualizar parcialmente un cliente
- `DELETE /api/v1/clients/:id` - Eliminar un cliente

### Mascotas

- `GET /api/v1/pets` - Obtener todas las mascotas
- `GET /api/v1/pets?client_id=X` - Obtener mascotas por ID de cliente
- `GET /api/v1/pets/:id` - Obtener una mascota por ID
- `POST /api/v1/pets` - Crear una nueva mascota
- `PUT /api/v1/pets/:id` - Actualizar una mascota
- `PATCH /api/v1/pets/:id` - Actualizar parcialmente una mascota
- `DELETE /api/v1/pets/:id` - Eliminar una mascota

## Ejemplos de uso

### Crear un cliente

```bash
curl -X POST http://localhost:8080/api/v1/clients \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Juan Pérez",
    "email": "juan@example.com",
    "phone": "123456789",
    "address": "Calle Principal 123"
  }'
```

### Crear una mascota

```bash
curl -X POST http://localhost:8080/api/v1/pets \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Rex",
    "species": "Perro",
    "breed": "Labrador",
    "birth_date": "2020-01-15T00:00:00Z",
    "client_id": 1,
    "description": "Perro muy amigable"
  }'
```

## Ejemplos de uso en Postman

### Crear un cliente en Postman

- URL: <http://localhost:8080/api/v1/clients>
- Method: POST
- Body: JSON

```json
{
    "name": "Alfrero Rodriguez",
    "email": "alfredrod@example.com",
    "phone": "8521478966",
    "address": "Calle 100"
}
```

### Crear una mascota en Postman

- URL: <http://localhost:8080/api/v1/pets>
- Method: POST
- Body: JSON

```json
{
    "name": "Rex",
    "species": "Perro",
    "breed": "Labrador",
    "birth_date": "2020-01-15T00:00:00Z",
    "client_id": 1,
    "description": "Perro muy amigable"
}
```

## Contribuir

1. Haz un fork del proyecto
2. Crea tu rama de características (`git checkout -b feature/amazing-feature`)
3. Haz commit de tus cambios (`git commit -m 'Add some amazing feature'`)
4. Haz push a la rama (`git push origin feature/amazing-feature`)
5. Abre un Pull Request

## Licencia

Este proyecto está licenciado bajo la licencia MIT - ver el archivo [LICENSE](LICENSE) para más detalles.