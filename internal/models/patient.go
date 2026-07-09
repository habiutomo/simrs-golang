package models

import "time"

type Patient struct {
	ID                   uint      `gorm:"primaryKey" json:"id"`
	MedicalRecordNumber  string    `gorm:"uniqueIndex;size:20;not null" json:"medical_record_number"`
	Name                 string    `gorm:"size:150;not null" json:"name"`
	NIK                  string    `gorm:"uniqueIndex;size:20;not null" json:"nik"`
	PlaceOfBirth         string    `gorm:"size:50" json:"place_of_birth"`
	DateOfBirth          string    `gorm:"size:10" json:"date_of_birth"`
	Gender               string    `gorm:"size:10;not null" json:"gender"`
	Address              string    `gorm:"size:255" json:"address"`
	Phone                string    `gorm:"size:20" json:"phone"`
	Email                string    `gorm:"size:100" json:"email"`
	BloodType            string    `gorm:"size:5" json:"blood_type"`
	Allergies            string    `gorm:"size:255" json:"allergies"`
	InsuranceName        string    `gorm:"size:100" json:"insurance_name"`
	InsuranceNumber      string    `gorm:"size:50" json:"insurance_number"`
	EmergencyContact     string    `gorm:"size:100" json:"emergency_contact"`
	EmergencyPhone       string    `gorm:"size:20" json:"emergency_phone"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

func (Patient) TableName() string {
	return "patients"
}
