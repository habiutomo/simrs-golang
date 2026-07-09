package models

import "time"

type Department struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:100;not null" json:"name"`
	Code        string `gorm:"size:10;uniqueIndex;not null" json:"code"`
	Description string `gorm:"size:255" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Department) TableName() string {
	return "departments"
}
