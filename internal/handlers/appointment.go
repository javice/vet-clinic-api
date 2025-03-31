package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/javice/vet-clinic-api/internal/models"
)

// GetAppointments obtiene todas las citas o filtra por mascota.
// @Summary Obtiene citas
// @Description Obtiene todas las citas o filtra por mascota si se especifica el parámetro `pet_id`.
// @Tags Appointments
// @Accept json
// @Produce json
// @Param pet_id query int false "ID de la mascota"
// @Success 200 {array} models.Appointment
// @Failure 400 {object} map[string]interface{} "Formato de ID inválido"
// @Failure 404 {object} map[string]interface{} "Mascota no encontrada"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /api/v1/appointments [get]
func (h *Handler) GetAppointments(c *gin.Context) {
    // Si se especifica pet_id, filtrar por mascota
    petID := c.Query("pet_id")
    if petID != "" {
        id, err := strconv.ParseUint(petID, 10, 32)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de ID mascota NO válido"})
            return
        }

		appointments, err := h.AppointmentRepo.GetByPetID(uint(id))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, appointments)
        return

        
    }

    // Si no hay mascota específica, devolver todas las citas
    appointments, err := h.AppointmentRepo.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las citas"})
        return
    }
    c.JSON(http.StatusOK, appointments)
}


// GetAppointment obtiene Una cita por ID.
// @Summary Obtiene cita
// @Description Obtiene una cita
// @Tags Appointments
// @Accept json
// @Produce json
// @Param id path int true "ID de la cita"
// @Success 200 {object} models.Appointment
// @Failure 400 {object} map[string]interface{} "Formato de ID inválido"
// @Failure 404 {object} map[string]interface{} "Cita no encontrada"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /api/v1/appointments/{id} [get]
func (h *Handler) GetAppointment(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Formato ID inválido"})
        return
    }

    appointment, err := h.AppointmentRepo.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Cita no encontrada"})
        return
    }

    c.JSON(http.StatusOK, appointment)
}


// CreateAppointment da de alta una nueva cita.
// @Summary Crea una cita
// @Description Crea una nueva cita
// @Tags Appointments
// @Accept json
// @Produce json
// @Param appointment body models.Appointment true "Datos de la cita"
// @Success 201 {object} models.Appointment
// @Failure 400 {object} map[string]interface{} "Error en los datos enviados o fecha inválida"
// @Failure 404 {object} map[string]interface{} "Mascota no encontrada"
// @Failure 409 {object} map[string]interface{} "Conflicto con otra cita"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /api/v1/appointments [post]
func (h *Handler) CreateAppointment(c *gin.Context) {
    var appointment models.Appointment
    if err := c.ShouldBindJSON(&appointment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validación básica
    if appointment.PetID == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID de mascota es requerido"})
        return
    }

    if appointment.Reason == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Motivo de la cita es requerido"})
        return
    }

    // Usar el servicio para crear la cita, que incluye todas las validaciones
    if err := h.AppointmentRepo.Create(&appointment); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la cita"})
        return // Este return debe estar dentro del bloque if
    }

    c.JSON(http.StatusCreated, appointment)
}


// UpdateAppointment actualiza una cita existente.
// @Summary Actualiza cita
// @Description Actualiza una cita existente
// @Tags Appointments
// @Accept json
// @Produce json
// @Param id path int true "ID de la cita"
// @Param appointment body models.Appointment true "Datos de la cita"
// @Success 200 {object} models.Appointment
// @Failure 400 {object} map[string]interface{} "Error en los datos enviados o fecha inválida"
// @Failure 404 {object} map[string]interface{} "Cita o mascota no encontrada"
// @Failure 409 {object} map[string]interface{} "Conflicto con otra cita"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /api/v1/appointments/{id} [put]
func (h *Handler) UpdateAppointment(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de ID inválido"})
        return
    }

    var appointment models.Appointment
    if err := c.ShouldBindJSON(&appointment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Establecer el ID de la cita a actualizar
    appointment.ID = uint(id)

    // Validaciones básicas
    if appointment.PetID == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID de mascota es requerido"})
        return
    }

    if appointment.Reason == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Motivo de la cita es requerido"})
        return
    }

    // Usar el repo para actualizar la cita
    if err := h.AppointmentRepo.Update(&appointment); err != nil { 
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la cita"})
        return
    }

    c.JSON(http.StatusOK, appointment)
}

// DeleteAppointment elimina una cita existente.
// @Summary Elimina cita
// @Description Elimina una cita existente
// @Tags Appointments
// @Accept json
// @Produce json
// @Param id path int true "ID de la cita"
// @Success 200 {object} map[string]interface{} "Cita eliminada exitosamente"
// @Failure 400 {object} map[string]interface{} "Formato de ID inválido"
// @Failure 404 {object} map[string]interface{} "Cita no encontrada"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /api/v1/appointments/{id} [delete]
func (h *Handler) DeleteAppointment(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de ID inválido"})
        return
    }

    if err := h.AppointmentRepo.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la cita"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Cita eliminada exitosamente"})
}

// GetAppointmentsByPet obtiene todas las citas de una mascota específica
// @Summary Obtiene citas por mascota
// @Description Obtiene todas las citas de una mascota específica
// @Tags Appointments
// @Accept json
// @Produce json
// @Param id path int true "ID de la mascota"
// @Success 200 {array} models.Appointment
// @Failure 400 {object} map[string]interface{} "Formato de ID inválido"
// @Failure 404 {object} map[string]interface{} "Mascota no encontrada"
// @Failure 500 {object} map[string]interface{} "Error interno del servidor"
// @Router /api/v1/pets/{id}/appointments [get]
func (h *Handler) GetAppointmentsByPet(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de ID inválido"})
        return
    }
    
    appointments, err := h.AppointmentRepo.GetByPetID(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las citas"})
        return
    }
    
    c.JSON(http.StatusOK, appointments)
}
