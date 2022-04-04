package main

import (
  "github.com/gin-gonic/gin"
	"hello_gin/db"
	"hello_gin/controller"
	"hello_gin/job"
	"net/http"
	"github.com/bamzi/jobrunner"
	"github.com/fvbock/endless"
)

func main() {
	db.Init()
	r := gin.Default()

	jobrunner.Start()
	jobrunner.Schedule("@every 5m", job.ContentfulJob{})

	articles := r.Group("/articles")
	{
		articles.GET("", controller.ArticleIndex)
		articles.GET("/:id", controller.ArticleShow)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	endless.ListenAndServe(":8080", r)
}
