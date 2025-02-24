package repository

import (
	"github.com/djohanmirza/go-struct/entity"
	"gorm.io/gorm"
)

type DiscountRepository interface {
	Create(discount *entity.Discount) error
	GetAll() ([]entity.Discount, error)
	GetByID(id string) (*entity.Discount, error)
	Update(id string, discount *entity.Discount) error
	Delete(id string) error
}

type discountRepositoryImpl struct {
	db *gorm.DB
}

func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	return &discountRepositoryImpl{db}
}

func (r *discountRepositoryImpl) Create(discount *entity.Discount) error {
	return r.db.Create(discount).Error
}

func (r *discountRepositoryImpl) GetAll() ([]entity.Discount, error) {
	var discounts []entity.Discount
	err := r.db.Find(&discounts).Error
	return discounts, err
}

func (r *discountRepositoryImpl) GetByID(id string) (*entity.Discount, error) {
	var discount entity.Discount
	err := r.db.First(&discount, "discount_id = ?", id).Error
	return &discount, err
}

func (r *discountRepositoryImpl) Update(id string, discount *entity.Discount) error {
	return r.db.Model(&entity.Discount{}).Where("discount_id = ?", id).Updates(discount).Error
}

func (r *discountRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&entity.Discount{}, "discount_id = ?", id).Error
}
