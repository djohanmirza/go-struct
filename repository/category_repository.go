package repository

import "github.com/djohanmirza/go-struct/entity"

type CategoryRepository interface {
	Create(category entity.Category) (entity.Category, error)
	FindById(id uint64) (entity.Category, error)
	FindAll() ([]entity.Category, error)
}
type CategoryRepositoryImpl struct {
	CategoryRepository CategoryRepository
}

func (repo *CategoryRepositoryImpl) Create(category entity.Category) (entity.Category, error) {
	return repo.CategoryRepository.Create(category)
}

func (repo *CategoryRepositoryImpl) FindById(id uint64) (entity.Category, error) {
	return repo.CategoryRepository.FindById(id)
}

func (repo *CategoryRepositoryImpl) FindAll() ([]entity.Category, error) {
	return repo.CategoryRepository.FindAll()
}
