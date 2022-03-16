package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Session struct {
	gorm.Model
	Token  string    `gorm:"primaryKey" json:"token"`
	Data   []byte    `gorm:"Not Null;type:bytea" json:"data"`
	Expiry time.Time `gorm:"index:sessions_expiry_idx;Not Null;type:TIMESTAMPTZ" json:"expiry"`
}

func RegisterSession(db *gorm.DB) {
	db.AutoMigrate(&Session{})
}
