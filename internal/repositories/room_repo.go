package repositories

import (
	"simrs-golang/internal/database"
	"simrs-golang/internal/models"
)

type RoomRepository struct{}

func NewRoomRepository() *RoomRepository {
	return &RoomRepository{}
}

func (r *RoomRepository) FindAll() ([]models.Room, error) {
	var rooms []models.Room
	err := database.DB.Order("number ASC").Find(&rooms).Error
	return rooms, err
}

func (r *RoomRepository) FindByID(id uint) (*models.Room, error) {
	var room models.Room
	err := database.DB.First(&room, id).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepository) FindAvailable() ([]models.Room, error) {
	var rooms []models.Room
	err := database.DB.Where("status = ?", "available").Order("number ASC").Find(&rooms).Error
	return rooms, err
}

func (r *RoomRepository) Update(room *models.Room) error {
	return database.DB.Save(room).Error
}

func (r *RoomRepository) CreateInpatient(inpatient *models.Inpatient) error {
	return database.DB.Create(inpatient).Error
}

func (r *RoomRepository) FindInpatientByID(id uint) (*models.Inpatient, error) {
	var inpatient models.Inpatient
	err := database.DB.Preload("Patient").Preload("Doctor.User").Preload("Room").
		First(&inpatient, id).Error
	return &inpatient, err
}

func (r *RoomRepository) FindActiveInpatients() ([]models.Inpatient, error) {
	var inpatients []models.Inpatient
	err := database.DB.Preload("Patient").Preload("Doctor.User").Preload("Room").
		Where("status = ?", "admitted").Order("admission_date DESC").Find(&inpatients).Error
	return inpatients, err
}

func (r *RoomRepository) UpdateInpatient(inpatient *models.Inpatient) error {
	return database.DB.Save(inpatient).Error
}

func (r *RoomRepository) GetOccupiedCount() (int64, error) {
	var total int64
	database.DB.Model(&models.Room{}).Where("status = ?", "occupied").Count(&total)
	return total, nil
}
