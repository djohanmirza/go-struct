package repository

import (
	"github.com/djohanmirza/go-struct/entity"
	"gorm.io/gorm"
)

type TaxRepository interface {
	Create(tax *entity.Tax) error
	GetAll() ([]entity.Tax, error)
	GetByID(id string) (*entity.Tax, error)
	Update(id string, tax *entity.Tax) error
	Delete(id string) error
}

type taxRepositoryImpl struct {
	db *gorm.DB
}

func NewTaxRepository(db *gorm.DB) TaxRepository {
	return &taxRepositoryImpl{db}
}

func (r *taxRepositoryImpl) Create(tax *entity.Tax) error {
	return r.db.Create(tax).Error
}

func (r *taxRepositoryImpl) GetAll() ([]entity.Tax, error) {
	var taxes []entity.Tax
	err := r.db.Find(&taxes).Error
	return taxes, err
}

func (r *taxRepositoryImpl) GetByID(id string) (*entity.Tax, error) {
	var tax entity.Tax
	err := r.db.First(&tax, "tax_id = ?", id).Error
	return &tax, err
}

func (r *taxRepositoryImpl) Update(id string, tax *entity.Tax) error {
	return r.db.Model(&entity.Tax{}).Where("tax_id = ?", id).Updates(tax).Error
}

func (r *taxRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&entity.Tax{}, "tax_id = ?", id).Error
}
