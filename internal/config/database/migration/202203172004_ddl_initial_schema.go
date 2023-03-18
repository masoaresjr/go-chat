package db

import (
	"chat/internal/common/entity"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// ID202303172004DDlInitialSchemas godoc
var ID202303172004DDlInitialSchemas = gormigrate.Migration{
	ID: "202303172004",
	Migrate: func(db *gorm.DB) error {

		// Create tables
		err := db.Migrator().AutoMigrate(
			&entity.User{},
			&entity.Room{},
			&entity.RoomUser{},
			&entity.Message{},
			&entity.MessageStatusType{},
			&entity.MessageStatusUser{},
		)

		if err != nil {
			log.Fatal(err)
			return err
		}

		return nil
	},
	Rollback: func(db *gorm.DB) error {

		// Rollback Drop tables
		err := db.Migrator().DropTable(
			&entity.User{},
			&entity.Room{},
			&entity.RoomUser{},
			&entity.Message{},
			&entity.MessageStatusType{},
			&entity.MessageStatusUser{},
		)

		if err != nil {
			log.Fatal(err)
			return err
		}

		return nil
	},
}
