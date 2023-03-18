package db

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/gorm"
)

var migrations = []*gormigrate.Migration{
	&ID202303172004DDlInitialSchemas,
	// &ID202303172028DmlAddMessageStatusTypes,
}

func RunMigrations(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
