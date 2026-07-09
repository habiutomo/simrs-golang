package services

import "simrs-golang/internal/repositories"

type DashboardService struct {
	patientRepo     *repositories.PatientRepository
	doctorRepo      *repositories.DoctorRepository
	appointmentRepo *repositories.AppointmentRepository
	medicalRepo     *repositories.MedicalRecordRepository
	billingRepo     *repositories.BillingRepository
	roomRepo        *repositories.RoomRepository
}

func NewDashboardService(
	patientRepo *repositories.PatientRepository,
	doctorRepo *repositories.DoctorRepository,
	appointmentRepo *repositories.AppointmentRepository,
	medicalRepo *repositories.MedicalRecordRepository,
	billingRepo *repositories.BillingRepository,
	roomRepo *repositories.RoomRepository,
) *DashboardService {
	return &DashboardService{
		patientRepo:     patientRepo,
		doctorRepo:      doctorRepo,
		appointmentRepo: appointmentRepo,
		medicalRepo:     medicalRepo,
		billingRepo:     billingRepo,
		roomRepo:        roomRepo,
	}
}

type DashboardResponse struct {
	TotalPatients     int64   `json:"total_patients"`
	TotalDoctors      int64   `json:"total_doctors"`
	TodayAppointments int64   `json:"today_appointments"`
	TotalRecords      int64   `json:"total_medical_records"`
	TotalRevenue      float64 `json:"total_revenue"`
	TotalUnpaid       float64 `json:"total_unpaid"`
	OccupiedRooms     int64   `json:"occupied_rooms"`
}

func (s *DashboardService) GetDashboard() (*DashboardResponse, error) {
	totalPatients, _ := s.patientRepo.GetTotalCount()
	totalDoctors, _ := s.doctorRepo.GetTotalCount()
	todayAppointments, _ := s.appointmentRepo.GetTodayCount()
	totalRecords, _ := s.medicalRepo.GetTotalCount()
	totalRevenue, _ := s.billingRepo.GetTotalRevenue()
	totalUnpaid, _ := s.billingRepo.GetTotalUnpaid()
	occupiedRooms, _ := s.roomRepo.GetOccupiedCount()

	return &DashboardResponse{
		TotalPatients:     totalPatients,
		TotalDoctors:      totalDoctors,
		TodayAppointments: todayAppointments,
		TotalRecords:      totalRecords,
		TotalRevenue:      totalRevenue,
		TotalUnpaid:       totalUnpaid,
		OccupiedRooms:     occupiedRooms,
	}, nil
}
