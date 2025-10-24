package request

import (
	"github.com/gin-gonic/gin"
)

func GetParamId(c *gin.Context) string {
	return c.Param("id")
}
