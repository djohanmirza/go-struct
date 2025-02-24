package repository

import (
	"github.com/djohanmirza/go-struct/entity"
	"gorm.io/gorm"
)

type ShipRepository interface {
	Create(ship *entity.Ship) error
	GetAll() ([]entity.Ship, error)
	GetByID(id uint64) (*entity.Ship, error)
	Update(id uint64, ship *entity.Ship) error
	Delete(id uint64) error
}

type shipRepositoryImpl struct {
	db *gorm.DB
}

func NewShipRepository(db *gorm.DB) ShipRepository {
	return &shipRepositoryImpl{db}
}

func (r *shipRepositoryImpl) Create(ship *entity.Ship) error {
	return r.db.Create(ship).Error
}

func (r *shipRepositoryImpl) GetAll() ([]entity.Ship, error) {
	var ships []entity.Ship
	err := r.db.Find(&ships).Error
	return ships, err
}

func (r *shipRepositoryImpl) GetByID(id uint64) (*entity.Ship, error) {
	var ship entity.Ship
	err := r.db.First(&ship, "id = ?", id).Error
	return &ship, err
}

func (r *shipRepositoryImpl) Update(id uint64, ship *entity.Ship) error {
	return r.db.Model(&entity.Ship{}).Where("id = ?", id).Updates(ship).Error
}

func (r *shipRepositoryImpl) Delete(id uint64) error {
	return r.db.Delete(&entity.Ship{}, "id = ?", id).Error
}
