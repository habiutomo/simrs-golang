package models

import "time"

type MedicalRecord struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	AppointmentID   *uint     `gorm:"index" json:"appointment_id"`
	Appointment     *Appointment `gorm:"foreignKey:AppointmentID" json:"appointment,omitempty"`
	PatientID       uint      `gorm:"not null;index" json:"patient_id"`
	Patient         Patient   `gorm:"foreignKey:PatientID" json:"patient"`
	DoctorID        uint      `gorm:"not null;index" json:"doctor_id"`
	Doctor          Doctor    `gorm:"foreignKey:DoctorID" json:"doctor"`
	Diagnosis       string    `gorm:"type:text" json:"diagnosis"`
	Complaint       string    `gorm:"type:text" json:"complaint"`
	Examination     string    `gorm:"type:text" json:"examination"`
	Treatment       string    `gorm:"type:text" json:"treatment"`
	Notes           string                `gorm:"type:text" json:"notes"`
	Prescriptions   []MedicalPrescription `gorm:"foreignKey:MedicalRecordID" json:"prescriptions,omitempty"`
	CreatedAt       time.Time             `json:"created_at"`
	UpdatedAt       time.Time             `json:"updated_at"`
}

func (MedicalRecord) TableName() string {
	return "medical_records"
}

type MedicalPrescription struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	MedicalRecordID  uint           `gorm:"not null;index" json:"medical_record_id"`
	MedicalRecord    MedicalRecord  `gorm:"foreignKey:MedicalRecordID" json:"medical_record"`
	MedicationID     uint           `gorm:"not null" json:"medication_id"`
	Medication       Medication     `gorm:"foreignKey:MedicationID" json:"medication"`
	Dosage           string         `gorm:"size:50;not null" json:"dosage"`
	Frequency        string         `gorm:"size:50;not null" json:"frequency"`
	Duration         string         `gorm:"size:50" json:"duration"`
	Quantity         int            `gorm:"not null" json:"quantity"`
	Notes            string         `gorm:"size:255" json:"notes"`
	CreatedAt        time.Time      `json:"created_at"`
}

func (MedicalPrescription) TableName() string {
	return "medical_prescriptions"
}
