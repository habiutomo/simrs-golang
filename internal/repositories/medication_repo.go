package repositories

import (
	"simrs-golang/internal/database"
	"simrs-golang/internal/models"
)

type MedicationRepository struct{}

func NewMedicationRepository() *MedicationRepository {
	return &MedicationRepository{}
}

func (r *MedicationRepository) Create(medication *models.Medication) error {
	return database.DB.Create(medication).Error
}

func (r *MedicationRepository) FindByID(id uint) (*models.Medication, error) {
	var medication models.Medication
	err := database.DB.First(&medication, id).Error
	if err != nil {
		return nil, err
	}
	return &medication, nil
}

func (r *MedicationRepository) FindAll(page, limit int, search string) ([]models.Medication, int64, error) {
	var medications []models.Medication
	var total int64

	query := database.DB.Model(&models.Medication{})
	if search != "" {
		query = query.Where("name LIKE ? OR category LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	query.Count(&total)

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Order("name ASC").Find(&medications).Error
	return medications, total, err
}

func (r *MedicationRepository) Update(medication *models.Medication) error {
	return database.DB.Save(medication).Error
}

func (r *MedicationRepository) Delete(id uint) error {
	return database.DB.Delete(&models.Medication{}, id).Error
}

func (r *MedicationRepository) ReduceStock(id uint, quantity int) error {
	return database.DB.Model(&models.Medication{}).Where("id = ?", id).
		UpdateColumn("stock", database.DB.Raw("stock - ?", quantity)).Error
}
