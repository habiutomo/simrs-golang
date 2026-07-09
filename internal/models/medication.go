package models

import "time"

type Medication struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:150;not null" json:"name"`
	Category    string    `gorm:"size:50" json:"category"`
	Stock       int       `gorm:"not null;default:0" json:"stock"`
	Price       float64   `gorm:"not null;default:0" json:"price"`
	Unit        string    `gorm:"size:20;not null" json:"unit"`
	Description string    `gorm:"size:255" json:"description"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Medication) TableName() string {
	return "medications"
}
