package handler

import (
	"net/http"
	"newsfeed/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetNewsfeed(repo model.NewsfeedStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idString := ctx.Param("id")
		id, _ := strconv.ParseInt(idString, 10, 64)

		post1 := repo.Read(id)
		if post1.ID > 0 {
			ctx.JSON(http.StatusOK, post1)
		} else {
			ctx.Status(http.StatusNotFound)
		}
	}
}
