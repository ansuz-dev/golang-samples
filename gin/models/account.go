package models

import (
  "time"
)

type Account struct {
  ID           int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
  Username     string    `gorm:"type:VARCHAR(256);NOT NULL" json:"username"`
  PasswordHash string    `gorm:"type:VARCHAR(64);NOT NULL" json:"-"`
  State        string    `gorm:"type:VARCHAR(64);NOT NULL" json:"state"`
  CreatedAt    time.Time `gorm:"NOT NULL" json:"createdAt"`
  UpdatedAt    time.Time `gorm:"NOT NULL" json:"updatedAt"`
  Posts        []Post    `gorm:"foreignkey:AccountId" json:"posts,omitempty"`
}
