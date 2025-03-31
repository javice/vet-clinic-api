// internal/repositories/appointment.go
package repositories

import (
    "github.com/javice/vet-clinic-api/internal/models"
    "gorm.io/gorm"
)

type AppointmentRepository struct {
    DB *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) *AppointmentRepository {
    return &AppointmentRepository{DB: db}
}

func (r *AppointmentRepository) GetAll() ([]models.Appointment, error) {
    var appointments []models.Appointment
    result := r.DB.Find(&appointments)
    return appointments, result.Error
}

func (r *AppointmentRepository) GetByID(id uint) (models.Appointment, error) {
    var appointment models.Appointment
    result := r.DB.First(&appointment, id)
    return appointment, result.Error
}

func (r *AppointmentRepository) GetByPetID(petID uint) ([]models.Appointment, error) {
    var appointments []models.Appointment
    result := r.DB.Where("pet_id = ?", petID).Find(&appointments)
    return appointments, result.Error
}

func (r *AppointmentRepository) Create(appointment *models.Appointment) error {
    return r.DB.Create(appointment).Error
}

func (r *AppointmentRepository) Update(appointment *models.Appointment) error {
    return r.DB.Save(appointment).Error
}

func (r *AppointmentRepository) Delete(id uint) error {
    return r.DB.Delete(&models.Appointment{}, id).Error
}