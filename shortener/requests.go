package shortener

type CreateReq struct {
	URL string `form:"url" binding:"required"`
}
