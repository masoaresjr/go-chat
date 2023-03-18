package main

import (
	db "chat/internal/config/database"
	migration "chat/internal/config/database/migration"
	"log"
)

func main() {

	postgresDB, err := db.Instance()
	if err != nil {
		log.Printf("Error DB: %s", err.Error())
	} else {
		log.Print("Ok")
	}

	migration.RunMigrations(postgresDB.GetDB())
}
