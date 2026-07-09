package repositories

import (
	"simrs-golang/internal/database"
	"simrs-golang/internal/models"
)

type DoctorRepository struct{}

func NewDoctorRepository() *DoctorRepository {
	return &DoctorRepository{}
}

func (r *DoctorRepository) Create(doctor *models.Doctor) error {
	return database.DB.Create(doctor).Error
}

func (r *DoctorRepository) FindByID(id uint) (*models.Doctor, error) {
	var doctor models.Doctor
	err := database.DB.Preload("User").Preload("Department").First(&doctor, id).Error
	if err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (r *DoctorRepository) FindAll(page, limit int) ([]models.Doctor, int64, error) {
	var doctors []models.Doctor
	var total int64

	database.DB.Model(&models.Doctor{}).Count(&total)
	offset := (page - 1) * limit
	err := database.DB.Preload("User").Preload("Department").Offset(offset).Limit(limit).Find(&doctors).Error
	return doctors, total, err
}

func (r *DoctorRepository) FindByDepartmentID(deptID uint) ([]models.Doctor, error) {
	var doctors []models.Doctor
	err := database.DB.Preload("User").Preload("Department").Where("department_id = ?", deptID).Find(&doctors).Error
	return doctors, err
}

func (r *DoctorRepository) Update(doctor *models.Doctor) error {
	return database.DB.Save(doctor).Error
}

func (r *DoctorRepository) Delete(id uint) error {
	return database.DB.Delete(&models.Doctor{}, id).Error
}

func (r *DoctorRepository) GetTotalCount() (int64, error) {
	var total int64
	err := database.DB.Model(&models.Doctor{}).Count(&total).Error
	return total, err
}

func (r *DoctorRepository) FindAllAvailable() ([]models.Doctor, error) {
	var doctors []models.Doctor
	err := database.DB.Preload("User").Preload("Department").Where("is_available = ?", true).Find(&doctors).Error
	return doctors, err
}

func (r *DoctorRepository) CreateSchedule(schedule *models.DoctorSchedule) error {
	return database.DB.Create(schedule).Error
}

func (r *DoctorRepository) FindSchedulesByDoctor(doctorID uint) ([]models.DoctorSchedule, error) {
	var schedules []models.DoctorSchedule
	err := database.DB.Where("doctor_id = ?", doctorID).Order("day_of_week ASC, start_time ASC").Find(&schedules).Error
	return schedules, err
}

func (r *DoctorRepository) DeleteSchedule(id uint) error {
	return database.DB.Delete(&models.DoctorSchedule{}, id).Error
}
