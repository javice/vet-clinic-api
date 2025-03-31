// internal/repositories/client.go
package repositories

import (
    "github.com/javice/vet-clinic-api/internal/models"
    "gorm.io/gorm"
)

type ClientRepository struct {
    DB *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
    return &ClientRepository{DB: db}
}

func (r *ClientRepository) GetAll() ([]models.Client, error) {
    var clients []models.Client
    result := r.DB.Find(&clients)
    return clients, result.Error
}

func (r *ClientRepository) GetByID(id uint) (models.Client, error) {
    var client models.Client
    result := r.DB.Preload("Pets").First(&client, id)
    return client, result.Error
}

func (r *ClientRepository) Create(client *models.Client) error {
    return r.DB.Create(client).Error
}

func (r *ClientRepository) Update(client *models.Client) error {
    return r.DB.Save(client).Error
}

func (r *ClientRepository) Delete(id uint) error {
    return r.DB.Delete(&models.Client{}, id).Error
}

