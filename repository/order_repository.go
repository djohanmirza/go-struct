package repository

import (
	"github.com/djohanmirza/go-struct/entity"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *entity.Orders) error
	GetAll() ([]entity.Orders, error)
	GetByID(orderID string) (*entity.Orders, error)
	Delete(orderID string) error
}

type orderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepositoryImpl{db}
}

func (r *orderRepositoryImpl) Create(order *entity.Orders) error {
	return r.db.Create(order).Error
}

func (r *orderRepositoryImpl) GetAll() ([]entity.Orders, error) {
	var orders []entity.Orders
	err := r.db.Preload("OrderItems").Find(&orders).Error
	return orders, err
}

func (r *orderRepositoryImpl) GetByID(orderID string) (*entity.Orders, error) {
	var order entity.Orders
	err := r.db.Preload("OrderItems").First(&order, "order_id = ?", orderID).Error
	return &order, err
}

func (r *orderRepositoryImpl) Delete(orderID string) error {
	return r.db.Delete(&entity.Orders{}, "order_id = ?", orderID).Error
}
