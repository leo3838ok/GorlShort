package shortener

import (
	"github.com/jinzhu/gorm"
)

type Repository interface {
	Create(m *Url) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(m *Url) error {
	return r.db.Table("url").Create(m).Error
}
