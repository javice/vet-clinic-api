# Makefile para Vet Clinic API

# Variables
APP_NAME = vet-clinic-api
BUILD_DIR = ./build
MAIN_FILE = ./cmd/api/main.go
BINARY = $(BUILD_DIR)/$(APP_NAME)

# Go commands
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOMOD = $(GOCMD) mod
GOGET = $(GOCMD) get

.PHONY: all build clean test run tidy deps help

all: clean build

# Crear el directorio de compilación si no existe
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# Compilar la aplicación
build: tidy $(BUILD_DIR)
	$(GOBUILD) -o $(BINARY) $(MAIN_FILE)

# Ejecutar la aplicación
run: build
	$(BINARY)

# Ejecutar tests
test:
	$(GOTEST) -v ./tests/...

# Ejecutar tests con cobertura
test-coverage:
	$(GOTEST) -v -cover ./tests/...

# Limpiar binarios generados
clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

# Actualizar dependencias
tidy:
	$(GOMOD) tidy

# Instalar dependencias
deps:
	$(GOGET) -u github.com/gin-gonic/gin
	$(GOGET) -u gorm.io/gorm
	$(GOGET) -u gorm.io/driver/sqlite
	$(GOGET) -u github.com/stretchr/testify

# Iniciar en modo desarrollo (con hot-reload si se tiene air instalado)
dev:
	which air > /dev/null && air || $(GOCMD) run $(MAIN_FILE)

# Mostrar ayuda
help:
	@echo "Opciones disponibles:"
	@echo "  make build         - Compila la aplicación"
	@echo "  make run           - Ejecuta la aplicación"
	@echo "  make test          - Ejecuta los tests"
	@echo "  make test-coverage - Ejecuta los tests con cobertura"
	@echo "  make clean         - Limpia los binarios"
	@echo "  make tidy          - Actualiza las dependencias en go.mod"
	@echo "  make deps          - Instala las dependencias necesarias"
	@echo "  make dev           - Ejecuta en modo desarrollo (con hot-reload si air está instalado)"