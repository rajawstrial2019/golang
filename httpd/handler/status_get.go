package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "Alive")
	}
}
