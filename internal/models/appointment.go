package models

import (
    "time"
)

type Appointment struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    PetID       uint      `json:"pet_id" binding:"required"`
    Date        time.Time `json:"date" binding:"required"`
    Reason      string    `json:"reason" binding:"required"`
    Notes       string    `json:"notes"`
    Completed   bool      `json:"completed" default:"false"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
	Duration   int       `json:"duration" binding:"required"` 
}