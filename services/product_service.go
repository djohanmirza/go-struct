package services

import (
	"github.com/djohanmirza/go-struct/entity"
	"github.com/djohanmirza/go-struct/repository"
)

type ProductService interface {
	CreateProduct(product *entity.Product) error
	GetAllProducts() ([]entity.Product, error)
	GetProductByID(id string) (*entity.Product, error)
	UpdateProduct(id string, product *entity.Product) error
	DeleteProduct(id string) error
}

type productServiceImpl struct {
	Repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productServiceImpl{repo}
}
func (s *productServiceImpl) CreateProduct(product *entity.Product) error {
	return s.Repo.Create(product)
}

func (s *productServiceImpl) GetAllProducts() ([]entity.Product, error) {
	return s.Repo.GetAll()
}

func (s *productServiceImpl) GetProductByID(id string) (*entity.Product, error) {
	return s.Repo.GetByID(id)
}

func (s *productServiceImpl) UpdateProduct(id string, product *entity.Product) error {
	return s.Repo.Update(id, product)
}

func (s *productServiceImpl) DeleteProduct(id string) error {
	return s.Repo.Delete(id)
}
