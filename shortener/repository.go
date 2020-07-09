package shortener

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	Create(m *Url) error
	GetUrl(id int) (string, error)
}

type repository struct {
	mySQL *gorm.DB
	redis *redis.Client
}

func NewRepository(mySQL *gorm.DB, redis *redis.Client) Repository {
	return &repository{
		mySQL: mySQL,
		redis: redis,
	}
}

func (r *repository) Create(m *Url) error {
	return r.mySQL.Table("url").Create(m).Error
}

func (r *repository) GetUrl(id int) (string, error) {
	key := fmt.Sprintf("url:%d", id)
	val := r.redis.Get(key).Val()
	if val != "" {
		return val, nil
	}

	// Query from database
	var url Url
	if err := r.mySQL.Table("url").Where("id = ?", id).First(&url).Error; err != nil {
		return "", err
	}

	// Set value into redis
	if err := r.redis.Set(key, url.URL, 0).Err(); err != nil {
		return "", err
	}

	return url.URL, nil
}
