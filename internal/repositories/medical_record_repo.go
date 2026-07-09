package repositories

import (
	"simrs-golang/internal/database"
	"simrs-golang/internal/models"
)

type MedicalRecordRepository struct{}

func NewMedicalRecordRepository() *MedicalRecordRepository {
	return &MedicalRecordRepository{}
}

func (r *MedicalRecordRepository) Create(record *models.MedicalRecord) error {
	return database.DB.Create(record).Error
}

func (r *MedicalRecordRepository) FindByID(id uint) (*models.MedicalRecord, error) {
	var record models.MedicalRecord
	err := database.DB.Preload("Patient").Preload("Doctor.User").Preload("Appointment").
		First(&record, id).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *MedicalRecordRepository) FindByPatientID(patientID uint) ([]models.MedicalRecord, error) {
	var records []models.MedicalRecord
	err := database.DB.Preload("Doctor.User").Preload("Appointment").
		Where("patient_id = ?", patientID).Order("created_at DESC").Find(&records).Error
	return records, err
}

func (r *MedicalRecordRepository) FindByDoctorID(doctorID uint) ([]models.MedicalRecord, error) {
	var records []models.MedicalRecord
	err := database.DB.Preload("Patient").Preload("Appointment").
		Where("doctor_id = ?", doctorID).Order("created_at DESC").Find(&records).Error
	return records, err
}

func (r *MedicalRecordRepository) FindAll(page, limit int) ([]models.MedicalRecord, int64, error) {
	var records []models.MedicalRecord
	var total int64

	database.DB.Model(&models.MedicalRecord{}).Count(&total)
	offset := (page - 1) * limit
	err := database.DB.Preload("Patient").Preload("Doctor.User").
		Offset(offset).Limit(limit).Order("created_at DESC").Find(&records).Error
	return records, total, err
}

func (r *MedicalRecordRepository) Update(record *models.MedicalRecord) error {
	return database.DB.Save(record).Error
}

func (r *MedicalRecordRepository) CreatePrescription(prescription *models.MedicalPrescription) error {
	return database.DB.Create(prescription).Error
}

func (r *MedicalRecordRepository) FindPrescriptionsByRecordID(recordID uint) ([]models.MedicalPrescription, error) {
	var prescriptions []models.MedicalPrescription
	err := database.DB.Preload("Medication").
		Where("medical_record_id = ?", recordID).Find(&prescriptions).Error
	return prescriptions, err
}

func (r *MedicalRecordRepository) GetTotalCount() (int64, error) {
	var total int64
	err := database.DB.Model(&models.MedicalRecord{}).Count(&total).Error
	return total, err
}
