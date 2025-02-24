package services

import (
	"github.com/djohanmirza/go-struct/entity"
	"github.com/djohanmirza/go-struct/repository"
)

type ShipService interface {
	CreateShip(ship *entity.Ship) error
	GetAllShips() ([]entity.Ship, error)
	GetShipByID(id uint64) (*entity.Ship, error)
	UpdateShip(id uint64, ship *entity.Ship) error
	DeleteShip(id uint64) error
}

type shipServiceImpl struct {
	repo repository.ShipRepository
}

func NewShipService(repo repository.ShipRepository) ShipService {
	return &shipServiceImpl{repo}
}

func (s *shipServiceImpl) CreateShip(ship *entity.Ship) error {
	return s.repo.Create(ship)
}

func (s *shipServiceImpl) GetAllShips() ([]entity.Ship, error) {
	return s.repo.GetAll()
}

func (s *shipServiceImpl) GetShipByID(id uint64) (*entity.Ship, error) {
	return s.repo.GetByID(id)
}

func (s *shipServiceImpl) UpdateShip(id uint64, ship *entity.Ship) error {
	return s.repo.Update(id, ship)
}

func (s *shipServiceImpl) DeleteShip(id uint64) error {
	return s.repo.Delete(id)
}
