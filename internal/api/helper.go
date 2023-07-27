package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type okResponse struct {
	OK bool `json:"ok"`
}

type CustomError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *CustomError) Error() string {
	return e.Msg
}

func NewError(code int, msg string) *CustomError {
	return &CustomError{
		Code: code,
		Msg:  msg,
	}
}

func errResponse(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, okResponse{
		OK: false,
	})
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, e := range c.Errors {
			err := e.Err
			if cusErr, ok := err.(*CustomError); ok {
				c.AbortWithStatusJSON(cusErr.Code, gin.H{
					"ok":  false,
					"msg": cusErr.Msg,
				})
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"ok":  false,
					"msg": "internal server error",
				})
			}
			return
		}
	}
}
