package services

import (
	"simrs-golang/internal/dto"
	"simrs-golang/internal/models"
	"simrs-golang/internal/repositories"
)

type MedicationService struct {
	repo *repositories.MedicationRepository
}

func NewMedicationService(repo *repositories.MedicationRepository) *MedicationService {
	return &MedicationService{repo: repo}
}

func (s *MedicationService) Create(req *dto.CreateMedicationRequest) (*models.Medication, error) {
	medication := &models.Medication{
		Name:        req.Name,
		Category:    req.Category,
		Stock:       req.Stock,
		Price:       req.Price,
		Unit:        req.Unit,
		Description: req.Description,
		IsActive:    true,
	}
	err := s.repo.Create(medication)
	return medication, err
}

func (s *MedicationService) GetByID(id uint) (*models.Medication, error) {
	return s.repo.FindByID(id)
}

func (s *MedicationService) GetAll(page, limit int, search string) ([]models.Medication, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	return s.repo.FindAll(page, limit, search)
}

func (s *MedicationService) Update(id uint, req *dto.CreateMedicationRequest) (*models.Medication, error) {
	medication, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	medication.Name = req.Name
	medication.Category = req.Category
	medication.Stock = req.Stock
	medication.Price = req.Price
	medication.Unit = req.Unit
	medication.Description = req.Description

	err = s.repo.Update(medication)
	return medication, err
}

func (s *MedicationService) Delete(id uint) error {
	return s.repo.Delete(id)
}
