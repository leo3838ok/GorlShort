package server

import (
	"github.com/gin-gonic/gin"
	"github.com/leo3838ok/go-url-shortener/internal/mysql"
	"github.com/leo3838ok/go-url-shortener/internal/redis"
	"github.com/leo3838ok/go-url-shortener/shortener"
)

func Register(r *gin.Engine) {
	handler := shortener.NewHandler(shortener.NewRepository(mysql.MySQL, redis.Redis))

	root := r.Group("/")
	{
		root.GET("/:code", handler.Resolve)
	}

	v1 := r.Group("/v1")
	{
		v1.POST("/create", handler.Create)
	}
}
