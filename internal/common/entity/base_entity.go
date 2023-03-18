package entity

import (
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	ID          int            `json:"id"            gorm:"primary_key;autoIncrement"`
	DateCreated time.Time      `json:"date_created"  gorm:"DEFAULT:current_timestamp;not null"`
	LastUpdeted time.Time      `json:"last_updated"  gorm:"DEFAULT:current_timestamp;not null"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
