package entity

type Article struct {
  Id uint   `json:"id"`
  Title string `json:"title";gorm:"size:128"`
  Description string `json:"description";gorm:"size:128"`
  Body string `json:"body";gorm:"size:128"`
}
