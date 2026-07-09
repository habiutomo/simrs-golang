package services

import (
	"fmt"

	"simrs-golang/internal/dto"
	"simrs-golang/internal/models"
	"simrs-golang/internal/repositories"
)

type AppointmentService struct {
	repo *repositories.AppointmentRepository
}

func NewAppointmentService(repo *repositories.AppointmentRepository) *AppointmentService {
	return &AppointmentService{repo: repo}
}

func (s *AppointmentService) Create(req *dto.CreateAppointmentRequest) (*models.Appointment, error) {
	appointment := &models.Appointment{
		PatientID:       req.PatientID,
		DoctorID:        req.DoctorID,
		AppointmentDate: req.Date,
		StartTime:       req.StartTime,
		EndTime:         req.EndTime,
		Status:          "scheduled",
		Complaint:       req.Complaint,
	}

	err := s.repo.Create(appointment)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByID(appointment.ID)
}

func (s *AppointmentService) GetByID(id uint) (*models.Appointment, error) {
	return s.repo.FindByID(id)
}

func (s *AppointmentService) GetAll(page, limit int, status string) ([]models.Appointment, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	return s.repo.FindAll(page, limit, status)
}

func (s *AppointmentService) GetByPatient(patientID uint) ([]models.Appointment, error) {
	return s.repo.FindByPatientID(patientID)
}

func (s *AppointmentService) GetByDoctor(doctorID uint, date string) ([]models.Appointment, error) {
	return s.repo.FindByDoctorID(doctorID, date)
}

func (s *AppointmentService) UpdateStatus(id uint, status string) (*models.Appointment, error) {
	appointment, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("appointment not found")
	}

	validStatuses := map[string]bool{
		"scheduled": true, "completed": true, "cancelled": true, "no_show": true,
	}
	if !validStatuses[status] {
		return nil, fmt.Errorf("invalid status: %s", status)
	}

	appointment.Status = status
	err = s.repo.Update(appointment)
	return appointment, err
}

func (s *AppointmentService) GetTodayCount() (int64, error) {
	return s.repo.GetTodayCount()
}
