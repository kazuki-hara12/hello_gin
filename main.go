package main

import (
  "github.com/gin-gonic/gin"
	"hello_gin/db"
	"net/http"
	"hello_gin/entity"
)

func main() {
	db.Init()
	r := gin.Default()

	r.GET("/articles", func(c *gin.Context) {
		a := db.GetDB()
		var article []entity.Article
		result := a.Find(&article)
		c.JSON(http.StatusOK, result)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
