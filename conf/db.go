package conf

import (
	"github.com/jimmyann2023/Gin/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func InitDB() (*gorm.DB, error) {
	loggerLeve := logger.Info
	if !viper.GetBool("mode.develop") {
		loggerLeve = logger.Error
	}

	db, err := gorm.Open(mysql.Open(viper.GetString("db.dsn")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_", // 表名前缀
			SingularTable: true,   // 取消表名复数形式 就是表名 +s
		},
		Logger: logger.Default.LogMode(loggerLeve), // 输出mysql 语句
	})

	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(viper.GetInt("db.maxIdleConn"))
	sqlDB.SetMaxOpenConns(viper.GetInt("db.maxOpenConn"))
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(
		&model.User{},
	)

	return db, nil
}
