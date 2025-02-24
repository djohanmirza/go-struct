package services

import (
	"github.com/djohanmirza/go-struct/entity"
	"github.com/djohanmirza/go-struct/repository"
)

type OrderService interface {
	CreateOrder(order *entity.Orders) error
	GetAllOrders() ([]entity.Orders, error)
	GetOrderByID(orderID string) (*entity.Orders, error)
	DeleteOrder(orderID string) error
}

type orderServiceImpl struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderServiceImpl{repo}
}

func (s *orderServiceImpl) CreateOrder(order *entity.Orders) error {
	return s.repo.Create(order)
}

func (s *orderServiceImpl) GetAllOrders() ([]entity.Orders, error) {
	return s.repo.GetAll()
}

func (s *orderServiceImpl) GetOrderByID(orderID string) (*entity.Orders, error) {
	return s.repo.GetByID(orderID)
}

func (s *orderServiceImpl) DeleteOrder(orderID string) error {
	return s.repo.Delete(orderID)
}
