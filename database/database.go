package database

import (
	"fmt"
	"github.com/eunice99x/jwt-gofiber/configs"
	"github.com/eunice99x/jwt-gofiber/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(migrate bool) *gorm.DB {
	configs.InitEnvConfigs()
	c := configs.EnvConfigs

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "localhost", c.DatabaseUser, c.DatabasePass, c.DatabaseName, "5432")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}
	if migrate {
		db.AutoMigrate(models.Product{}, models.User{}, models.User{})
	}
	return db
}
