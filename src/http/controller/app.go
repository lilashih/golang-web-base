package controller

import (
	"gbase/src/core/config"
	"gbase/src/core/http/response"
	"gbase/src/http/resource"

	"github.com/gin-gonic/gin"
)

type AppController struct {
	res resource.IResource
}

func NewAppController() *AppController {
	return &AppController{
		res: resource.AppResource{},
	}
}

func (ctrl *AppController) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/app/configs", ctrl.AppIndex)
}

// @Tags 系統配置
// @Router /app/configs [get]
// @Summary 獲取系統配置
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{data=resource.Configs} "code:200"
func (ctrl *AppController) AppIndex(c *gin.Context) {
	configs := map[string]string{
		"version": config.App.Version,
		"mode":    config.App.Mode,
	}

	response.Success(c, ctrl.res.Collection(resource.Pagination{}, configs))
}
