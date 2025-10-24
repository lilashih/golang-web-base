package controller

import (
	"fmt"
	"gbase/src/core/http/request"
	"gbase/src/core/http/response"
	"gbase/src/def"
	"gbase/src/http/resource"
	m "gbase/src/model"
	"gbase/src/repository"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type SettingController struct {
	repo *repository.SettingRepository
	res  resource.SettingResource
}

func NewSettingController() *SettingController {
	repo := repository.NewSettingRepository()
	res := resource.SettingResource{}

	return &SettingController{repo: repo, res: res}
}

func (ctrl *SettingController) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/settings", ctrl.SettingIndex)
	r.GET("/settings/:group", ctrl.SettingShow)
	r.PUT("/settings/:group", ctrl.SettingUpdate)
}

// @Tags Setting 設定
// @Router /settings [get]
// @Summary 獲取設定
// @Description 獲取系統設定，依權限回傳可存取的設定
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{data=resource.Settings} "code:200"
func (ctrl *SettingController) SettingIndex(c *gin.Context) {
	list, pagination, err := ctrl.repo.FindAll(nil)
	if err != nil {
		response.Error(c, nil, err.Error())
		return
	}

	response.Success(c, ctrl.res.Collection(pagination, list))
}

// @Tags Setting 設定
// @Router /settings/{group} [get]
// @Summary 獲取設定資訊
// @Accept  json
// @Produce  json
// @Param   group path string true "群組"
// @Success 200 {object} response.Response{data=resource.Settings} "code:200"
func (ctrl *SettingController) SettingShow(c *gin.Context) {
	group := c.Param("group")

	list, pagination, err := ctrl.repo.FindAll(map[string]string{"group": group})
	if err != nil {
		response.Error(c, nil, err.Error())
		return
	}

	response.Success(c, ctrl.res.Collection(pagination, list))
}

// @Tags Setting 設定
// @Router /settings/{group} [put]
// @Summary 編輯設定
// @Accept  json
// @Produce  json
// @Param   group path string true "群組"
// @Param  data body model.SettingInputs true "設定資訊"
// @Success 200 {object} response.Response{data=resource.Setting} "code:200, message:編輯成功"
// @Failure 400 {object} response.Response{} "code:400, message:編輯失敗"
// @Failure 404 {object} response.Response{} "code:404, message:查無該資料"
// @Failure 422 {object} response.Response{data=resource.ErrorValidation} "code:422 (表單資料有誤)"
func (ctrl *SettingController) SettingUpdate(c *gin.Context) {
	group := c.Param("group")

	// get body data
	body := ctrl.repo.NewInput().(m.SettingInputs)
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		response.Error(c, nil, fmt.Sprintf("%s: %v", def.BIND_BODY_FAILED, err))
		return
	}

	// 驗證
	for _, item := range body {
		if err := request.Validate.Struct(item); err != nil {
			response.ErrorValidation(c, err)
			return
		}
	}

	err := ctrl.repo.UpdateSettings(body, group)
	if err != nil {
		response.Error(c, nil, fmt.Sprintf("%s: %v", def.UPDATE_FAILED, err))
		return
	}

	response.Success(c, nil, def.UPDATE_SUCCESS)
}
