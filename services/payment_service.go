package services

import (
	"github.com/djohanmirza/go-struct/entity"
	"github.com/djohanmirza/go-struct/repository"
)

type PaymentService interface {
	CreatePayment(payment *entity.Payment) error
	GetAllPayments() ([]entity.Payment, error)
	GetPaymentByID(paymentID string) (*entity.Payment, error)
	DeletePayment(paymentID string) error
}

type paymentServiceImpl struct {
	repo repository.PaymentRepository
}

func NewPaymentService(repo repository.PaymentRepository) PaymentService {
	return &paymentServiceImpl{repo}
}

func (s *paymentServiceImpl) CreatePayment(payment *entity.Payment) error {
	return s.repo.Create(payment)
}

func (s *paymentServiceImpl) GetAllPayments() ([]entity.Payment, error) {
	return s.repo.GetAll()
}

func (s *paymentServiceImpl) GetPaymentByID(paymentID string) (*entity.Payment, error) {
	return s.repo.GetByID(paymentID)
}

func (s *paymentServiceImpl) DeletePayment(paymentID string) error {
	return s.repo.Delete(paymentID)
}
