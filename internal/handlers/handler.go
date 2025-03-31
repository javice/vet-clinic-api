package handlers

import (
    "github.com/javice/vet-clinic-api/internal/repositories"
)

type Handler struct {
    ClientRepo *repositories.ClientRepository
    PetRepo    *repositories.PetRepository
	AppointmentRepo *repositories.AppointmentRepository
}

func NewHandler(clientRepo *repositories.ClientRepository, petRepo *repositories.PetRepository, appointmentRepo *repositories.AppointmentRepository) *Handler {
    return &Handler{
        ClientRepo: clientRepo,
        PetRepo:    petRepo,
		AppointmentRepo: appointmentRepo,
    }
}