package handler

import (
	"net/http"
	"newsfeed/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteNewsfeed(repo model.NewsfeedStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idString := ctx.Param("id")
		id, _ := strconv.ParseInt(idString, 10, 64)

		if ok := repo.Delete(id); ok {
			ctx.Status(http.StatusOK)
		} else {
			ctx.Status(http.StatusNotFound)
		}
	}
}
