package services

import (
	"fmt"
	"time"

	"simrs-golang/internal/dto"
	"simrs-golang/internal/models"
	"simrs-golang/internal/repositories"
)

type RoomService struct {
	roomRepo       *repositories.RoomRepository
	patientRepo    *repositories.PatientRepository
	doctorRepo     *repositories.DoctorRepository
}

func NewRoomService(roomRepo *repositories.RoomRepository, patientRepo *repositories.PatientRepository, doctorRepo *repositories.DoctorRepository) *RoomService {
	return &RoomService{roomRepo: roomRepo, patientRepo: patientRepo, doctorRepo: doctorRepo}
}

func (s *RoomService) GetAll() ([]models.Room, error) {
	return s.roomRepo.FindAll()
}

func (s *RoomService) GetAvailable() ([]models.Room, error) {
	return s.roomRepo.FindAvailable()
}

func (s *RoomService) AdmitPatient(req *dto.AdmitPatientRequest) (*models.Inpatient, error) {
	room, err := s.roomRepo.FindByID(req.RoomID)
	if err != nil {
		return nil, fmt.Errorf("room not found")
	}
	if room.Status != "available" {
		return nil, fmt.Errorf("room is not available")
	}

	inpatient := &models.Inpatient{
		PatientID:     req.PatientID,
		DoctorID:      req.DoctorID,
		RoomID:        req.RoomID,
		AdmissionDate: time.Now(),
		Diagnosis:     req.Diagnosis,
		Status:        "admitted",
	}

	if err := s.roomRepo.CreateInpatient(inpatient); err != nil {
		return nil, err
	}

	room.Status = "occupied"
	s.roomRepo.Update(room)

	return s.roomRepo.FindInpatientByID(inpatient.ID)
}

func (s *RoomService) DischargePatient(inpatientID uint) (*models.Inpatient, error) {
	inpatient, err := s.roomRepo.FindInpatientByID(inpatientID)
	if err != nil {
		return nil, fmt.Errorf("inpatient not found")
	}

	now := time.Now()
	inpatient.DischargeDate = &now
	inpatient.Status = "discharged"

	if err := s.roomRepo.UpdateInpatient(inpatient); err != nil {
		return nil, err
	}

	room, _ := s.roomRepo.FindByID(inpatient.RoomID)
	if room != nil {
		room.Status = "available"
		s.roomRepo.Update(room)
	}

	return inpatient, nil
}

func (s *RoomService) GetActiveInpatients() ([]models.Inpatient, error) {
	return s.roomRepo.FindActiveInpatients()
}

func (s *RoomService) GetInpatientByID(id uint) (*models.Inpatient, error) {
	return s.roomRepo.FindInpatientByID(id)
}
