package models

import "time"

type Doctor struct {
	ID               uint       `gorm:"primaryKey" json:"id"`
	UserID           uint       `gorm:"uniqueIndex;not null" json:"user_id"`
	User             User       `gorm:"foreignKey:UserID" json:"user"`
	DepartmentID     uint       `gorm:"not null" json:"department_id"`
	Department       Department `gorm:"foreignKey:DepartmentID" json:"department"`
	LicenseNumber    string     `gorm:"size:50;uniqueIndex;not null" json:"license_number"`
	Specialization   string     `gorm:"size:100" json:"specialization"`
	ConsultationFee  float64    `gorm:"default:0" json:"consultation_fee"`
	IsAvailable      bool       `gorm:"default:true" json:"is_available"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

func (Doctor) TableName() string {
	return "doctors"
}

type DoctorSchedule struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	DoctorID  uint      `gorm:"not null;index" json:"doctor_id"`
	Doctor    Doctor    `gorm:"foreignKey:DoctorID" json:"doctor"`
	DayOfWeek int       `gorm:"not null" json:"day_of_week"` // 0=Sunday, 1=Monday...
	StartTime string    `gorm:"size:5;not null" json:"start_time"` // HH:MM
	EndTime   string    `gorm:"size:5;not null" json:"end_time"`   // HH:MM
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (DoctorSchedule) TableName() string {
	return "doctor_schedules"
}
