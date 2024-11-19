package database

import (
	"log"

	model "github.com/sidikimamsetiyawan/go-project-library-app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnectDB() {

	dsn := "root:@tcp(127.0.0.1:3306)/fiber_libraries_app?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("Database connection failed.")
	}

	log.Println("Connection succesfully.")

	db.AutoMigrate(new(model.Blog)) // Find uses model.blog link to github

	DBConn = db
}
