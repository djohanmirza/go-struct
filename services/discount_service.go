package services

import (
	"github.com/djohanmirza/go-struct/entity"
	"github.com/djohanmirza/go-struct/repository"
)

type DiscountService interface {
	CreateDiscount(discount *entity.Discount) error
	GetAllDiscounts() ([]entity.Discount, error)
	GetDiscountByID(id string) (*entity.Discount, error)
	UpdateDiscount(id string, discount *entity.Discount) error
	DeleteDiscount(id string) error
}

type discountServiceImpl struct {
	repo repository.DiscountRepository
}

func NewDiscountService(repo repository.DiscountRepository) DiscountService {
	return &discountServiceImpl{repo}
}

func (s *discountServiceImpl) CreateDiscount(discount *entity.Discount) error {
	return s.repo.Create(discount)
}

func (s *discountServiceImpl) GetAllDiscounts() ([]entity.Discount, error) {
	return s.repo.GetAll()
}

func (s *discountServiceImpl) GetDiscountByID(id string) (*entity.Discount, error) {
	return s.repo.GetByID(id)
}

func (s *discountServiceImpl) UpdateDiscount(id string, discount *entity.Discount) error {
	return s.repo.Update(id, discount)
}

func (s *discountServiceImpl) DeleteDiscount(id string) error {
	return s.repo.Delete(id)
}
