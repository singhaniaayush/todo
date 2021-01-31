package utils

import (
	"todo/server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Db - database variable
var Db *gorm.DB

//ConnectDataBase - database connection
func ConnectDataBase() error {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=todo port=5432"

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}
	Db.AutoMigrate(&models.Todos{})
	return nil
}
