package repository

import (
	"github.com/djohanmirza/go-struct/entity"
	"gorm.io/gorm"
)

type SupplierRepository interface {
	Create(supplier *entity.Supplier) error
	GetAll() ([]entity.Supplier, error)
	GetByID(id uint64) (*entity.Supplier, error)
	Update(id uint64, supplier *entity.Supplier) error
	Delete(id uint64) error
}

type supplierRepositoryImpl struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) SupplierRepository {
	return &supplierRepositoryImpl{db}
}

func (r *supplierRepositoryImpl) Create(supplier *entity.Supplier) error {
	return r.db.Create(supplier).Error
}

func (r *supplierRepositoryImpl) GetAll() ([]entity.Supplier, error) {
	var suppliers []entity.Supplier
	err := r.db.Find(&suppliers).Error
	return suppliers, err
}

func (r *supplierRepositoryImpl) GetByID(id uint64) (*entity.Supplier, error) {
	var supplier entity.Supplier
	err := r.db.First(&supplier, "id = ?", id).Error
	return &supplier, err
}

func (r *supplierRepositoryImpl) Update(id uint64, supplier *entity.Supplier) error {
	return r.db.Model(&entity.Supplier{}).Where("id = ?", id).Updates(supplier).Error
}

func (r *supplierRepositoryImpl) Delete(id uint64) error {
	return r.db.Delete(&entity.Supplier{}, "id = ?", id).Error
}
