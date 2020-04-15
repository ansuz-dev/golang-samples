package models

import (
  "time"
)

type Post struct {
  ID        int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
  AccountId int       `gorm:"NOT NULL" json:"accountId"`
  Content   string    `gorm:"NOT NULL" json:"content"`
  CreatedAt time.Time `gorm:"NOT NULL" json:"createdAt"`
  UpdatedAt time.Time `gorm:"NOT NULL" json:"updatedAt"`
}
