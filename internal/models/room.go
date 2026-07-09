package models

import "time"

type Room struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Number      string    `gorm:"uniqueIndex;size:10;not null" json:"number"`
	Type        string    `gorm:"size:50;not null" json:"type"`
	Class       string    `gorm:"size:10;not null" json:"class"` // VIP, 1, 2, 3
	Capacity    int       `gorm:"not null;default:1" json:"capacity"`
	PricePerDay float64   `gorm:"not null;default:0" json:"price_per_day"`
	Status      string    `gorm:"size:20;not null;default:'available'" json:"status"` // available, occupied, maintenance
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Room) TableName() string {
	return "rooms"
}

type Inpatient struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	PatientID    uint      `gorm:"not null;index" json:"patient_id"`
	Patient      Patient   `gorm:"foreignKey:PatientID" json:"patient"`
	DoctorID     uint      `gorm:"not null;index" json:"doctor_id"`
	Doctor       Doctor    `gorm:"foreignKey:DoctorID" json:"doctor"`
	RoomID       uint      `gorm:"not null;index" json:"room_id"`
	Room         Room      `gorm:"foreignKey:RoomID" json:"room"`
	AdmissionDate time.Time `gorm:"not null" json:"admission_date"`
	DischargeDate *time.Time `json:"discharge_date,omitempty"`
	Diagnosis    string    `gorm:"type:text" json:"diagnosis"`
	Status       string    `gorm:"size:20;not null;default:'admitted'" json:"status"` // admitted, discharged, transferred
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (Inpatient) TableName() string {
	return "inpatients"
}
