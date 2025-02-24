package repository

import (
	"github.com/djohanmirza/go-struct/entity"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *entity.Payment) error
	GetAll() ([]entity.Payment, error)
	GetByID(paymentID string) (*entity.Payment, error)
	Delete(paymentID string) error
}

type paymentRepositoryImpl struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepositoryImpl{db}
}

func (r *paymentRepositoryImpl) Create(payment *entity.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepositoryImpl) GetAll() ([]entity.Payment, error) {
	var payments []entity.Payment
	err := r.db.Find(&payments).Error
	return payments, err
}

func (r *paymentRepositoryImpl) GetByID(paymentID string) (*entity.Payment, error) {
	var payment entity.Payment
	err := r.db.First(&payment, "payment_id = ?", paymentID).Error
	return &payment, err
}

func (r *paymentRepositoryImpl) Delete(paymentID string) error {
	return r.db.Delete(&entity.Payment{}, "payment_id = ?", paymentID).Error
}
