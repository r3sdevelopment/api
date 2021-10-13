package database

import (
	"fmt"
	"log"
	"os"

	"api/config"
	"api/pkg/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func Connect(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%s sslmode=%s TimeZone=%s", cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port, cfg.DB.SslMode, cfg.DB.Timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	log.Println("running migrations")
	db.AutoMigrate(&entities.Post{})

	DB = DbInstance{
		Db: db,
	}
}
