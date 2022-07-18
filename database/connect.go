package database

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/rpparas/flight_log/config"
	"github.com/rpparas/flight_log/internals/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(strings.TrimSpace(p), 10, 32)

	if err != nil {
		log.Println("Invalid port")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Connection opened to database")

	// Migrate the database to the latest schema
	DB.AutoMigrate(&model.Flight{})
	DB.AutoMigrate(&model.Robot{})
	fmt.Println("Database Migrated")
}
