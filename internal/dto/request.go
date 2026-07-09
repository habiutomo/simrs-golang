package dto

import (
	"time"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreatePatientRequest struct {
	Name             string `json:"name" binding:"required"`
	NIK              string `json:"nik" binding:"required"`
	PlaceOfBirth     string `json:"place_of_birth"`
	DateOfBirth      string `json:"date_of_birth"`
	Gender           string `json:"gender" binding:"required"`
	Address          string `json:"address"`
	Phone            string `json:"phone"`
	Email            string `json:"email"`
	BloodType        string `json:"blood_type"`
	Allergies        string `json:"allergies"`
	InsuranceName    string `json:"insurance_name"`
	InsuranceNumber  string `json:"insurance_number"`
	EmergencyContact string `json:"emergency_contact"`
	EmergencyPhone   string `json:"emergency_phone"`
}

type UpdatePatientRequest struct {
	Name             string `json:"name"`
	PlaceOfBirth     string `json:"place_of_birth"`
	DateOfBirth      string `json:"date_of_birth"`
	Gender           string `json:"gender"`
	Address          string `json:"address"`
	Phone            string `json:"phone"`
	Email            string `json:"email"`
	BloodType        string `json:"blood_type"`
	Allergies        string `json:"allergies"`
	InsuranceName    string `json:"insurance_name"`
	InsuranceNumber  string `json:"insurance_number"`
	EmergencyContact string `json:"emergency_contact"`
	EmergencyPhone   string `json:"emergency_phone"`
}

type CreateAppointmentRequest struct {
	PatientID  uint   `json:"patient_id" binding:"required"`
	DoctorID   uint   `json:"doctor_id" binding:"required"`
	Date       string `json:"date" binding:"required"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Complaint  string `json:"complaint"`
}

type CreateMedicalRecordRequest struct {
	AppointmentID *uint  `json:"appointment_id"`
	PatientID     uint   `json:"patient_id" binding:"required"`
	DoctorID      uint   `json:"doctor_id" binding:"required"`
	Diagnosis     string `json:"diagnosis" binding:"required"`
	Complaint     string `json:"complaint"`
	Examination   string `json:"examination"`
	Treatment     string `json:"treatment"`
	Notes         string `json:"notes"`
	Prescriptions []CreatePrescriptionRequest `json:"prescriptions"`
}

type CreatePrescriptionRequest struct {
	MedicationID uint   `json:"medication_id" binding:"required"`
	Dosage       string `json:"dosage" binding:"required"`
	Frequency    string `json:"frequency" binding:"required"`
	Duration     string `json:"duration"`
	Quantity     int    `json:"quantity" binding:"required"`
	Notes        string `json:"notes"`
}

type CreateBillingRequest struct {
	PatientID     uint                  `json:"patient_id" binding:"required"`
	AppointmentID *uint                 `json:"appointment_id"`
	InpatientID   *uint                 `json:"inpatient_id"`
	Items         []CreateBillingItemRequest `json:"items" binding:"required"`
	PaymentMethod string                `json:"payment_method"`
	Notes         string                `json:"notes"`
}

type CreateBillingItemRequest struct {
	ItemType  string  `json:"item_type" binding:"required"`
	ItemName  string  `json:"item_name" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required"`
	UnitPrice float64 `json:"unit_price" binding:"required"`
}

type PayBillingRequest struct {
	Amount      float64 `json:"amount" binding:"required"`
	PaymentMethod string `json:"payment_method"`
}

type CreateDoctorRequest struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	FullName        string `json:"full_name" binding:"required"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	DepartmentID    uint   `json:"department_id" binding:"required"`
	LicenseNumber   string `json:"license_number" binding:"required"`
	Specialization  string `json:"specialization"`
	ConsultationFee float64 `json:"consultation_fee"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role" binding:"required"`
}

type UpdateDoctorScheduleRequest struct {
	DoctorID  uint   `json:"doctor_id" binding:"required"`
	DayOfWeek int    `json:"day_of_week" binding:"required"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime   string `json:"end_time" binding:"required"`
}

type DateRangeRequest struct {
	StartDate string `json:"start_date" form:"start_date"`
	EndDate   string `json:"end_date" form:"end_date"`
}

type AdmitPatientRequest struct {
	PatientID   uint   `json:"patient_id" binding:"required"`
	DoctorID    uint   `json:"doctor_id" binding:"required"`
	RoomID      uint   `json:"room_id" binding:"required"`
	Diagnosis   string `json:"diagnosis"`
}

type CreateMedicationRequest struct {
	Name        string  `json:"name" binding:"required"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Unit        string  `json:"unit" binding:"required"`
	Description string  `json:"description"`
}

type PaginationRequest struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

var _ = time.Time{}
