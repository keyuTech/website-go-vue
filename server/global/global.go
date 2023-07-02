package global

import (
	"gorm.io/gorm"
	"keyu.tech/website/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
