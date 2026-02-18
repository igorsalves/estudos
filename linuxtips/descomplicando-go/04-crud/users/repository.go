package users

import "gorm.io/gorm"

type GormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &GormRepository{db: db}
}

type Repository interface {
	Create(name string) uint
	Update(id uint, name string)
	Delete(id uint)
	List() []User
}
