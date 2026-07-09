package repositories

import (
	"simrs-golang/internal/database"
	"simrs-golang/internal/models"
)

type PatientRepository struct{}

func NewPatientRepository() *PatientRepository {
	return &PatientRepository{}
}

func (r *PatientRepository) Create(patient *models.Patient) error {
	return database.DB.Create(patient).Error
}

func (r *PatientRepository) FindByID(id uint) (*models.Patient, error) {
	var patient models.Patient
	err := database.DB.First(&patient, id).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *PatientRepository) FindByNIK(nik string) (*models.Patient, error) {
	var patient models.Patient
	err := database.DB.Where("nik = ?", nik).First(&patient).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *PatientRepository) FindByMedicalRecordNumber(mrn string) (*models.Patient, error) {
	var patient models.Patient
	err := database.DB.Where("medical_record_number = ?", mrn).First(&patient).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *PatientRepository) FindAll(page, limit int, search string) ([]models.Patient, int64, error) {
	var patients []models.Patient
	var total int64

	query := database.DB.Model(&models.Patient{})
	if search != "" {
		query = query.Where("name LIKE ? OR nik LIKE ? OR medical_record_number LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%")
	}
	query.Count(&total)

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&patients).Error
	return patients, total, err
}

func (r *PatientRepository) Update(patient *models.Patient) error {
	return database.DB.Save(patient).Error
}

func (r *PatientRepository) Delete(id uint) error {
	return database.DB.Delete(&models.Patient{}, id).Error
}

func (r *PatientRepository) GetTotalCount() (int64, error) {
	var total int64
	err := database.DB.Model(&models.Patient{}).Count(&total).Error
	return total, err
}
