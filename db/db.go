package db

import (
  "os"
	"gorm.io/driver/mysql"
  "gorm.io/gorm"
  "github.com/joho/godotenv"
  "hello_gin/entity"
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

  dsn := "root:password@tcp(localhost)/sample?charset=utf8&parseTime=True&loc=Local"

  // DB接続
  db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
    panic(err)
  }

  autoMigration()
}

// DB取得
func GetDB() *gorm.DB {
  return db
}

// マイグレーション
func autoMigration() {
  db.AutoMigrate(&entity.Article{})
}
