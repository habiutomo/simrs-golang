package services

import (
	"fmt"

	"simrs-golang/internal/dto"
	"simrs-golang/internal/models"
	"simrs-golang/internal/repositories"
)

type MedicalRecordService struct {
	recordRepo      *repositories.MedicalRecordRepository
	appointmentRepo *repositories.AppointmentRepository
	medicationRepo  *repositories.MedicationRepository
}

func NewMedicalRecordService(recordRepo *repositories.MedicalRecordRepository, appointmentRepo *repositories.AppointmentRepository, medicationRepo *repositories.MedicationRepository) *MedicalRecordService {
	return &MedicalRecordService{recordRepo: recordRepo, appointmentRepo: appointmentRepo, medicationRepo: medicationRepo}
}

func (s *MedicalRecordService) Create(req *dto.CreateMedicalRecordRequest) (*models.MedicalRecord, error) {
	record := &models.MedicalRecord{
		AppointmentID: req.AppointmentID,
		PatientID:     req.PatientID,
		DoctorID:      req.DoctorID,
		Diagnosis:     req.Diagnosis,
		Complaint:     req.Complaint,
		Examination:   req.Examination,
		Treatment:     req.Treatment,
		Notes:         req.Notes,
	}

	err := s.recordRepo.Create(record)
	if err != nil {
		return nil, err
	}

	if req.AppointmentID != nil {
		appointment, _ := s.appointmentRepo.FindByID(*req.AppointmentID)
		if appointment != nil {
			appointment.Status = "completed"
			s.appointmentRepo.Update(appointment)
		}
	}

	for _, p := range req.Prescriptions {
		prescription := &models.MedicalPrescription{
			MedicalRecordID: record.ID,
			MedicationID:    p.MedicationID,
			Dosage:          p.Dosage,
			Frequency:       p.Frequency,
			Duration:        p.Duration,
			Quantity:        p.Quantity,
			Notes:           p.Notes,
		}
		if err := s.recordRepo.CreatePrescription(prescription); err != nil {
			return nil, fmt.Errorf("failed to create prescription: %v", err)
		}

		s.medicationRepo.ReduceStock(p.MedicationID, p.Quantity)
	}

	return s.recordRepo.FindByID(record.ID)
}

func (s *MedicalRecordService) GetByID(id uint) (*models.MedicalRecord, error) {
	record, err := s.recordRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	prescriptions, _ := s.recordRepo.FindPrescriptionsByRecordID(id)
	record.Prescriptions = prescriptions
	return record, nil
}

func (s *MedicalRecordService) GetByPatient(patientID uint) ([]models.MedicalRecord, error) {
	records, err := s.recordRepo.FindByPatientID(patientID)
	if err != nil {
		return nil, err
	}
	for i := range records {
		prescriptions, _ := s.recordRepo.FindPrescriptionsByRecordID(records[i].ID)
		records[i].Prescriptions = prescriptions
	}
	return records, nil
}

func (s *MedicalRecordService) GetAll(page, limit int) ([]models.MedicalRecord, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	return s.recordRepo.FindAll(page, limit)
}

func (s *MedicalRecordService) AddPrescription(medicalRecordID uint, req *dto.CreatePrescriptionRequest) (*models.MedicalPrescription, error) {
	prescription := &models.MedicalPrescription{
		MedicalRecordID: medicalRecordID,
		MedicationID:    req.MedicationID,
		Dosage:          req.Dosage,
		Frequency:       req.Frequency,
		Duration:        req.Duration,
		Quantity:        req.Quantity,
		Notes:           req.Notes,
	}
	err := s.recordRepo.CreatePrescription(prescription)
	if err != nil {
		return nil, err
	}

	s.medicationRepo.ReduceStock(req.MedicationID, req.Quantity)
	return prescription, nil
}
