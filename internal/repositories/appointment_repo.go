package repositories

import (
	"simrs-golang/internal/database"
	"simrs-golang/internal/models"
)

type AppointmentRepository struct{}

func NewAppointmentRepository() *AppointmentRepository {
	return &AppointmentRepository{}
}

func (r *AppointmentRepository) Create(appointment *models.Appointment) error {
	return database.DB.Create(appointment).Error
}

func (r *AppointmentRepository) FindByID(id uint) (*models.Appointment, error) {
	var appointment models.Appointment
	err := database.DB.Preload("Patient").Preload("Doctor.User").Preload("Doctor.Department").First(&appointment, id).Error
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

func (r *AppointmentRepository) FindAll(page, limit int, status string) ([]models.Appointment, int64, error) {
	var appointments []models.Appointment
	var total int64

	query := database.DB.Model(&models.Appointment{})
	if status != "" {
		query = query.Where("status = ?", status)
	}
	query.Count(&total)

	offset := (page - 1) * limit
	err := query.Preload("Patient").Preload("Doctor.User").Preload("Doctor.Department").
		Offset(offset).Limit(limit).Order("appointment_date DESC, start_time DESC").Find(&appointments).Error
	return appointments, total, err
}

func (r *AppointmentRepository) FindByPatientID(patientID uint) ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := database.DB.Preload("Patient").Preload("Doctor.User").Preload("Doctor.Department").
		Where("patient_id = ?", patientID).Order("appointment_date DESC").Find(&appointments).Error
	return appointments, err
}

func (r *AppointmentRepository) FindByDoctorID(doctorID uint, date string) ([]models.Appointment, error) {
	var appointments []models.Appointment
	query := database.DB.Preload("Patient").Preload("Doctor.User").Where("doctor_id = ?", doctorID)
	if date != "" {
		query = query.Where("appointment_date = ?", date)
	}
	err := query.Order("start_time ASC").Find(&appointments).Error
	return appointments, err
}

func (r *AppointmentRepository) FindByDate(date string) ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := database.DB.Preload("Patient").Preload("Doctor.User").Preload("Doctor.Department").
		Where("appointment_date = ?", date).Order("start_time ASC").Find(&appointments).Error
	return appointments, err
}

func (r *AppointmentRepository) Update(appointment *models.Appointment) error {
	return database.DB.Save(appointment).Error
}

func (r *AppointmentRepository) GetTodayCount() (int64, error) {
	var total int64
	err := database.DB.Model(&models.Appointment{}).Where("appointment_date = CURDATE()").Count(&total).Error
	return total, err
}

func (r *AppointmentRepository) GetTotalCount() (int64, error) {
	var total int64
	err := database.DB.Model(&models.Appointment{}).Count(&total).Error
	return total, err
}

func (r *AppointmentRepository) GetCompletedCount() (int64, error) {
	var total int64
	err := database.DB.Model(&models.Appointment{}).Where("status = ?", "completed").Count(&total).Error
	return total, err
}
