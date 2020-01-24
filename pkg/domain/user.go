package domain

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model // DBメタ情報構造体
	Name       string
	Posts      []Post
}
