package entity

import (
  "time"
)

type Article struct {
  Id uint
  Title string `gorm:"size:128"`
  Category int
  Author string `gorm:"size:64"`
  CreatedAt time.Time
}
