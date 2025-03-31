package handlers

import (
    "errors"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/javice/vet-clinic-api/internal/models"
)

// Error messages
const (
    InvalidClientIDFormat   = "Formato de ID cliente NO válido"
    ClientNotFoundMessage   = "Cliente NO encontrado"
    ClientDeletedMessage    = "Cliente eliminado correctamente"
    InternalServerErrMsg    = "Error interno del servidor"
)

// GetClients obtiene todos los clientes.
// @Summary Obtiene clientes
// @Description Obtiene todos los clientes.
// @Tags Clients
// @Accept json
// @Produce json
// @Success 200 {array} models.Client
// @Failure 400 {object} map[string]interface{} "Formato de ID inválido"
// @Failure 404 {object} map[string]interface{} "Cliente no encontrado"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /api/v1/clients [get]
func (h *Handler) GetClients(c *gin.Context) {
    clients, err := h.ClientRepo.GetAll()
    if err != nil {
        statusCode := http.StatusInternalServerError
        errorMsg := err.Error()

        if errors.Is(err, h.ClientRepo.DB.Error) {
            // Already using the correct status code
        } else {
            statusCode = http.StatusBadRequest
        }

        c.JSON(statusCode, gin.H{"error": errorMsg})
        return
    }

    c.JSON(http.StatusOK, clients)
}

// GetClient obtiene un único cliente
// @Summary Obtiene cliente por ID
// @Description Obtiene los detalles de un cliente específico por su ID
// @Tags Clients
// @Accept json
// @Produce json
// @Param id path int true "ID del cliente"
// @Success 200 {object} models.Client
// @Failure 400 {object} map[string]interface{} "Formato de ID inválido"
// @Failure 404 {object} map[string]interface{} "Cliente no encontrada"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /api/v1/clients/{id} [get]
func (h *Handler) GetClient(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": InvalidClientIDFormat})
        return
    }

    client, err := h.ClientRepo.GetByID(uint(id))
    if err != nil {
        statusCode := http.StatusInternalServerError
        errorMsg := err.Error()
        c.JSON(statusCode, gin.H{"error": errorMsg})
        return
    }

    c.JSON(http.StatusOK, client)
}

// CreateClient da de alta un nuevo cliente
// @Summary Crea cliente
// @Description Crea un nuevo cliente
// @Tags Clients
// @Accept json
// @Produce json
// @Param pet body models.Client true "Datos del cliente"
// @Success 201 {object} models.Client
// @Failure 400 {object} map[string]interface{} "Error en los datos enviados"
// @Failure 404 {object} map[string]interface{} "Cliente no encontrado"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /api/v1/clients [post]
func (h *Handler) CreateClient(c *gin.Context) {
    var client models.Client

    if err := c.ShouldBindJSON(&client); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.ClientRepo.Create(&client); err != nil {
        statusCode := http.StatusInternalServerError
        errorMsg := err.Error()
        c.JSON(statusCode, gin.H{"error": errorMsg})
        return
    }

    c.JSON(http.StatusCreated, client)
}

// UpdateClient modifica total (PUT) o parcialmente (PATCH) un cliente
// @Summary Modifica cliente
// @Description Modifica un cliente total o parcialmente
// @Tags Clients
// @Accept json
// @Produce json
// @Param id path int true "ID del cliente"
// @Param pet body models.Client true "Datos del cliente"
// @Success 200 {object} models.Client
// @Failure 400 {object} map[string]interface{} "Error en los datos enviados"
// @Failure 404 {object} map[string]interface{} "Cliente no encontrado"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /api/v1/clients/{id} [put][patch]
func (h *Handler) UpdateClient(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": InvalidClientIDFormat})
        return
    }

    var client models.Client
    if err := c.ShouldBindJSON(&client); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Set the ID from the path parameter
    client.ID = uint(id)

    // Call the service to update the client
    if err := h.ClientRepo.Update(&client); err != nil {
        statusCode := http.StatusInternalServerError
        errorMsg := err.Error()
        if errors.Is(err, h.ClientRepo.DB.Error) {
            // Already using the correct status code
            errorMsg = InternalServerErrMsg
        } else {
            statusCode = http.StatusBadRequest
        }
        
        c.JSON(statusCode, gin.H{"error": errorMsg})
        return
    }

    c.JSON(http.StatusOK, client)
}

// DeleteClient elimina un cliente por ID
// @Summary Elimina cliente
// @Description Elimina un cliente por su ID
// @Tags Clients
// @Accept json
// @Produce json
// @Param id path int true "ID del cliente"
// @Success 200 {object} map[string]string "Cliente eliminado correctamente"
// @Failure 400 {object} map[string]interface{} "Formato de ID inválido"
// @Failure 404 {object} map[string]interface{} "Cliente no encontrado"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /api/v1/clients/{id} [delete]
func (h *Handler) DeleteClient(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": InvalidClientIDFormat})
        return
    }

    if err := h.ClientRepo.Delete(uint(id)); err != nil {
        statusCode := http.StatusInternalServerError
        errorMsg := err.Error()
        
        if errors.Is(err, h.ClientRepo.DB.Error) {
            // Already using the correct status code
            errorMsg = InternalServerErrMsg
        } else {
            errorMsg = ClientNotFoundMessage
        }
        
        c.JSON(statusCode, gin.H{"error": errorMsg})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": ClientDeletedMessage})
}

