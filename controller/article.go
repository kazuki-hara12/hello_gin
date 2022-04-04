package controller

import (
	"github.com/gin-gonic/gin"
  "hello_gin/db"
	"net/http"
	"hello_gin/entity"
)

func ArticleIndex(c *gin.Context) {
	a := db.GetDB()
	var article []entity.Article
	a.Find(&article)
	c.JSON(http.StatusOK, article)
}

func ArticleShow(c *gin.Context) {
	id := c.Params.ByName("id")
	a := db.GetDB()
	var article entity.Article
	a.Where("id = ?", id).First(&article)
	c.JSON(http.StatusOK, article)
}
