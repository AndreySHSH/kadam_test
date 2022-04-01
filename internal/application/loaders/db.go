package loaders

import (
	"fmt"
	"github.com/ivahaev/go-logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// NewDataBase - init database
func NewDataBase() (*gorm.DB, error) {
	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow`,
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DBNAME"),
		os.Getenv("POSTGRES_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	logger.Info("Connect DB")

	return db, nil
}
