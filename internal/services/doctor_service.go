package services

import (
	"fmt"

	"simrs-golang/internal/database"
	"simrs-golang/internal/dto"
	"simrs-golang/internal/models"
	"simrs-golang/internal/repositories"
)

type DoctorService struct {
	repo         *repositories.DoctorRepository
	userRepo     *repositories.UserRepository
	deptRepo     *repositories.DepartmentRepository
}

func NewDoctorService(repo *repositories.DoctorRepository, userRepo *repositories.UserRepository, deptRepo *repositories.DepartmentRepository) *DoctorService {
	return &DoctorService{repo: repo, userRepo: userRepo, deptRepo: deptRepo}
}

func (s *DoctorService) Create(req *dto.CreateDoctorRequest) (*models.Doctor, error) {
	existingUser, _ := s.userRepo.FindByUsername(req.Username)
	if existingUser != nil {
		return nil, fmt.Errorf("username %s already exists", req.Username)
	}

	hashedPassword, err := database.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: req.Username,
		Password: hashedPassword,
		FullName: req.FullName,
		Email:    req.Email,
		Phone:    req.Phone,
		Role:     "doctor",
		IsActive: true,
	}
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	doctor := &models.Doctor{
		UserID:          user.ID,
		DepartmentID:    req.DepartmentID,
		LicenseNumber:   req.LicenseNumber,
		Specialization:  req.Specialization,
		ConsultationFee: req.ConsultationFee,
		IsAvailable:     true,
	}

	if err := s.repo.Create(doctor); err != nil {
		return nil, err
	}

	return s.repo.FindByID(doctor.ID)
}

func (s *DoctorService) GetByID(id uint) (*models.Doctor, error) {
	return s.repo.FindByID(id)
}

func (s *DoctorService) GetAll(page, limit int) ([]models.Doctor, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	return s.repo.FindAll(page, limit)
}

func (s *DoctorService) GetByDepartment(deptID uint) ([]models.Doctor, error) {
	return s.repo.FindByDepartmentID(deptID)
}

func (s *DoctorService) GetAllAvailable() ([]models.Doctor, error) {
	return s.repo.FindAllAvailable()
}

func (s *DoctorService) Delete(id uint) error {
	doctor, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return s.userRepo.Delete(doctor.UserID)
}

func (s *DoctorService) AddSchedule(req *dto.UpdateDoctorScheduleRequest) (*models.DoctorSchedule, error) {
	schedule := &models.DoctorSchedule{
		DoctorID:  req.DoctorID,
		DayOfWeek: req.DayOfWeek,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		IsActive:  true,
	}
	err := s.repo.CreateSchedule(schedule)
	return schedule, err
}

func (s *DoctorService) GetSchedules(doctorID uint) ([]models.DoctorSchedule, error) {
	return s.repo.FindSchedulesByDoctor(doctorID)
}
