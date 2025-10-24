package controller

import (
	"gbase/src/core/http/response"
	"gbase/src/http/resource"
	"gbase/src/repository"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	repo *repository.PermissionRepository
	res  resource.MenuResource
}

func NewMenuController() *MenuController {
	repo := repository.NewPermissionRepository()
	res := resource.MenuResource{}

	return &MenuController{repo: repo, res: res}
}

func (ctrl *MenuController) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/menus", ctrl.MenuIndex)
}

// @Tags Menu 菜單
// @Router /menus [get]
// @Summary 獲取菜單
// @Description 獲取系統菜單，依權限回傳可存取的菜單
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{data=resource.Menus} "code:200"
func (ctrl *MenuController) MenuIndex(c *gin.Context) {
	list, err := ctrl.repo.FindMenu()
	if err != nil {
		response.Error(c, nil, err.Error())
		return
	}

	response.Success(c, ctrl.res.Single(list))
}
