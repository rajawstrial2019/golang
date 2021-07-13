package main

import (
	"newsfeed/httpd/handler"
	"newsfeed/model"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := model.New()

	router := gin.Default()

	router.GET("/", handler.GetStatus())
	router.GET("/newsfeed", handler.GetAllNewsfeed(repo))
	router.POST("/newsfeed", handler.PostNewsfeed(repo))
	router.GET("/newsfeed/:id", handler.GetNewsfeed(repo))
	router.PUT("/newsfeed/:id", handler.PutNewsfeed(repo))
	router.DELETE("/newsfeed/:id", handler.DeleteNewsfeed(repo))

	router.Run(":8080")
}
