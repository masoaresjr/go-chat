package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db *gorm.DB
}

func Instance() (*PostgresDB, error) {

	dsn := "host=localhost user=root password=password dbname=go-chat port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &PostgresDB{db: db}, nil
}

func (d *PostgresDB) Close() {
	postgresDB, err := d.db.DB()
	if err != nil {
		log.Fatal("Error closing DB: ", err.Error())
	}
	postgresDB.Close()
}

func (d *PostgresDB) GetDB() *gorm.DB {
	return d.db
}
