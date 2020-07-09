package shortener

import (
	"github.com/jinzhu/gorm"
)

type Repository interface {
	Create(m *Url) error
	Get(m *Url, id int) error
}

type repository struct {
	mySQL *gorm.DB
}

func NewRepository(mySQL *gorm.DB) Repository {
	return &repository{mySQL: mySQL}
}

func (r *repository) Create(m *Url) error {
	return r.mySQL.Table("url").Create(m).Error
}

func (r *repository) Get(m *Url, id int) error {
	return r.mySQL.Table("url").Where("id = ?", id).First(m).Error
}
