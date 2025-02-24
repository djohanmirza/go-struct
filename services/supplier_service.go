package services

import (
	"github.com/djohanmirza/go-struct/entity"
	"github.com/djohanmirza/go-struct/repository"
)

type SupplierService interface {
	CreateSupplier(supplier *entity.Supplier) error
	GetAllSuppliers() ([]entity.Supplier, error)
	GetSupplierByID(id uint64) (*entity.Supplier, error)
	UpdateSupplier(id uint64, supplier *entity.Supplier) error
	DeleteSupplier(id uint64) error
}

type supplierServiceImpl struct {
	repo repository.SupplierRepository
}

func NewSupplierService(repo repository.SupplierRepository) SupplierService {
	return &supplierServiceImpl{repo}
}

func (s *supplierServiceImpl) CreateSupplier(supplier *entity.Supplier) error {
	return s.repo.Create(supplier)
}

func (s *supplierServiceImpl) GetAllSuppliers() ([]entity.Supplier, error) {
	return s.repo.GetAll()
}

func (s *supplierServiceImpl) GetSupplierByID(id uint64) (*entity.Supplier, error) {
	return s.repo.GetByID(id)
}

func (s *supplierServiceImpl) UpdateSupplier(id uint64, supplier *entity.Supplier) error {
	return s.repo.Update(id, supplier)
}

func (s *supplierServiceImpl) DeleteSupplier(id uint64) error {
	return s.repo.Delete(id)
}
