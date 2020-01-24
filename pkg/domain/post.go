package domain

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model // DBメタ情報構造体
	comment    string
	UserID     uint // UserへのFK
}
