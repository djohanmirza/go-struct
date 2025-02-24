package services

import (
	"github.com/djohanmirza/go-struct/entity"
	"github.com/djohanmirza/go-struct/repository"
)

type CategoryService interface {
	CreateCategory(category *entity.Category) error
	GetAllCategories() ([]entity.Category, error)
	GetCategoryByID(id uint64) (*entity.Category, error)
	UpdateCategory(id uint64, category *entity.Category) error
	DeleteCategory(id uint64) error
}

type categoryServiceImpl struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryServiceImpl{repo}
}

func (s *categoryServiceImpl) CreateCategory(category *entity.Category) error {
	return s.repo.Create(category)
}

func (s *categoryServiceImpl) GetAllCategories() ([]entity.Category, error) {
	return s.repo.GetAll()
}

func (s *categoryServiceImpl) GetCategoryByID(id uint64) (*entity.Category, error) {
	return s.repo.GetByID(id)
}

func (s *categoryServiceImpl) UpdateCategory(id uint64, category *entity.Category) error {
	return s.repo.Update(id, category)
}

func (s *categoryServiceImpl) DeleteCategory(id uint64) error {
	return s.repo.Delete(id)
}
