package service

import (
	"rest-api/models"
	"rest-api/repository"
)

type EmployeeService interface {
	GetAllEmployees() ([]models.Employee, error)
	GetEmployeeByID(id int) (*models.Employee, error)
	AddEmployee(employee *models.Employee) error
	UpdateEmployee(employee *models.Employee) error
	DeleteEmployee(id int) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) GetAllEmployees() ([]models.Employee, error) {
	return s.repo.FindAll()
}

func (s *employeeService) GetEmployeeByID(id int) (*models.Employee, error) {
	return s.repo.FindByID(id)
}

func (s *employeeService) AddEmployee(employee *models.Employee) error {
	return s.repo.Create(employee)
}

func (s *employeeService) UpdateEmployee(employee *models.Employee) error {
	return s.repo.Update(employee)
}

func (s *employeeService) DeleteEmployee(id int) error {
	return s.repo.Delete(id)
}
