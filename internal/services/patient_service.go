package services

import (
	"fmt"
	"time"

	"simrs-golang/internal/database"
	"simrs-golang/internal/dto"
	"simrs-golang/internal/models"
	"simrs-golang/internal/repositories"
)

type PatientService struct {
	repo *repositories.PatientRepository
}

func NewPatientService(repo *repositories.PatientRepository) *PatientService {
	return &PatientService{repo: repo}
}

func generateMRN() string {
	var count int64
	database.DB.Model(&models.Patient{}).Count(&count)
	now := time.Now()
	return fmt.Sprintf("RM-%s%04d", now.Format("2006"), count+1)
}

func (s *PatientService) Create(req *dto.CreatePatientRequest) (*models.Patient, error) {
	existing, _ := s.repo.FindByNIK(req.NIK)
	if existing != nil {
		return nil, fmt.Errorf("patient with NIK %s already exists", req.NIK)
	}

	patient := &models.Patient{
		MedicalRecordNumber: generateMRN(),
		Name:                req.Name,
		NIK:                 req.NIK,
		PlaceOfBirth:        req.PlaceOfBirth,
		DateOfBirth:         req.DateOfBirth,
		Gender:              req.Gender,
		Address:             req.Address,
		Phone:               req.Phone,
		Email:               req.Email,
		BloodType:           req.BloodType,
		Allergies:           req.Allergies,
		InsuranceName:       req.InsuranceName,
		InsuranceNumber:     req.InsuranceNumber,
		EmergencyContact:    req.EmergencyContact,
		EmergencyPhone:      req.EmergencyPhone,
	}

	err := s.repo.Create(patient)
	if err != nil {
		return nil, err
	}
	return patient, nil
}

func (s *PatientService) GetByID(id uint) (*models.Patient, error) {
	return s.repo.FindByID(id)
}

func (s *PatientService) GetAll(page, limit int, search string) ([]models.Patient, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	return s.repo.FindAll(page, limit, search)
}

func (s *PatientService) Update(id uint, req *dto.UpdatePatientRequest) (*models.Patient, error) {
	patient, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("patient not found")
	}

	if req.Name != "" {
		patient.Name = req.Name
	}
	if req.PlaceOfBirth != "" {
		patient.PlaceOfBirth = req.PlaceOfBirth
	}
	if req.DateOfBirth != "" {
		patient.DateOfBirth = req.DateOfBirth
	}
	if req.Gender != "" {
		patient.Gender = req.Gender
	}
	if req.Address != "" {
		patient.Address = req.Address
	}
	if req.Phone != "" {
		patient.Phone = req.Phone
	}
	if req.Email != "" {
		patient.Email = req.Email
	}
	if req.BloodType != "" {
		patient.BloodType = req.BloodType
	}
	if req.Allergies != "" {
		patient.Allergies = req.Allergies
	}
	if req.InsuranceName != "" {
		patient.InsuranceName = req.InsuranceName
	}
	if req.InsuranceNumber != "" {
		patient.InsuranceNumber = req.InsuranceNumber
	}
	if req.EmergencyContact != "" {
		patient.EmergencyContact = req.EmergencyContact
	}
	if req.EmergencyPhone != "" {
		patient.EmergencyPhone = req.EmergencyPhone
	}

	err = s.repo.Update(patient)
	return patient, err
}

func (s *PatientService) Delete(id uint) error {
	return s.repo.Delete(id)
}
