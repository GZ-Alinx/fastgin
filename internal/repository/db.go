package repository

import (
	"fastgin/config"
	"fastgin/internal/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (error, *gorm.DB) {
	conf := config.GlobalConfig.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DBName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 自动迁移
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		panic("数据库迁移失败: " + err.Error())
	}

	// 创建默认管理员账户
	var adminCount int64
	DB.Model(&model.User{}).Where("role = ?", "admin").Count(&adminCount)
	if adminCount == 0 {
		admin := model.User{
			Username: "admin",
			Password: "123456",
			Role:     "admin",
		}
		admin.HashPassword()
		DB.Create(&admin)
	}

	return err, DB
}
