package db

import (
  "os"
	"gorm.io/driver/mysql"
  "gorm.io/gorm"
  "github.com/joho/godotenv"

  "entity"
)

var (
  db *gorm.DB
  err error
)

// DB初期化
func Init() {
  // 実行環境取得
  env := os.Getenv("ENV")

  if "production" == env {
    env = "production"
  } else {
    env = "development"
  }

  // 環境変数取得
  godotenv.Load(".env." + env)
  godotenv.Load()

  // DB接続
  db, err = gorm.Open("mysql", os.Getenv("CONNECT"))

  if err != nil {
    panic(err)
  }

  autoMigration()
}

// DB取得
func GetDB() *gorm.DB {
  return db
}

// DB接続終了
func Close() {
  if err := db.Close(); err != nil {
    panic(err)
  }
}

// マイグレーション
func autoMigration() {
  db.AutoMigrate(&entity.Article{})
}
