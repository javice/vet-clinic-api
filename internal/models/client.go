package models

import (
    "time"
)

type Client struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Name      string    `json:"name" gorm:"not null" binding:"required"`
    Email     string    `json:"email" binding:"required,email" gorm:"unique;not null"`
    Phone     string    `json:"phone" binding:"required"`
    Address   string    `json:"address"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Pets      []Pet     `json:"pets,omitempty" gorm:"foreignKey:ClientID"`
}