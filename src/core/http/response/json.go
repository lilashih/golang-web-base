package response

import (
	"gbase/src/core/http/request"
	"gbase/src/core/logger"
	"gbase/src/http/resource"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"` // 跟http code一樣
	Message string `json:"message"`
	Data    any    `json:"data"`
} //@name Response

func formatData(data any) any {
	switch data.(type) {
	case int, string, nil:
		return []string{}
	default:
		return data
	}
}

func getCode(defaultCode int, codes ...int) int {
	if len(codes) > 0 {
		return codes[0]
	}
	return defaultCode
}

func response(c *gin.Context, data any, message string, code int) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    formatData(data),
	})
}

func Success(c *gin.Context, data any, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	response(c, data, message, http.StatusOK)
}

func Error(c *gin.Context, data any, message string, codes ...int) {
	response(c, data, message, getCode(http.StatusBadRequest, codes...))
}

func ErrorNotFound(c *gin.Context) {
	response(c, nil, "查無該資料", http.StatusNotFound)
}

func ErrorValidation(c *gin.Context, errValidate error) {
	code := http.StatusUnprocessableEntity

	errorBag, err := request.FormateErrorBag(errValidate)
	if err != nil {
		logger.Log.Println(err)
		response(c, nil, err.Error(), code)
		return
	}

	response(c, resource.ErrorValidation{Errors: errorBag}, "", code)
}
