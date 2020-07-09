package shortener

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leo3838ok/go-url-shortener/internal/utils/base62"
)

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := Url{URL: req.URL}
	if err := h.repo.Create(&url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":  url.ID,
			"url": fmt.Sprintf("%s/%s", c.Request.Host, base62.Encode(url.ID)),
		},
	})
}
