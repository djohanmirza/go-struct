package repository

import (
	"github.com/djohanmirza/go-struct/entity"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *entity.Customer) error
	GetAll() ([]entity.Customer, error)
	GetByID(id string) (*entity.Customer, error)
	Update(id string, customer *entity.Customer) error
	Delete(id string) error
}

type CustomerRepositoryImpl struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &CustomerRepositoryImpl{db}
}

func (c *CustomerRepositoryImpl) Create(customer *entity.Customer) error {
	return c.db.Create(customer).Error
}

func (c *CustomerRepositoryImpl) GetAll() ([]entity.Customer, error) {
	var customers []entity.Customer
	err := c.db.Find(&customers).Error
	return customers, err
}

func (c *CustomerRepositoryImpl) GetByID(id string) (*entity.Customer, error) {
	var customer entity.Customer
	err := c.db.First(&customer, "customer_id = ?", id).Error
	return &customer, err
}

func (c *CustomerRepositoryImpl) Update(id string, customer *entity.Customer) error {
	return c.db.Model(&entity.Customer{}).Where("customer_id = ?", id).Updates(customer).Error
}

func (c *CustomerRepositoryImpl) Delete(id string) error {
	return c.db.Delete(&entity.Customer{}, "customer_id = ?", id).Error
}
