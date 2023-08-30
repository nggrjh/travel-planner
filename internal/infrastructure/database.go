package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type IDatabase interface {
	Close()
	AutoMigrate(values ...interface{})
}

type database struct {
	conn *gorm.DB
}

func NewDatabaseConnection() (IDatabase, error) {
	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s dbname=%s user=%s password=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
		),
	)
	if err != nil {
		return nil, err
	}
	return &database{conn: db}, nil
}

func (db *database) Close() {
	if err := db.conn.Close(); err != nil {
		log.Printf("Failed to close database: %s\n", err.Error())
	}
}

func (db *database) AutoMigrate(values ...interface{}) {
	db.conn = db.conn.AutoMigrate(values...)
}
