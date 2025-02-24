package services

import (
	"github.com/djohanmirza/go-struct/entity"
	"github.com/djohanmirza/go-struct/repository"
)

type InventoryService interface {
	CreateInventory(inventory *entity.Inventory) error
	GetAllInventories() ([]entity.Inventory, error)
	GetInventoryByProductID(productID string) (*entity.Inventory, error)
	UpdateInventory(productID string, inventory *entity.Inventory) error
	DeleteInventory(productID string) error
}

type inventoryServiceImpl struct {
	repo repository.InventoryRepository
}

func NewInventoryService(repo repository.InventoryRepository) InventoryService {
	return &inventoryServiceImpl{repo}
}

func (s *inventoryServiceImpl) CreateInventory(inventory *entity.Inventory) error {
	return s.repo.Create(inventory)
}

func (s *inventoryServiceImpl) GetAllInventories() ([]entity.Inventory, error) {
	return s.repo.GetAll()
}

func (s *inventoryServiceImpl) GetInventoryByProductID(productID string) (*entity.Inventory, error) {
	return s.repo.GetByProductID(productID)
}

func (s *inventoryServiceImpl) UpdateInventory(productID string, inventory *entity.Inventory) error {
	return s.repo.Update(productID, inventory)
}

func (s *inventoryServiceImpl) DeleteInventory(productID string) error {
	return s.repo.Delete(productID)
}
