package services

import (
	"github.com/djohanmirza/go-struct/entity"
	"github.com/djohanmirza/go-struct/repository"
)

type EmployeeService interface {
	CreateEmployee(employee *entity.Employee) error
	GetAllEmployees() ([]entity.Employee, error)
	GetEmployeeByID(id string) (*entity.Employee, error)
	UpdateEmployee(id string, employee *entity.Employee) error
	DeleteEmployee(id string) error
}

type employeeServiceImpl struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeServiceImpl{repo}
}

func (s *employeeServiceImpl) CreateEmployee(employee *entity.Employee) error {
	return s.repo.Create(employee)
}

func (s *employeeServiceImpl) GetAllEmployees() ([]entity.Employee, error) {
	return s.repo.GetAll()
}

func (s *employeeServiceImpl) GetEmployeeByID(id string) (*entity.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *employeeServiceImpl) UpdateEmployee(id string, employee *entity.Employee) error {
	return s.repo.Update(id, employee)
}

func (s *employeeServiceImpl) DeleteEmployee(id string) error {
	return s.repo.Delete(id)
}
