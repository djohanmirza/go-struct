package services

import (
	"github.com/djohanmirza/go-struct/entity"
	"github.com/djohanmirza/go-struct/repository"
)

type TaxService interface {
	CreateTax(tax *entity.Tax) error
	GetAllTaxes() ([]entity.Tax, error)
	GetTaxByID(id string) (*entity.Tax, error)
	UpdateTax(id string, tax *entity.Tax) error
	DeleteTax(id string) error
}

type taxServiceImpl struct {
	repo repository.TaxRepository
}

func NewTaxService(repo repository.TaxRepository) TaxService {
	return &taxServiceImpl{repo}
}

func (s *taxServiceImpl) CreateTax(tax *entity.Tax) error {
	return s.repo.Create(tax)
}

func (s *taxServiceImpl) GetAllTaxes() ([]entity.Tax, error) {
	return s.repo.GetAll()
}

func (s *taxServiceImpl) GetTaxByID(id string) (*entity.Tax, error) {
	return s.repo.GetByID(id)
}

func (s *taxServiceImpl) UpdateTax(id string, tax *entity.Tax) error {
	return s.repo.Update(id, tax)
}

func (s *taxServiceImpl) DeleteTax(id string) error {
	return s.repo.Delete(id)
}
