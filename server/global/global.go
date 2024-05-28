package global

import (
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	Once = &singleflight.Group{}
)
