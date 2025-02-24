package repository

import (
	"github.com/djohanmirza/go-struct/entity"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(employee *entity.Employee) error
	GetAll() ([]entity.Employee, error)
	GetByID(id string) (*entity.Employee, error)
	Update(id string, employee *entity.Employee) error
	Delete(id string) error
}

type employeeRepositoryImpl struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepositoryImpl{db}
}

func (r *employeeRepositoryImpl) Create(employee *entity.Employee) error {
	return r.db.Create(employee).Error
}

func (r *employeeRepositoryImpl) GetAll() ([]entity.Employee, error) {
	var employees []entity.Employee
	err := r.db.Find(&employees).Error
	return employees, err
}

func (r *employeeRepositoryImpl) GetByID(id string) (*entity.Employee, error) {
	var employee entity.Employee
	err := r.db.First(&employee, "employee_id = ?", id).Error
	return &employee, err
}

func (r *employeeRepositoryImpl) Update(id string, employee *entity.Employee) error {
	return r.db.Model(&entity.Employee{}).Where("employee_id = ?", id).Updates(employee).Error
}

func (r *employeeRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&entity.Employee{}, "employee_id = ?", id).Error
}
