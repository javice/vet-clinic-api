package handlers

import (
    "errors"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/javice/vet-clinic-api/internal/models"
    
)

const (
    InvalidIDFormat = "Formato de ID inválido"
    PetNotFound     = "Mascota no encontrada"
    ClientNotExists = "El cliente no existe"
    InvalidPetData  = "Datos de mascota inválidos"
    ServerError     = "Error interno del servidor"
    PetDeleted      = "Mascota eliminada correctamente"
)



// GetPets obtiene todas las mascotas o filtra por cliente.
// @Summary Obtiene mascotas
// @Description Obtiene todas las mascotas o filtra por cliente si se especifica el parámetro `client_id`.
// @Tags Pets
// @Accept json
// @Produce json
// @Param client_id query int false "ID del cliente"
// @Success 200 {array} models.Pet
// @Failure 400 {object} map[string]interface{} "Formato de ID inválido"
// @Failure 404 {object} map[string]interface{} "Cliente no encontrado"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /api/v1/pets [get]
func (h *Handler) GetPets(c *gin.Context) {
    // Si se especifica client_id, filtrar por cliente
    clientID := c.Query("client_id")
    if clientID != "" {
        id, err := strconv.ParseUint(clientID, 10, 32)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": InvalidIDFormat})
            return
        }

        pets, err := h.PetRepo.GetByClientID(uint(id))
        if err != nil {
            status := http.StatusInternalServerError
            message := ServerError

            if errors.Is(err, h.PetRepo.DB.Error) {
                status = http.StatusNotFound
                message = ClientNotExists
            }

            c.JSON(status, gin.H{"error": message})
            return
        }
        c.JSON(http.StatusOK, pets)
        return
    }

    // Si no hay cliente específico, devolver todas las mascotas
    pets, err := h.PetRepo.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": ServerError})
        return
    }
    c.JSON(http.StatusOK, pets)
}

// GetPet obtiene una mascota por su ID
// @Summary Obtiene mascota por ID
// @Description Obtiene los detalles de una mascota específica por su ID
// @Tags Pets
// @Accept json
// @Produce json
// @Param id path int true "ID de la mascota"
// @Success 200 {object} models.Pet
// @Failure 400 {object} map[string]interface{} "Formato de ID inválido"
// @Failure 404 {object} map[string]interface{} "Mascota no encontrada"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /api/v1/pets/{id} [get]
func (h *Handler) GetPet(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": InvalidIDFormat})
        return
    }

    pet, err := h.PetRepo.GetByID(uint(id))
    if err != nil {
        status := http.StatusInternalServerError
        message := ServerError

        if errors.Is(err, h.PetRepo.DB.Error) {
            status = http.StatusNotFound
            message = PetNotFound
        }

        c.JSON(status, gin.H{"error": message})
        return
    }

    c.JSON(http.StatusOK, pet)
}


// CreatePet da de alta una nueva mascota
// @Summary Crea mascota
// @Description Crea una nueva mascota
// @Tags Pets
// @Accept json
// @Produce json
// @Param pet body models.Pet true "Datos de la mascota"
// @Success 201 {object} models.Pet
// @Failure 400 {object} map[string]interface{} "Error en los datos enviados"
// @Failure 404 {object} map[string]interface{} "Cliente no encontrado"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /api/v1/pets [post]
func (h *Handler) CreatePet(c *gin.Context) {
    var pet models.Pet
    if err := c.ShouldBindJSON(&pet); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    
    if err := h.PetRepo.Create(&pet); err != nil {
        status := http.StatusInternalServerError
        message := ServerError

        if errors.Is(err, h.PetRepo.DB.Error) {
            status = http.StatusBadRequest
            message = InvalidPetData
        }

        c.JSON(status, gin.H{"error": message})
        return
    }

    c.JSON(http.StatusCreated, pet)
}


// UpdatePet modifica total (PUT) o parcialmente (PATCH) una mascota
// @Summary Modifica mascota
// @Description Modifica una mascota total o parcialmente
// @Tags Pets
// @Accept json
// @Produce json
// @Param id path int true "ID de la mascota"
// @Param pet body models.Pet true "Datos de la mascota"
// @Success 200 {object} models.Pet
// @Failure 400 {object} map[string]interface{} "Error en los datos enviados"
// @Failure 404 {object} map[string]interface{} "Mascota o cliente no encontrado"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /api/v1/pets/{id} [put][patch]
func (h *Handler) UpdatePet(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": InvalidIDFormat})
        return
    }

    var pet models.Pet
    if err := c.ShouldBindJSON(&pet); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    pet.ID = uint(id)
    if err := h.PetRepo.Update(&pet); err != nil {
        status := http.StatusInternalServerError
        message := ServerError

        if errors.Is(err, h.PetRepo.DB.Error) {
			status = http.StatusNotFound
			message = PetNotFound
		}

        c.JSON(status, gin.H{"error": message})
        return
    }

    c.JSON(http.StatusOK, pet)
}

// DeletePet elimina una mascota
// @Summary Elimina mascota
// @Description Elimina una mascota por su ID
// @Tags Pets
// @Accept json
// @Produce json
// @Param id path int true "ID de la mascota"
// @Success 200 {object} map[string]string "Mascota eliminada correctamente"
// @Failure 400 {object} map[string]interface{} "Formato de ID inválido"
// @Failure 404 {object} map[string]interface{} "Mascota no encontrada"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /api/v1/pets/{id} [delete]
func (h *Handler) DeletePet(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": InvalidIDFormat})
        return
    }

    if err := h.PetRepo.Delete(uint(id)); err != nil {
        status := http.StatusInternalServerError
        message := ServerError

        if errors.Is(err, h.PetRepo.DB.Error) {
            status = http.StatusNotFound
            message = PetNotFound
        }

        c.JSON(status, gin.H{"error": message})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": PetDeleted})
}

// GetPetsByClient obtiene todas las mascotas de un cliente específico
// @Summary Obtiene mascotas por cliente
// @Description Obtiene todas las mascotas de un cliente específico
// @Tags Pets
// @Accept json
// @Produce json
// @Param id path int true "ID del cliente"
// @Success 200 {array} models.Pet
// @Failure 400 {object} map[string]interface{} "Formato de ID inválido"
// @Failure 404 {object} map[string]interface{} "Cliente no encontrado"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /api/v1/clients/{id}/pets [get]
func (h *Handler) GetPetsByClient(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": InvalidIDFormat})
        return
    }
    
    pets, err := h.PetRepo.GetByClientID(uint(id))
    if err != nil {
        status := http.StatusInternalServerError
        message := ServerError
        
        if errors.Is(err, h.PetRepo.DB.Error) {
            status = http.StatusNotFound
            message = ClientNotExists
        }
        
        c.JSON(status, gin.H{"error": message})
        return
    }
    
    c.JSON(http.StatusOK, pets)
}
