package handler

import (
	"net/http"
	"newsfeed/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PutNewsfeed(repo model.NewsfeedStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idString := ctx.Param("id")
		id, _ := strconv.ParseInt(idString, 10, 64)

		requestBody := newsfeedRequest{}
		ctx.Bind(&requestBody)

		post := model.Post{
			Title: requestBody.Title,
			Post:  requestBody.Post,
			ID:    id,
		}

		if ok := repo.Update(post); ok {
			ctx.JSON(http.StatusOK, post)
		} else {
			ctx.Status(http.StatusNotFound)
		}
	}
}
