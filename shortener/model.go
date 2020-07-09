package shortener

type Url struct {
	ID  int    `gorm:"AUTO_INCREMENT"`
	URL string `gorm:"column:url"`
}
