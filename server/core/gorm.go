package core

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"keyu.tech/website/global"
)

func InitGorm() *gorm.DB {
	if global.Config.PostgreSQL.Host == "" {
		global.Log.Warnln("未配置postgres，取消gorm连接")
		return nil
	}

	dsn := global.Config.PostgreSQL.Dsn()

	var postgreLogger logger.Interface
	if global.Config.System.Env == "debug" {
		postgreLogger = logger.Default.LogMode(logger.Info)
	} else {
		postgreLogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: postgreLogger,
	})

	if err != nil {
		global.Log.Fatalf(fmt.Sprintf("[%s] postgress connect fail", dsn))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)

	return db

}
