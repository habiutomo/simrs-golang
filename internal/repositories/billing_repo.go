package repositories

import (
	"simrs-golang/internal/database"
	"simrs-golang/internal/models"
)

type BillingRepository struct{}

func NewBillingRepository() *BillingRepository {
	return &BillingRepository{}
}

func (r *BillingRepository) Create(billing *models.Billing) error {
	return database.DB.Create(billing).Error
}

func (r *BillingRepository) FindByID(id uint) (*models.Billing, error) {
	var billing models.Billing
	err := database.DB.Preload("Patient").Preload("Items").
		Preload("Appointment").Preload("Inpatient.Room").
		First(&billing, id).Error
	if err != nil {
		return nil, err
	}
	return &billing, nil
}

func (r *BillingRepository) FindByPatientID(patientID uint) ([]models.Billing, error) {
	var billings []models.Billing
	err := database.DB.Preload("Items").
		Where("patient_id = ?", patientID).Order("created_at DESC").Find(&billings).Error
	return billings, err
}

func (r *BillingRepository) FindAll(page, limit int, status string) ([]models.Billing, int64, error) {
	var billings []models.Billing
	var total int64

	query := database.DB.Model(&models.Billing{})
	if status != "" {
		query = query.Where("payment_status = ?", status)
	}
	query.Count(&total)

	offset := (page - 1) * limit
	err := query.Preload("Patient").Preload("Items").
		Offset(offset).Limit(limit).Order("created_at DESC").Find(&billings).Error
	return billings, total, err
}

func (r *BillingRepository) Update(billing *models.Billing) error {
	return database.DB.Save(billing).Error
}

func (r *BillingRepository) GetTotalRevenue() (float64, error) {
	var total float64
	err := database.DB.Model(&models.Billing{}).
		Where("payment_status = ?", "paid").
		Select("COALESCE(SUM(total_amount), 0)").Scan(&total).Error
	return total, err
}

func (r *BillingRepository) GetTotalUnpaid() (float64, error) {
	var total float64
	err := database.DB.Model(&models.Billing{}).
		Where("payment_status IN ?", []string{"unpaid", "partial"}).
		Select("COALESCE(SUM(total_amount - paid_amount), 0)").Scan(&total).Error
	return total, err
}

func (r *BillingRepository) CreateItem(item *models.BillingItem) error {
	return database.DB.Create(item).Error
}
