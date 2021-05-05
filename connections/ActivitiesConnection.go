package connections

import (
	"github.com/sam-samid/go-refview-activities-module/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	dsn := "root:root@(db_mysql_57)/refview?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("connection refused")
	}

	db.AutoMigrate(&models.Activity{})

	return db
}
