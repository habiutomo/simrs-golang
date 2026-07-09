package repositories

import (
	"simrs-golang/internal/database"
	"simrs-golang/internal/models"
)

type DepartmentRepository struct{}

func NewDepartmentRepository() *DepartmentRepository {
	return &DepartmentRepository{}
}

func (r *DepartmentRepository) FindAll() ([]models.Department, error) {
	var departments []models.Department
	err := database.DB.Order("name ASC").Find(&departments).Error
	return departments, err
}

func (r *DepartmentRepository) FindByID(id uint) (*models.Department, error) {
	var department models.Department
	err := database.DB.First(&department, id).Error
	if err != nil {
		return nil, err
	}
	return &department, nil
}

func (r *DepartmentRepository) Create(department *models.Department) error {
	return database.DB.Create(department).Error
}

func (r *DepartmentRepository) Update(department *models.Department) error {
	return database.DB.Save(department).Error
}

func (r *DepartmentRepository) Delete(id uint) error {
	return database.DB.Delete(&models.Department{}, id).Error
}
