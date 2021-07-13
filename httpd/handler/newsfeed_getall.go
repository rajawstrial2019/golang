package handler

import (
	"net/http"
	"newsfeed/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllNewsfeed(repo model.NewsfeedStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idString := ctx.Query("start")
		startId, _ := strconv.ParseInt(idString, 10, 64)

		idString = ctx.Query("end")
		endId, _ := strconv.ParseInt(idString, 10, 64)

		time.Sleep(500 * time.Millisecond)

		ctx.JSON(http.StatusOK, repo.GetAll(startId, endId))
	}
}
