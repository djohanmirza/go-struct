package repository

import (
	"github.com/djohanmirza/go-struct/entity"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *entity.Category) error
	GetAll() ([]entity.Category, error)
	GetByID(id uint64) (*entity.Category, error)
	Update(id uint64, category *entity.Category) error
	Delete(id uint64) error
}

type categoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepositoryImpl{db}
}

func (r *categoryRepositoryImpl) Create(category *entity.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepositoryImpl) GetAll() ([]entity.Category, error) {
	var categories []entity.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepositoryImpl) GetByID(id uint64) (*entity.Category, error) {
	var category entity.Category
	err := r.db.First(&category, "id = ?", id).Error
	return &category, err
}

func (r *categoryRepositoryImpl) Update(id uint64, category *entity.Category) error {
	return r.db.Model(&entity.Category{}).Where("id = ?", id).Updates(category).Error
}

func (r *categoryRepositoryImpl) Delete(id uint64) error {
	return r.db.Delete(&entity.Category{}, "id = ?", id).Error
}
