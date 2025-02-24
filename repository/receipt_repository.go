package repository

import (
	"github.com/djohanmirza/go-struct/entity"
	"gorm.io/gorm"
)

type ReceiptRepository interface {
	Create(receipt *entity.Receipt) error
	GetAll() ([]entity.Receipt, error)
	GetByID(receiptID string) (*entity.Receipt, error)
	Delete(receiptID string) error
}

type receiptRepositoryImpl struct {
	db *gorm.DB
}

func NewReceiptRepository(db *gorm.DB) ReceiptRepository {
	return &receiptRepositoryImpl{db}
}

func (r *receiptRepositoryImpl) Create(receipt *entity.Receipt) error {
	return r.db.Create(receipt).Error
}

func (r *receiptRepositoryImpl) GetAll() ([]entity.Receipt, error) {
	var receipts []entity.Receipt
	err := r.db.Find(&receipts).Error
	return receipts, err
}

func (r *receiptRepositoryImpl) GetByID(receiptID string) (*entity.Receipt, error) {
	var receipt entity.Receipt
	err := r.db.First(&receipt, "receipt_id = ?", receiptID).Error
	return &receipt, err
}

func (r *receiptRepositoryImpl) Delete(receiptID string) error {
	return r.db.Delete(&entity.Receipt{}, "receipt_id = ?", receiptID).Error
}
