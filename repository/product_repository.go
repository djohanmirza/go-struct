package repository

import (
	"github.com/djohanmirza/go-struct/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *entity.Product) error
	GetAll() ([]entity.Product, error)
	GetByID(id string) (*entity.Product, error)
	Update(id string, product *entity.Product) error
	Delete(id string) error
}

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryImpl{db}
}

func (r *productRepositoryImpl) Create(product *entity.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepositoryImpl) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepositoryImpl) GetByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := r.db.First(&product, "product_id = ?", id).Error
	return &product, err
}

func (r *productRepositoryImpl) Update(id string, product *entity.Product) error {
	return r.db.Model(&entity.Product{}).Where("product_id = ?", id).Updates(product).Error
}

func (r *productRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&entity.Product{}, "product_id = ?", id).Error
}
