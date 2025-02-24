package services

import (
	"github.com/djohanmirza/go-struct/entity"
	"github.com/djohanmirza/go-struct/repository"
)

type CustomerService interface {
	CreateCustomer(customer *entity.Customer) error
	GetAllCustomers() ([]entity.Customer, error)
	GetCustomerByID(id string) (*entity.Customer, error)
	UpdateCustomer(id string, customer *entity.Customer) error
	DeleteCustomer(id string) error
}

type customerServiceImpl struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerServiceImpl{repo}
}

func (s *customerServiceImpl) CreateCustomer(customer *entity.Customer) error {
	return s.customerRepo.Create(customer)
}

func (s *customerServiceImpl) GetAllCustomers() ([]entity.Customer, error) {
	return s.customerRepo.GetAll()
}

func (s *customerServiceImpl) GetCustomerByID(id string) (*entity.Customer, error) {
	return s.customerRepo.GetByID(id)
}

func (s *customerServiceImpl) UpdateCustomer(id string, customer *entity.Customer) error {
	return s.customerRepo.Update(id, customer)
}

func (s *customerServiceImpl) DeleteCustomer(id string) error {
	return s.customerRepo.Delete(id)
}
