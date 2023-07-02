package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"keyu.tech/website/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
