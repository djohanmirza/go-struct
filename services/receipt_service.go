package services

import (
	"github.com/djohanmirza/go-struct/entity"
	"github.com/djohanmirza/go-struct/repository"
)

type ReceiptService interface {
	CreateReceipt(receipt *entity.Receipt) error
	GetAllReceipts() ([]entity.Receipt, error)
	GetReceiptByID(receiptID string) (*entity.Receipt, error)
	DeleteReceipt(receiptID string) error
}

type receiptServiceImpl struct {
	repo repository.ReceiptRepository
}

func NewReceiptService(repo repository.ReceiptRepository) ReceiptService {
	return &receiptServiceImpl{repo}
}

func (s *receiptServiceImpl) CreateReceipt(receipt *entity.Receipt) error {
	return s.repo.Create(receipt)
}

func (s *receiptServiceImpl) GetAllReceipts() ([]entity.Receipt, error) {
	return s.repo.GetAll()
}

func (s *receiptServiceImpl) GetReceiptByID(receiptID string) (*entity.Receipt, error) {
	return s.repo.GetByID(receiptID)
}

func (s *receiptServiceImpl) DeleteReceipt(receiptID string) error {
	return s.repo.Delete(receiptID)
}
