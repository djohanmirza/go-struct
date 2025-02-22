package controller

import (
	"github.com/djohanmirza/go-struct/entity"
	"github.com/djohanmirza/go-struct/repository"
)

type CategoryController struct {
	repo repository.CategoryRepository
}

func (controller *CategoryController) create(category entity.Category) (entity.Category, error) {
	return controller.repo.Create(category)
}
