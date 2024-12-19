package repository

import (
	"rest-api/models"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindAll() ([]models.Employee, error)
	FindByID(id int) (*models.Employee, error)
	Create(employee *models.Employee) error
	Update(employee *models.Employee) error
	Delete(id int) error
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) FindAll() ([]models.Employee, error) {
	var employees []models.Employee
	err := r.db.Find(&employees).Error
	return employees, err
}

func (r *employeeRepository) FindByID(id int) (*models.Employee, error) {
	var employee models.Employee
	err := r.db.First(&employee, id).Error
	return &employee, err
}

func (r *employeeRepository) Create(employee *models.Employee) error {
	return r.db.Create(employee).Error
}

func (r *employeeRepository) Update(employee *models.Employee) error {
	return r.db.Save(employee).Error
}

func (r *employeeRepository) Delete(id int) error {
	return r.db.Delete(&models.Employee{}, id).Error
}
