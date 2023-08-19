package config

import (
	"fmt"
	"os"

	"github.com/anasmisbah/backend-inventory/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB()  {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
	os.Getenv("DB_ROOT"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_NAME"))

	var e error
	DB,e = gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if e!=nil {
		panic(e)
	}
	initMiggration()
}

func initMiggration()  {
	DB.AutoMigrate(&models.Product{},&models.Stock{})
}