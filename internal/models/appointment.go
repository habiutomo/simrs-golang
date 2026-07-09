package models

import "time"

type Appointment struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	PatientID    uint      `gorm:"not null;index" json:"patient_id"`
	Patient      Patient   `gorm:"foreignKey:PatientID" json:"patient"`
	DoctorID     uint      `gorm:"not null;index" json:"doctor_id"`
	Doctor       Doctor    `gorm:"foreignKey:DoctorID" json:"doctor"`
	AppointmentDate string `gorm:"size:10;not null" json:"appointment_date"`
	StartTime    string    `gorm:"size:5" json:"start_time"`
	EndTime      string    `gorm:"size:5" json:"end_time"`
	Status       string    `gorm:"size:20;not null;default:'scheduled'" json:"status"` // scheduled, completed, cancelled, no_show
	Complaint    string    `gorm:"size:255" json:"complaint"`
	Notes        string    `gorm:"size:255" json:"notes"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (Appointment) TableName() string {
	return "appointments"
}
