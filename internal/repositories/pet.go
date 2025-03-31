// internal/repositories/pet.go
package repositories

import (
    "github.com/javice/vet-clinic-api/internal/models"
    "gorm.io/gorm"
)

type PetRepository struct {
    DB *gorm.DB
}

func NewPetRepository(db *gorm.DB) *PetRepository {
    return &PetRepository{DB: db}
}

func (r *PetRepository) GetAll() ([]models.Pet, error) {
    var pets []models.Pet
    result := r.DB.Find(&pets)
    return pets, result.Error
}

func (r *PetRepository) GetByID(id uint) (models.Pet, error) {
    var pet models.Pet
    result := r.DB.First(&pet, id)
    return pet, result.Error
}

func (r *PetRepository) GetByClientID(clientID uint) ([]models.Pet, error) {
    var pets []models.Pet
    result := r.DB.Where("client_id = ?", clientID).Find(&pets)
    return pets, result.Error
}

func (r *PetRepository) Create(pet *models.Pet) error {
    return r.DB.Create(pet).Error
}

func (r *PetRepository) Update(pet *models.Pet) error {
    return r.DB.Save(pet).Error
}

func (r *PetRepository) Delete(id uint) error {
    return r.DB.Delete(&models.Pet{}, id).Error
}