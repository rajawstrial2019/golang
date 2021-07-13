package handler

import (
	"net/http"
	"newsfeed/model"

	"github.com/gin-gonic/gin"
)

type newsfeedRequest struct {
	Title string `json:"title"`
	Post  string `json:"message"`
}

func PostNewsfeed(repo model.NewsfeedStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestBody := newsfeedRequest{}
		ctx.Bind(&requestBody)

		post := model.Post{
			Title: requestBody.Title,
			Post:  requestBody.Post,
			ID:    repo.NextId(),
		}
		repo.Add(post)
		ctx.JSON(http.StatusCreated, post)
	}
}
