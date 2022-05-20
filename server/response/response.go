package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context, httpStatus int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": httpStatus,
		"data": data,
	})
}

func Success(c *gin.Context, data interface{}) {
	Response(c, http.StatusOK, data)
}

func Fail(c *gin.Context, data interface{}) {
	Response(c, http.StatusUnprocessableEntity, data)
}
