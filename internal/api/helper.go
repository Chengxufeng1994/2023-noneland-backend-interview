package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type okResponse struct {
	OK bool `json:"ok"`
}

func errResponse(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, okResponse{
		OK: false,
	})
}
