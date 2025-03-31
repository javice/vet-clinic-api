package models

import (
    "time"
)

type Pet struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Name        string    `json:"name" binding:"required"`
    Species     string    `json:"species" binding:"required"`
    Breed       string    `json:"breed"`
    BirthDate   time.Time `json:"birth_date"`
    Weight      float64   `json:"weight"`
    ClientID    uint      `json:"client_id" binding:"required"`
	Appointments      []Appointment     `json:"appointments,omitempty" gorm:"foreignKey:PetID"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

/* type Appointment struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    PetID       uint      `json:"pet_id" binding:"required"`
    Date        time.Time `json:"date" binding:"required"`
    Reason      string    `json:"reason" binding:"required"`
    Notes       string    `json:"notes"`
    Completed   bool      `json:"completed" default:"false"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
} */